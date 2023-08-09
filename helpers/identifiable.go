package helpers

import "google.golang.org/protobuf/proto"

type Identifiable interface {
	GetIdentifier() string
	SetIdentifier(string)
}

const UrlPrefix = "type.googleapis.com/"

func GetIdentifier(m proto.Message) string {
	if m, ok := m.(Identifiable); ok {
		return m.GetIdentifier()
	}
	sig, err := MessageSignature(m)
	if err == nil {
		return sig
	}
	return string(m.ProtoReflect().Descriptor().FullName())
}

func SetIdentifier(m proto.Message, identifier string) {
	mId, ok := m.(Identifiable)
	if !ok {
		return
	}
	mId.SetIdentifier(identifier)
}

func TypeUrl[T proto.Message]() string {
	var m T
	return TypeUrlFromMessage(m)
}

func TypeUrlFromMessage(m proto.Message) string {
	return UrlPrefix + string(m.ProtoReflect().Descriptor().FullName())
}
