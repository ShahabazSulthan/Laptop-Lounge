package routes

import (
	"Laptop_Lounge/pkg/api/handler"
	"Laptop_Lounge/pkg/api/middlewire"

	"github.com/gin-gonic/gin"
)

func UserRoutes(engin *gin.RouterGroup, user *handler.UserHandler, product *handler.ProductHandler, cart *handler.CartHandler, order *handler.OrderHandler, payment *handler.PaymentHandler, wishlist *handler.WishlistHandler, review *handler.ReviewHandler, helpdesk *handler.HelpDeskHandler) {

	engin.GET("/products", product.GetProduct)
	engin.GET("/product/:productid", product.GetAProduct)
	engin.GET("/product/HighToLow", product.GetAProductHightoLow)
	engin.GET("/product/LowToHigh", product.GetAProductLowtoHigh)
	engin.GET("/product/AtoZ", product.GetAProductAtoZ)
	engin.GET("/product/ZtoA", product.GetAProductZtoA)
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
			addressmanagement.GET("/", user.GetAddress)
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
			cartmanagement.POST("/", cart.CreateCart)
			cartmanagement.DELETE("/:productID", cart.DeleteProductFromCart)
			cartmanagement.PATCH("/increment/:productID", cart.IncrementQuantityCart)
			cartmanagement.PATCH("/decrement/:productID", cart.DecrementQuantityCart)
			cartmanagement.GET("/", cart.ShowCart)
		}

		ordermanagement := engin.Group("/order")
		{
			ordermanagement.POST("", order.NewOrder)
			ordermanagement.POST("/Address", user.NewAddress)
			ordermanagement.GET("/Address", user.GetAddress)
			ordermanagement.PATCH("/EditAddress", user.EditAddress)
			ordermanagement.GET("/", order.ShowAbstractOrders)
			ordermanagement.GET("/:orderItemID", order.SingleOrderDetails)
			ordermanagement.PATCH("/cancel/:orderItemID", order.CancelUserOrder)
			ordermanagement.PATCH("/return/:orderItemID", order.ReturnUserOrder)
			ordermanagement.GET("/invoice/:orderItemID", order.GetInvoice)
		}

		paymentmanagement := engin.Group("/payment")
		{
			paymentmanagement.POST("/verify", payment.VerifyOnlinePayment)
		}

		walletmenagement := engin.Group("/wallet")
		{
			walletmenagement.GET("/", payment.ViewWallet)
			walletmenagement.GET("/transaction", payment.GetWalletTransaction)
		}

		wishlistmenagement := engin.Group("/wishlist")
		{
			wishlistmenagement.POST("/:productID", wishlist.AddToWishlist)
			wishlistmenagement.GET("/", wishlist.GetWishlist)
			wishlistmenagement.DELETE("/:productID", wishlist.DeleteWishlist)
		}

		reviewmenagement := engin.Group("/review")
		{
			reviewmenagement.POST("/:productID", review.AddReview)
			reviewmenagement.GET("/", review.GetReviewsByProductID)
			reviewmenagement.DELETE("/:reviewID", review.DeleteReviewByID)
			reviewmenagement.GET("/:productID", review.GetAverageRating)

		}

		helpdeskmenagement := engin.Group("/helpdesk")
		{
			helpdeskmenagement.POST("/", helpdesk.CreateRequest)
			helpdeskmenagement.GET("/replayed", helpdesk.GetRepliedRequests)
			helpdeskmenagement.GET("/unreplayed", helpdesk.GetUnrepliedRequests)
		}
	}

}
