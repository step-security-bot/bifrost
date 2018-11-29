// Code generated by protoc-gen-go. DO NOT EDIT.
// source: github.com/aperturerobotics/bifrost/stream/grpc/dial/dial.proto

package stream_grpc_dial

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

// Config configures the grpcdial controller.
type Config struct {
	// PeerId is the remote peer ID to dial.
	PeerId string `protobuf:"bytes,1,opt,name=peer_id,json=peerId" json:"peer_id,omitempty"`
	// LocalPeerId is the peer ID to dial with.
	// Can be empty to accept any loaded peer.
	LocalPeerId string `protobuf:"bytes,2,opt,name=local_peer_id,json=localPeerId" json:"local_peer_id,omitempty"`
	// ProtocolId is the protocol ID to dial with.
	ProtocolId string `protobuf:"bytes,3,opt,name=protocol_id,json=protocolId" json:"protocol_id,omitempty"`
	// TransportId constrains the transport ID to dial with.
	// Can be empty.
	TransportId uint64 `protobuf:"varint,4,opt,name=transport_id,json=transportId" json:"transport_id,omitempty"`
	// Encrypted indicates the stream should be encrypted.
	Encrypted bool `protobuf:"varint,5,opt,name=encrypted" json:"encrypted,omitempty"`
	// Reliable indicates the stream should be reliable.
	Reliable             bool     `protobuf:"varint,6,opt,name=reliable" json:"reliable,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Config) Reset()         { *m = Config{} }
func (m *Config) String() string { return proto.CompactTextString(m) }
func (*Config) ProtoMessage()    {}
func (*Config) Descriptor() ([]byte, []int) {
	return fileDescriptor_dial_92d48260ffe2cca8, []int{0}
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

func (m *Config) GetPeerId() string {
	if m != nil {
		return m.PeerId
	}
	return ""
}

func (m *Config) GetLocalPeerId() string {
	if m != nil {
		return m.LocalPeerId
	}
	return ""
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

func (m *Config) GetEncrypted() bool {
	if m != nil {
		return m.Encrypted
	}
	return false
}

func (m *Config) GetReliable() bool {
	if m != nil {
		return m.Reliable
	}
	return false
}

// Request is a request packet.
type Request struct {
	// Data is packet data from the remote.
	Data                 []byte   `protobuf:"bytes,1,opt,name=data,proto3" json:"data,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Request) Reset()         { *m = Request{} }
func (m *Request) String() string { return proto.CompactTextString(m) }
func (*Request) ProtoMessage()    {}
func (*Request) Descriptor() ([]byte, []int) {
	return fileDescriptor_dial_92d48260ffe2cca8, []int{1}
}
func (m *Request) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Request.Unmarshal(m, b)
}
func (m *Request) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Request.Marshal(b, m, deterministic)
}
func (dst *Request) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Request.Merge(dst, src)
}
func (m *Request) XXX_Size() int {
	return xxx_messageInfo_Request.Size(m)
}
func (m *Request) XXX_DiscardUnknown() {
	xxx_messageInfo_Request.DiscardUnknown(m)
}

var xxx_messageInfo_Request proto.InternalMessageInfo

func (m *Request) GetData() []byte {
	if m != nil {
		return m.Data
	}
	return nil
}

// Response is a response packet.
type Response struct {
	// Data is packet data from the remote.
	Data                 []byte   `protobuf:"bytes,1,opt,name=data,proto3" json:"data,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Response) Reset()         { *m = Response{} }
func (m *Response) String() string { return proto.CompactTextString(m) }
func (*Response) ProtoMessage()    {}
func (*Response) Descriptor() ([]byte, []int) {
	return fileDescriptor_dial_92d48260ffe2cca8, []int{2}
}
func (m *Response) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Response.Unmarshal(m, b)
}
func (m *Response) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Response.Marshal(b, m, deterministic)
}
func (dst *Response) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Response.Merge(dst, src)
}
func (m *Response) XXX_Size() int {
	return xxx_messageInfo_Response.Size(m)
}
func (m *Response) XXX_DiscardUnknown() {
	xxx_messageInfo_Response.DiscardUnknown(m)
}

var xxx_messageInfo_Response proto.InternalMessageInfo

func (m *Response) GetData() []byte {
	if m != nil {
		return m.Data
	}
	return nil
}

func init() {
	proto.RegisterType((*Config)(nil), "stream.grpc.dial.Config")
	proto.RegisterType((*Request)(nil), "stream.grpc.dial.Request")
	proto.RegisterType((*Response)(nil), "stream.grpc.dial.Response")
}

func init() {
	proto.RegisterFile("github.com/aperturerobotics/bifrost/stream/grpc/dial/dial.proto", fileDescriptor_dial_92d48260ffe2cca8)
}

var fileDescriptor_dial_92d48260ffe2cca8 = []byte{
	// 257 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x6c, 0x8f, 0xbf, 0x4e, 0xfb, 0x30,
	0x10, 0x80, 0x95, 0xdf, 0xaf, 0xa4, 0xed, 0xa5, 0x48, 0xc8, 0x0b, 0x11, 0xe2, 0x4f, 0xc8, 0x94,
	0x29, 0x19, 0x78, 0x00, 0x06, 0xa6, 0x6c, 0xc8, 0x2f, 0x50, 0x39, 0xf6, 0x35, 0x58, 0x72, 0x63,
	0x73, 0xbe, 0x0c, 0x3c, 0x1f, 0x2f, 0x86, 0xe2, 0xaa, 0x65, 0x61, 0xb1, 0x7c, 0xdf, 0xf7, 0x49,
	0xf6, 0xc1, 0xeb, 0x68, 0xf9, 0x63, 0x1e, 0x5a, 0xed, 0x8f, 0x9d, 0x0a, 0x48, 0x3c, 0x13, 0x92,
	0x1f, 0x3c, 0x5b, 0x1d, 0xbb, 0xc1, 0x1e, 0xc8, 0x47, 0xee, 0x22, 0x13, 0xaa, 0x63, 0x37, 0x52,
	0xd0, 0x9d, 0xb1, 0xca, 0xa5, 0xa3, 0x0d, 0xe4, 0xd9, 0x8b, 0x9b, 0x93, 0x6c, 0x17, 0xd9, 0x2e,
	0xbc, 0xfe, 0xce, 0x20, 0x7f, 0xf3, 0xd3, 0xc1, 0x8e, 0xe2, 0x16, 0xd6, 0x01, 0x91, 0xf6, 0xd6,
	0x94, 0x59, 0x95, 0x35, 0x5b, 0x99, 0x2f, 0x63, 0x6f, 0x44, 0x0d, 0xd7, 0xce, 0x6b, 0xe5, 0xf6,
	0x67, 0xfd, 0x2f, 0xe9, 0x22, 0xc1, 0xf7, 0x53, 0xf3, 0x04, 0x45, 0x7a, 0x42, 0x7b, 0xb7, 0x14,
	0xff, 0x53, 0x01, 0x67, 0xd4, 0x1b, 0xf1, 0x0c, 0x3b, 0x26, 0x35, 0xc5, 0xe0, 0x89, 0x97, 0x62,
	0x55, 0x65, 0xcd, 0x4a, 0x16, 0x17, 0xd6, 0x1b, 0x71, 0x0f, 0x5b, 0x9c, 0x34, 0x7d, 0x05, 0x46,
	0x53, 0x5e, 0x55, 0x59, 0xb3, 0x91, 0xbf, 0x40, 0xdc, 0xc1, 0x86, 0xd0, 0x59, 0x35, 0x38, 0x2c,
	0xf3, 0x24, 0x2f, 0x73, 0xfd, 0x00, 0x6b, 0x89, 0x9f, 0x33, 0x46, 0x16, 0x02, 0x56, 0x46, 0xb1,
	0x4a, 0x2b, 0xec, 0x64, 0xba, 0xd7, 0x8f, 0xb0, 0x91, 0x18, 0x83, 0x9f, 0x22, 0xfe, 0xe5, 0x87,
	0x3c, 0xfd, 0xf3, 0xe5, 0x27, 0x00, 0x00, 0xff, 0xff, 0xd5, 0xd4, 0x78, 0x4b, 0x60, 0x01, 0x00,
	0x00,
}