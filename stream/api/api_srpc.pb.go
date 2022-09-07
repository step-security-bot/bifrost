// Code generated by protoc-gen-srpc. DO NOT EDIT.
// protoc-gen-srpc version: v0.10.0
// source: github.com/aperturerobotics/bifrost/stream/api/api.proto

package stream_api

import (
	context "context"

	srpc "github.com/aperturerobotics/starpc/srpc"
)

type SRPCStreamServiceClient interface {
	SRPCClient() srpc.Client

	ForwardStreams(ctx context.Context, in *ForwardStreamsRequest) (SRPCStreamService_ForwardStreamsClient, error)
	ListenStreams(ctx context.Context, in *ListenStreamsRequest) (SRPCStreamService_ListenStreamsClient, error)
	AcceptStream(ctx context.Context) (SRPCStreamService_AcceptStreamClient, error)
	DialStream(ctx context.Context) (SRPCStreamService_DialStreamClient, error)
}

type srpcStreamServiceClient struct {
	cc srpc.Client
}

func NewSRPCStreamServiceClient(cc srpc.Client) SRPCStreamServiceClient {
	return &srpcStreamServiceClient{cc}
}

func (c *srpcStreamServiceClient) SRPCClient() srpc.Client { return c.cc }

func (c *srpcStreamServiceClient) ForwardStreams(ctx context.Context, in *ForwardStreamsRequest) (SRPCStreamService_ForwardStreamsClient, error) {
	stream, err := c.cc.NewStream(ctx, "stream.api.StreamService", "ForwardStreams", in)
	if err != nil {
		return nil, err
	}
	strm := &srpcStreamService_ForwardStreamsClient{stream}
	if err := strm.CloseSend(); err != nil {
		return nil, err
	}
	return strm, nil
}

type SRPCStreamService_ForwardStreamsClient interface {
	srpc.Stream
	Recv() (*ForwardStreamsResponse, error)
	RecvTo(*ForwardStreamsResponse) error
}

type srpcStreamService_ForwardStreamsClient struct {
	srpc.Stream
}

func (x *srpcStreamService_ForwardStreamsClient) Recv() (*ForwardStreamsResponse, error) {
	m := new(ForwardStreamsResponse)
	if err := x.MsgRecv(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (x *srpcStreamService_ForwardStreamsClient) RecvTo(m *ForwardStreamsResponse) error {
	return x.MsgRecv(m)
}

func (c *srpcStreamServiceClient) ListenStreams(ctx context.Context, in *ListenStreamsRequest) (SRPCStreamService_ListenStreamsClient, error) {
	stream, err := c.cc.NewStream(ctx, "stream.api.StreamService", "ListenStreams", in)
	if err != nil {
		return nil, err
	}
	strm := &srpcStreamService_ListenStreamsClient{stream}
	if err := strm.CloseSend(); err != nil {
		return nil, err
	}
	return strm, nil
}

type SRPCStreamService_ListenStreamsClient interface {
	srpc.Stream
	Recv() (*ListenStreamsResponse, error)
	RecvTo(*ListenStreamsResponse) error
}

type srpcStreamService_ListenStreamsClient struct {
	srpc.Stream
}

func (x *srpcStreamService_ListenStreamsClient) Recv() (*ListenStreamsResponse, error) {
	m := new(ListenStreamsResponse)
	if err := x.MsgRecv(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (x *srpcStreamService_ListenStreamsClient) RecvTo(m *ListenStreamsResponse) error {
	return x.MsgRecv(m)
}

func (c *srpcStreamServiceClient) AcceptStream(ctx context.Context) (SRPCStreamService_AcceptStreamClient, error) {
	stream, err := c.cc.NewStream(ctx, "stream.api.StreamService", "AcceptStream", nil)
	if err != nil {
		return nil, err
	}
	strm := &srpcStreamService_AcceptStreamClient{stream}
	return strm, nil
}

type SRPCStreamService_AcceptStreamClient interface {
	srpc.Stream
	Send(*AcceptStreamRequest) error
	Recv() (*AcceptStreamResponse, error)
	RecvTo(*AcceptStreamResponse) error
}

type srpcStreamService_AcceptStreamClient struct {
	srpc.Stream
}

func (x *srpcStreamService_AcceptStreamClient) Send(m *AcceptStreamRequest) error {
	return x.MsgSend(m)
}

func (x *srpcStreamService_AcceptStreamClient) Recv() (*AcceptStreamResponse, error) {
	m := new(AcceptStreamResponse)
	if err := x.MsgRecv(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (x *srpcStreamService_AcceptStreamClient) RecvTo(m *AcceptStreamResponse) error {
	return x.MsgRecv(m)
}

func (c *srpcStreamServiceClient) DialStream(ctx context.Context) (SRPCStreamService_DialStreamClient, error) {
	stream, err := c.cc.NewStream(ctx, "stream.api.StreamService", "DialStream", nil)
	if err != nil {
		return nil, err
	}
	strm := &srpcStreamService_DialStreamClient{stream}
	return strm, nil
}

type SRPCStreamService_DialStreamClient interface {
	srpc.Stream
	Send(*DialStreamRequest) error
	Recv() (*DialStreamResponse, error)
	RecvTo(*DialStreamResponse) error
}

type srpcStreamService_DialStreamClient struct {
	srpc.Stream
}

func (x *srpcStreamService_DialStreamClient) Send(m *DialStreamRequest) error {
	return x.MsgSend(m)
}

func (x *srpcStreamService_DialStreamClient) Recv() (*DialStreamResponse, error) {
	m := new(DialStreamResponse)
	if err := x.MsgRecv(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (x *srpcStreamService_DialStreamClient) RecvTo(m *DialStreamResponse) error {
	return x.MsgRecv(m)
}

type SRPCStreamServiceServer interface {
	ForwardStreams(*ForwardStreamsRequest, SRPCStreamService_ForwardStreamsStream) error
	ListenStreams(*ListenStreamsRequest, SRPCStreamService_ListenStreamsStream) error
	AcceptStream(SRPCStreamService_AcceptStreamStream) error
	DialStream(SRPCStreamService_DialStreamStream) error
}

type SRPCStreamServiceUnimplementedServer struct{}

func (s *SRPCStreamServiceUnimplementedServer) ForwardStreams(*ForwardStreamsRequest, SRPCStreamService_ForwardStreamsStream) error {
	return srpc.ErrUnimplemented
}

func (s *SRPCStreamServiceUnimplementedServer) ListenStreams(*ListenStreamsRequest, SRPCStreamService_ListenStreamsStream) error {
	return srpc.ErrUnimplemented
}

func (s *SRPCStreamServiceUnimplementedServer) AcceptStream(SRPCStreamService_AcceptStreamStream) error {
	return srpc.ErrUnimplemented
}

func (s *SRPCStreamServiceUnimplementedServer) DialStream(SRPCStreamService_DialStreamStream) error {
	return srpc.ErrUnimplemented
}

const SRPCStreamServiceServiceID = "stream.api.StreamService"

type SRPCStreamServiceHandler struct {
	impl SRPCStreamServiceServer
}

func (SRPCStreamServiceHandler) GetServiceID() string { return SRPCStreamServiceServiceID }

func (SRPCStreamServiceHandler) GetMethodIDs() []string {
	return []string{
		"ForwardStreams",
		"ListenStreams",
		"AcceptStream",
		"DialStream",
	}
}

func (d *SRPCStreamServiceHandler) InvokeMethod(
	serviceID, methodID string,
	strm srpc.Stream,
) (bool, error) {
	if serviceID != "" && serviceID != d.GetServiceID() {
		return false, nil
	}

	switch methodID {
	case "ForwardStreams":
		return true, d.InvokeMethod_ForwardStreams(d.impl, strm)
	case "ListenStreams":
		return true, d.InvokeMethod_ListenStreams(d.impl, strm)
	case "AcceptStream":
		return true, d.InvokeMethod_AcceptStream(d.impl, strm)
	case "DialStream":
		return true, d.InvokeMethod_DialStream(d.impl, strm)
	default:
		return false, nil
	}
}

func (SRPCStreamServiceHandler) InvokeMethod_ForwardStreams(impl SRPCStreamServiceServer, strm srpc.Stream) error {
	req := new(ForwardStreamsRequest)
	if err := strm.MsgRecv(req); err != nil {
		return err
	}
	serverStrm := &srpcStreamService_ForwardStreamsStream{strm}
	return impl.ForwardStreams(req, serverStrm)
}

func (SRPCStreamServiceHandler) InvokeMethod_ListenStreams(impl SRPCStreamServiceServer, strm srpc.Stream) error {
	req := new(ListenStreamsRequest)
	if err := strm.MsgRecv(req); err != nil {
		return err
	}
	serverStrm := &srpcStreamService_ListenStreamsStream{strm}
	return impl.ListenStreams(req, serverStrm)
}

func (SRPCStreamServiceHandler) InvokeMethod_AcceptStream(impl SRPCStreamServiceServer, strm srpc.Stream) error {
	clientStrm := &srpcStreamService_AcceptStreamStream{strm}
	return impl.AcceptStream(clientStrm)
}

func (SRPCStreamServiceHandler) InvokeMethod_DialStream(impl SRPCStreamServiceServer, strm srpc.Stream) error {
	clientStrm := &srpcStreamService_DialStreamStream{strm}
	return impl.DialStream(clientStrm)
}

func SRPCRegisterStreamService(mux srpc.Mux, impl SRPCStreamServiceServer) error {
	return mux.Register(&SRPCStreamServiceHandler{impl: impl})
}

type SRPCStreamService_ForwardStreamsStream interface {
	srpc.Stream
	Send(*ForwardStreamsResponse) error
}

type srpcStreamService_ForwardStreamsStream struct {
	srpc.Stream
}

func (x *srpcStreamService_ForwardStreamsStream) Send(m *ForwardStreamsResponse) error {
	return x.MsgSend(m)
}

type SRPCStreamService_ListenStreamsStream interface {
	srpc.Stream
	Send(*ListenStreamsResponse) error
}

type srpcStreamService_ListenStreamsStream struct {
	srpc.Stream
}

func (x *srpcStreamService_ListenStreamsStream) Send(m *ListenStreamsResponse) error {
	return x.MsgSend(m)
}

type SRPCStreamService_AcceptStreamStream interface {
	srpc.Stream
	Send(*AcceptStreamResponse) error
	Recv() (*AcceptStreamRequest, error)
}

type srpcStreamService_AcceptStreamStream struct {
	srpc.Stream
}

func (x *srpcStreamService_AcceptStreamStream) Send(m *AcceptStreamResponse) error {
	return x.MsgSend(m)
}

func (x *srpcStreamService_AcceptStreamStream) Recv() (*AcceptStreamRequest, error) {
	m := new(AcceptStreamRequest)
	if err := x.MsgRecv(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (x *srpcStreamService_AcceptStreamStream) RecvTo(m *AcceptStreamRequest) error {
	return x.MsgRecv(m)
}

type SRPCStreamService_DialStreamStream interface {
	srpc.Stream
	Send(*DialStreamResponse) error
	Recv() (*DialStreamRequest, error)
}

type srpcStreamService_DialStreamStream struct {
	srpc.Stream
}

func (x *srpcStreamService_DialStreamStream) Send(m *DialStreamResponse) error {
	return x.MsgSend(m)
}

func (x *srpcStreamService_DialStreamStream) Recv() (*DialStreamRequest, error) {
	m := new(DialStreamRequest)
	if err := x.MsgRecv(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (x *srpcStreamService_DialStreamStream) RecvTo(m *DialStreamRequest) error {
	return x.MsgRecv(m)
}
