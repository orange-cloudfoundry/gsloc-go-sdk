package helpers

import (
	"bytes"
	"encoding/json"
	"fmt"
	"os"
	"reflect"

	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/types/known/anypb"
	"gopkg.in/yaml.v2"
)

type Validator interface {
	InitData() error
	Validate() error
}

func IsValidator(v interface{}) bool {
	_, ok := v.(Validator)
	return ok
}

func RecursiveValidate(v interface{}) error {
	if v == nil {
		return nil
	}

	vValue := reflect.ValueOf(v)
	vType := vValue.Type()
	vValueNoPtr := vValue
	isPtr := vType.Kind() == reflect.Ptr
	if isPtr {
		vType = vType.Elem()
		vValueNoPtr = vValue.Elem()
	}
	if vType.Kind() != reflect.Struct && vType.Kind() != reflect.Slice {
		return nil
	}
	kind := vType.Kind()
	switch kind {
	case reflect.Slice:
		for i := 0; i < vValueNoPtr.Len(); i++ {
			vElem := vValueNoPtr.Index(i)
			err := RecursiveValidate(vElem.Interface())
			if err != nil {
				return err
			}
		}
	case reflect.Struct:

		// do not inspect any type
		if vValue.Type() == reflect.TypeOf((*anypb.Any)(nil)) {
			break
		}
		if vValue.IsZero() {
			break
		}
		for i := 0; i < vValueNoPtr.NumField(); i++ {
			vField := vValueNoPtr.Field(i)
			if !vField.CanAddr() || !vField.CanSet() {
				continue
			}

			data := vField.Interface()
			err := RecursiveValidate(data)
			if err != nil {
				return err
			}
		}
	}
	if IsValidator(v) {
		initValid := v.(Validator)
		err := initValid.Validate()
		if err != nil {
			return fmt.Errorf("Could not validate data: %s", err.Error())
		}
	}
	return nil
}

func loadFileStructured(path string) ([]byte, error) {
	b, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}
	var raw interface{}
	err = yaml.Unmarshal(b, &raw)
	if err != nil {
		return nil, err
	}
	raw = cleanupMapValue(raw)
	return json.Marshal(raw)
}

func StoreProtoToFile(path string, v proto.Message) error {
	b, err := protojson.MarshalOptions{
		Multiline:       true,
		Indent:          "  ",
		AllowPartial:    true,
		UseProtoNames:   false,
		UseEnumNumbers:  false,
		EmitUnpopulated: true,
	}.Marshal(v)
	if err != nil {
		return err
	}
	return os.WriteFile(path, b, 0644)
}

func StoreListProtoToFile[T proto.Message](path string, messages []T) error {
	b, err := MarshalListProtoMessage[T](protojson.MarshalOptions{
		Multiline:       true,
		Indent:          "  ",
		AllowPartial:    true,
		UseProtoNames:   false,
		UseEnumNumbers:  false,
		EmitUnpopulated: true,
	}, messages)
	if err != nil {
		return err
	}
	return os.WriteFile(path, b, 0644)
}

func LoadProtoFromFile(path string, v proto.Message) error {
	b, err := loadFileStructured(path)
	if err != nil {
		if os.IsNotExist(err) {
			err := StoreProtoToFile(path, v)
			if err != nil {
				return err
			}
			return LoadProtoFromFile(path, v)
		}
		return err
	}
	err = protojson.UnmarshalOptions{
		AllowPartial:   true,
		DiscardUnknown: true,
		Resolver:       nil,
	}.Unmarshal(b, v)

	if err != nil {
		return err
	}
	err = RecursiveValidate(v)
	if err != nil {
		return err
	}
	return nil
}

func LoadListProtoFromFile[T proto.Message](path string) ([]T, error) {
	b, err := loadFileStructured(path)
	if err != nil {
		if os.IsNotExist(err) {
			err := StoreListProtoToFile[T](path, make([]T, 0))
			if err != nil {
				return nil, err
			}
			return LoadListProtoFromFile[T](path)
		}
		return []T{}, err
	}

	listRawMessages := make([]json.RawMessage, 0)
	err = json.Unmarshal(b, &listRawMessages)
	if err != nil {
		return []T{}, err
	}

	listMessages := make([]T, len(listRawMessages))
	for i, rawMessage := range listRawMessages {
		var message T
		message = reflect.New(reflect.TypeOf(message).Elem()).Interface().(T)
		err = protojson.UnmarshalOptions{
			AllowPartial:   true,
			DiscardUnknown: true,
			Resolver:       nil,
		}.Unmarshal(rawMessage, message)
		if err != nil {
			return []T{}, err
		}
		err = RecursiveValidate(message)
		if err != nil {
			return []T{}, err
		}
		listMessages[i] = message
	}
	return listMessages, nil
}

func MarshalListProtoMessage[T proto.Message](opts protojson.MarshalOptions, messages []T) ([]byte, error) {
	var err error
	elems := make([][]byte, len(messages))
	for i, m := range messages {
		elems[i], err = opts.Marshal(m)
		if err != nil {
			return nil, err
		}
	}
	joinedElems := bytes.Join(elems, []byte(","))
	final := make([]byte, len(joinedElems)+2)
	final[0] = '['
	copy(final[1:], joinedElems)
	final[len(final)-1] = ']'

	return final, nil
}

func cleanupInterfaceArray(in []interface{}) []interface{} {
	res := make([]interface{}, len(in))
	for i, v := range in {
		res[i] = cleanupMapValue(v)
	}
	return res
}

func cleanupInterfaceMap(in map[interface{}]interface{}) map[string]interface{} {
	res := make(map[string]interface{})
	for k, v := range in {
		res[fmt.Sprintf("%v", k)] = cleanupMapValue(v)
	}
	return res
}

func cleanupMapValue(v interface{}) interface{} {
	switch v := v.(type) {
	case []interface{}:
		return cleanupInterfaceArray(v)
	case map[interface{}]interface{}:
		return cleanupInterfaceMap(v)
	default:
		return v
	}
}
