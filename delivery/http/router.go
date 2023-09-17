package http

import (
	"net/http"

	"github.com/etilite/xlsx-builder/builder"
	"github.com/etilite/xlsx-builder/internal/model/invoice"
)

func NewRouter() *http.ServeMux {
	b := builder.NewBuilder()
	xh := NewXlsxHandler(b)

	mux := http.NewServeMux()

	f := func() Sheet {
		return invoice.New()
	}

	mux.Handle("/invoice/", xh.handleSheet(f))

	return mux
}
