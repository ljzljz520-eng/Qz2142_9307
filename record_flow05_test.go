package main

import (
	"mathrush/internal/service"
	"mathrush/internal/store"
	"os"
	"testing"
)

func TestRecordFlow05(t *testing.T) {
	p := "f05.db"
	defer os.Remove(p)
	s, _ := store.Open(p)
	defer s.Close()
	svc := service.New(s)
	r, e := svc.Register("flow05", "player", "1+4", 5)
	if e != nil || r.ID != "flow05" {
		t.Fatal(e)
	}
}
