package transport_controller

import (
	"io"

	"github.com/aperturerobotics/bifrost/protocol"
	"github.com/golang/protobuf/proto"
	"github.com/pkg/errors"
)

// NewStreamEstablish constructs a new StreamEstablish message.
func NewStreamEstablish(protocolID protocol.ID) *StreamEstablish {
	return &StreamEstablish{ProtocolId: string(protocolID)}
}

func writeStreamEstablishHeader(w io.Writer, msg *StreamEstablish) (int, error) {
	dat, err := proto.Marshal(msg)
	if err != nil {
		return 0, err
	}

	i, err := w.Write(proto.EncodeVarint(uint64(len(dat))))
	nw := i
	if err != nil {
		return nw, err
	}

	i, err = w.Write(dat)
	nw += i
	return nw, err
}

func readAtLeast(r io.Reader, n, min int, buf []byte) (int, error) {
	for n < min {
		nr, err := r.Read(buf[n:])
		if err != nil {
			return n, err
		}
		n += nr
	}
	return n, nil
}

func readStreamEstablishHeader(r io.Reader) (*StreamEstablish, error) {
	b := make([]byte, 4)
	n := 0
	var err error
	n, err = readAtLeast(r, n, 4, b)
	if err != nil {
		return nil, err
	}

	// Read the header length varint
	headerLen, headerLenBytes := proto.DecodeVarint(b)
	if headerLenBytes == 0 {
		return nil, err
	}

	// header len is at most 100,000 bytes
	if int(headerLen) > streamEstablishMaxPacketSize || headerLen == 0 {
		return nil, errors.Errorf(
			"stream establish header length invalid: %d (expected <= %d)",
			headerLen,
			streamEstablishMaxPacketSize,
		)
	}

	headerBuf := make([]byte, int(headerLen))
	copy(headerBuf, b[headerLenBytes:])
	n = len(b) - headerLenBytes
	if _, err := readAtLeast(r, n, int(headerLen), headerBuf); err != nil {
		return nil, err
	}

	// logrus.Infof("read stream establish: %v", headerBuf)
	// decode stream establish header
	estHeader := &StreamEstablish{}
	if err := proto.Unmarshal(headerBuf, estHeader); err != nil {
		return nil, err
	}

	return estHeader, nil
}
