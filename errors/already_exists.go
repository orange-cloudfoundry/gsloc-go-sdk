package errors

import (
	"fmt"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type ErrAlreadyExists struct {
	elemName string
	elemKey  string
}

func NewErrAlreadyExists(elemName, elemKey string) *ErrNotFound {
	return &ErrNotFound{
		elemName: elemName,
		elemKey:  elemKey,
	}
}

func (e ErrAlreadyExists) Error() string {
	return fmt.Sprintf("%s with key %s not found", e.elemName, e.elemKey)
}

func IsErrAlreadyExists(err error) bool {
	_, ok := err.(*ErrAlreadyExists)
	if ok {
		return true
	}
	sErr, ok := status.FromError(err)
	if !ok {
		return false
	}
	return sErr.Code() == codes.AlreadyExists
}
