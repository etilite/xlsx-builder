package http

import (
	"net/http"

	"github.com/etilite/xlsx-builder/internal/builder"
	"github.com/etilite/xlsx-builder/internal/model/invoice"
	"github.com/etilite/xlsx-builder/internal/model/table"
)

func NewRouter() *http.ServeMux {
	b := builder.NewBuilder()
	xh := NewXlsxHandler(b)

	mux := http.NewServeMux()

	invoiceFactory := func() Sheet {
		return invoice.New()
	}
	mux.Handle("/invoice/", xh.handleSheet(invoiceFactory))

	tableFactory := func() Sheet {
		return table.New()
	}
	mux.Handle("/table/", xh.handleSheet(tableFactory))

	return mux
}
