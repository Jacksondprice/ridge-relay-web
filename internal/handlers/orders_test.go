package handlers

import (
	"net/http/httptest"
	"testing"
)

func TestListOrders(t *testing.T) {
	rec := httptest.NewRecorder()
	ListOrders(rec, httptest.NewRequest("GET", "/api/orders", nil))
	if rec.Code != 200 {
		t.Fatalf("status %d", rec.Code)
	}
}
