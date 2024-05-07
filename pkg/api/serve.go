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
func NewServerHtttp(user *handler.UserHandler, seller *handler.SellerHandler, admin *handler.AdminHandler, category *handler.CategoryHandler, product *handler.ProductHandler,cart *handler.CartHandler,order *handler.OrderHandler,payment *handler.PaymentHandler,coupon *handler.CouponHandler) *ServerHttp {
	engin := gin.New()
	engin.Use(gin.Logger())

	// load htmlpages
	engin.LoadHTMLGlob("./template/*.html")

	// Set up routes for users and sellers
	routes.UserRoutes(engin.Group("/"), user, product,cart,order,payment)
	routes.SellerRoutes(engin.Group("/seller"), seller, product, category,order)
	routes.AdminRoutes(engin.Group("/admin"), admin, seller, user, category,coupon)

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
