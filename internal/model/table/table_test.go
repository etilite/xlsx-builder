package table

import (
	"reflect"
	"testing"
	"xlsx-builder/interfaces"
)

func TestRows(t *testing.T) {
	tests := map[string]struct {
		table interfaces.Sheet
		rows  [][]string
	}{
		"empty table from factory": {
			table: Factory()(),
			rows:  [][]string{nil},
		},
		"only header": {
			table: &Table{
				Header: []string{"col1", "col2", "col3", "col4"},
			},
			rows: [][]string{
				{"col1", "col2", "col3", "col4"},
			},
		},
		"2x2 table": {
			table: &Table{
				Header: []string{"col1", "col2"},
				Data:   [][]string{{"01.01.2023", "1"}, {"01.01.2023", "2"}},
			},
			rows: [][]string{
				{"col1", "col2"},
				{"01.01.2023", "1"},
				{"01.01.2023", "2"},
			},
		},
	}
	for name, tc := range tests {
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
