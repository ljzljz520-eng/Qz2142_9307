package service

import (
	"path/filepath"
	"sync"
	"testing"

	"mathrush/internal/domain"
	"mathrush/internal/store"
)

func newStore(t *testing.T) *store.Store {
	t.Helper()
	s, err := store.Open(filepath.Join(t.TempDir(), "test.db"))
	if err != nil {
		t.Fatalf("open store: %v", err)
	}
	t.Cleanup(func() { _ = s.Close() })
	return s
}

// TestConfirmConcurrentPreservesAllResults reproduces the reported bug: when
// several people view and confirm the same math material at the same time, the
// later write erased the earlier result. After the fix every confirmation must
// be preserved — the record Version increments once per confirm and one
// "confirmed" event is stored for each, on distinct keys.
func TestConfirmConcurrentPreservesAllResults(t *testing.T) {
	s := New(newStore(t))
	if _, err := s.Register("math-1", "alice", "2+3", 5); err != nil {
		t.Fatalf("register: %v", err)
	}

	const n = 8
	var wg sync.WaitGroup
	for i := 0; i < n; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			if _, err := s.Confirm("math-1"); err != nil {
				t.Errorf("confirm: %v", err)
			}
		}()
	}
	wg.Wait()

	rec, err := s.Store.GetRecord("math-1")
	if err != nil {
		t.Fatalf("get record: %v", err)
	}
	if rec.Version != 1+n {
		t.Errorf("Version = %d, want %d (each confirm must increment)", rec.Version, 1+n)
	}
	// 1 "registered" event + n "confirmed" events, each on its own key.
	if got := s.Store.Count("events"); got != 1+n {
		t.Errorf("events = %d, want %d (every confirm must be preserved, not overwritten)", got, 1+n)
	}
}

// TestConfirmSequentialPreservesAllResults is the deterministic counterpart:
// confirms in sequence must each leave a distinct event record rather than
// overwriting a single fixed slot.
func TestConfirmSequentialPreservesAllResults(t *testing.T) {
	s := New(newStore(t))
	if _, err := s.Register("math-2", "bob", "1+1", 5); err != nil {
		t.Fatalf("register: %v", err)
	}

	const n = 5
	for i := 0; i < n; i++ {
		if _, err := s.Confirm("math-2"); err != nil {
			t.Fatalf("confirm #%d: %v", i, err)
		}
	}

	rec, _ := s.Store.GetRecord("math-2")
	if rec.Version != 1+n {
		t.Errorf("Version = %d, want %d", rec.Version, 1+n)
	}
	if got := s.Store.Count("events"); got != 1+n {
		t.Errorf("events = %d, want %d", got, 1+n)
	}
}

// TestConcurrentConfirmPreservesAllResults drives the confirm flow through
// ConcurrentConfirm, the batched entry point, to make sure it preserves every
// result too.
func TestConcurrentConfirmPreservesAllResults(t *testing.T) {
	s := New(newStore(t))
	for i := 0; i < 4; i++ {
		id := domain.Record{ID: "math-" + rds[i]}.ID
		if _, err := s.Register(id, "carol", "7+0", 5); err != nil {
			t.Fatalf("register %s: %v", id, err)
		}
	}
	ids := []string{"math-3", "math-4", "math-5", "math-6"}
	// Confirm each record twice, concurrently, so the second confirm of each
	// would overwrite the first under the old fixed-key scheme.
	all := append(append([]string{}, ids...), ids...)
	if err := s.ConcurrentConfirm(all); err != nil {
		t.Fatalf("concurrent confirm: %v", err)
	}
	for _, id := range ids {
		rec, err := s.Store.GetRecord(id)
		if err != nil {
			t.Fatalf("get %s: %v", id, err)
		}
		if rec.Version != 3 { // registered=1 + 2 confirms
			t.Errorf("%s Version = %d, want 3", id, rec.Version)
		}
	}
	// 4 registered + 8 confirmed events, all distinct.
	if got, want := s.Store.Count("events"), 4+8; got != want {
		t.Errorf("events = %d, want %d", got, want)
	}
}

var rds = []string{"3", "4", "5", "6"}
