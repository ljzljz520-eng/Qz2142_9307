package main

import (
	"mathrush/internal/service"
	"mathrush/internal/store"
	"os"
	"testing"
)

func TestWorkflowOne(t *testing.T) {
	p := "w1.db"
	defer os.Remove(p)
	s, _ := store.Open(p)
	defer s.Close()
	svc := service.New(s)
	r, e := svc.SubmitAndReview("w1", "u", "2+3", 5)
	if e != nil || r.Status != "approved" {
		t.Fatalf("%v %s", e, r.Status)
	}
}
func TestWorkflowTwo(t *testing.T) {
	p := "w2.db"
	defer os.Remove(p)
	s, _ := store.Open(p)
	defer s.Close()
	svc := service.New(s)
	svc.Register("w2", "u", "2+3", 5)
	r, e := svc.ReviewAndArchive("w2", "reviewer")
	if e != nil || r.Status != "archived" {
		t.Fatal(e)
	}
}
func TestWorkflowThree(t *testing.T) {
	p := "w3.db"
	defer os.Remove(p)
	s, _ := store.Open(p)
	defer s.Close()
	svc := service.New(s)
	svc.Register("same", "u", "2+3", 5)
	if e := svc.ConcurrentConfirm([]string{"same", "same"}); e != nil {
		t.Fatal(e)
	}
	r, _ := s.GetRecord("same")
	if r.Version < 3 {
		t.Fatalf("lost update: version %d", r.Version)
	}
}
