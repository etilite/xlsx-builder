package testing

import (
	"net/http/httptest"
	"testing"
)

func AssertStatusCode(t *testing.T, w *httptest.ResponseRecorder, code int) {
	t.Helper()
	got := w.Code
	want := code
	if got != want {
		t.Errorf("Status Code = %v; want %v", got, want)
	}
}

func AssertHeader(t *testing.T, w *httptest.ResponseRecorder, header string, value string) {
	t.Helper()
	got := w.Result().Header.Get(header)
	want := value
	if got != want {
		t.Errorf("%v = %v; want %v", header, got, want)
	}
}

func AssertBody(t *testing.T, w *httptest.ResponseRecorder, body string) {
	t.Helper()
	got := w.Body.String()
	want := body
	if got != want {
		t.Errorf("Body = %v; want %v", got, want)
	}
}
