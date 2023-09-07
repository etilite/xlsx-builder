package builder

import (
	"bytes"
	"github.com/xuri/excelize/v2"
	"log"
)

type Builder struct{}

func NewBuilder() *Builder {
	return &Builder{}
}

func (b *Builder) Build(rows [][]string) (*bytes.Buffer, error) {
	f := excelize.NewFile()
	defer func() {
		if err := f.Close(); err != nil {
			log.Print(err)
		}
	}()
	err := fillRows(rows, f)
	if err != nil {
		return nil, err
	}

	buf, err := f.WriteToBuffer()
	if err != nil {
		return nil, err
	}

	return buf, nil
}

func fillRows(rows [][]string, f *excelize.File) error {
	for i, r := range rows {
		cell, err := excelize.CoordinatesToCellName(1, i+1)
		if err != nil {
			return err
		}
		err = f.SetSheetRow("Sheet1", cell, &r)
		if err != nil {
			return err
		}
	}
	return nil
}
