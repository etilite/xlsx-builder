package invoice

import (
	"reflect"
	"testing"
	"xlsx-builder/interfaces"
)

func TestRows(t *testing.T) {
	type test struct {
		name    string
		invoice interfaces.Sheet
		rows    [][]string
	}
	tests := []test{
		{
			name:    "empty table from factory",
			invoice: Factory()(),
			rows: [][]string{
				{"Список туристов по счету №", "", "Дата", ""},
				{"Заказчик:", "", ""},
				{"Контакты заказчика:", ""},
				nil,
				{"Сумма:", "", "", "", ""},
			},
		},
		{
			name: "2x2 table",
			invoice: &Invoice{
				Id:     "123",
				Date:   "18.07.2023",
				Amount: "3",
				Client: Client{
					FullName:  "Ivanov Petr Petrovich",
					AccountId: "RU001",
					Email:     "some@mail.ru",
				},
				Header: []string{"Дата", "Кол-во"},
				Data:   [][]string{{"01.01.2023", "1"}, {"01.01.2023", "2"}},
			},
			rows: [][]string{
				{"Список туристов по счету №", "123", "Дата", "18.07.2023"},
				{"Заказчик:", "Ivanov Petr Petrovich", "RU001"},
				{"Контакты заказчика:", "some@mail.ru"},
				{"Дата", "Кол-во"},
				{"01.01.2023", "1"},
				{"01.01.2023", "2"},
				{"Сумма:", "", "", "", "3"},
			},
		},
	}
	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			tc := tc
			t.Parallel()

			sheetRows := tc.invoice.Rows()

			if !reflect.DeepEqual(tc.rows, sheetRows) {
				t.Errorf("result mismatch: want %s, got %s", tc.rows, sheetRows)
			}
		})
	}
}
