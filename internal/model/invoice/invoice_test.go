package invoice

import (
	"reflect"
	"testing"
)

func TestRows(t *testing.T) {
	cases := map[string]struct {
		invoice *Invoice
		rows    [][]any
	}{
		"empty table from factory": {
			invoice: New(),
			rows: [][]any{
				{"Список туристов по счету №", "", "Дата", ""},
				{"Заказчик:", "", ""},
				{"Контакты заказчика:", ""},
				nil,
				{"Сумма:", "", "", "", 0},
			},
		},
		"2x2 table": {
			invoice: &Invoice{
				Id:     "123",
				Date:   "18.07.2023",
				Amount: 3,
				Client: Client{
					FullName:  "Ivanov Petr Petrovich",
					AccountId: "RU001",
					Email:     "some@mail.ru",
				},
				Header: []any{"Дата", "Кол-во"},
				Data:   [][]any{{"01.01.2023", 1}, {"01.01.2023", 2}},
			},
			rows: [][]any{
				{"Список туристов по счету №", "123", "Дата", "18.07.2023"},
				{"Заказчик:", "Ivanov Petr Petrovich", "RU001"},
				{"Контакты заказчика:", "some@mail.ru"},
				{"Дата", "Кол-во"},
				{"01.01.2023", 1},
				{"01.01.2023", 2},
				{"Сумма:", "", "", "", 3},
			},
		},
	}
	for name, tc := range cases {
		t.Run(name, func(t *testing.T) {
			tc := tc
			t.Parallel()

			sheetRows := tc.invoice.Rows()

			if !reflect.DeepEqual(tc.rows, sheetRows) {
				t.Errorf("result mismatch: want %s, got %s", tc.rows, sheetRows)
			}
		})
	}
}
