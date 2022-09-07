// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1-devel
// 	protoc        v3.19.3
// source: github.com/aperturerobotics/bifrost/stream/echo/echo.proto

package stream_echo

import (
	reflect "reflect"
	sync "sync"

	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// Config configures the echo controller.
type Config struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// PeerId is the peer ID to echo for.
	// Can be empty.
	PeerId string `protobuf:"bytes,1,opt,name=peer_id,json=peerId,proto3" json:"peer_id,omitempty"`
	// ProtocolId is the protocol ID to echo on.
	ProtocolId string `protobuf:"bytes,2,opt,name=protocol_id,json=protocolId,proto3" json:"protocol_id,omitempty"`
}

func (x *Config) Reset() {
	*x = Config{}
	if protoimpl.UnsafeEnabled {
		mi := &file_github_com_aperturerobotics_bifrost_stream_echo_echo_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Config) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Config) ProtoMessage() {}

func (x *Config) ProtoReflect() protoreflect.Message {
	mi := &file_github_com_aperturerobotics_bifrost_stream_echo_echo_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Config.ProtoReflect.Descriptor instead.
func (*Config) Descriptor() ([]byte, []int) {
	return file_github_com_aperturerobotics_bifrost_stream_echo_echo_proto_rawDescGZIP(), []int{0}
}

func (x *Config) GetPeerId() string {
	if x != nil {
		return x.PeerId
	}
	return ""
}

func (x *Config) GetProtocolId() string {
	if x != nil {
		return x.ProtocolId
	}
	return ""
}

var File_github_com_aperturerobotics_bifrost_stream_echo_echo_proto protoreflect.FileDescriptor

var file_github_com_aperturerobotics_bifrost_stream_echo_echo_proto_rawDesc = []byte{
	0x0a, 0x3a, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x61, 0x70, 0x65,
	0x72, 0x74, 0x75, 0x72, 0x65, 0x72, 0x6f, 0x62, 0x6f, 0x74, 0x69, 0x63, 0x73, 0x2f, 0x62, 0x69,
	0x66, 0x72, 0x6f, 0x73, 0x74, 0x2f, 0x73, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x2f, 0x65, 0x63, 0x68,
	0x6f, 0x2f, 0x65, 0x63, 0x68, 0x6f, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0b, 0x73, 0x74,
	0x72, 0x65, 0x61, 0x6d, 0x2e, 0x65, 0x63, 0x68, 0x6f, 0x22, 0x42, 0x0a, 0x06, 0x43, 0x6f, 0x6e,
	0x66, 0x69, 0x67, 0x12, 0x17, 0x0a, 0x07, 0x70, 0x65, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x70, 0x65, 0x65, 0x72, 0x49, 0x64, 0x12, 0x1f, 0x0a, 0x0b,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x0a, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x49, 0x64, 0x62, 0x06, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_github_com_aperturerobotics_bifrost_stream_echo_echo_proto_rawDescOnce sync.Once
	file_github_com_aperturerobotics_bifrost_stream_echo_echo_proto_rawDescData = file_github_com_aperturerobotics_bifrost_stream_echo_echo_proto_rawDesc
)

func file_github_com_aperturerobotics_bifrost_stream_echo_echo_proto_rawDescGZIP() []byte {
	file_github_com_aperturerobotics_bifrost_stream_echo_echo_proto_rawDescOnce.Do(func() {
		file_github_com_aperturerobotics_bifrost_stream_echo_echo_proto_rawDescData = protoimpl.X.CompressGZIP(file_github_com_aperturerobotics_bifrost_stream_echo_echo_proto_rawDescData)
	})
	return file_github_com_aperturerobotics_bifrost_stream_echo_echo_proto_rawDescData
}

var file_github_com_aperturerobotics_bifrost_stream_echo_echo_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_github_com_aperturerobotics_bifrost_stream_echo_echo_proto_goTypes = []interface{}{
	(*Config)(nil), // 0: stream.echo.Config
}
var file_github_com_aperturerobotics_bifrost_stream_echo_echo_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_github_com_aperturerobotics_bifrost_stream_echo_echo_proto_init() }
func file_github_com_aperturerobotics_bifrost_stream_echo_echo_proto_init() {
	if File_github_com_aperturerobotics_bifrost_stream_echo_echo_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_github_com_aperturerobotics_bifrost_stream_echo_echo_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Config); i {
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
			RawDescriptor: file_github_com_aperturerobotics_bifrost_stream_echo_echo_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_github_com_aperturerobotics_bifrost_stream_echo_echo_proto_goTypes,
		DependencyIndexes: file_github_com_aperturerobotics_bifrost_stream_echo_echo_proto_depIdxs,
		MessageInfos:      file_github_com_aperturerobotics_bifrost_stream_echo_echo_proto_msgTypes,
	}.Build()
	File_github_com_aperturerobotics_bifrost_stream_echo_echo_proto = out.File
	file_github_com_aperturerobotics_bifrost_stream_echo_echo_proto_rawDesc = nil
	file_github_com_aperturerobotics_bifrost_stream_echo_echo_proto_goTypes = nil
	file_github_com_aperturerobotics_bifrost_stream_echo_echo_proto_depIdxs = nil
}
