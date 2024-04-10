package di

import (
	https "Laptop_Lounge/pkg/api"
	"Laptop_Lounge/pkg/api/handler"
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

	userRepository := repository.NewUserRepository(DB)
	userUseCase := usecase.NewUserUseCase(userRepository, &cfg.Token)
	userHandler := handler.NewUserHandler(userUseCase)

	sellerRepository := repository.NewSellerRepository(DB)
	sellerUseCase := usecase.NewSellerUseCase(sellerRepository, &cfg.Token)
	sellerHandler := handler.NewSellerHandler(sellerUseCase)

	adminRepository := repository.NewAdminRepository(DB)
	adminUseCase := usecase.NewAdminUseCase(adminRepository, &cfg.Token)
	adminHandler := handler.NewAdminHandler(adminUseCase)

	serverHTTP := https.NewServerHtttp(
		userHandler,
		sellerHandler,
		adminHandler,
	)
	return serverHTTP, nil
}
