package scyllarepo

const msg = "errors detected in multi operation"

type MultiError struct {
	Detail []error
}

func (m *MultiError) Append(e error) {
	m.Detail = append(m.Detail, e)
}

func (m *MultiError) Error() string {
	return msg
}

func NewBufferedMultiError(size int) *MultiError {
	return &MultiError{
		Detail: make([]error, size),
	}
}

func NewDefaultMultiError(size int) *MultiError {
	return NewBufferedMultiError(0)
}
