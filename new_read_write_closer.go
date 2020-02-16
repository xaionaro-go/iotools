package iotools

import (
	"io"

	"github.com/xaionaro-go/errors"
)

type CustomReadWriteCloser struct {
	ReadFunction  func([]byte) (int, error)
	WriteFunction func([]byte) (int, error)
	CloseFunction func() error
}

var _ io.ReadWriteCloser = &CustomReadWriteCloser{}

func NewReadWriteCloser(
	readFunction func([]byte) (int, error),
	writeFunction func([]byte) (int, error),
	closeFunction func() error,
) *CustomReadWriteCloser {
	return &CustomReadWriteCloser{
		ReadFunction:  readFunction,
		WriteFunction: writeFunction,
		CloseFunction: closeFunction,
	}
}

func (rwc *CustomReadWriteCloser) Read(b []byte) (n int, err error) {
	defer func() { err = errors.Wrap(err) }()
	if rwc.ReadFunction == nil {
		return 0, &ErrReadFunctionNotDefined{}
	}
	return rwc.ReadFunction(b)
}

func (rwc *CustomReadWriteCloser) Write(b []byte) (n int, err error) {
	defer func() { err = errors.Wrap(err) }()
	if rwc.WriteFunction == nil {
		return 0, &ErrWriteFunctionNotDefined{}
	}
	return rwc.WriteFunction(b)
}

func (rwc *CustomReadWriteCloser) Close() (err error) {
	defer func() { err = errors.Wrap(err) }()
	if rwc.CloseFunction == nil {
		return &ErrCloseFunctionNotDefined{}
	}
	return rwc.CloseFunction()
}

func NewWriter(writeFunction func([]byte) (int, error)) *CustomReadWriteCloser {
	return &CustomReadWriteCloser{
		WriteFunction: writeFunction,
	}
}

func NewWriteCloser(
	writeFunction func([]byte) (int, error),
	closeFunction func() error,
) *CustomReadWriteCloser {
	return &CustomReadWriteCloser{
		WriteFunction: writeFunction,
		CloseFunction: closeFunction,
	}
}
