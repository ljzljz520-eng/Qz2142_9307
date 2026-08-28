package main

import (
	"log"
	"mathrush/internal/api"
	"mathrush/internal/config"
	"mathrush/internal/service"
	"mathrush/internal/store"
	"net/http"
)

func main() {
	c := config.Load()
	s, e := store.Open(c.DBPath)
	if e != nil {
		log.Fatal(e)
	}
	defer s.Close()
	log.Fatal(http.ListenAndServe(c.Addr, api.New(service.New(s)).Handler()))
}
