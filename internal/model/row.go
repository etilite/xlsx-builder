package model

type Row struct {
	Data []any `json:"data"`
}

func (r *Row) GetData() []any {
	return r.Data
}
