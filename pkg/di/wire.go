package di

import (
	https "Laptop_Lounge/pkg/api"
	"Laptop_Lounge/pkg/api/handler"
	"Laptop_Lounge/pkg/api/middlewire"
	"Laptop_Lounge/pkg/config"
	"Laptop_Lounge/pkg/db"
	"Laptop_Lounge/pkg/repository"
	"Laptop_Lounge/pkg/service"
	"Laptop_Lounge/pkg/usecase"
)

func InitializeAPI(cfg *config.Config) (*https.ServerHttp, error) {

	DB, err := db.ConnectDatabse(cfg.DB)
	if err != nil {
		return nil, err
	}

	service.OtpServices(cfg.Otp)
	middlewire.NewJwtTokenMiddleWire(cfg.Token)

	sellerRepository := repository.NewSellerRepository(DB)
	sellerUseCase := usecase.NewSellerUseCase(sellerRepository, &cfg.Token)
	sellerHandler := handler.NewSellerHandler(sellerUseCase)

	adminRepository := repository.NewAdminRepository(DB)
	adminUseCase := usecase.NewAdminUseCase(adminRepository, &cfg.Token)
	adminHandler := handler.NewAdminHandler(adminUseCase)

	categoryRepository := repository.NewCategoryRepository(DB)
	CategoryUseCase := usecase.NewCategoryUseCase(categoryRepository)
	categoryHandler := handler.NewCategoryHandler(CategoryUseCase)

	ProductRepository := repository.NewProductRepository(DB)
	ProductUseCase := usecase.NewProductUseCase(ProductRepository, &cfg.S3aws)
	ProductHandler := handler.NewProductHandler(ProductUseCase)

	cartRepository := repository.NewCartRepository(DB)
	cartUseCase := usecase.NewCartUseCase(cartRepository)
	cartHanlder := handler.NewCartHandler(cartUseCase)

	paymentRepository := repository.NewPaymentRepository(DB)
	paymentUseCase := usecase.NewPaymentUseCase(paymentRepository, &cfg.Razopay)
	paymentHandler := handler.NewPaymentHandler(paymentUseCase)

	userRepository := repository.NewUserRepository(DB)
	userUseCase := usecase.NewUserUseCase(userRepository, paymentRepository, &cfg.Token)
	userHandler := handler.NewUserHandler(userUseCase)

	couponRepository := repository.NewCoupenRepository(DB)
	couponUseCase := usecase.NewCouponUseCase(couponRepository)
	couponHandler := handler.NewCouponHandler(couponUseCase)

	wishlistRepository := repository.NewWishlistRepository(DB)
	wishlistUseCase := usecase.NewWishlisttUseCase(wishlistRepository)
	wishlistHandler := handler.NewwishlistHandler(wishlistUseCase)

	ReviewRepository := repository.NewReviewRepository(DB)
	ReviewUseCase := usecase.NewReviewtUseCase(ReviewRepository)
	ReviewHandler := handler.NewReviewHandler(ReviewUseCase)

	HelpDeskRepository := repository.NewHelpDeskRepository(DB)
	HelpDeskUseCase := usecase.NewHelpDeskUseCase(HelpDeskRepository)
	HelpDeskHandler := handler.NewHelpDeskHandler(HelpDeskUseCase)

	orderRepository := repository.NewOrderRepository(DB)
	orderUseCase := usecase.NewOrderUseCase(orderRepository, cartRepository, sellerRepository, paymentRepository, couponRepository, &cfg.Razopay)
	orderHandler := handler.NewOrderHandler(orderUseCase)

	serverHTTP := https.NewServerHtttp(
		userHandler,
		sellerHandler,
		adminHandler,
		categoryHandler,
		ProductHandler,
		cartHanlder,
		orderHandler,
		paymentHandler,
		couponHandler,
		wishlistHandler,
		ReviewHandler,
		HelpDeskHandler,
	)
	return serverHTTP, nil
}
