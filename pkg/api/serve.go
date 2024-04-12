package https

import (
	"Laptop_Lounge/pkg/api/handler"
	"Laptop_Lounge/pkg/routes"
	"fmt"
	"log"

	"github.com/gin-gonic/gin"
)

// ServerHttp represents an HTTP server using Gin.
type ServerHttp struct {
	engin *gin.Engine
}

// NewServerHttp creates a new HTTP server with the provided handlers.
func NewServerHtttp(user *handler.UserHandler, seller *handler.SellerHandler, admin *handler.AdminHandler) *ServerHttp {
	engin := gin.New()
	engin.Use(gin.Logger())

	// Set up routes for users and sellers
	routes.UserRoutes(engin.Group("/"), user)
	routes.SellerRoutes(engin.Group("/seller"), seller)
	routes.AdminRoutes(engin.Group("/admin"), admin, seller, user)

	return &ServerHttp{engin: engin}
}

// Start starts the HTTP server and listens for incoming requests.
func (server *ServerHttp) Start() {
	err := server.engin.Run(":8000")
	if err != nil {
		log.Fatal("gin engin couldn't start")
	}
	fmt.Println("gin engin start")
}
