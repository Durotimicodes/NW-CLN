package api

import (
	"net/http"
	"strconv"

	"github.com/durotimicodes/natwest-clone/user-service/models"
	"github.com/durotimicodes/natwest-clone/user-service/service"
	"github.com/gin-gonic/gin"
)

type UserHandler struct {
	UserService *service.UserService
}

// NewUserHandler creates a new UserHandler instance
func NewUserHandler(userService *service.UserService) *UserHandler {
	return &UserHandler{UserService: userService}
}

func (h *UserHandler) HeartBeat(ctx *gin.Context) {

	if ctx.Request.Method != http.MethodPost {
		ctx.JSON(http.StatusMethodNotAllowed, gin.H{
			"error": "Method not allowed",
		})
	}
	ctx.JSON(http.StatusOK, gin.H{
		"message": "pong",
	})
}

// RegisterUser handles user registration HTTP requests
func (h *UserHandler) RegisterUser(ctx *gin.Context) {
	var user models.User

	//Bind JSON request body to user struct
	if err := ctx.ShouldBindJSON(&user); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"err": "Invalid request data"})
		return
	}

	//Use channels to process request safely
	resultChan := make(chan error)

	go func() {
		newUsers, err := h.UserService.LoadUser()
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{
				"error": "failed to load user",
			})
			resultChan <- err
			return
		}
		user.ID = uint(len(newUsers) + 1)
		newUsers = append(newUsers, user)

		if err := h.UserService.SaveUser(newUsers); err != nil {
			resultChan <- err
			return
		}
	}()

	//Register user via service
	// if err := h.UserService.RegisterUser(&user); err != nil {
	// 	ctx.JSON(http.StatusConflict, gin.H{"error": err.Error()})
	// 	return
	// }

	if err := <-resultChan; err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": "failed to save user",
		})
	}

	ctx.JSON(http.StatusCreated, gin.H{
		"message": "User registered successfully",
		"data":    user,
	})
}

// GetUserByID fetches a user by ID
func (h *UserHandler) GetUserByIDHandler(ctx *gin.Context) {
	idParam := ctx.Param("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error":   "Invalid user ID",
			"message": nil,
		})
		return
	}

	user, err := h.UserService.GetUserByID(uint(id))
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{
			"error":   "User not found",
			"messgae": nil,
		})
		return
	}

	ctx.JSON(http.StatusOK, user)
}

// Authentication
func (h *UserHandler) LoginUser(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "LoginUser endpoint reached"})
}

func (h *UserHandler) LogoutUser(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "LogoutUser endpoint reached"})
}

func (h *UserHandler) RefreshToken(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "RefreshToken endpoint reached"})
}

// User Profile Management
func (h *UserHandler) GetUserProfile(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "GetUserProfile endpoint reached"})
}

func (h *UserHandler) UpdateUserProfile(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "UpdateUserProfile endpoint reached"})
}

func (h *UserHandler) ChangePassword(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "ChangePassword endpoint reached"})
}

func (h *UserHandler) UpdateEmail(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "UpdateEmail endpoint reached"})
}

// Account Management (Admin only)

func (h *UserHandler) GetAllUsers(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "GetAllUsers endpoint reached"})
}

func (h *UserHandler) DeleteUser(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "DeleteUser endpoint reached"})
}

// Security and Verification
func (h *UserHandler) VerifyEmail(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "VerifyEmail endpoint reached"})
}

func (h *UserHandler) ResendVerificationEmail(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "ResendVerificationEmail endpoint reached"})
}

func (h *UserHandler) ForgotPassword(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "ForgotPassword endpoint reached"})
}

func (h *UserHandler) ResetPassword(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "ResetPassword endpoint reached"})
}
