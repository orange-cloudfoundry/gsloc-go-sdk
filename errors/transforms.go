package errors

import (
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func ToStatus(err error) error {
	if err == nil {
		return nil
	}
	if IsErrNotFound(err) {
		return status.Errorf(codes.NotFound, err.Error())
	}
	if IsErrNotImplemented(err) {
		return status.Errorf(codes.Unimplemented, err.Error())
	}
	if IsErrNotSupported(err) {
		return status.Errorf(codes.Unimplemented, err.Error())
	}
	if IsErrAlreadyExists(err) {
		return status.Errorf(codes.AlreadyExists, err.Error())
	}
	return status.Errorf(codes.Internal, err.Error())
}
