package builder

import (
	"bytes"
	"fmt"
	"io"
	"strings"
	"testing"

	"github.com/etilite/xlsx-builder/internal/model"
	"github.com/gojuno/minimock/v3"
	"github.com/stretchr/testify/require"
	"github.com/xuri/excelize/v2"
)

func TestBuilder_Build(t *testing.T) {
	t.Parallel()

	cases := map[string]struct {
		body string
		want [][]string
	}{
		"empty table":   {body: `[]`, want: [][]string{}},
		"empty row":     {body: `[{"data": []}]`, want: [][]string{}},
		"one row table": {body: `[{"data": ["01.01.2023", 1, 10.5]}]`, want: [][]string{{"01.01.2023", "1", "10.5"}}},
		"2x4 table": {
			body: `[{"data": ["01.01.2023", 1, 10.5, "text"]}, {"data": ["08.06.2024", 2, 20.3, "текст"]}]`,
			want: [][]string{{"01.01.2023", "1", "10.5", "text"}, {"08.06.2024", "2", "20.3", "текст"}},
		},
	}

	for name, tc := range cases {
		tc := tc

		t.Run(name, func(t *testing.T) {
			t.Parallel()

			r := strings.NewReader(tc.body)
			w := bytes.NewBuffer(nil)

			b := NewBuilder()

			// act
			err := b.Build(r, w)
			require.NoError(t, err)

			f, err := excelize.OpenReader(w)
			require.NoError(t, err)

			// assert
			got, err := f.GetRows(sheetName)
			require.NoError(t, err)
			require.Equal(t, tc.want, got)
		})
	}
}

func TestBuilder_Build_errors(t *testing.T) {
	t.Parallel()

	t.Run("decode and process error", func(t *testing.T) {
		t.Parallel()

		mc := minimock.NewController(t)
		decoderMock := NewDecoderMock[model.Row](mc)
		decoderMock.DecodeAndProcessMock.Set(func(_ io.Reader, _ func(model.Row) error) (err error) {
			return fmt.Errorf("decoder error")
		})

		b := NewBuilder()
		b.decoder = decoderMock

		err := b.Build(nil, nil)
		require.ErrorContains(t, err, "decoder error")
	})

	t.Run("write error", func(t *testing.T) {
		t.Parallel()

		mc := minimock.NewController(t)
		decoderMock := NewDecoderMock[model.Row](mc)
		decoderMock.DecodeAndProcessMock.Set(func(_ io.Reader, _ func(model.Row) error) (err error) {
			return nil
		})

		writerMock := NewWriterMock(mc)
		writerMock.WriteMock.Set(func(_ []byte) (n int, err error) {
			return 0, fmt.Errorf("write error")
		})

		b := NewBuilder()
		b.decoder = decoderMock

		r := strings.NewReader("")

		err := b.Build(r, writerMock)
		require.ErrorContains(t, err, "write error")
	})
}
