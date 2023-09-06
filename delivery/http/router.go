package http

import (
	"net/http"
	"xlsx-builder/builder"
	"xlsx-builder/interfaces"
	"xlsx-builder/internal/model/invoice"
)

func NewRouter() *http.ServeMux {
	b := builder.NewBuilder()
	xh := NewXlsxHandler(b)

	mux := http.NewServeMux()

	f := func() interfaces.Sheet {
		return invoice.New()
	}

	mux.Handle("/invoice/", xh.handleSheet(f))

	return mux
}
