package websocket

import (
	"context"
	"net"
	"sync"

	"github.com/aperturerobotics/bifrost/handshake/identity"
	"github.com/aperturerobotics/bifrost/link"
	"github.com/aperturerobotics/bifrost/peer"
	"github.com/aperturerobotics/bifrost/stream"
	"github.com/aperturerobotics/bifrost/util/scrc"
	"github.com/sirupsen/logrus"
	"github.com/xtaci/smux"
)

// Link represents a WebSocket-based connection/link.
// This version is for the browser side of the connection.
type Link struct {
	// ctx is the context for this link
	ctx context.Context
	// ctxCancel cancels the context for this link.
	ctxCancel context.CancelFunc
	// le is the logrus entry
	le *logrus.Entry
	// url is the url we dialed/the remote addr
	url string
	// neg is the negotiated identity data
	neg *identity.Result
	// sharedSecret is the shared secret
	sharedSecret [32]byte
	// peerID is the remote peer id
	peerID peer.ID
	// mux is the reliable stream multiplexer
	mux *smux.Session
	// conn is the underlying connection
	conn net.Conn
	// closed is the closed callback
	closed func()
	// closedOnce guards closed
	closedOnce sync.Once

	uuid, transportUUID uint64
}

// NewLink builds a new link.
func NewLink(
	ctx context.Context,
	le *logrus.Entry,
	url string,
	transportUUID uint64,
	neg *identity.Result,
	sharedSecret [32]byte,
	conn net.Conn,
	initiator bool,
	closed func(),
) *Link {
	nctx, nctxCancel := context.WithCancel(ctx)
	pid, _ := peer.IDFromPublicKey(neg.Peer)

	// Construct the encrypted channel
	// TODO: add support for unencryoted streams
	conn = newEncConn(conn, sharedSecret)

	var sess *smux.Session
	if initiator {
		sess, _ = smux.Server(conn, smux.DefaultConfig())
	} else {
		sess, _ = smux.Client(conn, smux.DefaultConfig())
	}

	l := &Link{
		ctx:       nctx,
		ctxCancel: nctxCancel,

		le:     le,
		url:    url,
		neg:    neg,
		peerID: pid,
		mux:    sess,
		uuid:   newLinkUUID(url, pid),
		conn:   conn,
		closed: closed,

		sharedSecret:  sharedSecret,
		transportUUID: transportUUID,
	}

	go l.watchLinkContextCancel()
	go l.readPump()
	return l
}

// newLinkUUID builds the UUID for a link
func newLinkUUID(url string, peerID peer.ID) uint64 {
	return scrc.Crc64(
		[]byte("websocket"),
		[]byte(url),
		[]byte(peerID),
	)
}

// GetRemotePeer returns the identity of the remote peer.
func (l *Link) GetRemotePeer() peer.ID {
	return l.peerID
}

// GetUUID returns the host-unique ID.
// This should be repeatable between re-constructions of the same link.
func (l *Link) GetUUID() uint64 {
	return l.uuid
}

// GetTransportUUID returns the unique ID of the transport.
func (l *Link) GetTransportUUID() uint64 {
	return l.transportUUID
}

// OpenStream opens a stream on the link, with the given parameters.
// WebSocket is always reliable and always encrypted.
func (l *Link) OpenStream(opts stream.OpenOpts) (stream.Stream, error) {
	return l.mux.OpenStream()
}

// Close closes the link.
// Any blocked ReadFrom or WriteTo operations will be unblocked and return errors.
func (l *Link) Close() error {
	l.closedOnce.Do(func() {
		_ = l.conn.Close()
		if closed := l.closed; closed != nil {
			closed()
		}
		_ = l.mux.Close()
	})
	return nil
}

// watchLinkContextCancel cleans up the link after the context is canceled.
func (l *Link) watchLinkContextCancel() {
	<-l.ctx.Done()
	l.le.Debug("link context canceled, calling close")
	l.Close()
}

// readPump reads messages from the link.
func (l *Link) readPump() {
	for {
		strm, err := l.mux.AcceptStream()
		if err != nil {
			l.le.WithError(err).Warn("link reader exiting")
			l.Close()
			return
		}

		l.le.
			WithField("stream-id", strm.ID()).
			Debug("stream accepted")
		_ = strm
	}
}

// _ is a type assertion
var _ link.Link = ((*Link)(nil))
