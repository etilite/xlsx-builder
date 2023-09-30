package http

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"
)

type XlsxHandler struct {
	builder Builder
}

func NewXlsxHandler(b Builder) *XlsxHandler {
	return &XlsxHandler{builder: b}
}

func (h *XlsxHandler) handleSheet(newSheet func() Sheet) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		sheet := newSheet()
		//fmt.Printf("sheet: %T, &sheet: %T\n", sheet, &sheet)
		err := json.NewDecoder(r.Body).Decode(&sheet)
		//fmt.Printf("sheet: %v, &sheet: %v\n", sheet, &sheet)
		if err != nil {
			err = fmt.Errorf("failed to decode JSON: %w", err)
			log.Print(err)
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		buf, err := h.builder.Build(sheet.Rows())
		if err != nil {
			log.Print(err)
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		h.setHeaders(w, int64(buf.Len()))
		_, err = buf.WriteTo(w)
		if err != nil {
			log.Print(err)
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	}
}

func (h *XlsxHandler) setHeaders(w http.ResponseWriter, length int64) {
	w.Header().Set("Content-Type", "application/octet-stream")
	//w.Header().Set("Data-Type", "application/vnd.openxmlformats-officedocument.spreadsheetml.sheet")
	w.Header().Set("Content-Disposition", "attachment; filename=sheet.xlsx")
	w.Header().Set("Content-Transfer-Encoding", "binary")
	w.Header().Set("Expires", "0")
	w.Header().Set("Content-Length", strconv.FormatInt(length, 10))
}
