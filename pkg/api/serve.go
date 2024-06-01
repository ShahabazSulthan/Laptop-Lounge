package https

import (
	"Laptop_Lounge/pkg/api/handler"
	"Laptop_Lounge/pkg/api/middlewire"
	"Laptop_Lounge/pkg/routes"
	"fmt"
	"log"
	"net/http"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// ServerHttp represents an HTTP server using Gin.
type ServerHttp struct {
	engine *gin.Engine
}

// NewServerHttp creates a new HTTP server with the provided handlers.
func NewServerHtttp(user *handler.UserHandler, seller *handler.SellerHandler, admin *handler.AdminHandler, category *handler.CategoryHandler, product *handler.ProductHandler, cart *handler.CartHandler, order *handler.OrderHandler, payment *handler.PaymentHandler, coupon *handler.CouponHandler, wishlist *handler.WishlistHandler, review *handler.ReviewHandler, helpdesk *handler.HelpDeskHandler) *ServerHttp {
	engin := gin.New()
	engin.Use(gin.Logger())
	engin.Use(gin.Recovery())
	engin.Use(middlewire.Logger())

	// Enable CORS for all origins
	config := cors.DefaultConfig()
	config.AllowAllOrigins = true
	config.AllowMethods = []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"}
	config.AllowHeaders = []string{"Origin", "Content-Type"}
	engin.Use(cors.New(config))

	// Load HTML templates and static files
	engin.LoadHTMLGlob("./template/*.html")
	engin.Static("/static", "./static")

	// Define routes
	engin.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "create_order.html", nil)
	})

	engin.POST("/html/orders", order.OrderHtml)

	// use ginSwagger middleware to serve the API docs
	engin.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	// Set up routes for users and sellers
	routes.UserRoutes(engin.Group("/"), user, product, cart, order, payment, wishlist, review, helpdesk)
	routes.SellerRoutes(engin.Group("/seller"), seller, product, category, order)
	routes.AdminRoutes(engin.Group("/admin"), admin, seller, user, category, coupon, helpdesk)

	return &ServerHttp{engine: engin}
}

// Start starts the HTTP server and listens for incoming requests.
func (server *ServerHttp) Start() {
	err := server.engine.Run(":8000")
	if err != nil {
		log.Fatal("gin engin couldn't start")
	}
	fmt.Println("gin engin start")
}
