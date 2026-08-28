package main

import (
	"mathrush/internal/domain"
	"mathrush/internal/store"
	"os"
	"testing"
)

func TestPersistenceSurvivesReopen(t *testing.T) {
	p := "reopen.db"
	defer os.Remove(p)
	s, e := store.Open(p)
	if e != nil {
		t.Fatal(e)
	}
	s.SaveRecord(domain.NewRecord("persist", "u", "1+4", 5))
	s.Close()
	s, e = store.Open(p)
	if e != nil {
		t.Fatal(e)
	}
	defer s.Close()
	if _, e = s.GetRecord("persist"); e != nil {
		t.Fatal(e)
	}
}
