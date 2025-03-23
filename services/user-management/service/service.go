package service

import (
	"encoding/json"
	"errors"
	"os"

	"github.com/durotimicodes/natwest-clone/user-service/models"
	"github.com/durotimicodes/natwest-clone/user-service/repository"
	"github.com/durotimicodes/natwest-clone/user-service/utils"
	"golang.org/x/crypto/bcrypt"
)

type UserService struct {
	UserRepo repository.UserRepository
}

const filePath = "user.json"

// NewUserService creates an instance of a User service
func NewUserService(userRepo repository.UserRepository) *UserService {
	return &UserService{
		UserRepo: userRepo,
	}
}


// RegisterUser handles user registration with validation & password hashing
func (s *UserService) RegisterUser(user *models.User) error {
	// Check if user with email already exists
	existingUser, err := s.UserRepo.FindByEmail(user.Email)
	if err != nil {
		return err
	}
	if existingUser != nil {
		return errors.New("user with this email already exists")
	}

	// Check if user with phone number already exists
	existingUser2, err := s.UserRepo.FindUserByPhoneNumber(user.PhoneNumber)
	if err != nil {
		return err
	}
	if existingUser2 != nil {
		return errors.New("user with this phone number already exists")
	}

	// Hash password before saving
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

// File path for loading and saving users
const filePaths = "user.json"

// LoadUser loads the users from the file
func (s *UserService) LoadUser() ([]models.User, error) {
	file, err := os.Open(filePaths)
	if err != nil {
		if os.IsNotExist(err) {
			return []models.User{}, nil // Return an empty user list if the file doesn't exist
		}
		return nil, err
	}
	defer file.Close()

	var users []models.User
	decoder := json.NewDecoder(file)
	err = decoder.Decode(&users)
	if err != nil {
		return nil, errors.New("failed to decode users from file")
	}

	return users, nil
}

// SaveUser saves the users to the file
func (s *UserService) SaveUser(users []models.User) error {
	data, err := json.MarshalIndent(users, "", "  ")
	if err != nil {
		return err
	}

	return os.WriteFile(filePath, data, 0644)
}
