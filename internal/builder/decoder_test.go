package builder

import (
	"errors"
	"strings"
	"testing"

	"github.com/etilite/xlsx-builder/internal/model"
	"github.com/stretchr/testify/require"
)

func TestDecoder_DecodeAndProcess(t *testing.T) {
	t.Parallel()

	t.Run("success with row model", func(t *testing.T) {
		t.Parallel()

		body := `[
        	{"data": ["01.01.2023", 1, 10.5]},
        	{"data": ["02.01.2023", 2, 20.3]},
        	{"data": ["03.01.2023", 3, "33"]}
    	]`

		want := [][]any{
			{"01.01.2023", 1.0, 10.5},
			{"02.01.2023", 2.0, 20.3},
			{"03.01.2023", 3.0, "33"},
		}

		d := NewDecoder[model.Row]()

		r := make([][]any, 0, 3)

		err := d.DecodeAndProcess(strings.NewReader(body), func(elem model.Row) error {
			row := make([]any, 3)
			copy(row, elem.GetData())
			r = append(r, row)
			return nil
		})

		require.NoError(t, err)
		require.Equal(t, 3, len(r))
		require.Equal(t, want, r)
	})

	t.Run("error opening token decode", func(t *testing.T) {
		t.Parallel()
		body := ""

		d := NewDecoder[int]()
		err := d.DecodeAndProcess(strings.NewReader(body), func(elem int) error {
			return nil
		})

		require.Error(t, err)
		require.ErrorIs(t, err, ErrDecode)
	})

	t.Run("error decode first elem", func(t *testing.T) {
		t.Parallel()
		body := `{"a": 1}`

		d := NewDecoder[int]()
		err := d.DecodeAndProcess(strings.NewReader(body), func(elem int) error {
			return nil
		})

		require.Error(t, err)
		require.ErrorIs(t, err, ErrDecode)
	})

	t.Run("error process", func(t *testing.T) {
		t.Parallel()
		body := `[1]`

		errCustom := errors.New("failed process")
		d := NewDecoder[int]()
		err := d.DecodeAndProcess(strings.NewReader(body), func(elem int) error {
			return errCustom
		})

		require.Error(t, err)
		require.ErrorIs(t, err, errCustom)
	})
}
