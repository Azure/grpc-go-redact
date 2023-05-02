// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.30.0
// 	protoc        v4.22.3
// source: generator/gen/stringfunc.proto

package internal

import (
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

/*
type X struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Secret       string  `protobuf:"bytes,1,opt,name=secret,proto3" json:"secret,omitempty"`
	NonSecret    string  `protobuf:"bytes,2,opt,name=nonSecret,proto3" json:"nonSecret,omitempty" redact:"nonsecret"`
	SecretPtr    *string `protobuf:"bytes,3,opt,name=secretPtr,proto3,oneof" json:"secretPtr,omitempty"`
	NonSecretPtr *string `protobuf:"bytes,4,opt,name=nonSecretPtr,proto3,oneof" json:"nonSecretPtr,omitempty" redact:"nonsecret"`
}
*/

func (x *X) Reset() {
	*x = X{}
	if protoimpl.UnsafeEnabled {
		mi := &file_generator_gen_stringfunc_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

/*
func (x *X) String() string {
	return protoimpl.X.MessageStringOf(x)
}
*/

func (*X) ProtoMessage() {}

func (x *X) ProtoReflect() protoreflect.Message {
	mi := &file_generator_gen_stringfunc_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use X.ProtoReflect.Descriptor instead.
func (*X) Descriptor() ([]byte, []int) {
	return file_generator_gen_stringfunc_proto_rawDescGZIP(), []int{0}
}

func (x *X) GetSecret() string {
	if x != nil {
		return x.Secret
	}
	return ""
}

func (x *X) GetNonSecret() string {
	if x != nil {
		return x.NonSecret
	}
	return ""
}

func (x *X) GetSecretPtr() string {
	if x != nil && x.SecretPtr != nil {
		return *x.SecretPtr
	}
	return ""
}

func (x *X) GetNonSecretPtr() string {
	if x != nil && x.NonSecretPtr != nil {
		return *x.NonSecretPtr
	}
	return ""
}

type Xs struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Data []*X `protobuf:"bytes,1,rep,name=data,proto3" json:"data,omitempty"`
}

func (x *Xs) Reset() {
	*x = Xs{}
	if protoimpl.UnsafeEnabled {
		mi := &file_generator_gen_stringfunc_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Xs) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Xs) ProtoMessage() {}

func (x *Xs) ProtoReflect() protoreflect.Message {
	mi := &file_generator_gen_stringfunc_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Xs.ProtoReflect.Descriptor instead.
func (*Xs) Descriptor() ([]byte, []int) {
	return file_generator_gen_stringfunc_proto_rawDescGZIP(), []int{1}
}

func (x *Xs) GetData() []*X {
	if x != nil {
		return x.Data
	}
	return nil
}

var File_generator_gen_stringfunc_proto protoreflect.FileDescriptor

var file_generator_gen_stringfunc_proto_rawDesc = []byte{
	0x0a, 0x1e, 0x67, 0x65, 0x6e, 0x65, 0x72, 0x61, 0x74, 0x6f, 0x72, 0x2f, 0x67, 0x65, 0x6e, 0x2f,
	0x73, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x66, 0x75, 0x6e, 0x63, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x12, 0x04, 0x74, 0x65, 0x73, 0x74, 0x22, 0xa4, 0x01, 0x0a, 0x01, 0x58, 0x12, 0x16, 0x0a, 0x06,
	0x73, 0x65, 0x63, 0x72, 0x65, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x73, 0x65,
	0x63, 0x72, 0x65, 0x74, 0x12, 0x1c, 0x0a, 0x09, 0x6e, 0x6f, 0x6e, 0x53, 0x65, 0x63, 0x72, 0x65,
	0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x6e, 0x6f, 0x6e, 0x53, 0x65, 0x63, 0x72,
	0x65, 0x74, 0x12, 0x21, 0x0a, 0x09, 0x73, 0x65, 0x63, 0x72, 0x65, 0x74, 0x50, 0x74, 0x72, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x09, 0x48, 0x00, 0x52, 0x09, 0x73, 0x65, 0x63, 0x72, 0x65, 0x74, 0x50,
	0x74, 0x72, 0x88, 0x01, 0x01, 0x12, 0x27, 0x0a, 0x0c, 0x6e, 0x6f, 0x6e, 0x53, 0x65, 0x63, 0x72,
	0x65, 0x74, 0x50, 0x74, 0x72, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x48, 0x01, 0x52, 0x0c, 0x6e,
	0x6f, 0x6e, 0x53, 0x65, 0x63, 0x72, 0x65, 0x74, 0x50, 0x74, 0x72, 0x88, 0x01, 0x01, 0x42, 0x0c,
	0x0a, 0x0a, 0x5f, 0x73, 0x65, 0x63, 0x72, 0x65, 0x74, 0x50, 0x74, 0x72, 0x42, 0x0f, 0x0a, 0x0d,
	0x5f, 0x6e, 0x6f, 0x6e, 0x53, 0x65, 0x63, 0x72, 0x65, 0x74, 0x50, 0x74, 0x72, 0x22, 0x21, 0x0a,
	0x02, 0x58, 0x73, 0x12, 0x1b, 0x0a, 0x04, 0x64, 0x61, 0x74, 0x61, 0x18, 0x01, 0x20, 0x03, 0x28,
	0x0b, 0x32, 0x07, 0x2e, 0x74, 0x65, 0x73, 0x74, 0x2e, 0x58, 0x52, 0x04, 0x64, 0x61, 0x74, 0x61,
	0x42, 0x32, 0x5a, 0x30, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x73,
	0x61, 0x6d, 0x72, 0x6b, 0x65, 0x74, 0x65, 0x72, 0x2f, 0x67, 0x72, 0x70, 0x63, 0x2d, 0x67, 0x6f,
	0x2d, 0x72, 0x65, 0x61, 0x63, 0x74, 0x2f, 0x67, 0x65, 0x6e, 0x65, 0x72, 0x61, 0x74, 0x6f, 0x72,
	0x2f, 0x67, 0x65, 0x6e, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_generator_gen_stringfunc_proto_rawDescOnce sync.Once
	file_generator_gen_stringfunc_proto_rawDescData = file_generator_gen_stringfunc_proto_rawDesc
)

func file_generator_gen_stringfunc_proto_rawDescGZIP() []byte {
	file_generator_gen_stringfunc_proto_rawDescOnce.Do(func() {
		file_generator_gen_stringfunc_proto_rawDescData = protoimpl.X.CompressGZIP(file_generator_gen_stringfunc_proto_rawDescData)
	})
	return file_generator_gen_stringfunc_proto_rawDescData
}

var file_generator_gen_stringfunc_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_generator_gen_stringfunc_proto_goTypes = []interface{}{
	(*X)(nil),  // 0: test.X
	(*Xs)(nil), // 1: test.Xs
}
var file_generator_gen_stringfunc_proto_depIdxs = []int32{
	0, // 0: test.Xs.data:type_name -> test.X
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_generator_gen_stringfunc_proto_init() }
func file_generator_gen_stringfunc_proto_init() {
	if File_generator_gen_stringfunc_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_generator_gen_stringfunc_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*X); i {
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
		file_generator_gen_stringfunc_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Xs); i {
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
	file_generator_gen_stringfunc_proto_msgTypes[0].OneofWrappers = []interface{}{}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_generator_gen_stringfunc_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_generator_gen_stringfunc_proto_goTypes,
		DependencyIndexes: file_generator_gen_stringfunc_proto_depIdxs,
		MessageInfos:      file_generator_gen_stringfunc_proto_msgTypes,
	}.Build()
	File_generator_gen_stringfunc_proto = out.File
	file_generator_gen_stringfunc_proto_rawDesc = nil
	file_generator_gen_stringfunc_proto_goTypes = nil
	file_generator_gen_stringfunc_proto_depIdxs = nil
}
