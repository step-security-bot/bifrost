// Code generated by protoc-gen-go. DO NOT EDIT.
// source: github.com/aperturerobotics/bifrost/pubsub/api/api.proto

package pubsub_api

import (
	fmt "fmt"
	math "math"

	proto "github.com/golang/protobuf/proto"
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

// SubcribeRequest is a pubsub subscription request message.
type SubscribeRequest struct {
	// ChannelId is the channel id to subscribe to.
	// Must be sent before / with publish.
	// Cannot change the channel ID after first transmission.
	ChannelId string `protobuf:"bytes,1,opt,name=channel_id,json=channelId,proto3" json:"channel_id,omitempty"`
	// PeerId is the peer identifier of the publisher/subscriber.
	// The peer ID will be used to acquire the peer private key.
	PeerId string `protobuf:"bytes,2,opt,name=peer_id,json=peerId,proto3" json:"peer_id,omitempty"`
	// PrivKeyPem is an alternate to PeerId, specify private key inline.
	// Overrides PeerId if set.
	PrivKeyPem string `protobuf:"bytes,3,opt,name=priv_key_pem,json=privKeyPem,proto3" json:"priv_key_pem,omitempty"`
	// PublishRequest contains a publish message request.
	PublishRequest       *PublishRequest `protobuf:"bytes,4,opt,name=publish_request,json=publishRequest,proto3" json:"publish_request,omitempty"`
	XXX_NoUnkeyedLiteral struct{}        `json:"-"`
	XXX_unrecognized     []byte          `json:"-"`
	XXX_sizecache        int32           `json:"-"`
}

func (m *SubscribeRequest) Reset()         { *m = SubscribeRequest{} }
func (m *SubscribeRequest) String() string { return proto.CompactTextString(m) }
func (*SubscribeRequest) ProtoMessage()    {}
func (*SubscribeRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_aee1bde3e41e7af3, []int{0}
}

func (m *SubscribeRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SubscribeRequest.Unmarshal(m, b)
}
func (m *SubscribeRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SubscribeRequest.Marshal(b, m, deterministic)
}
func (m *SubscribeRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SubscribeRequest.Merge(m, src)
}
func (m *SubscribeRequest) XXX_Size() int {
	return xxx_messageInfo_SubscribeRequest.Size(m)
}
func (m *SubscribeRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_SubscribeRequest.DiscardUnknown(m)
}

var xxx_messageInfo_SubscribeRequest proto.InternalMessageInfo

func (m *SubscribeRequest) GetChannelId() string {
	if m != nil {
		return m.ChannelId
	}
	return ""
}

func (m *SubscribeRequest) GetPeerId() string {
	if m != nil {
		return m.PeerId
	}
	return ""
}

func (m *SubscribeRequest) GetPrivKeyPem() string {
	if m != nil {
		return m.PrivKeyPem
	}
	return ""
}

func (m *SubscribeRequest) GetPublishRequest() *PublishRequest {
	if m != nil {
		return m.PublishRequest
	}
	return nil
}

// PublishRequest is a message published via the subscribe channel.
type PublishRequest struct {
	// Data is the published data.
	Data []byte `protobuf:"bytes,1,opt,name=data,proto3" json:"data,omitempty"`
	// Identifier is a uint32 identifier to use for outgoing status.
	// If zero, no outgoing status response will be sent.
	Identifier           uint32   `protobuf:"varint,2,opt,name=identifier,proto3" json:"identifier,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *PublishRequest) Reset()         { *m = PublishRequest{} }
func (m *PublishRequest) String() string { return proto.CompactTextString(m) }
func (*PublishRequest) ProtoMessage()    {}
func (*PublishRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_aee1bde3e41e7af3, []int{1}
}

func (m *PublishRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PublishRequest.Unmarshal(m, b)
}
func (m *PublishRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PublishRequest.Marshal(b, m, deterministic)
}
func (m *PublishRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PublishRequest.Merge(m, src)
}
func (m *PublishRequest) XXX_Size() int {
	return xxx_messageInfo_PublishRequest.Size(m)
}
func (m *PublishRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_PublishRequest.DiscardUnknown(m)
}

var xxx_messageInfo_PublishRequest proto.InternalMessageInfo

func (m *PublishRequest) GetData() []byte {
	if m != nil {
		return m.Data
	}
	return nil
}

func (m *PublishRequest) GetIdentifier() uint32 {
	if m != nil {
		return m.Identifier
	}
	return 0
}

// SubcribeResponse is a pubsub subscription response message.
type SubscribeResponse struct {
	// IncomingMessage is an incoming message.
	IncomingMessage *IncomingMessage `protobuf:"bytes,1,opt,name=incoming_message,json=incomingMessage,proto3" json:"incoming_message,omitempty"`
	// OutgoingStatus is status of an outgoing message.
	// Sent when a Publish request finishes.
	OutgoingStatus *OutgoingStatus `protobuf:"bytes,2,opt,name=outgoing_status,json=outgoingStatus,proto3" json:"outgoing_status,omitempty"`
	// SubscriptionStatus is the status of the subscription
	SubscriptionStatus   *SubscriptionStatus `protobuf:"bytes,3,opt,name=subscription_status,json=subscriptionStatus,proto3" json:"subscription_status,omitempty"`
	XXX_NoUnkeyedLiteral struct{}            `json:"-"`
	XXX_unrecognized     []byte              `json:"-"`
	XXX_sizecache        int32               `json:"-"`
}

func (m *SubscribeResponse) Reset()         { *m = SubscribeResponse{} }
func (m *SubscribeResponse) String() string { return proto.CompactTextString(m) }
func (*SubscribeResponse) ProtoMessage()    {}
func (*SubscribeResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_aee1bde3e41e7af3, []int{2}
}

func (m *SubscribeResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SubscribeResponse.Unmarshal(m, b)
}
func (m *SubscribeResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SubscribeResponse.Marshal(b, m, deterministic)
}
func (m *SubscribeResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SubscribeResponse.Merge(m, src)
}
func (m *SubscribeResponse) XXX_Size() int {
	return xxx_messageInfo_SubscribeResponse.Size(m)
}
func (m *SubscribeResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_SubscribeResponse.DiscardUnknown(m)
}

var xxx_messageInfo_SubscribeResponse proto.InternalMessageInfo

func (m *SubscribeResponse) GetIncomingMessage() *IncomingMessage {
	if m != nil {
		return m.IncomingMessage
	}
	return nil
}

func (m *SubscribeResponse) GetOutgoingStatus() *OutgoingStatus {
	if m != nil {
		return m.OutgoingStatus
	}
	return nil
}

func (m *SubscribeResponse) GetSubscriptionStatus() *SubscriptionStatus {
	if m != nil {
		return m.SubscriptionStatus
	}
	return nil
}

// SubscripionStatus is the status of the subscription handle.
type SubscriptionStatus struct {
	// Subscribed indicates the subscription is established.
	Subscribed           bool     `protobuf:"varint,1,opt,name=subscribed,proto3" json:"subscribed,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *SubscriptionStatus) Reset()         { *m = SubscriptionStatus{} }
func (m *SubscriptionStatus) String() string { return proto.CompactTextString(m) }
func (*SubscriptionStatus) ProtoMessage()    {}
func (*SubscriptionStatus) Descriptor() ([]byte, []int) {
	return fileDescriptor_aee1bde3e41e7af3, []int{3}
}

func (m *SubscriptionStatus) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SubscriptionStatus.Unmarshal(m, b)
}
func (m *SubscriptionStatus) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SubscriptionStatus.Marshal(b, m, deterministic)
}
func (m *SubscriptionStatus) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SubscriptionStatus.Merge(m, src)
}
func (m *SubscriptionStatus) XXX_Size() int {
	return xxx_messageInfo_SubscriptionStatus.Size(m)
}
func (m *SubscriptionStatus) XXX_DiscardUnknown() {
	xxx_messageInfo_SubscriptionStatus.DiscardUnknown(m)
}

var xxx_messageInfo_SubscriptionStatus proto.InternalMessageInfo

func (m *SubscriptionStatus) GetSubscribed() bool {
	if m != nil {
		return m.Subscribed
	}
	return false
}

// OutgoingStatus is status of an outgoing message.
type OutgoingStatus struct {
	// Identifier is the request-provided identifier for the message.
	Identifier uint32 `protobuf:"varint,1,opt,name=identifier,proto3" json:"identifier,omitempty"`
	// Sent indicates if the message was sent.
	Sent                 bool     `protobuf:"varint,2,opt,name=sent,proto3" json:"sent,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *OutgoingStatus) Reset()         { *m = OutgoingStatus{} }
func (m *OutgoingStatus) String() string { return proto.CompactTextString(m) }
func (*OutgoingStatus) ProtoMessage()    {}
func (*OutgoingStatus) Descriptor() ([]byte, []int) {
	return fileDescriptor_aee1bde3e41e7af3, []int{4}
}

func (m *OutgoingStatus) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_OutgoingStatus.Unmarshal(m, b)
}
func (m *OutgoingStatus) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_OutgoingStatus.Marshal(b, m, deterministic)
}
func (m *OutgoingStatus) XXX_Merge(src proto.Message) {
	xxx_messageInfo_OutgoingStatus.Merge(m, src)
}
func (m *OutgoingStatus) XXX_Size() int {
	return xxx_messageInfo_OutgoingStatus.Size(m)
}
func (m *OutgoingStatus) XXX_DiscardUnknown() {
	xxx_messageInfo_OutgoingStatus.DiscardUnknown(m)
}

var xxx_messageInfo_OutgoingStatus proto.InternalMessageInfo

func (m *OutgoingStatus) GetIdentifier() uint32 {
	if m != nil {
		return m.Identifier
	}
	return 0
}

func (m *OutgoingStatus) GetSent() bool {
	if m != nil {
		return m.Sent
	}
	return false
}

// IncomingMessage implements Message with a proto object.
type IncomingMessage struct {
	// FromPeerId is the peer identifier of the sender.
	FromPeerId string `protobuf:"bytes,1,opt,name=from_peer_id,json=fromPeerId,proto3" json:"from_peer_id,omitempty"`
	// Authenticated indicates if the message is verified to be from the sender.
	Authenticated bool `protobuf:"varint,2,opt,name=authenticated,proto3" json:"authenticated,omitempty"`
	// Data is the inner data.
	Data                 []byte   `protobuf:"bytes,3,opt,name=data,proto3" json:"data,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *IncomingMessage) Reset()         { *m = IncomingMessage{} }
func (m *IncomingMessage) String() string { return proto.CompactTextString(m) }
func (*IncomingMessage) ProtoMessage()    {}
func (*IncomingMessage) Descriptor() ([]byte, []int) {
	return fileDescriptor_aee1bde3e41e7af3, []int{5}
}

func (m *IncomingMessage) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_IncomingMessage.Unmarshal(m, b)
}
func (m *IncomingMessage) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_IncomingMessage.Marshal(b, m, deterministic)
}
func (m *IncomingMessage) XXX_Merge(src proto.Message) {
	xxx_messageInfo_IncomingMessage.Merge(m, src)
}
func (m *IncomingMessage) XXX_Size() int {
	return xxx_messageInfo_IncomingMessage.Size(m)
}
func (m *IncomingMessage) XXX_DiscardUnknown() {
	xxx_messageInfo_IncomingMessage.DiscardUnknown(m)
}

var xxx_messageInfo_IncomingMessage proto.InternalMessageInfo

func (m *IncomingMessage) GetFromPeerId() string {
	if m != nil {
		return m.FromPeerId
	}
	return ""
}

func (m *IncomingMessage) GetAuthenticated() bool {
	if m != nil {
		return m.Authenticated
	}
	return false
}

func (m *IncomingMessage) GetData() []byte {
	if m != nil {
		return m.Data
	}
	return nil
}

func init() {
	proto.RegisterType((*SubscribeRequest)(nil), "pubsub.api.SubscribeRequest")
	proto.RegisterType((*PublishRequest)(nil), "pubsub.api.PublishRequest")
	proto.RegisterType((*SubscribeResponse)(nil), "pubsub.api.SubscribeResponse")
	proto.RegisterType((*SubscriptionStatus)(nil), "pubsub.api.SubscriptionStatus")
	proto.RegisterType((*OutgoingStatus)(nil), "pubsub.api.OutgoingStatus")
	proto.RegisterType((*IncomingMessage)(nil), "pubsub.api.IncomingMessage")
}

func init() {
	proto.RegisterFile("github.com/aperturerobotics/bifrost/pubsub/api/api.proto", fileDescriptor_aee1bde3e41e7af3)
}

var fileDescriptor_aee1bde3e41e7af3 = []byte{
	// 461 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x74, 0x53, 0x4d, 0x6f, 0xd3, 0x40,
	0x10, 0xc5, 0xa4, 0x2a, 0xcd, 0xb4, 0x49, 0xca, 0x72, 0x20, 0x2a, 0xb4, 0x8a, 0x2c, 0x0e, 0x39,
	0x39, 0x28, 0x70, 0xe0, 0x0e, 0x42, 0x8a, 0x10, 0x34, 0x5a, 0xff, 0x00, 0x6b, 0xd7, 0x9e, 0x24,
	0x2b, 0x6a, 0xef, 0xb2, 0x1f, 0x95, 0xfa, 0xb3, 0xf8, 0x69, 0xfc, 0x03, 0xb4, 0xbb, 0x4e, 0x6a,
	0xa7, 0x70, 0x88, 0xb4, 0xf3, 0xe6, 0xe9, 0xe5, 0xcd, 0x9b, 0x31, 0x7c, 0xda, 0x0a, 0xbb, 0x73,
	0x3c, 0x2b, 0x65, 0xbd, 0x60, 0x0a, 0xb5, 0x75, 0x1a, 0xb5, 0xe4, 0xd2, 0x8a, 0xd2, 0x2c, 0xb8,
	0xd8, 0x68, 0x69, 0xec, 0x42, 0x39, 0x6e, 0x1c, 0x5f, 0x30, 0x25, 0xfc, 0x2f, 0x53, 0x5a, 0x5a,
	0x49, 0x20, 0xa2, 0x19, 0x53, 0x22, 0xfd, 0x9d, 0xc0, 0x65, 0xee, 0xb8, 0x29, 0xb5, 0xe0, 0x48,
	0xf1, 0x97, 0x43, 0x63, 0xc9, 0x35, 0x40, 0xb9, 0x63, 0x4d, 0x83, 0x77, 0x85, 0xa8, 0xa6, 0xc9,
	0x2c, 0x99, 0x0f, 0xe9, 0xb0, 0x45, 0x56, 0x15, 0x79, 0x0d, 0x2f, 0x14, 0xa2, 0xf6, 0xbd, 0xe7,
	0xa1, 0x77, 0xea, 0xcb, 0x55, 0x45, 0x66, 0x70, 0xa1, 0xb4, 0xb8, 0x2f, 0x7e, 0xe2, 0x43, 0xa1,
	0xb0, 0x9e, 0x0e, 0x42, 0x17, 0x3c, 0xf6, 0x0d, 0x1f, 0xd6, 0x58, 0x93, 0xcf, 0x30, 0x51, 0x8e,
	0xdf, 0x09, 0xb3, 0x2b, 0x74, 0xfc, 0xb3, 0xe9, 0xc9, 0x2c, 0x99, 0x9f, 0x2f, 0xaf, 0xb2, 0x47,
	0x53, 0xd9, 0x3a, 0x52, 0x5a, 0x3b, 0x74, 0xac, 0x7a, 0x75, 0xfa, 0x05, 0xc6, 0x7d, 0x06, 0x21,
	0x70, 0x52, 0x31, 0xcb, 0x82, 0xd5, 0x0b, 0x1a, 0xde, 0xe4, 0x06, 0x40, 0x54, 0xd8, 0x58, 0xb1,
	0x11, 0xa8, 0x83, 0xd1, 0x11, 0xed, 0x20, 0xe9, 0x9f, 0x04, 0x5e, 0x76, 0x26, 0x37, 0x4a, 0x36,
	0x06, 0xc9, 0x57, 0xb8, 0x14, 0x4d, 0x29, 0x6b, 0xd1, 0x6c, 0x8b, 0x1a, 0x8d, 0x61, 0x5b, 0x0c,
	0xaa, 0xe7, 0xcb, 0x37, 0x5d, 0x87, 0xab, 0x96, 0xf3, 0x3d, 0x52, 0xe8, 0x44, 0xf4, 0x01, 0x3f,
	0xa8, 0x74, 0x76, 0x2b, 0xbd, 0x8e, 0xb1, 0xcc, 0x3a, 0x13, 0x2c, 0x1c, 0x0d, 0x7a, 0xdb, 0x52,
	0xf2, 0xc0, 0xa0, 0x63, 0xd9, 0xab, 0xc9, 0x2d, 0xbc, 0x32, 0xd1, 0xa1, 0xb2, 0x42, 0x36, 0x7b,
	0xa1, 0x41, 0x10, 0xba, 0xe9, 0x0a, 0xe5, 0x1d, 0x5a, 0x2b, 0x46, 0xcc, 0x13, 0x2c, 0xfd, 0x08,
	0xe4, 0x29, 0xd3, 0x27, 0x65, 0xf6, 0x41, 0xc4, 0x75, 0x9f, 0xd1, 0x0e, 0xe2, 0xf3, 0xee, 0x1b,
	0x3d, 0xca, 0x36, 0x39, 0xce, 0xd6, 0xef, 0xc3, 0x60, 0x63, 0xc3, 0xc8, 0x67, 0x34, 0xbc, 0xd3,
	0x1a, 0x26, 0x47, 0xa9, 0xf9, 0x7b, 0xd9, 0x68, 0x59, 0x17, 0xfb, 0x6b, 0x8a, 0x97, 0x06, 0x1e,
	0x5b, 0xc7, 0x8b, 0x7a, 0x07, 0x23, 0xe6, 0xec, 0xce, 0x0b, 0x97, 0xcc, 0x62, 0xd5, 0x2a, 0xf6,
	0xc1, 0xc3, 0xfa, 0x07, 0x8f, 0xeb, 0x5f, 0x16, 0x30, 0x5a, 0x3b, 0x9e, 0x3b, 0x9e, 0xa3, 0xbe,
	0x17, 0x25, 0x92, 0x1f, 0x30, 0x3c, 0xac, 0x9b, 0xbc, 0xfd, 0x47, 0x78, 0x87, 0xfb, 0xbf, 0xba,
	0xfe, 0x4f, 0x37, 0xde, 0x48, 0xfa, 0x6c, 0x9e, 0xbc, 0x4f, 0xf8, 0x69, 0xf8, 0x98, 0x3e, 0xfc,
	0x0d, 0x00, 0x00, 0xff, 0xff, 0x64, 0x77, 0x7a, 0xe8, 0x88, 0x03, 0x00, 0x00,
}