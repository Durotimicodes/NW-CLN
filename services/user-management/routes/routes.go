package routes

import "github.com/gin-gonic/gin"

func SetUpUserRoutes(routes *gin.Engine) {

	//Import the service and the handler here

	//Group the routes
	userGroup := routes.Group("/api/v1/users")

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
		userGroup.GET("/:id")          //get user by id (admin only)
		userGroup.GET("")              //get all users (admin only)
		userGroup.DELETE("delete/:id") // delete a users (admin only)
	}

	//Security and Verification
	{
		userGroup.POST("/verify-email")           //verify email with OTP
		userGroup.POST("/send-verification-email") //resend verification email
		userGroup.POST("/forgot-password")        //send password reset link
		userGroup.POST("/reset-password")         //reset password with token
	}

}
