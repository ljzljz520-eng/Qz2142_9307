package main

import (
	"mathrush/internal/service"
	"mathrush/internal/store"
	"os"
	"testing"
)

func TestServiceRegister(t *testing.T) {
	p := "svc.db"
	defer os.Remove(p)
	s, _ := store.Open(p)
	defer s.Close()
	svc := service.New(s)
	r, e := svc.Register("r1", "alice", "2+3", 5)
	if e != nil || r.Score != 10 {
		t.Fatalf("%v", e)
	}
}
