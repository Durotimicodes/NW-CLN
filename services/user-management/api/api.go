package api

import (
	"context"
	"log"
	"net/http"
	"strconv"
	"time"

	"github.com/durotimicodes/natwest-clone/user-service/models"
	"github.com/durotimicodes/natwest-clone/user-service/service"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
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

	// Bind JSON request body to user struct with proper error handling
	if err := ctx.ShouldBindJSON(&user); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request data"})
		return
	}

	// Use context with timeout for better request handling
	ctxWithTimeout, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel() //defer cancel to prevent memory leaks

	// Use goroutine with error channel for concurrent execution
	errChan := make(chan error, 1)

	go func() {
		defer close(errChan)

		// Load existing users
		newUsers, err := h.UserService.LoadUser()
		if err != nil {
			log.Println("Failed to load users:", err)
			errChan <- err
			return
		}

		// Assign unique ID and append the new user
		user.ID = uint(len(newUsers) + 1)
		newUsers = append(newUsers, user)

		// Save users with proper error handling
		if err := h.UserService.SaveUser(newUsers); err != nil {
			log.Println("Failed to save users:", err)
			errChan <- err
			return
		}

		errChan <- nil
	}()

	// Listen for error or timeout
	select {
	case err := <-errChan:
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to save user"})
			return
		}
		ctx.JSON(http.StatusCreated, gin.H{
			"message": "User registered successfully",
			"data":    user,
		})
	case <-ctxWithTimeout.Done():
		ctx.JSON(http.StatusGatewayTimeout, gin.H{"error": "Request timed out"})
	}

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

// Secret key for JWT generation (in production, store this securely)
var jwtSecret = []byte("your_secret_key")

// Authentication
func (h *UserHandler) LoginUser(c *gin.Context) {
	var req models.LoginRequest

	//Bind JSON request to struct
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request data"})
		return
	}

	//Mock authentication (replace with real DB lookup)
	user, err := h.UserService.AuthenticateUser(req.Email, req.Password)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid credentials"})
		return
	}

	//Generate JWT token
	expirationTime := time.Now().Add(24 * time.Hour) //1-day token expiration
	claims := &models.Claims{
		Email: user.Email,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(expirationTime),
		},
	}

	// Generate token
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString(jwtSecret)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to generate token"})
		return
	}

	// Respond with token
	c.JSON(http.StatusOK, gin.H{
		"message": "Login successful",
		"token":   tokenString,
		"user": gin.H{
			"id":    user.ID,
			"email": user.Email,
			"name":  user.FullName,
		},
	})

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
