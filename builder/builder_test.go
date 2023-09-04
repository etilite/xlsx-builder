package builder

import (
	"github.com/xuri/excelize/v2"
	"reflect"
	"strings"
	"testing"
)

type mockSheet struct {
	rows [][]string
}

func (s mockSheet) Rows() [][]string {
	return s.rows
}

func TestBuild(t *testing.T) {
	tests := map[string]struct {
		sheet mockSheet
	}{
		"empty table": {sheet: mockSheet{rows: [][]string{}}},
		"2x2 table":   {sheet: mockSheet{rows: [][]string{{"a", "b"}, {"c", "d"}}}},
	}
	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {
			tc := tc
			t.Parallel()

			b := NewBuilder()

			buf, err := b.Build(tc.sheet)
			if err != nil {
				t.Error(err)
			}

			f, err := excelize.OpenReader(strings.NewReader(buf.String()))
			if err != nil {
				t.Error(err)
			}

			rows, err := f.GetRows("Sheet1")
			if err != nil {
				t.Error("got error getting rows")
			}

			if !reflect.DeepEqual(rows, tc.sheet.Rows()) {
				t.Errorf("result mismatch in test %s: want %s, got %s", name, tc.sheet.Rows(), rows)
			}
		})
	}
}
