package http

import (
	"context"
	"errors"
	"io"
	"log/slog"
	"net/http"

	"github.com/etilite/xlsx-builder/internal/builder"
)

type Builder interface {
	Build(ctx context.Context, r io.Reader, w io.Writer) error
}

type XlsxHandler struct {
	builder Builder
}

func NewXlsxHandler(b Builder) *XlsxHandler {
	return &XlsxHandler{builder: b}
}

func (h *XlsxHandler) handlerFn() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		h.setHeaders(w)
		if err := h.builder.Build(r.Context(), r.Body, w); err != nil {
			slog.Error("handler: failed to build xlsx", "error", err)
			// if w.Write() is called, there will be error msg in log
			// starting with "http: superfluous response.WriteHeader call..."
			// this happens cause w.Write() sets http.StatusOK
			http.Error(w, err.Error(), h.getErrorHTTPStatus(err))
			return
		}
	}
}

func (h *XlsxHandler) setHeaders(w http.ResponseWriter) {
	w.Header().Set("Content-Type", "application/octet-stream")
	w.Header().Set("Content-Disposition", "attachment; filename=sheet.xlsx")
	w.Header().Set("Content-Transfer-Encoding", "binary")
	w.Header().Set("Expires", "0")
}

func (h *XlsxHandler) getErrorHTTPStatus(err error) int {
	switch true {
	case errors.Is(err, builder.ErrDecode):
		return http.StatusBadRequest
	default:
		return http.StatusInternalServerError
	}
}
