// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.30.0
// 	protoc        (unknown)
// source: gsloc/api/config/core/v1/base.proto

package core

import (
	_ "github.com/envoyproxy/protoc-gen-validate/validate"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// Header name/value pair.
type HeaderValue struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Header name.
	Key string `protobuf:"bytes,1,opt,name=key,proto3" json:"key,omitempty"`
	// Header value.
	Value string `protobuf:"bytes,2,opt,name=value,proto3" json:"value,omitempty"`
}

func (x *HeaderValue) Reset() {
	*x = HeaderValue{}
	if protoimpl.UnsafeEnabled {
		mi := &file_gsloc_api_config_core_v1_base_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *HeaderValue) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*HeaderValue) ProtoMessage() {}

func (x *HeaderValue) ProtoReflect() protoreflect.Message {
	mi := &file_gsloc_api_config_core_v1_base_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use HeaderValue.ProtoReflect.Descriptor instead.
func (*HeaderValue) Descriptor() ([]byte, []int) {
	return file_gsloc_api_config_core_v1_base_proto_rawDescGZIP(), []int{0}
}

func (x *HeaderValue) GetKey() string {
	if x != nil {
		return x.Key
	}
	return ""
}

func (x *HeaderValue) GetValue() string {
	if x != nil {
		return x.Value
	}
	return ""
}

// Header name/value pair plus option to control append behavior.
type HeaderValueOption struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Header name/value pair that this option applies to.
	Header *HeaderValue `protobuf:"bytes,1,opt,name=header,proto3" json:"header,omitempty"`
	// Should the value be appended? If true (default), the value is appended to
	// existing values. Otherwise it replaces any existing values.
	Append bool `protobuf:"varint,2,opt,name=append,proto3" json:"append,omitempty"`
}

func (x *HeaderValueOption) Reset() {
	*x = HeaderValueOption{}
	if protoimpl.UnsafeEnabled {
		mi := &file_gsloc_api_config_core_v1_base_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *HeaderValueOption) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*HeaderValueOption) ProtoMessage() {}

func (x *HeaderValueOption) ProtoReflect() protoreflect.Message {
	mi := &file_gsloc_api_config_core_v1_base_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use HeaderValueOption.ProtoReflect.Descriptor instead.
func (*HeaderValueOption) Descriptor() ([]byte, []int) {
	return file_gsloc_api_config_core_v1_base_proto_rawDescGZIP(), []int{1}
}

func (x *HeaderValueOption) GetHeader() *HeaderValue {
	if x != nil {
		return x.Header
	}
	return nil
}

func (x *HeaderValueOption) GetAppend() bool {
	if x != nil {
		return x.Append
	}
	return false
}

var File_gsloc_api_config_core_v1_base_proto protoreflect.FileDescriptor

var file_gsloc_api_config_core_v1_base_proto_rawDesc = []byte{
	0x0a, 0x23, 0x67, 0x73, 0x6c, 0x6f, 0x63, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x63, 0x6f, 0x6e, 0x66,
	0x69, 0x67, 0x2f, 0x63, 0x6f, 0x72, 0x65, 0x2f, 0x76, 0x31, 0x2f, 0x62, 0x61, 0x73, 0x65, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x18, 0x67, 0x73, 0x6c, 0x6f, 0x63, 0x2e, 0x61, 0x70, 0x69,
	0x2e, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2e, 0x63, 0x6f, 0x72, 0x65, 0x2e, 0x76, 0x31, 0x1a,
	0x17, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x2f, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61,
	0x74, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x59, 0x0a, 0x0b, 0x48, 0x65, 0x61, 0x64,
	0x65, 0x72, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x12, 0x23, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x42, 0x11, 0xfa, 0x42, 0x0e, 0x72, 0x0c, 0x10, 0x01, 0x28, 0x80, 0x80,
	0x01, 0xc8, 0x01, 0x00, 0xc0, 0x01, 0x01, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x25, 0x0a, 0x05,
	0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x42, 0x0f, 0xfa, 0x42, 0x0c,
	0x72, 0x0a, 0x28, 0x80, 0x80, 0x01, 0xc8, 0x01, 0x00, 0xc0, 0x01, 0x02, 0x52, 0x05, 0x76, 0x61,
	0x6c, 0x75, 0x65, 0x22, 0x74, 0x0a, 0x11, 0x48, 0x65, 0x61, 0x64, 0x65, 0x72, 0x56, 0x61, 0x6c,
	0x75, 0x65, 0x4f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x47, 0x0a, 0x06, 0x68, 0x65, 0x61, 0x64,
	0x65, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x25, 0x2e, 0x67, 0x73, 0x6c, 0x6f, 0x63,
	0x2e, 0x61, 0x70, 0x69, 0x2e, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2e, 0x63, 0x6f, 0x72, 0x65,
	0x2e, 0x76, 0x31, 0x2e, 0x48, 0x65, 0x61, 0x64, 0x65, 0x72, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x42,
	0x08, 0xfa, 0x42, 0x05, 0x8a, 0x01, 0x02, 0x10, 0x01, 0x52, 0x06, 0x68, 0x65, 0x61, 0x64, 0x65,
	0x72, 0x12, 0x16, 0x0a, 0x06, 0x61, 0x70, 0x70, 0x65, 0x6e, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x08, 0x52, 0x06, 0x61, 0x70, 0x70, 0x65, 0x6e, 0x64, 0x42, 0x4b, 0x5a, 0x49, 0x67, 0x69, 0x74,
	0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x6f, 0x72, 0x61, 0x6e, 0x67, 0x65, 0x2d, 0x63,
	0x6c, 0x6f, 0x75, 0x64, 0x66, 0x6f, 0x75, 0x6e, 0x64, 0x72, 0x79, 0x2f, 0x67, 0x73, 0x6c, 0x6f,
	0x63, 0x2d, 0x67, 0x6f, 0x2d, 0x73, 0x64, 0x6b, 0x2f, 0x67, 0x73, 0x6c, 0x6f, 0x63, 0x2f, 0x61,
	0x70, 0x69, 0x2f, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2f, 0x63, 0x6f, 0x72, 0x65, 0x2f, 0x76,
	0x31, 0x3b, 0x63, 0x6f, 0x72, 0x65, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_gsloc_api_config_core_v1_base_proto_rawDescOnce sync.Once
	file_gsloc_api_config_core_v1_base_proto_rawDescData = file_gsloc_api_config_core_v1_base_proto_rawDesc
)

func file_gsloc_api_config_core_v1_base_proto_rawDescGZIP() []byte {
	file_gsloc_api_config_core_v1_base_proto_rawDescOnce.Do(func() {
		file_gsloc_api_config_core_v1_base_proto_rawDescData = protoimpl.X.CompressGZIP(file_gsloc_api_config_core_v1_base_proto_rawDescData)
	})
	return file_gsloc_api_config_core_v1_base_proto_rawDescData
}

var file_gsloc_api_config_core_v1_base_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_gsloc_api_config_core_v1_base_proto_goTypes = []interface{}{
	(*HeaderValue)(nil),       // 0: gsloc.api.config.core.v1.HeaderValue
	(*HeaderValueOption)(nil), // 1: gsloc.api.config.core.v1.HeaderValueOption
}
var file_gsloc_api_config_core_v1_base_proto_depIdxs = []int32{
	0, // 0: gsloc.api.config.core.v1.HeaderValueOption.header:type_name -> gsloc.api.config.core.v1.HeaderValue
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_gsloc_api_config_core_v1_base_proto_init() }
func file_gsloc_api_config_core_v1_base_proto_init() {
	if File_gsloc_api_config_core_v1_base_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_gsloc_api_config_core_v1_base_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*HeaderValue); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_gsloc_api_config_core_v1_base_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*HeaderValueOption); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_gsloc_api_config_core_v1_base_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_gsloc_api_config_core_v1_base_proto_goTypes,
		DependencyIndexes: file_gsloc_api_config_core_v1_base_proto_depIdxs,
		MessageInfos:      file_gsloc_api_config_core_v1_base_proto_msgTypes,
	}.Build()
	File_gsloc_api_config_core_v1_base_proto = out.File
	file_gsloc_api_config_core_v1_base_proto_rawDesc = nil
	file_gsloc_api_config_core_v1_base_proto_goTypes = nil
	file_gsloc_api_config_core_v1_base_proto_depIdxs = nil
}