package main

import (
	"log"
	"net/http"
	"os"

	"github.com/Jacksondprice/ridge-relay-web/internal/handlers"
)

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/api/orders", handlers.ListOrders)
	addr := os.Getenv("ADDR")
	if addr == "" {
		addr = ":8080"
	}
	log.Fatal(http.ListenAndServe(addr, mux))
}
