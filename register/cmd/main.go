package main

import (
	"log"

	"register/internal/app"
)

func main() {
	if err := app.Run(); err != nil {
		log.Fatalf("failed to init register: %s", err.Error())
	}
}
