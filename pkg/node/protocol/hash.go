package protocol

import (
	"encoding/hex"
)

type Hash [32]byte

func (h *Hash) String() string {
	return "0x" + hex.EncodeToString(h[:])
}

func (h *Hash) Bytes() []byte {
	return h[:]
}

func (h *Hash) Encode(enc *MessageEncoder) (em EncodedMessage, err error) {
	return enc.Encode(h)
}
