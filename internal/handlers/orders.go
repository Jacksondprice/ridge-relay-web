package handlers

import (
	"encoding/json"
	"net/http"
)

type Order struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

func ListOrders(w http.ResponseWriter, _ *http.Request) {
	items := []Order{{ID: 1, Name: "sample"}}
	w.Header().Set("Content-Type", "application/json")
	_ = json.NewEncoder(w).Encode(items)
}
