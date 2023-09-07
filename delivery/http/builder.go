package http

import (
	"bytes"
)

type Builder interface {
	Build(rows [][]string) (*bytes.Buffer, error)
}
