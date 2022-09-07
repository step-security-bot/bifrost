// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1-devel
// 	protoc        v3.19.3
// source: github.com/aperturerobotics/bifrost/transport/common/kcp/packet_type.proto

package kcp

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

// PacketType is a one-byte trailer indicating the type of packet.
type PacketType int32

const (
	PacketType_PacketType_HANDSHAKE  PacketType = 0
	PacketType_PacketType_RAW        PacketType = 1
	PacketType_PacketType_KCP_SMUX   PacketType = 2
	PacketType_PacketType_CLOSE_LINK PacketType = 3
)

// Enum value maps for PacketType.
var (
	PacketType_name = map[int32]string{
		0: "PacketType_HANDSHAKE",
		1: "PacketType_RAW",
		2: "PacketType_KCP_SMUX",
		3: "PacketType_CLOSE_LINK",
	}
	PacketType_value = map[string]int32{
		"PacketType_HANDSHAKE":  0,
		"PacketType_RAW":        1,
		"PacketType_KCP_SMUX":   2,
		"PacketType_CLOSE_LINK": 3,
	}
)

func (x PacketType) Enum() *PacketType {
	p := new(PacketType)
	*p = x
	return p
}

func (x PacketType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (PacketType) Descriptor() protoreflect.EnumDescriptor {
	return file_github_com_aperturerobotics_bifrost_transport_common_kcp_packet_type_proto_enumTypes[0].Descriptor()
}

func (PacketType) Type() protoreflect.EnumType {
	return &file_github_com_aperturerobotics_bifrost_transport_common_kcp_packet_type_proto_enumTypes[0]
}

func (x PacketType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use PacketType.Descriptor instead.
func (PacketType) EnumDescriptor() ([]byte, []int) {
	return file_github_com_aperturerobotics_bifrost_transport_common_kcp_packet_type_proto_rawDescGZIP(), []int{0}
}

var File_github_com_aperturerobotics_bifrost_transport_common_kcp_packet_type_proto protoreflect.FileDescriptor

var file_github_com_aperturerobotics_bifrost_transport_common_kcp_packet_type_proto_rawDesc = []byte{
	0x0a, 0x4a, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x61, 0x70, 0x65,
	0x72, 0x74, 0x75, 0x72, 0x65, 0x72, 0x6f, 0x62, 0x6f, 0x74, 0x69, 0x63, 0x73, 0x2f, 0x62, 0x69,
	0x66, 0x72, 0x6f, 0x73, 0x74, 0x2f, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x70, 0x6f, 0x72, 0x74, 0x2f,
	0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2f, 0x6b, 0x63, 0x70, 0x2f, 0x70, 0x61, 0x63, 0x6b, 0x65,
	0x74, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x03, 0x6b, 0x63,
	0x70, 0x2a, 0x6e, 0x0a, 0x0a, 0x50, 0x61, 0x63, 0x6b, 0x65, 0x74, 0x54, 0x79, 0x70, 0x65, 0x12,
	0x18, 0x0a, 0x14, 0x50, 0x61, 0x63, 0x6b, 0x65, 0x74, 0x54, 0x79, 0x70, 0x65, 0x5f, 0x48, 0x41,
	0x4e, 0x44, 0x53, 0x48, 0x41, 0x4b, 0x45, 0x10, 0x00, 0x12, 0x12, 0x0a, 0x0e, 0x50, 0x61, 0x63,
	0x6b, 0x65, 0x74, 0x54, 0x79, 0x70, 0x65, 0x5f, 0x52, 0x41, 0x57, 0x10, 0x01, 0x12, 0x17, 0x0a,
	0x13, 0x50, 0x61, 0x63, 0x6b, 0x65, 0x74, 0x54, 0x79, 0x70, 0x65, 0x5f, 0x4b, 0x43, 0x50, 0x5f,
	0x53, 0x4d, 0x55, 0x58, 0x10, 0x02, 0x12, 0x19, 0x0a, 0x15, 0x50, 0x61, 0x63, 0x6b, 0x65, 0x74,
	0x54, 0x79, 0x70, 0x65, 0x5f, 0x43, 0x4c, 0x4f, 0x53, 0x45, 0x5f, 0x4c, 0x49, 0x4e, 0x4b, 0x10,
	0x03, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_github_com_aperturerobotics_bifrost_transport_common_kcp_packet_type_proto_rawDescOnce sync.Once
	file_github_com_aperturerobotics_bifrost_transport_common_kcp_packet_type_proto_rawDescData = file_github_com_aperturerobotics_bifrost_transport_common_kcp_packet_type_proto_rawDesc
)

func file_github_com_aperturerobotics_bifrost_transport_common_kcp_packet_type_proto_rawDescGZIP() []byte {
	file_github_com_aperturerobotics_bifrost_transport_common_kcp_packet_type_proto_rawDescOnce.Do(func() {
		file_github_com_aperturerobotics_bifrost_transport_common_kcp_packet_type_proto_rawDescData = protoimpl.X.CompressGZIP(file_github_com_aperturerobotics_bifrost_transport_common_kcp_packet_type_proto_rawDescData)
	})
	return file_github_com_aperturerobotics_bifrost_transport_common_kcp_packet_type_proto_rawDescData
}

var file_github_com_aperturerobotics_bifrost_transport_common_kcp_packet_type_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_github_com_aperturerobotics_bifrost_transport_common_kcp_packet_type_proto_goTypes = []interface{}{
	(PacketType)(0), // 0: kcp.PacketType
}
var file_github_com_aperturerobotics_bifrost_transport_common_kcp_packet_type_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_github_com_aperturerobotics_bifrost_transport_common_kcp_packet_type_proto_init() }
func file_github_com_aperturerobotics_bifrost_transport_common_kcp_packet_type_proto_init() {
	if File_github_com_aperturerobotics_bifrost_transport_common_kcp_packet_type_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_github_com_aperturerobotics_bifrost_transport_common_kcp_packet_type_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   0,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_github_com_aperturerobotics_bifrost_transport_common_kcp_packet_type_proto_goTypes,
		DependencyIndexes: file_github_com_aperturerobotics_bifrost_transport_common_kcp_packet_type_proto_depIdxs,
		EnumInfos:         file_github_com_aperturerobotics_bifrost_transport_common_kcp_packet_type_proto_enumTypes,
	}.Build()
	File_github_com_aperturerobotics_bifrost_transport_common_kcp_packet_type_proto = out.File
	file_github_com_aperturerobotics_bifrost_transport_common_kcp_packet_type_proto_rawDesc = nil
	file_github_com_aperturerobotics_bifrost_transport_common_kcp_packet_type_proto_goTypes = nil
	file_github_com_aperturerobotics_bifrost_transport_common_kcp_packet_type_proto_depIdxs = nil
}
