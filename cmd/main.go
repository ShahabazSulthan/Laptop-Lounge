package main

import (
	"Laptop_Lounge/pkg/config"
	"Laptop_Lounge/pkg/di"
	"log"
)

func main() {
	config, err := config.LoadConfig()
	if err != nil {
		log.Fatal("error at loading the env file using viper", err)
	}

	server, err := di.InitializeAPI(config)
	if err != nil {
		log.Fatal("Error For Server Creation ", err)
	}
	server.Start()
}
