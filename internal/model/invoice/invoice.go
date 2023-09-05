package invoice

import (
	"xlsx-builder/interfaces"
)

type Invoice struct {
	Id     string     `json:"id"`
	Date   string     `json:"date"`
	Amount string     `json:"amount"`
	Client Client     `json:"client"`
	Header []string   `json:"header"`
	Data   [][]string `json:"data"`
}

type Client struct {
	FullName  string `json:"fullName"`
	AccountId string `json:"accountId"`
	Email     string `json:"email"`
}

func Factory() func() interfaces.Sheet {
	return func() interfaces.Sheet {
		return &Invoice{}
	}
}

func (i *Invoice) Rows() [][]string {
	r := append(i.buildHeader(), i.buildTable()...)
	r = append(r, i.buildFooter()...)
	return r
}

func (i *Invoice) buildHeader() [][]string {
	return [][]string{
		{"Список туристов по счету №", i.Id, "Дата", i.Date},
		{"Заказчик:", i.Client.FullName, i.Client.AccountId},
		{"Контакты заказчика:", i.Client.Email},
	}
}

func (i *Invoice) buildTable() [][]string {
	var r [][]string
	r = append(r, i.Header)
	return append(r, i.Data...)
}

func (i *Invoice) buildFooter() [][]string {
	return [][]string{
		{"Сумма:", "", "", "", i.Amount},
	}
}
