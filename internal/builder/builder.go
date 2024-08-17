package builder

import (
	"context"
	"io"
	"log/slog"

	"github.com/etilite/xlsx-builder/internal/model"
	"github.com/xuri/excelize/v2"
)

const (
	sheetName = "Sheet1"
	startCol  = 1
	startRow  = 1
)

type decoder[T any] interface {
	DecodeAndProcess(ctx context.Context, r io.Reader, process func(ctx context.Context, elem T) error) error
}

type xlsxFile interface {
	Write(w io.Writer, opts ...excelize.Options) error
	NewStreamWriter(sheet string) (*excelize.StreamWriter, error)
	Close() error
}

type streamWriter interface {
	SetRow(cell string, values []interface{}, opts ...excelize.RowOpts) error
	Flush() error
}

type processFn[T any] func(ctx context.Context, elem T) error

type Builder struct {
	decoder             decoder[model.Row]
	fileFactory         func() xlsxFile
	streamWriterFactory func(f xlsxFile) (streamWriter, error)
	processor           func(sw streamWriter, startCol, startRow int) processFn[model.Row]
	startCol            int
	startRow            int
}

func NewBuilder() *Builder {
	return &Builder{
		decoder:             NewDecoder[model.Row](),
		fileFactory:         newXlsxFile,
		streamWriterFactory: newStreamWriter,
		processor:           processStreamWrite,
		startCol:            startCol,
		startRow:            startRow,
	}
}

func (b *Builder) Build(ctx context.Context, r io.Reader, w io.Writer) error {
	f := b.fileFactory()
	defer func() {
		if err := f.Close(); err != nil {
			slog.Error("builder: failed to close file", "error", err)
		}
	}()

	sw, err := b.streamWriterFactory(f)
	if err != nil {
		return err
	}

	if err = b.streamWriteRows(ctx, r, sw); err != nil {
		return err
	}

	if err = f.Write(w); err != nil {
		return err
	}

	return nil
}

func (b *Builder) streamWriteRows(ctx context.Context, r io.Reader, sw streamWriter) error {
	if err := b.decoder.DecodeAndProcess(ctx, r, b.processor(sw, b.startCol, b.startRow)); err != nil {
		return err
	}

	if err := sw.Flush(); err != nil {
		return err
	}

	return nil
}

func processStreamWrite(sw streamWriter, startCol, startRow int) processFn[model.Row] {
	return func(_ context.Context, elem model.Row) error {
		cell, err := excelize.CoordinatesToCellName(startCol, startRow)
		if err != nil {
			return err
		}

		if err = sw.SetRow(cell, elem.GetData()); err != nil {
			return err
		}
		startRow++

		return nil
	}
}

func newXlsxFile() xlsxFile {
	return excelize.NewFile()
}

func newStreamWriter(f xlsxFile) (streamWriter, error) {
	return f.NewStreamWriter(sheetName)
}
