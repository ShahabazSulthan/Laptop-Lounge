package routes

import (
	"Laptop_Lounge/pkg/api/handler"

	"github.com/gin-gonic/gin"
)

func AdminRoutes(engin *gin.RouterGroup, admin *handler.AdminHandler) {

	engin.POST("/login",admin.AdminLogin)

}