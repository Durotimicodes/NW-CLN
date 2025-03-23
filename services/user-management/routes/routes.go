package routes

import (
	"github.com/durotimicodes/natwest-clone/user-service/api"
	"github.com/durotimicodes/natwest-clone/user-service/repository"
	"github.com/durotimicodes/natwest-clone/user-service/service"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func SetUpUserRoutes(db *gorm.DB) *gin.Engine {
	router := gin.Default()

	// Initialize dependencies
	userRepo := repository.NewUserRepository(db)
	userService := service.NewUserService(*userRepo)
	userHandler := api.NewUserHandler(userService)

	// Group the routes
	userGroup := router.Group("/api/v1/user")

	// User Authentication & Management
	{
		userGroup.POST("/ping", userHandler.HeartBeat)
		userGroup.POST("/register", userHandler.RegisterUser)      // Register a new user
		userGroup.POST("/login", userHandler.LoginUser)            // Authenticate user & generate token
		userGroup.POST("/logout", userHandler.LogoutUser)          // Logout user
		userGroup.POST("/refresh-token", userHandler.RefreshToken) // Refresh JWT token
	}

	// User Profile Management
	{
		userGroup.GET("/profile", userHandler.GetUserProfile)                   // Get user profile
		userGroup.PUT("/update-profile", userHandler.UpdateUserProfile)         // Update user profile
		userGroup.PATCH("/profile/change-password", userHandler.ChangePassword) // Change user password
		userGroup.PATCH("/profile/update-email", userHandler.UpdateEmail)       // Update user email
	}

	// Account Management (Admin only)
	{
		userGroup.GET("/:id", userHandler.GetUserByIDHandler)   // Get user by ID (Admin only)
		userGroup.GET("/", userHandler.GetAllUsers)             // Get all users (Admin only)
		userGroup.DELETE("/delete/:id", userHandler.DeleteUser) // Delete a user (Admin only)
	}

	// Security and Verification
	{
		userGroup.POST("/verify-email", userHandler.VerifyEmail)                        // Verify email with OTP
		userGroup.POST("/send-verification-email", userHandler.ResendVerificationEmail) // Resend verification email
		userGroup.POST("/forgot-password", userHandler.ForgotPassword)                  // Send password reset link
		userGroup.POST("/reset-password", userHandler.ResetPassword)                    // Reset password with token
	}

	//Start the server
	return router
}
