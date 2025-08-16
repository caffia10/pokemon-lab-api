package mderrors

import (
	"fmt"
	"path/filepath"
	"runtime"
)

type Metadata struct {
	Key   string
	Value any
}

func MakeIdMetaData(value string) Metadata {
	return Metadata{Key: "Id", Value: value}
}

type MetadataError struct {
	inner error
	Data  []Metadata
	Stack []string
}

func (md *MetadataError) Error() string {
	return md.inner.Error()
}

func (md *MetadataError) Unwrap() error {
	return md.inner
}

func NewMetadataError(e error, metadata ...Metadata) error {

	mde, ok := e.(*MetadataError)

	if ok {
		mde.Data = append(mde.Data, metadata...)
		mde.Stack = append(mde.Stack, FileAndFunc())
		return mde
	}

	return &MetadataError{
		inner: e,
		Data:  metadata,
		Stack: []string{FileAndFunc()},
	}
}

func FileAndFunc() string {
	pc, file, _, ok := runtime.Caller(2)
	if !ok {
		return "unkown"
	}

	funcName := runtime.FuncForPC(pc).Name()
	return fmt.Sprintf("%s:%s", filepath.Base(file), funcName)
}
