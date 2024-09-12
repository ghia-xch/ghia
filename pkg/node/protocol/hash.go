package protocol

import (
	"encoding/hex"
	"errors"
	"github.com/ghia-xch/ghia/pkg/node/protocol/codec"
)

type Hash [32]byte

func (h *Hash) String() string {
	return "0x" + hex.EncodeToString(h[:])
}

func (h *Hash) Bytes() []byte {
	return h[:]
}

func (h *Hash) Encode(enc []byte) ([]byte, error) {
	return codec.EncodeRaw(h[:], enc)
}

func (h *Hash) Decode(b []byte) ([]byte, error) {

	if len(b) < 32 {
		return nil, errors.New("invalid hash length")
	}

	copy(h[0:32], b[0:32])

	return b[32:], nil
}
