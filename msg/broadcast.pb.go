// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0-devel
// 	protoc        v3.12.3
// source: broadcast.proto

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

type MsgType int32

const (
	MsgType_TEXT  MsgType = 0
	MsgType_IMAGE MsgType = 1
	MsgType_VIDEO MsgType = 2
)

// Enum value maps for MsgType.
var (
	MsgType_name = map[int32]string{
		0: "TEXT",
		1: "IMAGE",
		2: "VIDEO",
	}
	MsgType_value = map[string]int32{
		"TEXT":  0,
		"IMAGE": 1,
		"VIDEO": 2,
	}
)

func (x MsgType) Enum() *MsgType {
	p := new(MsgType)
	*p = x
	return p
}

func (x MsgType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (MsgType) Descriptor() protoreflect.EnumDescriptor {
	return file_broadcast_proto_enumTypes[0].Descriptor()
}

func (MsgType) Type() protoreflect.EnumType {
	return &file_broadcast_proto_enumTypes[0]
}

func (x MsgType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use MsgType.Descriptor instead.
func (MsgType) EnumDescriptor() ([]byte, []int) {
	return file_broadcast_proto_rawDescGZIP(), []int{0}
}

type CSBroadcastMultiple struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	MsgType MsgType `protobuf:"varint,1,opt,name=msgType,proto3,enum=msg.MsgType" json:"msgType,omitempty"`
	Text    string  `protobuf:"bytes,2,opt,name=text,proto3" json:"text,omitempty"`
	Image   string  `protobuf:"bytes,3,opt,name=image,proto3" json:"image,omitempty"`
	Video   string  `protobuf:"bytes,4,opt,name=video,proto3" json:"video,omitempty"`
}

func (x *CSBroadcastMultiple) Reset() {
	*x = CSBroadcastMultiple{}
	if protoimpl.UnsafeEnabled {
		mi := &file_broadcast_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CSBroadcastMultiple) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CSBroadcastMultiple) ProtoMessage() {}

func (x *CSBroadcastMultiple) ProtoReflect() protoreflect.Message {
	mi := &file_broadcast_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CSBroadcastMultiple.ProtoReflect.Descriptor instead.
func (*CSBroadcastMultiple) Descriptor() ([]byte, []int) {
	return file_broadcast_proto_rawDescGZIP(), []int{0}
}

func (x *CSBroadcastMultiple) GetMsgType() MsgType {
	if x != nil {
		return x.MsgType
	}
	return MsgType_TEXT
}

func (x *CSBroadcastMultiple) GetText() string {
	if x != nil {
		return x.Text
	}
	return ""
}

func (x *CSBroadcastMultiple) GetImage() string {
	if x != nil {
		return x.Image
	}
	return ""
}

func (x *CSBroadcastMultiple) GetVideo() string {
	if x != nil {
		return x.Video
	}
	return ""
}

type SCBroadcastMultiple struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	MsgType MsgType `protobuf:"varint,1,opt,name=msgType,proto3,enum=msg.MsgType" json:"msgType,omitempty"`
	Source  string  `protobuf:"bytes,2,opt,name=source,proto3" json:"source,omitempty"`
	Text    string  `protobuf:"bytes,3,opt,name=text,proto3" json:"text,omitempty"`
	Image   string  `protobuf:"bytes,4,opt,name=image,proto3" json:"image,omitempty"`
	Video   string  `protobuf:"bytes,5,opt,name=video,proto3" json:"video,omitempty"`
}

func (x *SCBroadcastMultiple) Reset() {
	*x = SCBroadcastMultiple{}
	if protoimpl.UnsafeEnabled {
		mi := &file_broadcast_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SCBroadcastMultiple) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SCBroadcastMultiple) ProtoMessage() {}

func (x *SCBroadcastMultiple) ProtoReflect() protoreflect.Message {
	mi := &file_broadcast_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SCBroadcastMultiple.ProtoReflect.Descriptor instead.
func (*SCBroadcastMultiple) Descriptor() ([]byte, []int) {
	return file_broadcast_proto_rawDescGZIP(), []int{1}
}

func (x *SCBroadcastMultiple) GetMsgType() MsgType {
	if x != nil {
		return x.MsgType
	}
	return MsgType_TEXT
}

func (x *SCBroadcastMultiple) GetSource() string {
	if x != nil {
		return x.Source
	}
	return ""
}

func (x *SCBroadcastMultiple) GetText() string {
	if x != nil {
		return x.Text
	}
	return ""
}

func (x *SCBroadcastMultiple) GetImage() string {
	if x != nil {
		return x.Image
	}
	return ""
}

func (x *SCBroadcastMultiple) GetVideo() string {
	if x != nil {
		return x.Video
	}
	return ""
}

var File_broadcast_proto protoreflect.FileDescriptor

var file_broadcast_proto_rawDesc = []byte{
	0x0a, 0x0f, 0x62, 0x72, 0x6f, 0x61, 0x64, 0x63, 0x61, 0x73, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x12, 0x03, 0x6d, 0x73, 0x67, 0x22, 0x7d, 0x0a, 0x13, 0x43, 0x53, 0x42, 0x72, 0x6f, 0x61,
	0x64, 0x63, 0x61, 0x73, 0x74, 0x4d, 0x75, 0x6c, 0x74, 0x69, 0x70, 0x6c, 0x65, 0x12, 0x26, 0x0a,
	0x07, 0x6d, 0x73, 0x67, 0x54, 0x79, 0x70, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x0c,
	0x2e, 0x6d, 0x73, 0x67, 0x2e, 0x4d, 0x73, 0x67, 0x54, 0x79, 0x70, 0x65, 0x52, 0x07, 0x6d, 0x73,
	0x67, 0x54, 0x79, 0x70, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x74, 0x65, 0x78, 0x74, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x04, 0x74, 0x65, 0x78, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x69, 0x6d, 0x61,
	0x67, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x69, 0x6d, 0x61, 0x67, 0x65, 0x12,
	0x14, 0x0a, 0x05, 0x76, 0x69, 0x64, 0x65, 0x6f, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05,
	0x76, 0x69, 0x64, 0x65, 0x6f, 0x22, 0x95, 0x01, 0x0a, 0x13, 0x53, 0x43, 0x42, 0x72, 0x6f, 0x61,
	0x64, 0x63, 0x61, 0x73, 0x74, 0x4d, 0x75, 0x6c, 0x74, 0x69, 0x70, 0x6c, 0x65, 0x12, 0x26, 0x0a,
	0x07, 0x6d, 0x73, 0x67, 0x54, 0x79, 0x70, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x0c,
	0x2e, 0x6d, 0x73, 0x67, 0x2e, 0x4d, 0x73, 0x67, 0x54, 0x79, 0x70, 0x65, 0x52, 0x07, 0x6d, 0x73,
	0x67, 0x54, 0x79, 0x70, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x12, 0x12, 0x0a,
	0x04, 0x74, 0x65, 0x78, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x74, 0x65, 0x78,
	0x74, 0x12, 0x14, 0x0a, 0x05, 0x69, 0x6d, 0x61, 0x67, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x05, 0x69, 0x6d, 0x61, 0x67, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x69, 0x64, 0x65, 0x6f,
	0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x76, 0x69, 0x64, 0x65, 0x6f, 0x2a, 0x29, 0x0a,
	0x07, 0x4d, 0x73, 0x67, 0x54, 0x79, 0x70, 0x65, 0x12, 0x08, 0x0a, 0x04, 0x54, 0x45, 0x58, 0x54,
	0x10, 0x00, 0x12, 0x09, 0x0a, 0x05, 0x49, 0x4d, 0x41, 0x47, 0x45, 0x10, 0x01, 0x12, 0x09, 0x0a,
	0x05, 0x56, 0x49, 0x44, 0x45, 0x4f, 0x10, 0x02, 0x42, 0x07, 0x5a, 0x05, 0x2e, 0x3b, 0x6d, 0x73,
	0x67, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_broadcast_proto_rawDescOnce sync.Once
	file_broadcast_proto_rawDescData = file_broadcast_proto_rawDesc
)

func file_broadcast_proto_rawDescGZIP() []byte {
	file_broadcast_proto_rawDescOnce.Do(func() {
		file_broadcast_proto_rawDescData = protoimpl.X.CompressGZIP(file_broadcast_proto_rawDescData)
	})
	return file_broadcast_proto_rawDescData
}

var file_broadcast_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_broadcast_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_broadcast_proto_goTypes = []interface{}{
	(MsgType)(0),                // 0: msg.MsgType
	(*CSBroadcastMultiple)(nil), // 1: msg.CSBroadcastMultiple
	(*SCBroadcastMultiple)(nil), // 2: msg.SCBroadcastMultiple
}
var file_broadcast_proto_depIdxs = []int32{
	0, // 0: msg.CSBroadcastMultiple.msgType:type_name -> msg.MsgType
	0, // 1: msg.SCBroadcastMultiple.msgType:type_name -> msg.MsgType
	2, // [2:2] is the sub-list for method output_type
	2, // [2:2] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_broadcast_proto_init() }
func file_broadcast_proto_init() {
	if File_broadcast_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_broadcast_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CSBroadcastMultiple); i {
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
		file_broadcast_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SCBroadcastMultiple); i {
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
			RawDescriptor: file_broadcast_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_broadcast_proto_goTypes,
		DependencyIndexes: file_broadcast_proto_depIdxs,
		EnumInfos:         file_broadcast_proto_enumTypes,
		MessageInfos:      file_broadcast_proto_msgTypes,
	}.Build()
	File_broadcast_proto = out.File
	file_broadcast_proto_rawDesc = nil
	file_broadcast_proto_goTypes = nil
	file_broadcast_proto_depIdxs = nil
}