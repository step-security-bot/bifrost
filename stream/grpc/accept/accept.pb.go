// Code generated by protoc-gen-go. DO NOT EDIT.
// source: github.com/aperturerobotics/bifrost/stream/grpc/accept/accept.proto

package stream_grpc_accept

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

// Config configures the grpcaccept controller.
type Config struct {
	// LocalPeerId is the peer ID to accept incoming connections with.
	// Can be empty to accept any peer.
	LocalPeerId string `protobuf:"bytes,1,opt,name=local_peer_id,json=localPeerId" json:"local_peer_id,omitempty"`
	// RemotePeerIds are peer IDs to accept incoming connections from.
	// Can be empty to accept any remote peer IDs.
	RemotePeerIds []string `protobuf:"bytes,2,rep,name=remote_peer_ids,json=remotePeerIds" json:"remote_peer_ids,omitempty"`
	// ProtocolId is the protocol ID to accept.
	ProtocolId string `protobuf:"bytes,3,opt,name=protocol_id,json=protocolId" json:"protocol_id,omitempty"`
	// TransportId constrains the transport ID to accept from.
	// Can be empty.
	TransportId          uint64   `protobuf:"varint,4,opt,name=transport_id,json=transportId" json:"transport_id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Config) Reset()         { *m = Config{} }
func (m *Config) String() string { return proto.CompactTextString(m) }
func (*Config) ProtoMessage()    {}
func (*Config) Descriptor() ([]byte, []int) {
	return fileDescriptor_accept_7dd275a7d522964b, []int{0}
}
func (m *Config) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Config.Unmarshal(m, b)
}
func (m *Config) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Config.Marshal(b, m, deterministic)
}
func (dst *Config) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Config.Merge(dst, src)
}
func (m *Config) XXX_Size() int {
	return xxx_messageInfo_Config.Size(m)
}
func (m *Config) XXX_DiscardUnknown() {
	xxx_messageInfo_Config.DiscardUnknown(m)
}

var xxx_messageInfo_Config proto.InternalMessageInfo

func (m *Config) GetLocalPeerId() string {
	if m != nil {
		return m.LocalPeerId
	}
	return ""
}

func (m *Config) GetRemotePeerIds() []string {
	if m != nil {
		return m.RemotePeerIds
	}
	return nil
}

func (m *Config) GetProtocolId() string {
	if m != nil {
		return m.ProtocolId
	}
	return ""
}

func (m *Config) GetTransportId() uint64 {
	if m != nil {
		return m.TransportId
	}
	return 0
}

func init() {
	proto.RegisterType((*Config)(nil), "stream.grpc.accept.Config")
}

func init() {
	proto.RegisterFile("github.com/aperturerobotics/bifrost/stream/grpc/accept/accept.proto", fileDescriptor_accept_7dd275a7d522964b)
}

var fileDescriptor_accept_7dd275a7d522964b = []byte{
	// 205 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x3c, 0xcd, 0xbf, 0x6a, 0xc3, 0x30,
	0x10, 0xc7, 0x71, 0xd4, 0x84, 0x40, 0xe4, 0x86, 0x82, 0xa6, 0x6c, 0x75, 0x33, 0x14, 0x4f, 0xd2,
	0xd0, 0x47, 0xc8, 0xe4, 0xad, 0xf8, 0x05, 0x82, 0xfe, 0x5c, 0x5c, 0x81, 0x9d, 0x13, 0xa7, 0xf3,
	0xbb, 0xf4, 0x71, 0x8b, 0xa5, 0x3a, 0x93, 0xe0, 0xab, 0x0f, 0xbf, 0x93, 0xd7, 0x31, 0xf2, 0xcf,
	0xe2, 0xb4, 0xc7, 0xd9, 0xd8, 0x04, 0xc4, 0x0b, 0x01, 0xa1, 0x43, 0x8e, 0x3e, 0x1b, 0x17, 0xef,
	0x84, 0x99, 0x4d, 0x66, 0x02, 0x3b, 0x9b, 0x91, 0x92, 0x37, 0xd6, 0x7b, 0x48, 0xfc, 0xff, 0xe8,
	0x44, 0xc8, 0xa8, 0x54, 0x05, 0x7a, 0x05, 0xba, 0xfe, 0x5c, 0x7e, 0x85, 0x3c, 0x5c, 0xf1, 0x71,
	0x8f, 0xa3, 0xba, 0xc8, 0xd3, 0x84, 0xde, 0x4e, 0xb7, 0x04, 0x40, 0xb7, 0x18, 0xce, 0xa2, 0x15,
	0xdd, 0x71, 0x68, 0x4a, 0xfc, 0x06, 0xa0, 0x3e, 0xa8, 0x4f, 0xf9, 0x46, 0x30, 0x23, 0xc3, 0x86,
	0xf2, 0xf9, 0xa5, 0xdd, 0x75, 0xc7, 0xe1, 0x54, 0x73, 0x65, 0x59, 0xbd, 0xcb, 0xa6, 0xdc, 0xf4,
	0x38, 0xad, 0x4b, 0xbb, 0xb2, 0x24, 0xb7, 0xd4, 0x07, 0xf5, 0x21, 0x5f, 0x99, 0xec, 0x23, 0x27,
	0x24, 0x5e, 0xc5, 0xbe, 0x15, 0xdd, 0x7e, 0x68, 0x9e, 0xad, 0x0f, 0xee, 0x50, 0xf8, 0xd7, 0x5f,
	0x00, 0x00, 0x00, 0xff, 0xff, 0xe5, 0x18, 0xad, 0x8f, 0xfc, 0x00, 0x00, 0x00,
}
