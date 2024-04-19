package routes

import (
	"Laptop_Lounge/pkg/api/handler"

	"github.com/gin-gonic/gin"
)

func SellerRoutes(engin *gin.RouterGroup, seller *handler.SellerHandler, Product *handler.ProductHandler,category *handler.CategoryHandler) {
	engin.POST("/signup", seller.SellerSignup)
	engin.POST("/login", seller.SellerLogin)
	engin.GET("/profile/:SellerID", seller.GetSellerProfile)
	engin.PATCH("/profile/:SellerID", seller.EditSellerProfile)

	Productmanagement := engin.Group("/products")
	{
		Productmanagement.POST("/:SellerID", Product.AddProduct)
		Productmanagement.GET("/seller/:SellerID", Product.GetSellerIProduct)
		Productmanagement.GET("/:productid", Product.GetAProduct)
		Productmanagement.PATCH("/", Product.EditProduct)
		Productmanagement.DELETE("/:productid", Product.DeleteProduct)
		Productmanagement.PATCH("/:productid/block", Product.BlockProduct)
		Productmanagement.PATCH("/:productid/unblock", Product.UnblockProduct)
	}

	categorymanagement := engin.Group("/categoryoffer")
		{
			categorymanagement.GET("/brand/:page", category.FetchAllBrand)
			categorymanagement.GET("/categor/:page", category.FetchAllCatogry)
			categorymanagement.GET("/:SellerID", category.GetAllCategoryOffer)
			categorymanagement.POST("/:seller_id", category.CreateCategoryOffer)
			categorymanagement.PATCH("/:seller_id", category.EditCategoryOffer)
			categorymanagement.PATCH("/block/:categoryOfferID", category.BlockCategoryOffer)
			categorymanagement.PATCH("/unblock/:categoryOfferID", category.UnBlockCategoryOffer)
			categorymanagement.DELETE("/delete/:categoryOfferID", category.DeleteCategoryOffer)
		}

}
