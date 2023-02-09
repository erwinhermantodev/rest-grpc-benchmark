package main

import (
	"log"

	"github.com/erwinhermanto31/service2/transport/grpc"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error load .env")
	}

	grpc.Run()
}
