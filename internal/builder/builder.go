package builder

import (
	"io"
	"log"

	"github.com/etilite/xlsx-builder/internal/model"
	"github.com/xuri/excelize/v2"
)

const sheetName = "Sheet1"

type decoder[T any] interface {
	DecodeAndProcess(r io.Reader, process func(T) error) error
}

type processFn func(elem model.Row) error

type Builder struct {
	decoder   decoder[model.Row]
	processor func(sw *excelize.StreamWriter) processFn
}

func NewBuilder() *Builder {
	return &Builder{
		decoder:   NewDecoder[model.Row](),
		processor: processStreamWrite,
	}
}

func (b *Builder) Build(r io.Reader, w io.Writer) error {
	f := excelize.NewFile()
	defer func() {
		if err := f.Close(); err != nil {
			log.Print(err)
		}
	}()

	if err := b.streamWriteRows(r, f); err != nil {
		return err
	}

	if err := f.Write(w); err != nil {
		return err
	}

	return nil
}

func (b *Builder) streamWriteRows(r io.Reader, f *excelize.File) error {
	sw, err := f.NewStreamWriter(sheetName)
	if err != nil {
		return err
	}

	if err = b.decoder.DecodeAndProcess(r, b.processor(sw)); err != nil {
		return err
	}

	if err = sw.Flush(); err != nil {
		return err
	}

	return nil
}

func processStreamWrite(sw *excelize.StreamWriter) processFn {
	row := 1
	return func(elem model.Row) error {
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
