package app

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewConfigFromEnv(t *testing.T) {
	t.Run("default http address", func(t *testing.T) {
		cfg := NewConfigFromEnv()

		assert.Equal(t, ":8080", cfg.HTTPAddr)
	})

	t.Run("http address from env", func(t *testing.T) {
		t.Setenv("HTTP_ADDR", ":7777")
		cfg := NewConfigFromEnv()

		assert.Equal(t, ":7777", cfg.HTTPAddr)
	})
}
