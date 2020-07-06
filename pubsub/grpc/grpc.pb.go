// Code generated by protoc-gen-go. DO NOT EDIT.
// source: github.com/aperturerobotics/bifrost/pubsub/grpc/grpc.proto

package pubsub_grpc

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
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

// SubcribeRequest is a pubsub subscription request message.
type SubscribeRequest struct {
	// ChannelId is the channel id to subscribe to.
	// Must be sent before / with publish.
	// Cannot change the channel ID after first transmission.
	ChannelId string `protobuf:"bytes,1,opt,name=channel_id,json=channelId,proto3" json:"channel_id,omitempty"`
	// PublishRequest contains a publish message request.
	PublishRequest       *PublishRequest `protobuf:"bytes,2,opt,name=publish_request,json=publishRequest,proto3" json:"publish_request,omitempty"`
	XXX_NoUnkeyedLiteral struct{}        `json:"-"`
	XXX_unrecognized     []byte          `json:"-"`
	XXX_sizecache        int32           `json:"-"`
}

func (m *SubscribeRequest) Reset()         { *m = SubscribeRequest{} }
func (m *SubscribeRequest) String() string { return proto.CompactTextString(m) }
func (*SubscribeRequest) ProtoMessage()    {}
func (*SubscribeRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_d572159280ec2d84, []int{0}
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

func (m *SubscribeRequest) GetPublishRequest() *PublishRequest {
	if m != nil {
		return m.PublishRequest
	}
	return nil
}

// PublishRequest is a message published via the subscribe channel.
type PublishRequest struct {
	// FromPeerId is the peer identifier of the sender.
	// The peer ID will be used to acquire the peer private key and sign the mesasge.
	// The peer should be loaded with Identify in the peer service.
	FromPeerId string `protobuf:"bytes,1,opt,name=from_peer_id,json=fromPeerId,proto3" json:"from_peer_id,omitempty"`
	// Data is the published data.
	Data []byte `protobuf:"bytes,2,opt,name=data,proto3" json:"data,omitempty"`
	// Identifier is a uint32 identifier to use for outgoing status.
	// If zero, no outgoing status response will be sent.
	Identifier           uint32   `protobuf:"varint,3,opt,name=identifier,proto3" json:"identifier,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *PublishRequest) Reset()         { *m = PublishRequest{} }
func (m *PublishRequest) String() string { return proto.CompactTextString(m) }
func (*PublishRequest) ProtoMessage()    {}
func (*PublishRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_d572159280ec2d84, []int{1}
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

func (m *PublishRequest) GetFromPeerId() string {
	if m != nil {
		return m.FromPeerId
	}
	return ""
}

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
	return fileDescriptor_d572159280ec2d84, []int{2}
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
	return fileDescriptor_d572159280ec2d84, []int{3}
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
	return fileDescriptor_d572159280ec2d84, []int{4}
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
	return fileDescriptor_d572159280ec2d84, []int{5}
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
	proto.RegisterType((*SubscribeRequest)(nil), "pubsub.grpc.SubscribeRequest")
	proto.RegisterType((*PublishRequest)(nil), "pubsub.grpc.PublishRequest")
	proto.RegisterType((*SubscribeResponse)(nil), "pubsub.grpc.SubscribeResponse")
	proto.RegisterType((*SubscriptionStatus)(nil), "pubsub.grpc.SubscriptionStatus")
	proto.RegisterType((*OutgoingStatus)(nil), "pubsub.grpc.OutgoingStatus")
	proto.RegisterType((*IncomingMessage)(nil), "pubsub.grpc.IncomingMessage")
}

func init() {
	proto.RegisterFile("github.com/aperturerobotics/bifrost/pubsub/grpc/grpc.proto", fileDescriptor_d572159280ec2d84)
}

var fileDescriptor_d572159280ec2d84 = []byte{
	// 429 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x93, 0xcf, 0x8e, 0xd3, 0x30,
	0x10, 0xc6, 0x09, 0x45, 0x68, 0x3b, 0xdd, 0xb6, 0x8b, 0xb9, 0x54, 0xc0, 0x2e, 0x55, 0xc4, 0xa1,
	0xa7, 0x14, 0x15, 0x4e, 0x9c, 0x57, 0x42, 0x3d, 0x20, 0x22, 0xe7, 0x01, 0x2a, 0x3b, 0x99, 0x26,
	0x96, 0x36, 0xb1, 0xf1, 0xd8, 0xf0, 0xe2, 0x3c, 0x00, 0xca, 0x9f, 0xed, 0xc6, 0xdd, 0x45, 0xda,
	0x4b, 0x64, 0x7f, 0xf9, 0xf2, 0x8d, 0xfd, 0x9b, 0x09, 0x7c, 0x2b, 0x95, 0xab, 0xbc, 0x4c, 0x72,
	0x5d, 0x6f, 0x85, 0x41, 0xeb, 0xbc, 0x45, 0xab, 0xa5, 0x76, 0x2a, 0xa7, 0xad, 0x54, 0x47, 0xab,
	0xc9, 0x6d, 0x8d, 0x97, 0xe4, 0xe5, 0xb6, 0xb4, 0x26, 0xef, 0x1e, 0x89, 0xb1, 0xda, 0x69, 0x36,
	0xeb, 0xf5, 0xa4, 0x95, 0xe2, 0x3f, 0x70, 0x95, 0x79, 0x49, 0xb9, 0x55, 0x12, 0x39, 0xfe, 0xf2,
	0x48, 0x8e, 0x5d, 0x03, 0xe4, 0x95, 0x68, 0x1a, 0xbc, 0x3b, 0xa8, 0x62, 0x15, 0xad, 0xa3, 0xcd,
	0x94, 0x4f, 0x07, 0x65, 0x5f, 0xb0, 0x5b, 0x58, 0x1a, 0x2f, 0xef, 0x14, 0x55, 0x07, 0xdb, 0x7f,
	0xb1, 0x7a, 0xb9, 0x8e, 0x36, 0xb3, 0xdd, 0xfb, 0x64, 0x94, 0x9c, 0xa4, 0xbd, 0x67, 0x08, 0xe5,
	0x0b, 0x13, 0xec, 0xe3, 0x23, 0x2c, 0x42, 0x07, 0x5b, 0xc3, 0xe5, 0xd1, 0xea, 0xfa, 0x60, 0x10,
	0xed, 0x43, 0x61, 0x68, 0xb5, 0x14, 0xd1, 0xee, 0x0b, 0xc6, 0xe0, 0x55, 0x21, 0x9c, 0xe8, 0xca,
	0x5d, 0xf2, 0x6e, 0xcd, 0x6e, 0x00, 0x54, 0x81, 0x8d, 0x53, 0x47, 0x85, 0x76, 0x35, 0x59, 0x47,
	0x9b, 0x39, 0x1f, 0x29, 0xf1, 0xdf, 0x08, 0xde, 0x8c, 0x6e, 0x48, 0x46, 0x37, 0x84, 0xec, 0x3b,
	0x5c, 0xa9, 0x26, 0xd7, 0xb5, 0x6a, 0xca, 0x43, 0x8d, 0x44, 0xa2, 0xc4, 0xae, 0xde, 0x6c, 0xf7,
	0x21, 0xb8, 0xc4, 0x7e, 0x30, 0xfd, 0xe8, 0x3d, 0x7c, 0xa9, 0x42, 0xa1, 0x85, 0xa1, 0xbd, 0x2b,
	0x75, 0x1b, 0x44, 0x4e, 0x38, 0x4f, 0x4f, 0xc2, 0xf8, 0x39, 0x78, 0xb2, 0xce, 0xc2, 0x17, 0x3a,
	0xd8, 0xb3, 0x14, 0xde, 0x52, 0x7f, 0x46, 0xe3, 0x94, 0x6e, 0xee, 0x93, 0x26, 0x5d, 0xd2, 0xc7,
	0x20, 0x29, 0x1b, 0xf9, 0x86, 0x34, 0x46, 0x8f, 0xb4, 0xf8, 0x2b, 0xb0, 0xc7, 0xce, 0x16, 0x16,
	0xdd, 0xb3, 0xe8, 0x01, 0x5f, 0xf0, 0x91, 0x12, 0xdf, 0xc2, 0x22, 0x3c, 0xe9, 0x19, 0xde, 0xe8,
	0x1c, 0x6f, 0xdb, 0x12, 0xc2, 0xa6, 0x9f, 0x80, 0x0b, 0xde, 0xad, 0xe3, 0x1a, 0x96, 0x67, 0xdc,
	0x9e, 0xd1, 0xdb, 0x4f, 0x30, 0x17, 0xde, 0x55, 0x6d, 0x70, 0x2e, 0x1c, 0x16, 0x43, 0x62, 0x28,
	0x9e, 0x26, 0x60, 0xf2, 0x30, 0x01, 0x3b, 0x01, 0xf3, 0xd4, 0xcb, 0xcc, 0xcb, 0x0c, 0xed, 0x6f,
	0x95, 0x23, 0x4b, 0x61, 0x7a, 0xea, 0x38, 0xbb, 0x7e, 0x8a, 0xde, 0x69, 0xd6, 0xdf, 0xdd, 0xfc,
	0xef, 0x75, 0x3f, 0x28, 0xf1, 0x8b, 0x4d, 0xf4, 0x39, 0x92, 0xaf, 0xbb, 0x3f, 0xe7, 0xcb, 0xbf,
	0x00, 0x00, 0x00, 0xff, 0xff, 0x68, 0x70, 0x4b, 0xae, 0x77, 0x03, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// PubSubServiceClient is the client API for PubSubService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type PubSubServiceClient interface {
	// Subscribe subscribes to a channel, allowing the subscriber to publish
	// messages over the same channel.
	Subscribe(ctx context.Context, opts ...grpc.CallOption) (PubSubService_SubscribeClient, error)
}

type pubSubServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewPubSubServiceClient(cc grpc.ClientConnInterface) PubSubServiceClient {
	return &pubSubServiceClient{cc}
}

func (c *pubSubServiceClient) Subscribe(ctx context.Context, opts ...grpc.CallOption) (PubSubService_SubscribeClient, error) {
	stream, err := c.cc.NewStream(ctx, &_PubSubService_serviceDesc.Streams[0], "/pubsub.grpc.PubSubService/Subscribe", opts...)
	if err != nil {
		return nil, err
	}
	x := &pubSubServiceSubscribeClient{stream}
	return x, nil
}

type PubSubService_SubscribeClient interface {
	Send(*SubscribeRequest) error
	Recv() (*SubscribeResponse, error)
	grpc.ClientStream
}

type pubSubServiceSubscribeClient struct {
	grpc.ClientStream
}

func (x *pubSubServiceSubscribeClient) Send(m *SubscribeRequest) error {
	return x.ClientStream.SendMsg(m)
}

func (x *pubSubServiceSubscribeClient) Recv() (*SubscribeResponse, error) {
	m := new(SubscribeResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// PubSubServiceServer is the server API for PubSubService service.
type PubSubServiceServer interface {
	// Subscribe subscribes to a channel, allowing the subscriber to publish
	// messages over the same channel.
	Subscribe(PubSubService_SubscribeServer) error
}

// UnimplementedPubSubServiceServer can be embedded to have forward compatible implementations.
type UnimplementedPubSubServiceServer struct {
}

func (*UnimplementedPubSubServiceServer) Subscribe(srv PubSubService_SubscribeServer) error {
	return status.Errorf(codes.Unimplemented, "method Subscribe not implemented")
}

func RegisterPubSubServiceServer(s *grpc.Server, srv PubSubServiceServer) {
	s.RegisterService(&_PubSubService_serviceDesc, srv)
}

func _PubSubService_Subscribe_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(PubSubServiceServer).Subscribe(&pubSubServiceSubscribeServer{stream})
}

type PubSubService_SubscribeServer interface {
	Send(*SubscribeResponse) error
	Recv() (*SubscribeRequest, error)
	grpc.ServerStream
}

type pubSubServiceSubscribeServer struct {
	grpc.ServerStream
}

func (x *pubSubServiceSubscribeServer) Send(m *SubscribeResponse) error {
	return x.ServerStream.SendMsg(m)
}

func (x *pubSubServiceSubscribeServer) Recv() (*SubscribeRequest, error) {
	m := new(SubscribeRequest)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

var _PubSubService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "pubsub.grpc.PubSubService",
	HandlerType: (*PubSubServiceServer)(nil),
	Methods:     []grpc.MethodDesc{},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "Subscribe",
			Handler:       _PubSubService_Subscribe_Handler,
			ServerStreams: true,
			ClientStreams: true,
		},
	},
	Metadata: "github.com/aperturerobotics/bifrost/pubsub/grpc/grpc.proto",
}
