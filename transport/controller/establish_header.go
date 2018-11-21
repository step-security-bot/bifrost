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

func readStreamEstablishHeader(r io.Reader) (*StreamEstablish, error) {
	b := make([]byte, 1500)
	n, err := r.Read(b)
	if err != nil {
		return nil, err
	}
	b = b[:n]

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

	headerBuf := make([]byte, headerLen)
	copy(headerBuf, b[headerLenBytes:])
	remainingBytes := int(headerLen) - (len(b) - headerLenBytes)
	if remainingBytes > 0 {
		_, err := io.ReadFull(r, headerBuf[(len(b)-headerLenBytes):])
		if err != nil {
			return nil, err
		}
	}

	// decode stream establish header
	estHeader := &StreamEstablish{}
	if err := proto.Unmarshal(headerBuf, estHeader); err != nil {
		return nil, err
	}

	return estHeader, nil
}