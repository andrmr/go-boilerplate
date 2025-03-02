package main

import (
	"golnib/config"
	"log"
)

func main() {
	cfg, err := config.New()
	if err != nil {
		log.Fatalf("Error: %v", err)
	}

	println(cfg.Name)
}
