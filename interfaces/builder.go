package interfaces

import (
	"bytes"
)

type Builder interface {
	Build(s Sheet) (*bytes.Buffer, error)
}
