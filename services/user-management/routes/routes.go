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

	//Initialize dependencies
	userRepo := repository.NewUserRepository(db)
	userService := service.NewUserService(*userRepo)
	userHandler := api.NewUserHandler(userService)

	//Group the routes
	userGroup := router.Group("/api/v1/users")

	//User Authentication & Management
	{
		userGroup.POST("/register")      //register a new user
		userGroup.POST("/login")         //Authenticate user & generate token
		userGroup.POST("/logout")        //Logout user
		userGroup.POST("/refresh-token") //refresh JWT token
	}

	//User Profile Management
	{
		userGroup.GET("/profile")                   //get user profile
		userGroup.PUT("/update-profile")            // update user profile
		userGroup.PATCH("/profile/change-password") //change user password
		userGroup.PATCH("/profile/update-email")    //update user email
	}

	//Account Management
	{
		userGroup.GET("/:id", userHandler.GetUserByIDHandler) //get user by id (admin only)
		userGroup.GET("")                                     //get all users (admin only)
		userGroup.DELETE("delete/:id")                        // delete a users (admin only)
	}

	//Security and Verification
	{
		userGroup.POST("/verify-email")            //verify email with OTP
		userGroup.POST("/send-verification-email") //resend verification email
		userGroup.POST("/forgot-password")         //send password reset link
		userGroup.POST("/reset-password")          //reset password with token
	}

	return router

}
