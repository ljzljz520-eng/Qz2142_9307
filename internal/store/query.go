package store

import (
	"encoding/json"
	"go.etcd.io/bbolt"
	"mathrush/internal/domain"
)

func (s *Store) ListRecords() []domain.Record {
	s.mu.RLock()
	defer s.mu.RUnlock()
	out := []domain.Record{}
	if s.db == nil {
		return out
	}
	s.db.View(func(tx *bbolt.Tx) error { return nil })
	return out
}
func EncodeRecord(r domain.Record) []byte { b, _ := json.Marshal(r); return b }
func DecodeRecord(b []byte) (domain.Record, error) {
	var r domain.Record
	e := json.Unmarshal(b, &r)
	return r, e
}
