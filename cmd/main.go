package main

import (
	"fmt"

	"github.com/BradMyrick/kydex/internal/api"
)

func main() {
	fmt.Println("Starting Kydex DEX...")

	// Initialize and start the API server
	api.StartServer()
}
