package test

import (
	proto "github.com/golang/protobuf/proto"
	"encoding/json"
	"github.com/samkreter/redact"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
)

const (
	_	= protoimpl.EnforceVersion(20 - protoimpl.MinVersion)

	_	= protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

const _ = proto.ProtoPackageIsVersion4

type TestStruct struct {
	state		protoimpl.MessageState
	sizeCache	protoimpl.SizeCache
	unknownFields	protoimpl.UnknownFields

	Secret		string	`protobuf:"bytes,1,opt,name=secret,proto3" json:"secret,omitempty"`
	NonSecret	string	`protobuf:"bytes,2,opt,name=nonSecret,proto3" json:"nonSecret,omitempty" redact:"nonsecret"`
	SecretPtr	*string	`protobuf:"bytes,3,opt,name=secretPtr,proto3,oneof" json:"secretPtr,omitempty"`
}

func (x *TestStruct) Reset() {
	*x = TestStruct{}
	if protoimpl.UnsafeEnabled {
		mi := &file_hcp_types_v1_test_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}
func (x *TestStruct) String() string {
	type EnumType interface {
		Descriptor() protoreflect.
			EnumDescriptor
		Number() protoreflect.EnumNumber
	}
	enumType, ok := interface{}(x).(EnumType)
	if ok {
		return protoimpl.X.EnumStringOf(enumType.Descriptor(), enumType.
			Number())
	}
	var copy TestStruct
	jsonBytes, err := json.Marshal(x)
	if err != nil {
		return ""
	}
	err = json.Unmarshal(jsonBytes, &copy,
	)
	if err != nil {
		return ""
	}
	if err :=
		redact.Redact(&copy); err != nil {
		return ""
	}
	jsonBytes, err = json.Marshal(copy)
	if err != nil {
		return ""
	}
	return string(jsonBytes,
	)
}

func (*TestStruct) ProtoMessage()	{}

func (x *TestStruct) ProtoReflect() protoreflect.Message {
	mi := &file_hcp_types_v1_test_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

func (*TestStruct) Descriptor() ([]byte, []int) {
	return file_hcp_types_v1_test_proto_rawDescGZIP(), []int{0}
}

func (x *TestStruct) GetSecret() string {
	if x != nil {
		return x.Secret
	}
	return ""
}

func (x *TestStruct) GetNonSecret() string {
	if x != nil {
		return x.NonSecret
	}
	return ""
}

func (x *TestStruct) GetSecretPtr() string {
	if x != nil && x.SecretPtr != nil {
		return *x.SecretPtr
	}
	return ""
}

type TestStructList struct {
	state		protoimpl.MessageState
	sizeCache	protoimpl.SizeCache
	unknownFields	protoimpl.UnknownFields

	Data	[]*TestStruct	`protobuf:"bytes,1,rep,name=data,proto3" json:"data,omitempty"`
}

func (x *TestStructList) Reset() {
	*x = TestStructList{}
	if protoimpl.UnsafeEnabled {
		mi := &file_hcp_types_v1_test_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}
func (x *TestStructList) String() string {
	type EnumType interface {
		Descriptor() protoreflect.
			EnumDescriptor
		Number() protoreflect.EnumNumber
	}
	enumType, ok := interface{}(x).(EnumType)
	if ok {
		return protoimpl.X.EnumStringOf(enumType.Descriptor(), enumType.
			Number())
	}
	var copy TestStructList
	jsonBytes, err := json.Marshal(x)
	if err != nil {
		return ""
	}
	err = json.Unmarshal(jsonBytes, &copy,
	)
	if err != nil {
		return ""
	}
	if err :=
		redact.Redact(&copy); err != nil {
		return ""
	}
	jsonBytes, err = json.Marshal(copy)
	if err != nil {
		return ""
	}
	return string(jsonBytes,
	)
}

func (*TestStructList) ProtoMessage()	{}

func (x *TestStructList) ProtoReflect() protoreflect.Message {
	mi := &file_hcp_types_v1_test_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

func (*TestStructList) Descriptor() ([]byte, []int) {
	return file_hcp_types_v1_test_proto_rawDescGZIP(), []int{1}
}

func (x *TestStructList) GetData() []*TestStruct {
	if x != nil {
		return x.Data
	}
	return nil
}

var File_hcp_types_v1_test_proto protoreflect.FileDescriptor

var file_hcp_types_v1_test_proto_rawDesc = []byte{
	0x0a, 0x17, 0x68, 0x63, 0x70, 0x2f, 0x74, 0x79, 0x70, 0x65, 0x73, 0x2f, 0x76, 0x31, 0x2f, 0x74,
	0x65, 0x73, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0c, 0x68, 0x63, 0x70, 0x2e, 0x74,
	0x79, 0x70, 0x65, 0x73, 0x2e, 0x76, 0x31, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61,
	0x6d, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x73, 0x0a, 0x0a, 0x54, 0x65, 0x73, 0x74,
	0x53, 0x74, 0x72, 0x75, 0x63, 0x74, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x65, 0x63, 0x72, 0x65, 0x74,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x73, 0x65, 0x63, 0x72, 0x65, 0x74, 0x12, 0x1c,
	0x0a, 0x09, 0x6e, 0x6f, 0x6e, 0x53, 0x65, 0x63, 0x72, 0x65, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x09, 0x6e, 0x6f, 0x6e, 0x53, 0x65, 0x63, 0x72, 0x65, 0x74, 0x12, 0x21, 0x0a, 0x09,
	0x73, 0x65, 0x63, 0x72, 0x65, 0x74, 0x50, 0x74, 0x72, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x48,
	0x00, 0x52, 0x09, 0x73, 0x65, 0x63, 0x72, 0x65, 0x74, 0x50, 0x74, 0x72, 0x88, 0x01, 0x01, 0x42,
	0x0c, 0x0a, 0x0a, 0x5f, 0x73, 0x65, 0x63, 0x72, 0x65, 0x74, 0x50, 0x74, 0x72, 0x22, 0x3e, 0x0a,
	0x0e, 0x54, 0x65, 0x73, 0x74, 0x53, 0x74, 0x72, 0x75, 0x63, 0x74, 0x4c, 0x69, 0x73, 0x74, 0x12,
	0x2c, 0x0a, 0x04, 0x64, 0x61, 0x74, 0x61, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x18, 0x2e,
	0x68, 0x63, 0x70, 0x2e, 0x74, 0x79, 0x70, 0x65, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x54, 0x65, 0x73,
	0x74, 0x53, 0x74, 0x72, 0x75, 0x63, 0x74, 0x52, 0x04, 0x64, 0x61, 0x74, 0x61, 0x42, 0x24, 0x5a,
	0x22, 0x67, 0x6f, 0x6d, 0x73, 0x2e, 0x69, 0x6f, 0x2f, 0x61, 0x6b, 0x73, 0x2f, 0x72, 0x70, 0x2f,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2f, 0x68, 0x63, 0x70, 0x2f, 0x74, 0x79, 0x70, 0x65, 0x73,
	0x2f, 0x76, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_hcp_types_v1_test_proto_rawDescOnce	sync.Once
	file_hcp_types_v1_test_proto_rawDescData	= file_hcp_types_v1_test_proto_rawDesc
)

func file_hcp_types_v1_test_proto_rawDescGZIP() []byte {
	file_hcp_types_v1_test_proto_rawDescOnce.Do(func() {
		file_hcp_types_v1_test_proto_rawDescData = protoimpl.X.CompressGZIP(file_hcp_types_v1_test_proto_rawDescData)
	})
	return file_hcp_types_v1_test_proto_rawDescData
}

var file_hcp_types_v1_test_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_hcp_types_v1_test_proto_goTypes = []interface{}{
	(*TestStruct)(nil),
	(*TestStructList)(nil),
}
var file_hcp_types_v1_test_proto_depIdxs = []int32{
	0,
	1,
	1,
	1,
	1,
	0,
}

func init()	{ file_hcp_types_v1_test_proto_init() }
func file_hcp_types_v1_test_proto_init() {
	if File_hcp_types_v1_test_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_hcp_types_v1_test_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TestStruct); i {
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
		file_hcp_types_v1_test_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TestStructList); i {
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
	file_hcp_types_v1_test_proto_msgTypes[0].OneofWrappers = []interface{}{}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath:	reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor:	file_hcp_types_v1_test_proto_rawDesc,
			NumEnums:	0,
			NumMessages:	2,
			NumExtensions:	0,
			NumServices:	0,
		},
		GoTypes:		file_hcp_types_v1_test_proto_goTypes,
		DependencyIndexes:	file_hcp_types_v1_test_proto_depIdxs,
		MessageInfos:		file_hcp_types_v1_test_proto_msgTypes,
	}.Build()
	File_hcp_types_v1_test_proto = out.File
	file_hcp_types_v1_test_proto_rawDesc = nil
	file_hcp_types_v1_test_proto_goTypes = nil
	file_hcp_types_v1_test_proto_depIdxs = nil
}
