package main

import (
	"fmt"
	"log"
	"os"

	"github.com/alwindoss/wink/internal/server"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	host := os.Getenv("DB_HOST")
	port := os.Getenv("DB_PORT")
	if port == "" {
		port = "8080"
	}
	addr := fmt.Sprintf("%s:%s", host, port)
	cfg := server.Config{
		Addr: addr,
	}
	err = server.Run(&cfg)
	if err != nil {
		log.Fatal(err)
	}
	log.Println("Exiting wink server...")
}
