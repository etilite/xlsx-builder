package http

import (
	"errors"
	"log"
	"net/http"

	"github.com/etilite/xlsx-builder/internal/builder"
)

type XlsxHandler struct {
	builder Builder
}

func NewXlsxHandler(b Builder) *XlsxHandler {
	return &XlsxHandler{builder: b}
}

func (h *XlsxHandler) handlerFn() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		h.setHeaders(w)
		if err := h.builder.Build(r.Body, w); err != nil {
			log.Print(err)
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
