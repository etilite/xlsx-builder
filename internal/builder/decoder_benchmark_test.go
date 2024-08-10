package builder

import (
	"context"
	"strings"
	"testing"

	"github.com/etilite/xlsx-builder/internal/model"
)

func BenchmarkDecoder(b *testing.B) {
	body := `[
        {"data": ["01.01.2023", 1, 10.5]},
        {"data": ["02.01.2023", 2, 20.3]},
        {"data": ["03.01.2023", 3, "33"]}
    ]`

	for i := 0; i < b.N; i++ {
		d := NewDecoder[model.Row]()
		err := d.DecodeAndProcess(context.Background(), strings.NewReader(body), func(_ context.Context, _ model.Row) error {
			return nil
		})
		if err != nil {
			b.Fatal(err)
		}
	}
}
