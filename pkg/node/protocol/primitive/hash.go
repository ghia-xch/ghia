package primitive

import (
	"encoding/hex"
	"github.com/ghia-xch/ghia/pkg/node/protocol"
)

type Hash [32]byte

func (h *Hash) String() string {
	return "0x" + hex.EncodeToString(h[:])
}

func (h *Hash) Encode(enc *protocol.MessageEncoder) (em protocol.EncodedMessage, err error) {
	return enc.Encode(h)
}

func (h *Hash) Decode(dec *protocol.MessageDecoder) (err error) {

	var hash []byte

	if hash, err = dec.ParseBytes(32); err != nil {
		return
	}

	copy(h[:], hash[0:31])

	return
}

func NewHash(str string) String {
	return String{&str}
}
