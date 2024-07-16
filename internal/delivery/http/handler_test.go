package http

import (
	"bytes"
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/etilite/xlsx-builder/internal/builder"
	"github.com/gojuno/minimock/v3"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestHandlerFn(t *testing.T) {
	t.Parallel()

	t.Run("success", func(t *testing.T) {
		t.Parallel()

		mc := minimock.NewController(t)
		builderMock := NewBuilderMock(mc)
		builderMock.BuildMock.Set(func(r io.Reader, w io.Writer) (err error) {
			buf := &bytes.Buffer{}
			_, readErr := buf.ReadFrom(r)
			if readErr != nil {
				return readErr
			}
			_, writeErr := buf.WriteTo(w)
			if writeErr != nil {
				return writeErr
			}

			return nil
		})
		server := NewXlsxHandler(builderMock)
		requestBody := strings.NewReader("body-file-contents")
		request, _ := http.NewRequest(http.MethodPost, "/v2", requestBody)
		response := httptest.NewRecorder()
		handleFunc := server.handlerFn()

		handleFunc(response, request)
		defer request.Body.Close()

		require.Equal(t, http.StatusOK, response.Code)
		require.Equal(t, "body-file-contents", response.Body.String())
		assertHeaders(t, response)
	})

	t.Run("bad request", func(t *testing.T) {
		t.Parallel()

		mc := minimock.NewController(t)
		builderMock := NewBuilderMock(mc)
		builderMock.BuildMock.Set(func(r io.Reader, w io.Writer) (err error) {
			return builder.ErrDecode
		})
		server := NewXlsxHandler(builderMock)

		require.HTTPStatusCode(t, server.handlerFn(), http.MethodPost, "/v2", nil, http.StatusBadRequest)
		require.HTTPBodyContains(t, server.handlerFn(), http.MethodPost, "/v2", nil, "decode failed")
	})

	t.Run("internal server error", func(t *testing.T) {
		t.Parallel()

		mc := minimock.NewController(t)
		builderMock := NewBuilderMock(mc)
		builderMock.BuildMock.Set(func(r io.Reader, w io.Writer) (err error) {
			return fmt.Errorf("some error")
		})
		server := NewXlsxHandler(builderMock)

		require.HTTPStatusCode(t, server.handlerFn(), http.MethodPost, "/v2", nil, http.StatusInternalServerError)
		require.HTTPBodyContains(t, server.handlerFn(), http.MethodPost, "/v2", nil, "some error")
	})
}

func assertHeaders(t *testing.T, response *httptest.ResponseRecorder) {
	assertHeader(t, response, "Content-Disposition", "attachment; filename=sheet.xlsx")
	assertHeader(t, response, "Content-Type", "application/octet-stream")
	assertHeader(t, response, "Content-Transfer-Encoding", "binary")
	assertHeader(t, response, "Expires", "0")
}

func assertHeader(t *testing.T, w *httptest.ResponseRecorder, header string, value string) {
	t.Helper()

	assert.Equal(t, value, w.Result().Header.Get(header))
}
