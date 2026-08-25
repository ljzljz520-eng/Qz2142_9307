package service

import "mathrush/internal/domain"

func (s *Service) SubmitAndReview(id, user, expr string, answer int) (domain.Record, error) {
	r, e := s.Register(id, user, expr, answer)
	if e != nil {
		return r, e
	}
	return s.Review(id, user)
}
func (s *Service) ReviewAndArchive(id, actor string) (domain.Record, error) {
	if _, e := s.Review(id, actor); e != nil {
		var z domain.Record
		return z, e
	}
	return s.Archive(id, actor)
}
func (s *Service) Track(id string) (domain.Record, string, error) {
	r, e := s.Store.GetRecord(id)
	if e != nil {
		return r, "", e
	}
	return r, domain.Rank(r.Score), nil
}
func (s *Service) EnsureUser(id, name string) error {
	u := domain.User{ID: id, Name: name}
	return s.Store.SaveUser(u)
}
func (s *Service) EventCount() int { return s.Store.Count("events") }
