// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.0
// 	protoc        v3.21.2
// source: errors/errors.proto

package errors

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

type ErrorType int32

const (
	ErrorType_Fatal ErrorType = 0
)

// Enum value maps for ErrorType.
var (
	ErrorType_name = map[int32]string{
		0: "Fatal",
	}
	ErrorType_value = map[string]int32{
		"Fatal": 0,
	}
)

func (x ErrorType) Enum() *ErrorType {
	p := new(ErrorType)
	*p = x
	return p
}

func (x ErrorType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (ErrorType) Descriptor() protoreflect.EnumDescriptor {
	return file_errors_errors_proto_enumTypes[0].Descriptor()
}

func (ErrorType) Type() protoreflect.EnumType {
	return &file_errors_errors_proto_enumTypes[0]
}

func (x ErrorType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use ErrorType.Descriptor instead.
func (ErrorType) EnumDescriptor() ([]byte, []int) {
	return file_errors_errors_proto_rawDescGZIP(), []int{0}
}

type Error struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Type  ErrorType `protobuf:"varint,1,opt,name=type,proto3,enum=errors.ErrorType" json:"type,omitempty"`
	Value string    `protobuf:"bytes,2,opt,name=value,proto3" json:"value,omitempty"`
}

func (x *Error) Reset() {
	*x = Error{}
	if protoimpl.UnsafeEnabled {
		mi := &file_errors_errors_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Error) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Error) ProtoMessage() {}

func (x *Error) ProtoReflect() protoreflect.Message {
	mi := &file_errors_errors_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Error.ProtoReflect.Descriptor instead.
func (*Error) Descriptor() ([]byte, []int) {
	return file_errors_errors_proto_rawDescGZIP(), []int{0}
}

func (x *Error) GetType() ErrorType {
	if x != nil {
		return x.Type
	}
	return ErrorType_Fatal
}

func (x *Error) GetValue() string {
	if x != nil {
		return x.Value
	}
	return ""
}

var File_errors_errors_proto protoreflect.FileDescriptor

var file_errors_errors_proto_rawDesc = []byte{
	0x0a, 0x13, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x73, 0x2f, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x73, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x06, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x73, 0x22, 0x44, 0x0a,
	0x05, 0x45, 0x72, 0x72, 0x6f, 0x72, 0x12, 0x25, 0x0a, 0x04, 0x74, 0x79, 0x70, 0x65, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x0e, 0x32, 0x11, 0x2e, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x73, 0x2e, 0x45, 0x72,
	0x72, 0x6f, 0x72, 0x54, 0x79, 0x70, 0x65, 0x52, 0x04, 0x74, 0x79, 0x70, 0x65, 0x12, 0x14, 0x0a,
	0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x76, 0x61,
	0x6c, 0x75, 0x65, 0x2a, 0x16, 0x0a, 0x09, 0x45, 0x72, 0x72, 0x6f, 0x72, 0x54, 0x79, 0x70, 0x65,
	0x12, 0x09, 0x0a, 0x05, 0x46, 0x61, 0x74, 0x61, 0x6c, 0x10, 0x00, 0x42, 0x27, 0x5a, 0x25, 0x67,
	0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x76, 0x70, 0x6f, 0x6c, 0x2f, 0x64,
	0x61, 0x69, 0x6c, 0x79, 0x2d, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2d, 0x67, 0x6f, 0x2f, 0x65, 0x72,
	0x72, 0x6f, 0x72, 0x73, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_errors_errors_proto_rawDescOnce sync.Once
	file_errors_errors_proto_rawDescData = file_errors_errors_proto_rawDesc
)

func file_errors_errors_proto_rawDescGZIP() []byte {
	file_errors_errors_proto_rawDescOnce.Do(func() {
		file_errors_errors_proto_rawDescData = protoimpl.X.CompressGZIP(file_errors_errors_proto_rawDescData)
	})
	return file_errors_errors_proto_rawDescData
}

var file_errors_errors_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_errors_errors_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_errors_errors_proto_goTypes = []interface{}{
	(ErrorType)(0), // 0: errors.ErrorType
	(*Error)(nil),  // 1: errors.Error
}
var file_errors_errors_proto_depIdxs = []int32{
	0, // 0: errors.Error.type:type_name -> errors.ErrorType
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_errors_errors_proto_init() }
func file_errors_errors_proto_init() {
	if File_errors_errors_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_errors_errors_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Error); i {
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
			RawDescriptor: file_errors_errors_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_errors_errors_proto_goTypes,
		DependencyIndexes: file_errors_errors_proto_depIdxs,
		EnumInfos:         file_errors_errors_proto_enumTypes,
		MessageInfos:      file_errors_errors_proto_msgTypes,
	}.Build()
	File_errors_errors_proto = out.File
	file_errors_errors_proto_rawDesc = nil
	file_errors_errors_proto_goTypes = nil
	file_errors_errors_proto_depIdxs = nil
}
