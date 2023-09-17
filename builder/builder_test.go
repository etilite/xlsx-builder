package builder

import (
	"reflect"
	"strings"
	"testing"

	"github.com/xuri/excelize/v2"
)

func TestBuild(t *testing.T) {
	tests := map[string]struct {
		rows [][]string
	}{
		"empty table": {rows: [][]string{}},
		"2x2 table":   {rows: [][]string{{"a", "b"}, {"c", "d"}}},
	}
	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {
			tc := tc
			t.Parallel()

			b := NewBuilder()

			buf, err := b.Build(tc.rows)
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

			if !reflect.DeepEqual(rows, tc.rows) {
				t.Errorf("result mismatch in test %s: want %s, got %s", name, tc.rows, rows)
			}
		})
	}
}
