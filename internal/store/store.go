package store

import (
	"encoding/json"
	"fmt"
	"go.etcd.io/bbolt"
	"mathrush/internal/domain"
	"path/filepath"
	"sync"
)

var buckets = []byte("records")

type Store struct {
	db *bbolt.DB
	mu sync.RWMutex
}

func Open(path string) (*Store, error) {
	db, err := bbolt.Open(filepath.Clean(path), 0600, nil)
	if err != nil {
		return nil, err
	}
	err = db.Update(func(tx *bbolt.Tx) error {
		for _, n := range []string{"records", "users", "events", "audits"} {
			if _, e := tx.CreateBucketIfNotExists([]byte(n)); e != nil {
				return e
			}
		}
		return nil
	})
	if err != nil {
		db.Close()
		return nil, err
	}
	return &Store{db: db}, nil
}
func (s *Store) Close() error {
	s.mu.Lock()
	defer s.mu.Unlock()
	if s.db == nil {
		return nil
	}
	err := s.db.Close()
	s.db = nil
	return err
}
func put(tx *bbolt.Tx, bucket, key string, v any) error {
	b := tx.Bucket([]byte(bucket))
	raw, e := json.Marshal(v)
	if e != nil {
		return e
	}
	return b.Put([]byte(key), raw)
}
func get(tx *bbolt.Tx, bucket, key string, v any) error {
	b := tx.Bucket([]byte(bucket))
	raw := b.Get([]byte(key))
	if raw == nil {
		return fmt.Errorf("not found")
	}
	return json.Unmarshal(raw, v)
}
func (s *Store) SaveRecord(r domain.Record) error {
	s.mu.RLock()
	defer s.mu.RUnlock()
	return s.db.Update(func(tx *bbolt.Tx) error { return put(tx, "records", r.ID, r) })
}
func (s *Store) GetRecord(id string) (domain.Record, error) {
	s.mu.RLock()
	defer s.mu.RUnlock()
	var r domain.Record
	e := s.db.View(func(tx *bbolt.Tx) error { return get(tx, "records", id, &r) })
	return r, e
}
func (s *Store) SaveUser(u domain.User) error {
	s.mu.RLock()
	defer s.mu.RUnlock()
	return s.db.Update(func(tx *bbolt.Tx) error { return put(tx, "users", u.ID, u) })
}
func (s *Store) GetUser(id string) (domain.User, error) {
	s.mu.RLock()
	defer s.mu.RUnlock()
	var u domain.User
	e := s.db.View(func(tx *bbolt.Tx) error { return get(tx, "users", id, &u) })
	return u, e
}
func (s *Store) SaveEvent(e domain.Event) error {
	s.mu.RLock()
	defer s.mu.RUnlock()
	return s.db.Update(func(tx *bbolt.Tx) error { return put(tx, "events", e.ID, e) })
}
func (s *Store) SaveAudit(a domain.Audit) error {
	s.mu.RLock()
	defer s.mu.RUnlock()
	return s.db.Update(func(tx *bbolt.Tx) error { return put(tx, "audits", a.ID, a) })
}
func (s *Store) Count(bucket string) int {
	s.mu.RLock()
	defer s.mu.RUnlock()
	n := 0
	if s.db != nil {
		s.db.View(func(tx *bbolt.Tx) error {
			b := tx.Bucket([]byte(bucket))
			if b != nil {
				b.ForEach(func(k, v []byte) error { n++; return nil })
			}
			return nil
		})
	}
	return n
}
