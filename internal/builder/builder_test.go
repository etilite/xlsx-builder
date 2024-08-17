package builder

import (
	"bytes"
	"context"
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

			ctx := context.Background()
			r := strings.NewReader(tc.body)
			w := bytes.NewBuffer(nil)

			b := NewBuilder()

			// act
			err := b.Build(ctx, r, w)
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

	t.Run("new stream writer error and close file error", func(t *testing.T) {
		t.Parallel()

		mc := minimock.NewController(t)
		xlsxFileMock := NewXlsxFileMock(mc)
		xlsxFileMock.NewStreamWriterMock.Expect(sheetName).Return(nil, fmt.Errorf("new stream writer error"))
		xlsxFileMock.CloseMock.Expect().Return(fmt.Errorf("close file error"))

		b := NewBuilder()
		b.fileFactory = func() xlsxFile {
			return xlsxFileMock
		}

		err := b.Build(context.Background(), nil, nil)

		require.ErrorContains(t, err, "new stream writer error")
	})

	t.Run("coordinates to cell name error", func(t *testing.T) {
		t.Parallel()

		mc := minimock.NewController(t)
		decoderMock := NewDecoderMock[model.Row](mc)
		decoderMock.DecodeAndProcessMock.Set(
			func(_ context.Context, _ io.Reader, p func(ctx context.Context, elem model.Row) error) error {
				row := model.Row{}
				return p(context.Background(), row)
			},
		)

		b := NewBuilder()
		b.decoder = decoderMock
		b.startCol = -1

		err := b.Build(context.Background(), nil, nil)

		require.ErrorContains(t, err, "invalid cell reference [-1, 1]")
	})

	t.Run("set row error", func(t *testing.T) {
		t.Parallel()

		mc := minimock.NewController(t)
		streamWriterMock := NewStreamWriterMock(mc)
		streamWriterMock.SetRowMock.Expect("A1", nil).Return(fmt.Errorf("set row error"))
		xlsxFileMock := NewXlsxFileMock(mc)
		xlsxFileMock.CloseMock.Expect().Return(nil)
		decoderMock := NewDecoderMock[model.Row](mc)
		decoderMock.DecodeAndProcessMock.Set(
			func(_ context.Context, _ io.Reader, p func(ctx context.Context, elem model.Row) error) error {
				row := model.Row{}
				return p(context.Background(), row)
			},
		)

		b := NewBuilder()
		b.fileFactory = func() xlsxFile {
			return xlsxFileMock
		}
		b.streamWriterFactory = func(_ xlsxFile) (streamWriter, error) {
			return streamWriterMock, nil
		}
		b.decoder = decoderMock

		err := b.Build(context.Background(), nil, nil)

		require.ErrorContains(t, err, "set row error")
	})

	t.Run("flush error", func(t *testing.T) {
		t.Parallel()

		mc := minimock.NewController(t)
		streamWriterMock := NewStreamWriterMock(mc)
		streamWriterMock.SetRowMock.Expect("A1", nil).Return(nil)
		streamWriterMock.FlushMock.Expect().Return(fmt.Errorf("flush error"))
		xlsxFileMock := NewXlsxFileMock(mc)
		xlsxFileMock.CloseMock.Expect().Return(nil)
		decoderMock := NewDecoderMock[model.Row](mc)
		decoderMock.DecodeAndProcessMock.Set(
			func(_ context.Context, _ io.Reader, p func(ctx context.Context, elem model.Row) error) error {
				row := model.Row{}
				return p(context.Background(), row)
			},
		)

		b := NewBuilder()
		b.fileFactory = func() xlsxFile {
			return xlsxFileMock
		}
		b.streamWriterFactory = func(_ xlsxFile) (streamWriter, error) {
			return streamWriterMock, nil
		}
		b.decoder = decoderMock

		err := b.Build(context.Background(), nil, nil)

		require.ErrorContains(t, err, "flush error")
	})

	t.Run("decode and process error", func(t *testing.T) {
		t.Parallel()

		mc := minimock.NewController(t)
		decoderMock := NewDecoderMock[model.Row](mc)
		decoderMock.DecodeAndProcessMock.Set(func(_ context.Context, _ io.Reader, _ func(ctx context.Context, elem model.Row) error) (err error) {
			return fmt.Errorf("decoder error")
		})

		b := NewBuilder()
		b.decoder = decoderMock

		err := b.Build(context.Background(), nil, nil)

		require.ErrorContains(t, err, "decoder error")
	})

	t.Run("write error", func(t *testing.T) {
		t.Parallel()

		mc := minimock.NewController(t)
		decoderMock := NewDecoderMock[model.Row](mc)
		decoderMock.DecodeAndProcessMock.Set(func(_ context.Context, _ io.Reader, _ func(ctx context.Context, elem model.Row) error) (err error) {
			return nil
		})

		writerMock := NewWriterMock(mc)
		writerMock.WriteMock.Set(func(_ []byte) (n int, err error) {
			return 0, fmt.Errorf("write error")
		})

		b := NewBuilder()
		b.decoder = decoderMock

		r := strings.NewReader("")

		err := b.Build(context.Background(), r, writerMock)

		require.ErrorContains(t, err, "write error")
	})
}
