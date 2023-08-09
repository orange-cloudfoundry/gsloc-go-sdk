package errors

import (
	"fmt"
	"strings"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type ErrNotSupportedConnector struct {
	funcName      string
	connectorName string
}

func NewErrNotSupported(funcName string, connectorName string) *ErrNotSupportedConnector {
	return &ErrNotSupportedConnector{
		funcName:      strings.ToLower(funcName),
		connectorName: connectorName,
	}
}

func (e ErrNotSupportedConnector) Error() string {
	return fmt.Sprintf("functionnality %s not supported by connector %s", e.funcName, e.connectorName)
}

func IsErrNotSupported(err error) bool {
	_, ok := err.(*ErrNotSupportedConnector)
	if ok {
		return true
	}
	sErr, ok := status.FromError(err)
	if !ok {
		return false
	}
	return sErr.Code() == codes.Unimplemented
}

func IsErrNotSupportedFunc(err error, funcName string) bool {
	errNotSupp, ok := err.(*ErrNotSupportedConnector)
	if !ok {
		return false
	}
	return errNotSupp.funcName == strings.ToLower(funcName)
}
