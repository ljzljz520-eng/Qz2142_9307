package main

import (
	"mathrush/internal/domain"
	"mathrush/internal/notify"
	"testing"
)

func TestNotifyHub(t *testing.T) {
	h := notify.New()
	h.Publish(domain.Event{Kind: "registered"})
	if h.Size() != 1 || !notify.HasKind(h.Snapshot(), "registered") {
		t.Fatal()
	}
}
