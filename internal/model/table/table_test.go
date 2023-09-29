package table

import (
	"reflect"
	"testing"
)

func TestRows(t *testing.T) {
	cases := map[string]struct {
		table *Table
		rows  [][]any
	}{
		"empty table from factory": {
			table: New(),
			rows:  [][]any{nil},
		},
		"only header": {
			table: &Table{
				Header: []any{"col1", "col2", "col3", "col4"},
			},
			rows: [][]any{
				{"col1", "col2", "col3", "col4"},
			},
		},
		"2x2 table": {
			table: &Table{
				Header: []any{"col1", "col2"},
				Data:   [][]any{{"01.01.2023", 1}, {"01.01.2023", 2.1}},
			},
			rows: [][]any{
				{"col1", "col2"},
				{"01.01.2023", 1},
				{"01.01.2023", 2.1},
			},
		},
	}
	for name, tc := range cases {
		t.Run(name, func(t *testing.T) {
			tc := tc
			t.Parallel()

			sheetRows := tc.table.Rows()

			if !reflect.DeepEqual(tc.rows, sheetRows) {
				t.Errorf("result mismatch: want %s, got %s", tc.rows, sheetRows)
			}
		})
	}
}
