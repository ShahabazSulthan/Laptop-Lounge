package main

import (
	"Laptop_Lounge/docs"
	"Laptop_Lounge/pkg/config"
	"Laptop_Lounge/pkg/di"
	"log"
	"os"
)

// @title Laptop Lounge API
// @description Laptop Lounge - 🚀 Your One-Stop Destination for Ultimate Laptop Shopping!
// @description 🛒💻 Browse, compare, and buy top-notch laptops effortlessly.
// @description Powered by cutting-edge technology, we bring you a seamless shopping experience.
// @description 🌟 Dive into the future of laptop shopping with Laptop Lounge! 🚀🔥
// @contact.name API Support
// @contact.email shahabazsulthan4@gmail.com
// @contact.name Creator: Shahabaz Sulthan
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
	// Load the configuration
	cfg, err := config.LoadConfig()
	if err != nil {
		log.Fatal("Error loading the env file using viper: ", err)
	}

	// Set the Swagger documentation info
	docs.SwaggerInfo.Title = "Laptop Lounge"
	docs.SwaggerInfo.Host = "laptoplounge.shahabazsulthan.cloud"
	docs.SwaggerInfo.BasePath = "/"

	// Initialize the server
	server, err := di.InitializeAPI(cfg)
	if err != nil {
		log.Fatal("Error creating server: ", err)
	}

	// Set up log output to both stdout and a file
	file, err := os.OpenFile("app.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		log.Fatal("Failed to open log file: ", err)
	}
	log.SetOutput(os.Stdout)
	log.SetOutput(file)

	log.Println("Starting the server on :8000")
	server.Start()
}
