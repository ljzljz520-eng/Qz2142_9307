package api

import "net/http"

func (a *Server) Routes() []string                      { return []string{"/health", "/register", "/record"} }
func writeJSON(w http.ResponseWriter, v any)            { w.Header().Set("Content-Type", "application/json") }
func methodAllowed(r *http.Request, method string) bool { return r.Method == method }
