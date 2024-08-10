package builder

import (
	"context"
	"io"
	"log/slog"

	"github.com/etilite/xlsx-builder/internal/model"
	"github.com/xuri/excelize/v2"
)

const sheetName = "Sheet1"

type decoder[T any] interface {
	DecodeAndProcess(ctx context.Context, r io.Reader, process func(ctx context.Context, elem T) error) error
}

type processFn[T any] func(ctx context.Context, elem T) error

type Builder struct {
	decoder   decoder[model.Row]
	processor func(sw *excelize.StreamWriter) processFn[model.Row]
}

func NewBuilder() *Builder {
	return &Builder{
		decoder:   NewDecoder[model.Row](),
		processor: processStreamWrite,
	}
}

func (b *Builder) Build(ctx context.Context, r io.Reader, w io.Writer) error {
	f := excelize.NewFile()
	defer func() {
		if err := f.Close(); err != nil {
			slog.Error("builder: failed to close file", "error", err)
		}
	}()

	if err := b.streamWriteRows(ctx, r, f); err != nil {
		return err
	}

	if err := f.Write(w); err != nil {
		return err
	}

	return nil
}

func (b *Builder) streamWriteRows(ctx context.Context, r io.Reader, f *excelize.File) error {
	sw, err := f.NewStreamWriter(sheetName)
	if err != nil {
		return err
	}

	if err = b.decoder.DecodeAndProcess(ctx, r, b.processor(sw)); err != nil {
		return err
	}

	if err = sw.Flush(); err != nil {
		return err
	}

	return nil
}

func processStreamWrite(sw *excelize.StreamWriter) processFn[model.Row] {
	row := 1
	return func(_ context.Context, elem model.Row) error {
		cell, err := excelize.CoordinatesToCellName(1, row)
		if err != nil {
			return err
		}

		if err = sw.SetRow(cell, elem.GetData()); err != nil {
			return err
		}
		row++

		return nil
	}
}
