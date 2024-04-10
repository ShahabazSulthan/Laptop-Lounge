package routes

import (
	"Laptop_Lounge/pkg/api/handler"

	"github.com/gin-gonic/gin"
)

func UserRoutes(engin *gin.RouterGroup, user *handler.UserHandler) {

	// User-related routes

	engin.POST("/signup", user.UserSignup)
	engin.POST("/verifyOTP", user.VerifyOTP)
	engin.POST("/sendOTP", user.SendOtp)
	engin.POST("/login", user.UserLogin)
	engin.POST("/forgetpassword", user.ForgotPassword)

}
