package routes

import (
	"Laptop_Lounge/pkg/api/handler"

	"github.com/gin-gonic/gin"
)

func SellerRoutes(engin *gin.RouterGroup, seller *handler.SellerHandler, Product *handler.ProductHandler, category *handler.CategoryHandler, order *handler.OrderHandler) {
	engin.POST("/signup", seller.SellerSignup)
	engin.POST("/login", seller.SellerLogin)

	engin.GET("/:SellerID", seller.SellerDashbord)
	engin.GET("/profile/:SellerID", seller.GetSellerProfile)
	engin.PATCH("/profile/:SellerID", seller.EditSellerProfile)

	Productmanagement := engin.Group("/products")
	{
		Productmanagement.POST("/:SellerID", Product.AddProduct)
		Productmanagement.GET("/seller/:SellerID/:page", Product.GetSellerIProduct)
		Productmanagement.GET("/:productid", Product.GetAProduct)
		Productmanagement.PATCH("/:SellerID", Product.EditProduct)
		Productmanagement.DELETE("/:SellerID/:productid", Product.DeleteProduct)
		Productmanagement.PATCH("/:SellerID/:productid/block", Product.BlockProduct)
		Productmanagement.PATCH("/:SellerID/:productid/unblock", Product.UnblockProduct)
	}

	categorymanagement := engin.Group("/categoryoffer")
	{
		categorymanagement.GET("/brand", category.FetchAllBrand)
		categorymanagement.GET("/category", category.FetchAllCatogry)
		categorymanagement.GET("/:SellerID", category.GetAllCategoryOffer)
		categorymanagement.POST("/:seller_id", category.CreateCategoryOffer)
		categorymanagement.PATCH("/:seller_id", category.EditCategoryOffer)
		categorymanagement.PATCH("/block/:categoryOfferID", category.BlockCategoryOffer)
		categorymanagement.PATCH("/unblock/:categoryOfferID", category.UnBlockCategoryOffer)
		categorymanagement.DELETE("/delete/:categoryOfferID", category.DeleteCategoryOffer)
	}

	ordermanagenent := engin.Group("/order")
	{
		ordermanagenent.GET("/:SellerID", order.GetSellerOrders)
		ordermanagenent.GET("/processing/:SellerID", order.GetSellerOrdersProcessing)
		ordermanagenent.GET("/delivered/:SellerID", order.GetSellerOrdersDeliverd)
		ordermanagenent.GET("/cancelled/:SellerID", order.GetSellerOrdersCancelled)
		ordermanagenent.PATCH("/:SellerID/:orderItemID", order.ConfirmDeliverd)
		ordermanagenent.PATCH("/cancel/:SellerID/:orderID", order.CancelOrder)
	}

	salesreportmanagement := engin.Group("/report")
	{
		// salesreportmanagement.GET("", order.SalesReportByYear)
		// salesreportmanagement.GET("/month", order.SalesReportByMonth)
		// salesreportmanagement.GET("/week", order.SalesReportByWeek)
		salesreportmanagement.GET("/day/:SellerID", order.SalesReport)
		salesreportmanagement.GET("/days/:SellerID/:days", order.SalesReportCustomDays)
		salesreportmanagement.GET("/xlsx/:SellerID", order.SalesReportXlSX)
		salesreportmanagement.GET("/pdf/:SellerID", order.GenerateSalesReportPDF)

	}

}
