package table

type Table struct {
	Header []any   `json:"header"`
	Data   [][]any `json:"data"`
}

func New() *Table {
	return &Table{}
}

func (t *Table) Rows() [][]any {
	var r [][]any
	r = append(r, t.Header)
	return append(r, t.Data...)
}
