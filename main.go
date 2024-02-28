package main

import (
	"context"
	"log"
	"os"

	"github.com/nuttapat-swd/go-intial-microservice/config"
)

func main() {
	ctx := context.Background()
	_ = ctx

	// Initialize configuration
	cfg := config.LoadConfig(func() string {
		if len(os.Args) < 2 {
			log.Fatal("Error .env path is required.")
		}
		return os.Args[1]
	}())
	log.Println(cfg)
}
