package http

import (
	"net/http"
	"xlsx-builder/builder"
	"xlsx-builder/internal/model/invoice"
)

func NewRouter() *http.ServeMux {
	b := builder.NewBuilder()
	xh := NewXlsxHandler(b)

	mux := http.NewServeMux()
	mux.Handle("/invoice/", xh.handleSheet(invoice.Factory()))

	return mux
}
