package main

import (
	"mathrush/internal/domain"
	"mathrush/internal/store"
	"os"
	"testing"
)

func TestStoreRoundTrip(t *testing.T) {
	p := "store-test.db"
	defer os.Remove(p)
	s, e := store.Open(p)
	if e != nil {
		t.Fatal(e)
	}
	r := domain.NewRecord("r", "u", "2+3", 5)
	if e = s.SaveRecord(r); e != nil {
		t.Fatal(e)
	}
	if _, e = s.GetRecord("r"); e != nil {
		t.Fatal(e)
	}
	s.Close()
}
