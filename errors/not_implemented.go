package errors

import (
	"fmt"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type ErrNotImplemented struct {
	funcName string
}

func NewErrNotImplemented(funcName string) *ErrNotImplemented {
	return &ErrNotImplemented{
		funcName: funcName,
	}
}

func (e ErrNotImplemented) Error() string {
	return fmt.Sprintf("function %s not implemented", e.funcName)
}

func IsErrNotImplemented(err error) bool {
	_, ok := err.(*ErrNotImplemented)
	if ok {
		return true
	}
	sErr, ok := status.FromError(err)
	if !ok {
		return false
	}
	return sErr.Code() == codes.Unimplemented
}
