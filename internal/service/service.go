package service

import (
	"fmt"
	"mathrush/internal/domain"
	"mathrush/internal/store"
	"sync"
)

type Service struct {
	Store         *store.Store
	mu            sync.Mutex
	notifications []domain.Event
}

func New(s *store.Store) *Service { return &Service{Store: s} }
func (s *Service) Register(id, user, expr string, answer int) (domain.Record, error) {
	r := domain.NewRecord(id, user, expr, answer)
	if e := domain.ValidateRecord(r); e != nil {
		return r, e
	}
	u, _ := s.Store.GetUser(user)
	u.ID = user
	u.Name = domain.NormalizeName(user)
	u = domain.MergeUser(u, r)
	if e := s.Store.SaveRecord(r); e != nil {
		return r, e
	}
	s.Store.SaveUser(u)
	s.Store.SaveEvent(domain.EventFor(r, "registered"))
	return r, nil
}
func (s *Service) Review(id, actor string) (domain.Record, error) {
	r, e := s.Store.GetRecord(id)
	if e != nil {
		return r, e
	}
	if e = r.Approve(); e != nil {
		return r, e
	}
	s.Store.SaveRecord(r)
	s.Store.SaveAudit(domain.AuditFor(r, actor, "approve"))
	s.Store.SaveEvent(domain.EventFor(r, "approved"))
	return r, nil
}
func (s *Service) Archive(id, actor string) (domain.Record, error) {
	r, e := s.Store.GetRecord(id)
	if e != nil {
		return r, e
	}
	if e = r.Archive(); e != nil {
		return r, e
	}
	s.Store.SaveRecord(r)
	s.Store.SaveAudit(domain.AuditFor(r, actor, "archive"))
	return r, nil
}

// Confirm applies a confirm transition to a record. The whole read-modify-write
// is serialized per record so concurrent confirmers cannot both observe the same
// Version and then clobber each other's write: each Confirm sees the result of the
// previous one and increments Version exactly once. The resulting event is keyed
// by Version (see domain.EventFor), so every confirmation is preserved as its own
// record rather than overwriting an earlier result on a fixed slot.
func (s *Service) Confirm(id string) (domain.Record, error) {
	s.mu.Lock()
	defer s.mu.Unlock()
	r, e := s.Store.GetRecord(id)
	if e != nil {
		return r, e
	}
	if e = r.Confirm(); e != nil {
		return r, e
	}
	s.Store.SaveRecord(r)
	s.Store.SaveEvent(domain.EventFor(r, "confirmed"))
	return r, nil
}
func (s *Service) ConcurrentConfirm(ids []string) error {
	var wg sync.WaitGroup
	errs := make(chan error, len(ids))
	for _, id := range ids {
		wg.Add(1)
		go func(v string) {
			defer wg.Done()
			_, e := s.Confirm(v)
			if e != nil {
				errs <- e
			}
		}(id)
	}
	wg.Wait()
	close(errs)
	for e := range errs {
		if e != nil {
			return fmt.Errorf("confirm: %w", e)
		}
	}
	return nil
}
func (s *Service) Notifications() []domain.Event {
	s.mu.Lock()
	defer s.mu.Unlock()
	return append([]domain.Event(nil), s.notifications...)
}
