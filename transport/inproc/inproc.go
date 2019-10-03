package inproc

import (
	"context"
	"net"
	"sync"

	"github.com/aperturerobotics/bifrost/peer"
	"github.com/aperturerobotics/bifrost/transport"
	"github.com/aperturerobotics/bifrost/transport/common/pconn"
	"github.com/blang/semver"
	"github.com/libp2p/go-libp2p-core/crypto"
	"github.com/sirupsen/logrus"
)

// TransportID is the transport identifier
const TransportID = "inproc"

// ControllerID is the controller identifier.
const ControllerID = "bifrost/inproc/1"

// Version is the version of the inproc implementation.
var Version = semver.MustParse("0.0.1")

// Inproc implements a Inproc transport.
type Inproc struct {
	// Transport is the packet transport
	*pconn.Transport

	// le is the logger
	le *logrus.Entry
	// packetConn is the packet conn
	packetConn *packetConn
	// localAddr is the local addr
	localAddr net.Addr

	// mtx guards below
	mtx sync.Mutex
	// remotes are the currently known remotes
	// map is from string (net.addr.String()) to *packetConn
	remotes map[string]*packetConn
}

// NewInproc builds a new Inproc transport.
// Yields Links to other Inproc transports.
func NewInproc(
	ctx context.Context,
	le *logrus.Entry,
	opts *Config,
	pKey crypto.PrivKey,
	c transport.TransportHandler,
) (transport.Transport, error) {
	peerID, err := peer.IDFromPrivateKey(pKey)
	if err != nil {
		return nil, err
	}

	localAddr := newAddr(peerID)
	ip := &Inproc{
		le:        le,
		localAddr: localAddr,
		remotes:   make(map[string]*packetConn),
	}
	npc := newPacketConn(
		ctx,
		localAddr,
		ip.writeToAddr,
	)
	ip.Transport, err = pconn.New(
		le,
		npc,
		pKey,
		parseAddr,
		c,
		opts.GetPacketOpts(),
	)
	if err != nil {
		return nil, err
	}
	ip.packetConn = npc
	return ip, nil
}

// ConnectToInproc connects the inproc to a remote inproc.
// Will overwrite any existing connection
func (t *Inproc) ConnectToInproc(ctx context.Context, other *Inproc) {
	oa := other.localAddr.String()
	t.mtx.Lock()
	t.remotes[oa] = other.packetConn
	t.mtx.Unlock()
}

// DisconnectInproc disconnects a previously connected inproc.
func (t *Inproc) DisconnectInproc(ctx context.Context, other *Inproc) {
	oa := other.localAddr.String()
	t.mtx.Lock()
	delete(t.remotes, oa)
	t.mtx.Unlock()
}

// writeToAddr routes outgoing packets.
func (t *Inproc) writeToAddr(ctx context.Context, p []byte, addr net.Addr) (int, error) {
	oa := addr.String()
	t.mtx.Lock()
	out, outOk := t.remotes[oa]
	t.mtx.Unlock()
	if !outOk {
		return 0, &net.AddrError{
			Addr: oa,
			Err:  "remote transport not connected",
		}
	}
	return out.HandlePacket(ctx, p, t.localAddr)
}

// _ is a type assertion.
var _ transport.Transport = ((*Inproc)(nil))
