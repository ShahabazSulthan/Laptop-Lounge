package routes

import (
	"Laptop_Lounge/pkg/api/handler"
	"Laptop_Lounge/pkg/api/middlewire"

	"github.com/gin-gonic/gin"
)

func AdminRoutes(engin *gin.RouterGroup, admin *handler.AdminHandler, seller *handler.SellerHandler, user *handler.UserHandler, category *handler.CategoryHandler, coupon *handler.CouponHandler, helpdesk *handler.HelpDeskHandler) {

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
		categorymanagement := engin.Group("/category")
		{
			categorymanagement.POST("/", category.NewCategory)
			categorymanagement.GET("/", category.FetchAllCatogry)
			categorymanagement.PATCH("/", category.UpdateCategory)
			categorymanagement.DELETE("/:id", category.DeleteCategory)
		}
		brandmanagement := engin.Group("/brand")
		{
			brandmanagement.POST("/", category.CreateBrand)
			brandmanagement.GET("/", category.FetchAllBrand)
			brandmanagement.PATCH("/", category.UpdateBrand)
			brandmanagement.DELETE("/:id", category.DeleteBrand)
		}

		couponmanagment := engin.Group("/coupon")
		{
			couponmanagment.POST("/", coupon.CreateCoupon)
			couponmanagment.GET("/", coupon.GetCoupons)
			couponmanagment.PATCH("/unblock/:couponID", coupon.UnblockCoupon)
			couponmanagment.PATCH("/block/:couponID", coupon.BlockCoupon)
		}

		helpdeskmenagement := engin.Group("/helpdesk")
		{
			helpdeskmenagement.PATCH("/:requestID", helpdesk.UpdateAnswer)
			helpdeskmenagement.GET("/Updated", helpdesk.GetRepliedRequests)
			helpdeskmenagement.GET("/Pending", helpdesk.GetUnrepliedRequests)
		}
	}
}
