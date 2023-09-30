package http

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	. "github.com/etilite/xlsx-builder/internal/testing"
)

func TestNewRouter(t *testing.T) {
	t.Parallel()
	cases := map[string]struct {
		path string
	}{
		"invoice": {
			path: "/invoice/",
		},
		"table": {
			path: "/table/",
		},
	}
	for name, tc := range cases {
		t.Run(name, func(t *testing.T) {
			tc := tc
			t.Parallel()

			router := NewRouter()

			requestBody := strings.NewReader("{}")
			r, _ := http.NewRequest(http.MethodPost, tc.path, requestBody)
			w := httptest.NewRecorder()

			router.ServeHTTP(w, r)
			w.Flush()

			AssertStatusCode(t, w, http.StatusOK)
		})
	}
}
