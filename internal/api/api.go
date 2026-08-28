package api

import (
	"encoding/json"
	"mathrush/internal/service"
	"net/http"
)

type Server struct{ Svc *service.Service }

func New(s *service.Service) *Server { return &Server{Svc: s} }
func (a *Server) Health(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("ok"))
}
func (a *Server) Register(w http.ResponseWriter, r *http.Request) {
	var in struct {
		ID, User, Expression string
		Answer               int
	}
	if json.NewDecoder(r.Body).Decode(&in) != nil {
		http.Error(w, "bad request", 400)
		return
	}
	rec, e := a.Svc.Register(in.ID, in.User, in.Expression, in.Answer)
	if e != nil {
		http.Error(w, e.Error(), 400)
		return
	}
	json.NewEncoder(w).Encode(rec)
}
func (a *Server) Get(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Query().Get("id")
	rec, e := a.Svc.Store.GetRecord(id)
	if e != nil {
		http.NotFound(w, r)
		return
	}
	json.NewEncoder(w).Encode(rec)
}
func (a *Server) Handler() http.Handler {
	m := http.NewServeMux()
	m.HandleFunc("/health", a.Health)
	m.HandleFunc("/register", a.Register)
	m.HandleFunc("/record", a.Get)
	return m
}
