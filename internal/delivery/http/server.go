package http

import (
	"context"
	"errors"
	"log/slog"
	"net/http"
	"time"
)

type httpServer interface {
	ListenAndServe() error
	Shutdown(ctx context.Context) error
}

type Server struct {
	http httpServer
	addr string
}

func NewServer(addr string) *Server {
	return &Server{
		http: &http.Server{
			Addr:    addr,
			Handler: NewMux(),
		},
		addr: addr,
	}
}

func (s *Server) Run(ctx context.Context) error {
	errCh := make(chan error, 1)
	go func() {
		<-ctx.Done()

		errCh <- s.shutdown()
	}()

	slog.Info("http-server: starting web server", "address", s.addr)

	if err := s.http.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
		slog.Error("http-server: failed to serve", "error", err)
		return err
	}

	slog.Info("http-server: shutting down")

	if err := <-errCh; err != nil {
		slog.Error("http-server: error stopping server", "error", err)
		return err
	}

	slog.Info("http-server: stopped gracefully")

	return nil
}

func (s *Server) shutdown() error {
	shutdownCtx, done := context.WithTimeout(context.Background(), 5*time.Second)
	defer done()

	if err := s.http.Shutdown(shutdownCtx); err != nil {
		return err
	}

	return nil
}
