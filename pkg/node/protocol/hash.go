package protocol

import (
	"encoding/hex"
	"github.com/ghia-xch/ghia/pkg/node/protocol/message"
)

type Hash [32]byte

func (h *Hash) String() string {
	return "0x" + hex.EncodeToString(h[:])
}

func (h *Hash) Bytes() []byte {
	return h[:]
}

func (h *Hash) Encode(enc *message.MessageEncoder) (em message.EncodedMessage, err error) {
	return enc.Encode(h)
}
