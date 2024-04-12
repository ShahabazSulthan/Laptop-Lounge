package routes

import (
	"Laptop_Lounge/pkg/api/handler"
	"Laptop_Lounge/pkg/api/middlewire"

	"github.com/gin-gonic/gin"
)

func AdminRoutes(engin *gin.RouterGroup, admin *handler.AdminHandler, seller *handler.SellerHandler, user *handler.UserHandler) {

	engin.POST("/login", admin.AdminLogin)

	engin.Use(middlewire.AdminAuthorization)

	{

		engin.GET("/", admin.AdminDashBord)

		usermanagement := engin.Group("/users")
		{
			usermanagement.GET("/getuser", user.GetUser)
			usermanagement.PATCH("/block/:userID", user.BlockUser)
			usermanagement.PATCH("/unblock/:userID", user.UnblockUser)
		}

		sellermanagement := engin.Group("/sellers")
		{
			sellermanagement.GET("/getsellers", seller.GetSellers)
			sellermanagement.PATCH("/block/:sellerID", seller.BlockSeller)
			sellermanagement.PATCH("/unblock/:sellerID", seller.UnblockSeller)
			sellermanagement.GET("/pending", seller.GetPendingSellers)
			sellermanagement.GET("/singleview/:sellerID", seller.FetchSingleSeller)
			sellermanagement.PATCH("/verify/:sellerID", seller.VerifySeller)
		}
	}
}
