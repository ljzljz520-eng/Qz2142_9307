package main

import (
	"mathrush/internal/config"
	"testing"
)

func TestConfigDefaults(t *testing.T) {
	c := config.Default()
	if !c.IsValid() {
		t.Fatal()
	}
}
