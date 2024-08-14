package protocol

import (
	"encoding/hex"
)

type Hash [32]byte

func (h *Hash) String() string {
	return "0x" + hex.EncodeToString(h[:])
}

func (h *Hash) Encode(enc *MessageEncoder) (em EncodedMessage, err error) {
	return enc.Encode(h)
}

func (h *Hash) Decode(dec *MessageDecoder) (err error) {

	var hash []byte

	if hash, err = dec.ParseBytes(); err != nil {
		return
	}

	copy(h[:], hash[0:31])

	return
}
