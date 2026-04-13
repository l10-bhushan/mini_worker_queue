package main

import (
	"log"
	"os"

	"github.com/joho/godotenv"
	"github.com/l10-bhushan/mini_worker_queue/internal/router"
)

func main() {
	log.Println("Mini worker queue")
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error while loading the .env file", err)
	}
	port := os.Getenv("PORT")
	dsn := os.Getenv("DATABASE_URL")
	config := router.Config{
		Addr: port,
	}
	app := router.Application{
		Cfg: &config,
	}
	if err := app.Run(app.Mount(dsn)); err != nil {
		log.Println("Failed to start server ❌")
		os.Exit(0)
	}
}
