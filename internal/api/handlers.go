package api

import (
	"fmt"
	"net/http"
)

func StartServer() {
	http.HandleFunc("/swap", handleSwap)

	fmt.Println("API server started on :8080")
	http.ListenAndServe(":8080", nil)
}

func handleSwap(w http.ResponseWriter, r *http.Request) {
	// Handle token swap requests here
}
