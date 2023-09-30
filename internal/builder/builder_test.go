package builder

import (
	"reflect"
	"strings"
	"testing"

	"github.com/xuri/excelize/v2"
)

func TestBuild(t *testing.T) {
	cases := map[string]struct {
		rows   [][]any
		result [][]string
	}{
		"empty table": {rows: [][]any{}, result: [][]string{}},
		"2x2 table":   {rows: [][]any{{"a", 1}, {"c", 2.1}}, result: [][]string{{"a", "1"}, {"c", "2.1"}}},
	}
	for name, tc := range cases {
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

			if !reflect.DeepEqual(rows, tc.result) {
				t.Errorf("result mismatch in test %s: want %s, got %s", name, tc.result, rows)
			}
		})
	}
}
