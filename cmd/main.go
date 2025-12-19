package main

import (
	"log"
	"os"
)

func main() {
	cfg := config{
		addr: ":8080",
		db:   dbConfig{},
	}

	api := application{
		config: cfg,
	}

	if err := api.Run(api.Mount()); err != nil {
		log.Printf("server has failed to start. err: %s", err)
		os.Exit(1)
	}
}
