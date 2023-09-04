package http

import (
	"log"
	"net/http"
)

type Server struct {
	http *http.Server
}

func NewServer(addr string, h http.Handler) *Server {
	return &Server{
		http: &http.Server{
			Addr:    addr,
			Handler: h,
		},
	}
}

func (s *Server) Start() {
	log.Printf("Starting web server on port %s", s.http.Addr)
	log.Fatal(s.http.ListenAndServe())
}
