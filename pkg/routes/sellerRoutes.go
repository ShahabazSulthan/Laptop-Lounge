package routes

import (
	"Laptop_Lounge/pkg/api/handler"

	"github.com/gin-gonic/gin"
)

func SellerRoutes(engin *gin.RouterGroup,seller *handler.SellerHandler) {
	engin.POST("/signup",seller.SellerSignup)
	engin.POST("/login",seller.SellerLogin)

}