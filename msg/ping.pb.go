// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0-devel
// 	protoc        v3.12.3
// source: ping.proto

package msg

import (
	proto "github.com/golang/protobuf/proto"
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

// This is a compile-time assertion that a sufficiently up-to-date version
// of the legacy proto package is being used.
const _ = proto.ProtoPackageIsVersion4

type CSPing struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *CSPing) Reset() {
	*x = CSPing{}
	if protoimpl.UnsafeEnabled {
		mi := &file_ping_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CSPing) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CSPing) ProtoMessage() {}

func (x *CSPing) ProtoReflect() protoreflect.Message {
	mi := &file_ping_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CSPing.ProtoReflect.Descriptor instead.
func (*CSPing) Descriptor() ([]byte, []int) {
	return file_ping_proto_rawDescGZIP(), []int{0}
}

type SCPong struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Timestamp int64 `protobuf:"varint,1,opt,name=timestamp,proto3" json:"timestamp,omitempty"`
}

func (x *SCPong) Reset() {
	*x = SCPong{}
	if protoimpl.UnsafeEnabled {
		mi := &file_ping_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SCPong) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SCPong) ProtoMessage() {}

func (x *SCPong) ProtoReflect() protoreflect.Message {
	mi := &file_ping_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SCPong.ProtoReflect.Descriptor instead.
func (*SCPong) Descriptor() ([]byte, []int) {
	return file_ping_proto_rawDescGZIP(), []int{1}
}

func (x *SCPong) GetTimestamp() int64 {
	if x != nil {
		return x.Timestamp
	}
	return 0
}

var File_ping_proto protoreflect.FileDescriptor

var file_ping_proto_rawDesc = []byte{
	0x0a, 0x0a, 0x70, 0x69, 0x6e, 0x67, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x03, 0x6d, 0x73,
	0x67, 0x22, 0x08, 0x0a, 0x06, 0x43, 0x53, 0x50, 0x69, 0x6e, 0x67, 0x22, 0x26, 0x0a, 0x06, 0x53,
	0x43, 0x50, 0x6f, 0x6e, 0x67, 0x12, 0x1c, 0x0a, 0x09, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61,
	0x6d, 0x70, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x09, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74,
	0x61, 0x6d, 0x70, 0x42, 0x07, 0x5a, 0x05, 0x2e, 0x3b, 0x6d, 0x73, 0x67, 0x62, 0x06, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_ping_proto_rawDescOnce sync.Once
	file_ping_proto_rawDescData = file_ping_proto_rawDesc
)

func file_ping_proto_rawDescGZIP() []byte {
	file_ping_proto_rawDescOnce.Do(func() {
		file_ping_proto_rawDescData = protoimpl.X.CompressGZIP(file_ping_proto_rawDescData)
	})
	return file_ping_proto_rawDescData
}

var file_ping_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_ping_proto_goTypes = []interface{}{
	(*CSPing)(nil), // 0: msg.CSPing
	(*SCPong)(nil), // 1: msg.SCPong
}
var file_ping_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_ping_proto_init() }
func file_ping_proto_init() {
	if File_ping_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_ping_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CSPing); i {
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
		file_ping_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SCPong); i {
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
			RawDescriptor: file_ping_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_ping_proto_goTypes,
		DependencyIndexes: file_ping_proto_depIdxs,
		MessageInfos:      file_ping_proto_msgTypes,
	}.Build()
	File_ping_proto = out.File
	file_ping_proto_rawDesc = nil
	file_ping_proto_goTypes = nil
	file_ping_proto_depIdxs = nil
}
