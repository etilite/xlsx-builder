package http

import (
	"io"
)

type Builder interface {
	Build(r io.Reader, w io.Writer) error
}
