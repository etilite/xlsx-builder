package http

import (
	"context"
	"errors"
	"log"
	"net/http"
	"time"
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

func (s *Server) Run(ctx context.Context) {
	go func() {
		<-ctx.Done()
		shutdownCtx, done := context.WithTimeout(context.Background(), 5*time.Second)
		defer done()

		log.Printf("http-server: shutting down")
		if err := s.http.Shutdown(shutdownCtx); err != nil {
			log.Printf("http-server: shutdown error: %v", err)
		}
	}()

	log.Printf("http-server: starting web server on port %s", s.http.Addr)
	if err := s.http.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
		log.Printf("http-server: failed to serve: %v", err)
	}

	log.Printf("http-server: stopped gracefully")
}
