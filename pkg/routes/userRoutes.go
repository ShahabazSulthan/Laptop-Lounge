package routes

import (
	"Laptop_Lounge/pkg/api/handler"
	"Laptop_Lounge/pkg/api/middlewire"

	"github.com/gin-gonic/gin"
)

func UserRoutes(engin *gin.RouterGroup, user *handler.UserHandler, product *handler.ProductHandler, cart *handler.CartHandler, order *handler.OrderHandler, payment *handler.PaymentHandler) {

	engin.GET("/products", product.GetProduct)
	engin.GET("/product/:productid", product.GetAProduct)
	engin.GET("/product/HighToLow", product.GetAProductHightoLow)
	engin.GET("/product/LowToHigh", product.GetAProductLowtoHigh)
	engin.GET("/filter", product.FilterProduct)

	engin.GET("/razopay", payment.OnlinePayment)
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
			addressmanagement.GET("/:UserID/:page", user.GetAddress)
			addressmanagement.PATCH("/", user.EditAddress)
			addressmanagement.DELETE("/:id", user.DeleteAddress)
		}

		profilemanagement := engin.Group("/profile")
		{
			profilemanagement.GET("/", user.GetProfile)
			profilemanagement.PATCH("/", user.EditProfile)
		}

		cartmanagement := engin.Group("/cart")
		{
			cartmanagement.POST("/:UserID", cart.CreateCart)
			cartmanagement.DELETE("/:productID/:UserID", cart.DeleteProductFromCart)
			cartmanagement.PATCH("/increment/:productID", cart.IncrementQuantityCart)
			cartmanagement.PATCH("/decrement/:productID", cart.DecrementQuantityCart)
			cartmanagement.GET("/", cart.ShowCart)
		}

		ordermanagement := engin.Group("/order")
		{
			ordermanagement.POST("/", order.NewOrder)
			ordermanagement.GET("/", order.ShowAbstractOrders)
			ordermanagement.GET("/:orderID", order.SingleOrderDetails)
			ordermanagement.PATCH("/:orderID", order.CancelUserOrder)
			ordermanagement.PATCH("/return/:orderID", order.ReturnUserOrder)
			ordermanagement.GET("/invoice/:OrderID", order.GetInvoice)
		}

		paymentmanagement := engin.Group("/payment")
		{
			paymentmanagement.POST("/verify", payment.VerifyOnlinePayment)
		}

		// walletmenagement := engin.Group("/wallet")
		// {
		// 	walletmenagement.GET("/", payment.ViewWallet)
		// 	walletmenagement.GET("/transaction", payment.GetWalletTransaction)
		// }

	}

}
