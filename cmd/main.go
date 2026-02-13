package main

import (
	"log"
	"os"
)

func main() {
	cfg := config{
		addr: ":6767",
		db:   dbConfig{},
	}

	api := application{
		config: cfg,
	}

	if err := api.run(api.mount()); err != nil {
		log.Printf("Server failes to start, err: %s", err)
		os.Exit(1)
	}
}
