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

// RegisterUser handles user registration HTTP requests
func (h *UserHandler) RegisterUser(ctx *gin.Context) {
	var user models.User

	//Bind JSON request body to user struct
	if err := ctx.ShouldBindJSON(&user); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"err": "Invalid request data"})
		return
	}

	//Register user via service
	if err := h.UserService.RegisterUser(&user); err != nil {
		ctx.JSON(http.StatusConflict, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusCreated, gin.H{
		"message": "User registered successfully",
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
