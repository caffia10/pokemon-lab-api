package scyllarepo

import (
	"fmt"
	"strings"
)

const msg = "errors detected in multi operation"

type MultiError struct {
	detail []error
}

func (m *MultiError) Append(e error) {
	m.detail = append(m.detail, e)
}

func (m *MultiError) Error() string {
	var sb strings.Builder
	for i, d := range m.detail {
		sb.WriteString(fmt.Sprintf("- index: %d - detail: %s ", i, d.Error()))
	}
	return sb.String()
}

func NewBufferedMultiError(size int) *MultiError {
	return &MultiError{
		detail: make([]error, size),
	}
}

func NewDefaultMultiError(size int) *MultiError {
	return NewBufferedMultiError(0)
}
