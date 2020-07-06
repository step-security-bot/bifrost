// Code generated by protoc-gen-go. DO NOT EDIT.
// source: github.com/aperturerobotics/bifrost/transport/websocket/packet_type.proto

package websocket

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	math "math"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

// PacketType is a one-byte trailer indicating the type of packet.
type PacketType int32

const (
	PacketType_PacketType_HANDSHAKE  PacketType = 0
	PacketType_PacketType_RAW        PacketType = 1
	PacketType_PacketType_CRYPT_SMUX PacketType = 2
)

var PacketType_name = map[int32]string{
	0: "PacketType_HANDSHAKE",
	1: "PacketType_RAW",
	2: "PacketType_CRYPT_SMUX",
}

var PacketType_value = map[string]int32{
	"PacketType_HANDSHAKE":  0,
	"PacketType_RAW":        1,
	"PacketType_CRYPT_SMUX": 2,
}

func (x PacketType) String() string {
	return proto.EnumName(PacketType_name, int32(x))
}

func (PacketType) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_e395887c1e05dc29, []int{0}
}

func init() {
	proto.RegisterEnum("websocket.PacketType", PacketType_name, PacketType_value)
}

func init() {
	proto.RegisterFile("github.com/aperturerobotics/bifrost/transport/websocket/packet_type.proto", fileDescriptor_e395887c1e05dc29)
}

var fileDescriptor_e395887c1e05dc29 = []byte{
	// 163 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xf2, 0x4c, 0xcf, 0x2c, 0xc9,
	0x28, 0x4d, 0xd2, 0x4b, 0xce, 0xcf, 0xd5, 0x4f, 0x2c, 0x48, 0x2d, 0x2a, 0x29, 0x2d, 0x4a, 0x2d,
	0xca, 0x4f, 0xca, 0x2f, 0xc9, 0x4c, 0x2e, 0xd6, 0x4f, 0xca, 0x4c, 0x2b, 0xca, 0x2f, 0x2e, 0xd1,
	0x2f, 0x29, 0x4a, 0xcc, 0x2b, 0x2e, 0xc8, 0x2f, 0x2a, 0xd1, 0x2f, 0x4f, 0x4d, 0x2a, 0xce, 0x4f,
	0xce, 0x4e, 0x2d, 0xd1, 0x2f, 0x48, 0x04, 0x51, 0xf1, 0x25, 0x95, 0x05, 0xa9, 0x7a, 0x05, 0x45,
	0xf9, 0x25, 0xf9, 0x42, 0x9c, 0x70, 0x49, 0xad, 0x50, 0x2e, 0xae, 0x00, 0xb0, 0x7c, 0x48, 0x65,
	0x41, 0xaa, 0x90, 0x04, 0x97, 0x08, 0x82, 0x17, 0xef, 0xe1, 0xe8, 0xe7, 0x12, 0xec, 0xe1, 0xe8,
	0xed, 0x2a, 0xc0, 0x20, 0x24, 0xc4, 0xc5, 0x87, 0x24, 0x13, 0xe4, 0x18, 0x2e, 0xc0, 0x28, 0x24,
	0xc9, 0x25, 0x8a, 0x24, 0xe6, 0x1c, 0x14, 0x19, 0x10, 0x12, 0x1f, 0xec, 0x1b, 0x1a, 0x21, 0xc0,
	0x94, 0xc4, 0x06, 0xb6, 0xc8, 0x18, 0x10, 0x00, 0x00, 0xff, 0xff, 0x1d, 0x7d, 0xac, 0xe5, 0xb5,
	0x00, 0x00, 0x00,
}
