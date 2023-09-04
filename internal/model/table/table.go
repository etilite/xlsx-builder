package table

import (
	"xlsx-builder/interfaces"
)

type Table struct {
	Header []string   `json:"header"`
	Data   [][]string `json:"data"`
}

func Factory() func() interfaces.Sheet {
	return func() interfaces.Sheet {
		return &Table{}
	}
}

func (t *Table) Rows() [][]string {
	var r [][]string
	r = append(r, t.Header)
	return append(r, t.Data...)
}
