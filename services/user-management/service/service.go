package service

import (
	"errors"

	"github.com/durotimicodes/natwest-clone/user-service/models"
	"github.com/durotimicodes/natwest-clone/user-service/repository"
	"github.com/durotimicodes/natwest-clone/user-service/utils"
	"golang.org/x/crypto/bcrypt"
)

type UserService struct {
	UserRepo repository.UserRepository
}

// NewUserService creates an instance of a User service
func NewUserService(userRepo repository.UserRepository) *UserService {
	return &UserService{
		UserRepo: userRepo,
	}
}

// RegisterUser handles user registeration with validation & password hashing
func (s *UserService) RegisterUser(user *models.User) error {

	//Check if user with email already exists
	existingUser, _ := s.UserRepo.FindByEmail(user.Email)
	if existingUser != nil {
		return errors.New("user with this email already exists")
	}

	//Hash password before saving
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		return errors.New("failed to hash password")
	}

	user.Password = string(hashedPassword)

	// Generate Account Number & Sort Code
	user.AccountNumber = utils.GenerateAccountNumber()
	user.SortCode = utils.GenerateSortCode()

	// Generate IBAN
	user.IBAN = utils.GenerateIBAN(user.SortCode, user.AccountNumber)

	// Encrypt Account Number, Sort Code, and IBAN before storing
	encryptedAccountNumber, err := utils.EncryptData(user.AccountNumber)
	if err != nil {
		return errors.New("failed to encrypt account number")
	}
	user.AccountNumber = encryptedAccountNumber

	encryptedSortCode, err := utils.EncryptData(user.SortCode)
	if err != nil {
		return errors.New("failed to encrypt sort code")
	}
	user.SortCode = encryptedSortCode

	encryptedIBAN, err := utils.EncryptData(user.IBAN)
	if err != nil {
		return errors.New("failed to encrypt IBAN")
	}
	user.IBAN = encryptedIBAN

	return s.UserRepo.CreateUser(user)

}

// GetUserByID service gets user by ID returns a user or an error
func (s *UserService) GetUserByID(id uint) (*models.User, error) {
	return s.UserRepo.FindUserByID(id)
}
