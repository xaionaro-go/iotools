package iotools

type ErrReadFunctionNotDefined struct{}

func (err *ErrReadFunctionNotDefined) Error() string {
	return "function Read() is not defined"
}

type ErrWriteFunctionNotDefined struct{}

func (err *ErrWriteFunctionNotDefined) Error() string {
	return "function Write() is not defined"
}

type ErrCloseFunctionNotDefined struct{}

func (err *ErrCloseFunctionNotDefined) Error() string {
	return "function Close() is not defined"
}
