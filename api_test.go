package main

import (
	"mathrush/internal/api"
	"mathrush/internal/service"
	"mathrush/internal/store"
	"net/http/httptest"
	"os"
	"testing"
)

func TestHealth(t *testing.T) {
	p := "api.db"
	defer os.Remove(p)
	s, _ := store.Open(p)
	defer s.Close()
	w := httptest.NewRecorder()
	api.New(service.New(s)).Health(w, httptest.NewRequest("GET", "/health", nil))
	if w.Code != 200 {
		t.Fatal(w.Code)
	}
}
