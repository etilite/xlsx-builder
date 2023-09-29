package invoice

type Invoice struct {
	Id     string  `json:"id"`
	Date   string  `json:"date"`
	Amount int     `json:"amount"`
	Client Client  `json:"client"`
	Header []any   `json:"header"`
	Data   [][]any `json:"data"`
}

type Client struct {
	FullName  string `json:"fullName"`
	AccountId string `json:"accountId"`
	Email     string `json:"email"`
}

func New() *Invoice {
	return &Invoice{}
}

func (i *Invoice) Rows() [][]any {
	r := append(i.buildHeader(), i.buildTable()...)
	r = append(r, i.buildFooter()...)
	return r
}

func (i *Invoice) buildHeader() [][]any {
	return [][]any{
		{"Список туристов по счету №", i.Id, "Дата", i.Date},
		{"Заказчик:", i.Client.FullName, i.Client.AccountId},
		{"Контакты заказчика:", i.Client.Email},
	}
}

func (i *Invoice) buildTable() [][]any {
	var r [][]any
	r = append(r, i.Header)
	return append(r, i.Data...)
}

func (i *Invoice) buildFooter() [][]any {
	return [][]any{
		{"Сумма:", "", "", "", i.Amount},
	}
}
