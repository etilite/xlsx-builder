package http

import (
	"bytes"
)

type Builder interface {
	Build(rows [][]any) (*bytes.Buffer, error)
}
