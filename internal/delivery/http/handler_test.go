package http

import (
	"bytes"
	"errors"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	. "github.com/etilite/xlsx-builder/internal/testing"
)

type mockBuilder struct {
	buf *bytes.Buffer
	err error
}

func (b *mockBuilder) Build([][]any) (*bytes.Buffer, error) {
	return b.buf, b.err
}

type mockSheet struct {
	rows [][]any
}

func (s mockSheet) Rows() [][]any {
	return s.rows
}

func TestCreateInvoiceXlsx(t *testing.T) {
	mockFactory := func() Sheet {
		return &mockSheet{}
	}
	t.Run("success", func(t *testing.T) {
		t.Parallel()

		buf := &bytes.Buffer{}
		buf.Write([]byte("body-file-contents"))
		builder := &mockBuilder{buf: buf}
		server := NewXlsxHandler(builder)
		handleFunc := server.handleSheet(mockFactory)

		requestBody := strings.NewReader("{}")
		request, _ := http.NewRequest(http.MethodPost, "/invoice/", requestBody)
		response := httptest.NewRecorder()

		handleFunc(response, request)

		AssertStatusCode(t, response, http.StatusOK)
		AssertBody(t, response, "body-file-contents")
		assertHeaders(t, response)
	})

	t.Run("bad request", func(t *testing.T) {
		t.Parallel()

		builder := &mockBuilder{}
		server := NewXlsxHandler(builder)
		handleFunc := server.handleSheet(mockFactory)

		requestBody := strings.NewReader("")
		request, _ := http.NewRequest(http.MethodPost, "/invoice/", requestBody)
		response := httptest.NewRecorder()

		handleFunc(response, request)

		AssertStatusCode(t, response, http.StatusBadRequest)
		AssertBody(t, response, "failed to decode JSON: EOF\n")
	})

	t.Run("internal server error", func(t *testing.T) {
		t.Parallel()

		builder := &mockBuilder{err: errors.New("internal server error")}
		server := NewXlsxHandler(builder)
		handleFunc := server.handleSheet(mockFactory)

		requestBody := strings.NewReader("{}")
		request, _ := http.NewRequest(http.MethodPost, "/invoice/", requestBody)
		response := httptest.NewRecorder()

		handleFunc(response, request)

		AssertStatusCode(t, response, http.StatusInternalServerError)
		AssertBody(t, response, "internal server error\n")
	})
}

func assertHeaders(t *testing.T, response *httptest.ResponseRecorder) {
	AssertHeader(t, response, "Content-Disposition", "attachment; filename=sheet.xlsx")
	AssertHeader(t, response, "Content-Type", "application/octet-stream")
	//AssertHeader(t, response, "Content-Type", "application/vnd.openxmlformats-officedocument.spreadsheetml.sheet")
	AssertHeader(t, response, "Content-Transfer-Encoding", "binary")
	AssertHeader(t, response, "Expires", "0")
	AssertHeader(t, response, "Content-Length", "18")
}
