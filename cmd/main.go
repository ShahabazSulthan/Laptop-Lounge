package main

import (
	"Laptop_Lounge/docs"
	"Laptop_Lounge/pkg/config"
	"Laptop_Lounge/pkg/di"
	"log"
)

// @title Laptop Lounge API
// @description Laptop Lounge - 🚀 Your One-Stop Destination for Ultimate Laptop Shopping! 🛒💻 Browse, compare, and buy top-notch laptops effortlessly. Powered by cutting-edge technology, we bring you a seamless shopping experience. 🌟 Dive into the future of laptop shopping with Laptop Lounge! 🚀🔥
// @contact.name API Support
// @contact.email shahabazsulthan4@gmail.com
// @securityDefinitions.apikey BearerTokenAuth
// @in header
// @name Authorization
// @description Use your Bearer token for authentication. Example: "Bearer {token}"
// @securityDefinitions.apikey RefreshTokenAuth
// @in header
// @name Refreshtoken
// @description Use your Refresh token to obtain a new Bearer token. Example: "{token}"
// @BasePath /
// @query.collection.format multi

func main() {
	cfg, err := config.LoadConfig()
	if err != nil {
		log.Fatal("Error loading the env file using viper: ", err)
	}

	docs.SwaggerInfo.Title = "Laptop Lounge"
	// docs.SwaggerInfo.Host = "laptoplounge.shahabaz.tech"
	docs.SwaggerInfo.Host = "localhost:8000"

	server, err := di.InitializeAPI(cfg)
	if err != nil {
		log.Fatal("Error creating server: ", err)
	}
	server.Start()
}
