package http

import (
	"net/http"

	"github.com/etilite/xlsx-builder/internal/builder"
)

func NewMux() *http.ServeMux {
	b := builder.NewBuilder()
	xh := NewXlsxHandler(b)

	mux := http.NewServeMux()

	mux.Handle("POST /api/build", xh.handlerFn())

	return mux
}
