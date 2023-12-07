package api

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/ethereum/go-ethereum/crypto"
	"github.com/gorilla/mux"
)


type SwapRequest struct {
	Message   string `json:"message"`
	Signature string `json:"signature"`
	Account   string `json:"account"`
}

func StartServer() {
	r := mux.NewRouter()
	r.HandleFunc("/swap", startDex).Methods("GET", "POST")
	r.HandleFunc("/", handleStatus).Methods("GET")
	addr := "localhost:8080"
	fmt.Println("API server started on localhost:8080")
	log.Fatal(http.ListenAndServe(addr, r))
}

func startDex(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		fmt.Fprintf(w, "GET request received")
		
	case "POST":
		var swapReq SwapRequest
		err := json.NewDecoder(r.Body).Decode(&swapReq)
		if err != nil {
			http.Error(w, "Error decoding request body", http.StatusBadRequest)
			return
		}

		signature := crypto.Keccak256([]byte(swapReq.Message))
		publicKey, err := crypto.Ecrecover(signature, []byte(swapReq.Signature))
		if err != nil {
			http.Error(w, "Error recovering public key", http.StatusInternalServerError)
			return
		}
		isValid := crypto.VerifySignature(publicKey, signature, []byte(swapReq.Signature))
		if isValid {
			fmt.Fprintf(w, "POST request received and signature verified")
		} else {
			http.Error(w, "Invalid signature", http.StatusUnauthorized)
		}
	default:
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}
}

func handleStatus(w http.ResponseWriter, r *http.Request) {
    if r.URL.Path != "/" {
        http.NotFound(w, r)
        return
    }
    http.ServeFile(w, r, "index.html")
}
