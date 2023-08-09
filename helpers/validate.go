package helpers

import "google.golang.org/protobuf/proto"

type validable interface {
	ValidateAll() error
}

type postValidable interface {
	PostValidate() error
}

func Validate(message proto.Message) error {
	if message == nil {
		return nil
	}
	if v, ok := message.(validable); ok {
		err := v.ValidateAll()
		if err != nil {
			return err
		}
	}
	if v, ok := message.(postValidable); ok {
		err := v.PostValidate()
		if err != nil {
			return err
		}
	}
	return nil
}
