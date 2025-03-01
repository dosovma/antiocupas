package app

import (
	"fmt"

	"github.com/google/uuid"
)

func Run() error {
	fmt.Println("it works")

	uuid.NewUUID()

	return nil
}
