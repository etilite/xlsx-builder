package http

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewMux(t *testing.T) {
	t.Parallel()
	tests := map[string]struct {
		method string
		path   string
	}{
		"api/build": {
			method: http.MethodPost,
			path:   "/api/build",
		},
	}
	for name, tt := range tests {
		tt := tt

		t.Run(name, func(t *testing.T) {
			t.Parallel()

			mux := NewMux()

			requestBody := strings.NewReader(`[]`)
			request := httptest.NewRequest(tt.method, tt.path, requestBody)
			response := httptest.NewRecorder()

			mux.ServeHTTP(response, request)

			assert.Equal(t, http.StatusOK, response.Code)
		})
	}
}
