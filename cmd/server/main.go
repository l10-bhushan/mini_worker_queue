package main

import (
	"log"
	"os"
)

func main() {
	log.Println("Mini worker queue")
	config := config{
		addr: ":8080",
	}
	app := application{
		cfg: &config,
	}
	if err := app.run(app.mount()); err != nil {
		log.Println("Failed to start server ❌")
		os.Exit(0)
	}
}
