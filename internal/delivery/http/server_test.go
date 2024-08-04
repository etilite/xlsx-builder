package http

import (
	"context"
	"errors"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestServer(t *testing.T) {
	t.Parallel()

	t.Run("success run and shutdown", func(t *testing.T) {
		server := NewServer(":12345")
		server.http = newHttpServerMock(false, false)

		ctx, cancel := context.WithTimeout(context.Background(), 32*time.Millisecond)
		defer cancel()

		err := server.Run(ctx)

		assert.NoError(t, err)
	})

	t.Run("error run", func(t *testing.T) {
		server := NewServer(":12345")
		server.http = newHttpServerMock(true, false)

		err := server.Run(context.Background())

		assert.Error(t, err)
		assert.ErrorContains(t, err, "listen error")
	})

	t.Run("error shutdown", func(t *testing.T) {
		server := NewServer(":12345")
		server.http = newHttpServerMock(false, true)

		ctx, cancel := context.WithTimeout(context.Background(), 32*time.Millisecond)
		defer cancel()

		err := server.Run(ctx)

		assert.Error(t, err)
		assert.ErrorContains(t, err, "shutdown error")
	})

}

func newHttpServerMock(listenErr, shutdownErr bool) *mockHttpServer {
	return &mockHttpServer{
		ch:          make(chan bool),
		listenErr:   listenErr,
		shutdownErr: shutdownErr,
	}
}

type mockHttpServer struct {
	ch          chan bool
	listenErr   bool
	shutdownErr bool
}

func (s *mockHttpServer) ListenAndServe() error {
	if s.listenErr {
		return errors.New("listen error")
	}

	<-s.ch

	return nil
}
func (s *mockHttpServer) Shutdown(_ context.Context) error {
	s.ch <- true

	if s.shutdownErr {
		return errors.New("shutdown error")
	}

	return nil
}
