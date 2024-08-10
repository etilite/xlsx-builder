package builder

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"log/slog"
)

// ErrDecode is type of error returned when decoding fails.
var ErrDecode = errors.New("decode failed")

type Decoder[T any] struct{}

func NewDecoder[T any]() *Decoder[T] {
	return &Decoder[T]{}
}

// DecodeAndProcess decodes json stream and applies process func to each element.
// If [T] has any slice fields underneath arrays will be overwritten with each iteration
// so u should copy this slice for further processing.
// returns an error wrapping ErrDecode when decoding is failed due to invalid input data.
func (d *Decoder[T]) DecodeAndProcess(ctx context.Context, r io.Reader, process func(ctx context.Context, elem T) error) error {
	jsonDecoder := json.NewDecoder(r)

	if _, err := jsonDecoder.Token(); err != nil {
		return fmt.Errorf("%w: invalid opening token: %v", ErrDecode, err)
	}

	var in T
	slog.Info("decoder: cycle")
	for jsonDecoder.More() {
		select {
		case <-ctx.Done():
			return ctx.Err()
		default:
		}

		if err := jsonDecoder.Decode(&in); err != nil {
			return fmt.Errorf("%w: %v", ErrDecode, err)
		}
		if err := process(ctx, in); err != nil {
			return fmt.Errorf("err process elem %v: %w", in, err)
		}
	}

	return nil
}
