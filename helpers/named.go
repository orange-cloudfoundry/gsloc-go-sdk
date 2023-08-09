package helpers

import "google.golang.org/protobuf/proto"

type Named interface {
	GetName() string
}

func GetName(m proto.Message) string {
	if m, ok := m.(Named); ok {
		return m.GetName()
	}
	return ""
}
