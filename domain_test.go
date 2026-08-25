package main

import (
	"mathrush/internal/domain"
	"testing"
)

func TestDomainValidation(t *testing.T) {
	if domain.ValidateRecord(domain.NewRecord("x", "u", "2+3", 5)) != nil {
		t.Fatal()
	}
	if domain.ScoreAnswer(5) != 10 {
		t.Fatal()
	}
}
