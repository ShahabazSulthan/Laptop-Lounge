package routes

import (
	"Laptop_Lounge/pkg/api/handler"
	"Laptop_Lounge/pkg/api/middlewire"

	"github.com/gin-gonic/gin"
)

func UserRoutes(engin *gin.RouterGroup, user *handler.UserHandler, product *handler.ProductHandler) {

	engin.GET("/", product.GetProduct)
	engin.GET("/:productid", product.GetAProduct)
	engin.GET("/filter", product.FilterProduct)

	// User-related routes

	engin.POST("/signup", user.UserSignup)
	engin.POST("/verifyOTP", user.VerifyOTP)
	engin.POST("/sendOTP", user.SendOtp)
	engin.POST("/login", user.UserLogin)
	engin.POST("/forgetpassword", user.ForgotPassword)

	engin.Use(middlewire.UserAuthorization)
	{
		addressmanagement := engin.Group("/address")
		{
			addressmanagement.POST("/", user.NewAddress)
			addressmanagement.GET("/", user.GetAddress)
			addressmanagement.PATCH("/", user.EditAddress)
			addressmanagement.DELETE("/:id", user.DeleteAddress)
		}

		profilemanagement := engin.Group("/profile")
		{
			profilemanagement.GET("/", user.GetProfile)
			profilemanagement.PATCH("/", user.EditProfile)
		}

	}

}
