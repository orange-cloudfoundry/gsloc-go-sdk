package errors

import (
	"fmt"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type ErrNotFound struct {
	elemName string
	elemKey  string
}

func NewErrNotFound(elemName, elemKey string) *ErrNotFound {
	return &ErrNotFound{
		elemName: elemName,
		elemKey:  elemKey,
	}
}

func (e ErrNotFound) Error() string {
	return fmt.Sprintf("%s with key %s not found", e.elemName, e.elemKey)
}

func IsErrNotFound(err error) bool {
	_, ok := err.(*ErrNotFound)
	if ok {
		return true
	}
	sErr, ok := status.FromError(err)
	if !ok {
		return false
	}
	return sErr.Code() == codes.NotFound
}
