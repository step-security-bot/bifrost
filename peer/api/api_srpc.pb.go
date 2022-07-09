// Code generated by protoc-gen-srpc. DO NOT EDIT.
// protoc-gen-srpc version: v0.8.1
// source: github.com/aperturerobotics/bifrost/peer/api/api.proto

package peer_api

import (
	context "context"

	srpc "github.com/aperturerobotics/starpc/srpc"
)

type SRPCPeerServiceClient interface {
	SRPCClient() srpc.Client

	Identify(ctx context.Context, in *IdentifyRequest) (SRPCPeerService_IdentifyClient, error)
	GetPeerInfo(ctx context.Context, in *GetPeerInfoRequest) (*GetPeerInfoResponse, error)
}

type srpcPeerServiceClient struct {
	cc srpc.Client
}

func NewSRPCPeerServiceClient(cc srpc.Client) SRPCPeerServiceClient {
	return &srpcPeerServiceClient{cc}
}

func (c *srpcPeerServiceClient) SRPCClient() srpc.Client { return c.cc }

func (c *srpcPeerServiceClient) Identify(ctx context.Context, in *IdentifyRequest) (SRPCPeerService_IdentifyClient, error) {
	stream, err := c.cc.NewStream(ctx, "peer.api.PeerService", "Identify", in)
	if err != nil {
		return nil, err
	}
	strm := &srpcPeerService_IdentifyClient{stream}
	if err := strm.CloseSend(); err != nil {
		return nil, err
	}
	return strm, nil
}

type SRPCPeerService_IdentifyClient interface {
	srpc.Stream
	Recv() (*IdentifyResponse, error)
	RecvTo(*IdentifyResponse) error
}

type srpcPeerService_IdentifyClient struct {
	srpc.Stream
}

func (x *srpcPeerService_IdentifyClient) Recv() (*IdentifyResponse, error) {
	m := new(IdentifyResponse)
	if err := x.MsgRecv(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (x *srpcPeerService_IdentifyClient) RecvTo(m *IdentifyResponse) error {
	return x.MsgRecv(m)
}

func (c *srpcPeerServiceClient) GetPeerInfo(ctx context.Context, in *GetPeerInfoRequest) (*GetPeerInfoResponse, error) {
	out := new(GetPeerInfoResponse)
	err := c.cc.Invoke(ctx, "peer.api.PeerService", "GetPeerInfo", in, out)
	if err != nil {
		return nil, err
	}
	return out, nil
}

type SRPCPeerServiceServer interface {
	Identify(*IdentifyRequest, SRPCPeerService_IdentifyStream) error
	GetPeerInfo(context.Context, *GetPeerInfoRequest) (*GetPeerInfoResponse, error)
}

type SRPCPeerServiceUnimplementedServer struct{}

func (s *SRPCPeerServiceUnimplementedServer) Identify(*IdentifyRequest, SRPCPeerService_IdentifyStream) error {
	return srpc.ErrUnimplemented
}

func (s *SRPCPeerServiceUnimplementedServer) GetPeerInfo(context.Context, *GetPeerInfoRequest) (*GetPeerInfoResponse, error) {
	return nil, srpc.ErrUnimplemented
}

const SRPCPeerServiceServiceID = "peer.api.PeerService"

type SRPCPeerServiceHandler struct {
	impl SRPCPeerServiceServer
}

func (SRPCPeerServiceHandler) GetServiceID() string { return SRPCPeerServiceServiceID }

func (SRPCPeerServiceHandler) GetMethodIDs() []string {
	return []string{
		"Identify",
		"GetPeerInfo",
	}
}

func (d *SRPCPeerServiceHandler) InvokeMethod(
	serviceID, methodID string,
	strm srpc.Stream,
) (bool, error) {
	if serviceID != "" && serviceID != d.GetServiceID() {
		return false, nil
	}

	switch methodID {
	case "Identify":
		return true, d.InvokeMethod_Identify(d.impl, strm)
	case "GetPeerInfo":
		return true, d.InvokeMethod_GetPeerInfo(d.impl, strm)
	default:
		return false, nil
	}
}

func (SRPCPeerServiceHandler) InvokeMethod_Identify(impl SRPCPeerServiceServer, strm srpc.Stream) error {
	req := new(IdentifyRequest)
	if err := strm.MsgRecv(req); err != nil {
		return err
	}
	serverStrm := &srpcPeerService_IdentifyStream{strm}
	return impl.Identify(req, serverStrm)
}

func (SRPCPeerServiceHandler) InvokeMethod_GetPeerInfo(impl SRPCPeerServiceServer, strm srpc.Stream) error {
	req := new(GetPeerInfoRequest)
	if err := strm.MsgRecv(req); err != nil {
		return err
	}
	out, err := impl.GetPeerInfo(strm.Context(), req)
	if err != nil {
		return err
	}
	return strm.MsgSend(out)
}

func SRPCRegisterPeerService(mux srpc.Mux, impl SRPCPeerServiceServer) error {
	return mux.Register(&SRPCPeerServiceHandler{impl: impl})
}

type SRPCPeerService_IdentifyStream interface {
	srpc.Stream
	Send(*IdentifyResponse) error
}

type srpcPeerService_IdentifyStream struct {
	srpc.Stream
}

func (x *srpcPeerService_IdentifyStream) Send(m *IdentifyResponse) error {
	return x.MsgSend(m)
}

type SRPCPeerService_GetPeerInfoStream interface {
	srpc.Stream
	SendAndClose(*GetPeerInfoResponse) error
}

type srpcPeerService_GetPeerInfoStream struct {
	srpc.Stream
}

func (x *srpcPeerService_GetPeerInfoStream) SendAndClose(m *GetPeerInfoResponse) error {
	if err := x.MsgSend(m); err != nil {
		return err
	}
	return x.CloseSend()
}