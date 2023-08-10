// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.30.0
// 	protoc        (unknown)
// source: gsloc/api/config/entries/v1/entry.proto

package entries

import (
	_ "github.com/envoyproxy/protoc-gen-validate/validate"
	v11 "github.com/orange-cloudfoundry/gsloc-go-sdk/gsloc/api/config/healthchecks/v1"
	v1 "github.com/orange-cloudfoundry/gsloc-go-sdk/gsloc/api/config/permission/v1"
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

type LBAlgo int32

const (
	LBAlgo_ROUND_ROBIN LBAlgo = 0
	LBAlgo_TOPOLOGY    LBAlgo = 1
	LBAlgo_RATIO       LBAlgo = 2
	LBAlgo_RANDOM      LBAlgo = 3
)

// Enum value maps for LBAlgo.
var (
	LBAlgo_name = map[int32]string{
		0: "ROUND_ROBIN",
		1: "TOPOLOGY",
		2: "RATIO",
		3: "RANDOM",
	}
	LBAlgo_value = map[string]int32{
		"ROUND_ROBIN": 0,
		"TOPOLOGY":    1,
		"RATIO":       2,
		"RANDOM":      3,
	}
)

func (x LBAlgo) Enum() *LBAlgo {
	p := new(LBAlgo)
	*p = x
	return p
}

func (x LBAlgo) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (LBAlgo) Descriptor() protoreflect.EnumDescriptor {
	return file_gsloc_api_config_entries_v1_entry_proto_enumTypes[0].Descriptor()
}

func (LBAlgo) Type() protoreflect.EnumType {
	return &file_gsloc_api_config_entries_v1_entry_proto_enumTypes[0]
}

func (x LBAlgo) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use LBAlgo.Descriptor instead.
func (LBAlgo) EnumDescriptor() ([]byte, []int) {
	return file_gsloc_api_config_entries_v1_entry_proto_rawDescGZIP(), []int{0}
}

type Entry struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Fqdn              string                  `protobuf:"bytes,1,opt,name=fqdn,proto3" json:"fqdn,omitempty"`
	LbAlgoPreferred   LBAlgo                  `protobuf:"varint,2,opt,name=lb_algo_preferred,json=lbAlgoPreferred,proto3,enum=gsloc.api.config.entries.v1.LBAlgo" json:"lb_algo_preferred,omitempty"`
	LbAlgoAlternate   LBAlgo                  `protobuf:"varint,3,opt,name=lb_algo_alternate,json=lbAlgoAlternate,proto3,enum=gsloc.api.config.entries.v1.LBAlgo" json:"lb_algo_alternate,omitempty"`
	LbAlgoFallback    LBAlgo                  `protobuf:"varint,4,opt,name=lb_algo_fallback,json=lbAlgoFallback,proto3,enum=gsloc.api.config.entries.v1.LBAlgo" json:"lb_algo_fallback,omitempty"`
	MaxAnswerReturned uint32                  `protobuf:"varint,5,opt,name=max_answer_returned,json=maxAnswerReturned,proto3" json:"max_answer_returned,omitempty"`
	MembersIpv4       []*Member               `protobuf:"bytes,6,rep,name=members_ipv4,json=membersIpv4,proto3" json:"members_ipv4,omitempty"`
	MembersIpv6       []*Member               `protobuf:"bytes,7,rep,name=members_ipv6,json=membersIpv6,proto3" json:"members_ipv6,omitempty"`
	Ttl               uint32                  `protobuf:"varint,8,opt,name=ttl,proto3" json:"ttl,omitempty"`
	Permissions       []*v1.ElementPermission `protobuf:"bytes,9,rep,name=permissions,proto3" json:"permissions,omitempty"`
	Tags              []string                `protobuf:"bytes,10,rep,name=tags,proto3" json:"tags,omitempty"`
}

func (x *Entry) Reset() {
	*x = Entry{}
	if protoimpl.UnsafeEnabled {
		mi := &file_gsloc_api_config_entries_v1_entry_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Entry) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Entry) ProtoMessage() {}

func (x *Entry) ProtoReflect() protoreflect.Message {
	mi := &file_gsloc_api_config_entries_v1_entry_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Entry.ProtoReflect.Descriptor instead.
func (*Entry) Descriptor() ([]byte, []int) {
	return file_gsloc_api_config_entries_v1_entry_proto_rawDescGZIP(), []int{0}
}

func (x *Entry) GetFqdn() string {
	if x != nil {
		return x.Fqdn
	}
	return ""
}

func (x *Entry) GetLbAlgoPreferred() LBAlgo {
	if x != nil {
		return x.LbAlgoPreferred
	}
	return LBAlgo_ROUND_ROBIN
}

func (x *Entry) GetLbAlgoAlternate() LBAlgo {
	if x != nil {
		return x.LbAlgoAlternate
	}
	return LBAlgo_ROUND_ROBIN
}

func (x *Entry) GetLbAlgoFallback() LBAlgo {
	if x != nil {
		return x.LbAlgoFallback
	}
	return LBAlgo_ROUND_ROBIN
}

func (x *Entry) GetMaxAnswerReturned() uint32 {
	if x != nil {
		return x.MaxAnswerReturned
	}
	return 0
}

func (x *Entry) GetMembersIpv4() []*Member {
	if x != nil {
		return x.MembersIpv4
	}
	return nil
}

func (x *Entry) GetMembersIpv6() []*Member {
	if x != nil {
		return x.MembersIpv6
	}
	return nil
}

func (x *Entry) GetTtl() uint32 {
	if x != nil {
		return x.Ttl
	}
	return 0
}

func (x *Entry) GetPermissions() []*v1.ElementPermission {
	if x != nil {
		return x.Permissions
	}
	return nil
}

func (x *Entry) GetTags() []string {
	if x != nil {
		return x.Tags
	}
	return nil
}

type Member struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Ip       string `protobuf:"bytes,1,opt,name=ip,proto3" json:"ip,omitempty"`
	Ratio    uint32 `protobuf:"varint,2,opt,name=ratio,proto3" json:"ratio,omitempty"`
	Dc       string `protobuf:"bytes,3,opt,name=dc,proto3" json:"dc,omitempty"`
	Disabled bool   `protobuf:"varint,4,opt,name=disabled,proto3" json:"disabled,omitempty"`
}

func (x *Member) Reset() {
	*x = Member{}
	if protoimpl.UnsafeEnabled {
		mi := &file_gsloc_api_config_entries_v1_entry_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Member) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Member) ProtoMessage() {}

func (x *Member) ProtoReflect() protoreflect.Message {
	mi := &file_gsloc_api_config_entries_v1_entry_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Member.ProtoReflect.Descriptor instead.
func (*Member) Descriptor() ([]byte, []int) {
	return file_gsloc_api_config_entries_v1_entry_proto_rawDescGZIP(), []int{1}
}

func (x *Member) GetIp() string {
	if x != nil {
		return x.Ip
	}
	return ""
}

func (x *Member) GetRatio() uint32 {
	if x != nil {
		return x.Ratio
	}
	return 0
}

func (x *Member) GetDc() string {
	if x != nil {
		return x.Dc
	}
	return ""
}

func (x *Member) GetDisabled() bool {
	if x != nil {
		return x.Disabled
	}
	return false
}

type SignedEntry struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Entry       *Entry           `protobuf:"bytes,1,opt,name=entry,proto3" json:"entry,omitempty"`
	Signature   string           `protobuf:"bytes,2,opt,name=signature,proto3" json:"signature,omitempty"`
	Healthcheck *v11.HealthCheck `protobuf:"bytes,3,opt,name=healthcheck,proto3" json:"healthcheck,omitempty"`
}

func (x *SignedEntry) Reset() {
	*x = SignedEntry{}
	if protoimpl.UnsafeEnabled {
		mi := &file_gsloc_api_config_entries_v1_entry_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SignedEntry) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SignedEntry) ProtoMessage() {}

func (x *SignedEntry) ProtoReflect() protoreflect.Message {
	mi := &file_gsloc_api_config_entries_v1_entry_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SignedEntry.ProtoReflect.Descriptor instead.
func (*SignedEntry) Descriptor() ([]byte, []int) {
	return file_gsloc_api_config_entries_v1_entry_proto_rawDescGZIP(), []int{2}
}

func (x *SignedEntry) GetEntry() *Entry {
	if x != nil {
		return x.Entry
	}
	return nil
}

func (x *SignedEntry) GetSignature() string {
	if x != nil {
		return x.Signature
	}
	return ""
}

func (x *SignedEntry) GetHealthcheck() *v11.HealthCheck {
	if x != nil {
		return x.Healthcheck
	}
	return nil
}

var File_gsloc_api_config_entries_v1_entry_proto protoreflect.FileDescriptor

var file_gsloc_api_config_entries_v1_entry_proto_rawDesc = []byte{
	0x0a, 0x27, 0x67, 0x73, 0x6c, 0x6f, 0x63, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x63, 0x6f, 0x6e, 0x66,
	0x69, 0x67, 0x2f, 0x65, 0x6e, 0x74, 0x72, 0x69, 0x65, 0x73, 0x2f, 0x76, 0x31, 0x2f, 0x65, 0x6e,
	0x74, 0x72, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x1b, 0x67, 0x73, 0x6c, 0x6f, 0x63,
	0x2e, 0x61, 0x70, 0x69, 0x2e, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2e, 0x65, 0x6e, 0x74, 0x72,
	0x69, 0x65, 0x73, 0x2e, 0x76, 0x31, 0x1a, 0x17, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65,
	0x2f, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a,
	0x2f, 0x67, 0x73, 0x6c, 0x6f, 0x63, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x63, 0x6f, 0x6e, 0x66, 0x69,
	0x67, 0x2f, 0x70, 0x65, 0x72, 0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x2f, 0x76, 0x31, 0x2f,
	0x70, 0x65, 0x72, 0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x1a, 0x32, 0x67, 0x73, 0x6c, 0x6f, 0x63, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x63, 0x6f, 0x6e, 0x66,
	0x69, 0x67, 0x2f, 0x68, 0x65, 0x61, 0x6c, 0x74, 0x68, 0x63, 0x68, 0x65, 0x63, 0x6b, 0x73, 0x2f,
	0x76, 0x31, 0x2f, 0x68, 0x65, 0x61, 0x6c, 0x74, 0x68, 0x63, 0x68, 0x65, 0x63, 0x6b, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x22, 0xd9, 0x04, 0x0a, 0x05, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x1b,
	0x0a, 0x04, 0x66, 0x71, 0x64, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x42, 0x07, 0xfa, 0x42,
	0x04, 0x72, 0x02, 0x10, 0x01, 0x52, 0x04, 0x66, 0x71, 0x64, 0x6e, 0x12, 0x4f, 0x0a, 0x11, 0x6c,
	0x62, 0x5f, 0x61, 0x6c, 0x67, 0x6f, 0x5f, 0x70, 0x72, 0x65, 0x66, 0x65, 0x72, 0x72, 0x65, 0x64,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x23, 0x2e, 0x67, 0x73, 0x6c, 0x6f, 0x63, 0x2e, 0x61,
	0x70, 0x69, 0x2e, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2e, 0x65, 0x6e, 0x74, 0x72, 0x69, 0x65,
	0x73, 0x2e, 0x76, 0x31, 0x2e, 0x4c, 0x42, 0x41, 0x6c, 0x67, 0x6f, 0x52, 0x0f, 0x6c, 0x62, 0x41,
	0x6c, 0x67, 0x6f, 0x50, 0x72, 0x65, 0x66, 0x65, 0x72, 0x72, 0x65, 0x64, 0x12, 0x4f, 0x0a, 0x11,
	0x6c, 0x62, 0x5f, 0x61, 0x6c, 0x67, 0x6f, 0x5f, 0x61, 0x6c, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x74,
	0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x23, 0x2e, 0x67, 0x73, 0x6c, 0x6f, 0x63, 0x2e,
	0x61, 0x70, 0x69, 0x2e, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2e, 0x65, 0x6e, 0x74, 0x72, 0x69,
	0x65, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x4c, 0x42, 0x41, 0x6c, 0x67, 0x6f, 0x52, 0x0f, 0x6c, 0x62,
	0x41, 0x6c, 0x67, 0x6f, 0x41, 0x6c, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x74, 0x65, 0x12, 0x4d, 0x0a,
	0x10, 0x6c, 0x62, 0x5f, 0x61, 0x6c, 0x67, 0x6f, 0x5f, 0x66, 0x61, 0x6c, 0x6c, 0x62, 0x61, 0x63,
	0x6b, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x23, 0x2e, 0x67, 0x73, 0x6c, 0x6f, 0x63, 0x2e,
	0x61, 0x70, 0x69, 0x2e, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2e, 0x65, 0x6e, 0x74, 0x72, 0x69,
	0x65, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x4c, 0x42, 0x41, 0x6c, 0x67, 0x6f, 0x52, 0x0e, 0x6c, 0x62,
	0x41, 0x6c, 0x67, 0x6f, 0x46, 0x61, 0x6c, 0x6c, 0x62, 0x61, 0x63, 0x6b, 0x12, 0x37, 0x0a, 0x13,
	0x6d, 0x61, 0x78, 0x5f, 0x61, 0x6e, 0x73, 0x77, 0x65, 0x72, 0x5f, 0x72, 0x65, 0x74, 0x75, 0x72,
	0x6e, 0x65, 0x64, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0d, 0x42, 0x07, 0xfa, 0x42, 0x04, 0x2a, 0x02,
	0x10, 0x64, 0x52, 0x11, 0x6d, 0x61, 0x78, 0x41, 0x6e, 0x73, 0x77, 0x65, 0x72, 0x52, 0x65, 0x74,
	0x75, 0x72, 0x6e, 0x65, 0x64, 0x12, 0x46, 0x0a, 0x0c, 0x6d, 0x65, 0x6d, 0x62, 0x65, 0x72, 0x73,
	0x5f, 0x69, 0x70, 0x76, 0x34, 0x18, 0x06, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x23, 0x2e, 0x67, 0x73,
	0x6c, 0x6f, 0x63, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2e, 0x65,
	0x6e, 0x74, 0x72, 0x69, 0x65, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x4d, 0x65, 0x6d, 0x62, 0x65, 0x72,
	0x52, 0x0b, 0x6d, 0x65, 0x6d, 0x62, 0x65, 0x72, 0x73, 0x49, 0x70, 0x76, 0x34, 0x12, 0x46, 0x0a,
	0x0c, 0x6d, 0x65, 0x6d, 0x62, 0x65, 0x72, 0x73, 0x5f, 0x69, 0x70, 0x76, 0x36, 0x18, 0x07, 0x20,
	0x03, 0x28, 0x0b, 0x32, 0x23, 0x2e, 0x67, 0x73, 0x6c, 0x6f, 0x63, 0x2e, 0x61, 0x70, 0x69, 0x2e,
	0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2e, 0x65, 0x6e, 0x74, 0x72, 0x69, 0x65, 0x73, 0x2e, 0x76,
	0x31, 0x2e, 0x4d, 0x65, 0x6d, 0x62, 0x65, 0x72, 0x52, 0x0b, 0x6d, 0x65, 0x6d, 0x62, 0x65, 0x72,
	0x73, 0x49, 0x70, 0x76, 0x36, 0x12, 0x10, 0x0a, 0x03, 0x74, 0x74, 0x6c, 0x18, 0x08, 0x20, 0x01,
	0x28, 0x0d, 0x52, 0x03, 0x74, 0x74, 0x6c, 0x12, 0x53, 0x0a, 0x0b, 0x70, 0x65, 0x72, 0x6d, 0x69,
	0x73, 0x73, 0x69, 0x6f, 0x6e, 0x73, 0x18, 0x09, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x31, 0x2e, 0x67,
	0x73, 0x6c, 0x6f, 0x63, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2e,
	0x70, 0x65, 0x72, 0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x2e, 0x76, 0x31, 0x2e, 0x45, 0x6c,
	0x65, 0x6d, 0x65, 0x6e, 0x74, 0x50, 0x65, 0x72, 0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x52,
	0x0b, 0x70, 0x65, 0x72, 0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x73, 0x12, 0x12, 0x0a, 0x04,
	0x74, 0x61, 0x67, 0x73, 0x18, 0x0a, 0x20, 0x03, 0x28, 0x09, 0x52, 0x04, 0x74, 0x61, 0x67, 0x73,
	0x22, 0x6c, 0x0a, 0x06, 0x4d, 0x65, 0x6d, 0x62, 0x65, 0x72, 0x12, 0x17, 0x0a, 0x02, 0x69, 0x70,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x42, 0x07, 0xfa, 0x42, 0x04, 0x72, 0x02, 0x70, 0x01, 0x52,
	0x02, 0x69, 0x70, 0x12, 0x14, 0x0a, 0x05, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x0d, 0x52, 0x05, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x12, 0x17, 0x0a, 0x02, 0x64, 0x63, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x09, 0x42, 0x07, 0xfa, 0x42, 0x04, 0x72, 0x02, 0x10, 0x01, 0x52, 0x02,
	0x64, 0x63, 0x12, 0x1a, 0x0a, 0x08, 0x64, 0x69, 0x73, 0x61, 0x62, 0x6c, 0x65, 0x64, 0x18, 0x04,
	0x20, 0x01, 0x28, 0x08, 0x52, 0x08, 0x64, 0x69, 0x73, 0x61, 0x62, 0x6c, 0x65, 0x64, 0x22, 0xd3,
	0x01, 0x0a, 0x0b, 0x53, 0x69, 0x67, 0x6e, 0x65, 0x64, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x42,
	0x0a, 0x05, 0x65, 0x6e, 0x74, 0x72, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x22, 0x2e,
	0x67, 0x73, 0x6c, 0x6f, 0x63, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67,
	0x2e, 0x65, 0x6e, 0x74, 0x72, 0x69, 0x65, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x45, 0x6e, 0x74, 0x72,
	0x79, 0x42, 0x08, 0xfa, 0x42, 0x05, 0x8a, 0x01, 0x02, 0x10, 0x01, 0x52, 0x05, 0x65, 0x6e, 0x74,
	0x72, 0x79, 0x12, 0x25, 0x0a, 0x09, 0x73, 0x69, 0x67, 0x6e, 0x61, 0x74, 0x75, 0x72, 0x65, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x09, 0x42, 0x07, 0xfa, 0x42, 0x04, 0x72, 0x02, 0x10, 0x01, 0x52, 0x09,
	0x73, 0x69, 0x67, 0x6e, 0x61, 0x74, 0x75, 0x72, 0x65, 0x12, 0x59, 0x0a, 0x0b, 0x68, 0x65, 0x61,
	0x6c, 0x74, 0x68, 0x63, 0x68, 0x65, 0x63, 0x6b, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x2d,
	0x2e, 0x67, 0x73, 0x6c, 0x6f, 0x63, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x63, 0x6f, 0x6e, 0x66, 0x69,
	0x67, 0x2e, 0x68, 0x65, 0x61, 0x6c, 0x74, 0x68, 0x63, 0x68, 0x65, 0x63, 0x6b, 0x73, 0x2e, 0x76,
	0x31, 0x2e, 0x48, 0x65, 0x61, 0x6c, 0x74, 0x68, 0x43, 0x68, 0x65, 0x63, 0x6b, 0x42, 0x08, 0xfa,
	0x42, 0x05, 0x8a, 0x01, 0x02, 0x10, 0x01, 0x52, 0x0b, 0x68, 0x65, 0x61, 0x6c, 0x74, 0x68, 0x63,
	0x68, 0x65, 0x63, 0x6b, 0x2a, 0x3e, 0x0a, 0x06, 0x4c, 0x42, 0x41, 0x6c, 0x67, 0x6f, 0x12, 0x0f,
	0x0a, 0x0b, 0x52, 0x4f, 0x55, 0x4e, 0x44, 0x5f, 0x52, 0x4f, 0x42, 0x49, 0x4e, 0x10, 0x00, 0x12,
	0x0c, 0x0a, 0x08, 0x54, 0x4f, 0x50, 0x4f, 0x4c, 0x4f, 0x47, 0x59, 0x10, 0x01, 0x12, 0x09, 0x0a,
	0x05, 0x52, 0x41, 0x54, 0x49, 0x4f, 0x10, 0x02, 0x12, 0x0a, 0x0a, 0x06, 0x52, 0x41, 0x4e, 0x44,
	0x4f, 0x4d, 0x10, 0x03, 0x42, 0x51, 0x5a, 0x4f, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63,
	0x6f, 0x6d, 0x2f, 0x6f, 0x72, 0x61, 0x6e, 0x67, 0x65, 0x2d, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x66,
	0x6f, 0x75, 0x6e, 0x64, 0x72, 0x79, 0x2f, 0x67, 0x73, 0x6c, 0x6f, 0x63, 0x2d, 0x67, 0x6f, 0x2d,
	0x73, 0x64, 0x6b, 0x2f, 0x67, 0x73, 0x6c, 0x6f, 0x63, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x63, 0x6f,
	0x6e, 0x66, 0x69, 0x67, 0x2f, 0x65, 0x6e, 0x74, 0x72, 0x69, 0x65, 0x73, 0x2f, 0x76, 0x31, 0x3b,
	0x65, 0x6e, 0x74, 0x72, 0x69, 0x65, 0x73, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_gsloc_api_config_entries_v1_entry_proto_rawDescOnce sync.Once
	file_gsloc_api_config_entries_v1_entry_proto_rawDescData = file_gsloc_api_config_entries_v1_entry_proto_rawDesc
)

func file_gsloc_api_config_entries_v1_entry_proto_rawDescGZIP() []byte {
	file_gsloc_api_config_entries_v1_entry_proto_rawDescOnce.Do(func() {
		file_gsloc_api_config_entries_v1_entry_proto_rawDescData = protoimpl.X.CompressGZIP(file_gsloc_api_config_entries_v1_entry_proto_rawDescData)
	})
	return file_gsloc_api_config_entries_v1_entry_proto_rawDescData
}

var file_gsloc_api_config_entries_v1_entry_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_gsloc_api_config_entries_v1_entry_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_gsloc_api_config_entries_v1_entry_proto_goTypes = []interface{}{
	(LBAlgo)(0),                  // 0: gsloc.api.config.entries.v1.LBAlgo
	(*Entry)(nil),                // 1: gsloc.api.config.entries.v1.Entry
	(*Member)(nil),               // 2: gsloc.api.config.entries.v1.Member
	(*SignedEntry)(nil),          // 3: gsloc.api.config.entries.v1.SignedEntry
	(*v1.ElementPermission)(nil), // 4: gsloc.api.config.permission.v1.ElementPermission
	(*v11.HealthCheck)(nil),      // 5: gsloc.api.config.healthchecks.v1.HealthCheck
}
var file_gsloc_api_config_entries_v1_entry_proto_depIdxs = []int32{
	0, // 0: gsloc.api.config.entries.v1.Entry.lb_algo_preferred:type_name -> gsloc.api.config.entries.v1.LBAlgo
	0, // 1: gsloc.api.config.entries.v1.Entry.lb_algo_alternate:type_name -> gsloc.api.config.entries.v1.LBAlgo
	0, // 2: gsloc.api.config.entries.v1.Entry.lb_algo_fallback:type_name -> gsloc.api.config.entries.v1.LBAlgo
	2, // 3: gsloc.api.config.entries.v1.Entry.members_ipv4:type_name -> gsloc.api.config.entries.v1.Member
	2, // 4: gsloc.api.config.entries.v1.Entry.members_ipv6:type_name -> gsloc.api.config.entries.v1.Member
	4, // 5: gsloc.api.config.entries.v1.Entry.permissions:type_name -> gsloc.api.config.permission.v1.ElementPermission
	1, // 6: gsloc.api.config.entries.v1.SignedEntry.entry:type_name -> gsloc.api.config.entries.v1.Entry
	5, // 7: gsloc.api.config.entries.v1.SignedEntry.healthcheck:type_name -> gsloc.api.config.healthchecks.v1.HealthCheck
	8, // [8:8] is the sub-list for method output_type
	8, // [8:8] is the sub-list for method input_type
	8, // [8:8] is the sub-list for extension type_name
	8, // [8:8] is the sub-list for extension extendee
	0, // [0:8] is the sub-list for field type_name
}

func init() { file_gsloc_api_config_entries_v1_entry_proto_init() }
func file_gsloc_api_config_entries_v1_entry_proto_init() {
	if File_gsloc_api_config_entries_v1_entry_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_gsloc_api_config_entries_v1_entry_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Entry); i {
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
		file_gsloc_api_config_entries_v1_entry_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Member); i {
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
		file_gsloc_api_config_entries_v1_entry_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SignedEntry); i {
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
			RawDescriptor: file_gsloc_api_config_entries_v1_entry_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_gsloc_api_config_entries_v1_entry_proto_goTypes,
		DependencyIndexes: file_gsloc_api_config_entries_v1_entry_proto_depIdxs,
		EnumInfos:         file_gsloc_api_config_entries_v1_entry_proto_enumTypes,
		MessageInfos:      file_gsloc_api_config_entries_v1_entry_proto_msgTypes,
	}.Build()
	File_gsloc_api_config_entries_v1_entry_proto = out.File
	file_gsloc_api_config_entries_v1_entry_proto_rawDesc = nil
	file_gsloc_api_config_entries_v1_entry_proto_goTypes = nil
	file_gsloc_api_config_entries_v1_entry_proto_depIdxs = nil
}
