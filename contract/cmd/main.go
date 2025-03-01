package main

import (
	"log"

	"contract/internal/app"
)

func main() {
	if err := app.Run(); err != nil {
		log.Fatalf("failed to init service-contract: %s", err.Error())
	}
}
