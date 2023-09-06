package table

type Table struct {
	Header []string   `json:"header"`
	Data   [][]string `json:"data"`
}

func New() *Table {
	return &Table{}
}

func (t *Table) Rows() [][]string {
	var r [][]string
	r = append(r, t.Header)
	return append(r, t.Data...)
}
