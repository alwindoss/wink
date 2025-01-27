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
	host := os.Getenv("WINK_HOST")
	port := os.Getenv("WINK_PORT")
	if port == "" {
		port = "8080"
	}
	addr := fmt.Sprintf("%s:%s", host, port)
	dsn := os.Getenv("DB_URL")
	cfg := server.Config{
		Addr: addr,
		DSN:  dsn,
	}
	err = server.Run(&cfg)
	if err != nil {
		log.Fatal(err)
	}
	log.Println("Exiting wink server...")
}
