package main

import (
	"mathrush/internal/domain"
	"mathrush/internal/report"
	"testing"
)

func TestReportTotals(t *testing.T) {
	r := []domain.Record{domain.NewRecord("1", "u", "x", 5), domain.NewRecord("2", "u", "x", 2)}
	if report.Total(r) != 10 {
		t.Fatal()
	}
	if report.PerfectCount(r) != 1 {
		t.Fatal()
	}
}
