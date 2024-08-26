package coin

import (
	"crypto/sha256"
	"encoding/binary"
	"github.com/davecgh/go-spew/spew"
	"github.com/ghia-xch/ghia/pkg/node/protocol"
)

type Coin struct {
	Parent     protocol.Hash
	PuzzleHash protocol.Hash
	Amount     uint64
}

func (c *Coin) Id() protocol.Hash {

	hash := sha256.New()

	hash.Write(c.Parent.Bytes())
	hash.Write(c.PuzzleHash.Bytes())

	uiBytes := make([]byte, 8)
	binary.BigEndian.PutUint64(uiBytes, c.Amount)

	spew.Dump(uiBytes)

	hash.Write(uiBytes)

	res := make([]byte, 32)
	res = hash.Sum(nil)

	return protocol.Hash(res[:])
}

func i64tob(val uint64) []byte {
	r := make([]byte, 8)
	for i := uint64(0); i < 8; i++ {
		r[i] = byte((val >> (i * 8)) & 0xff)
	}
	return r
}

func btoi64(val []byte) uint64 {
	r := uint64(0)
	for i := uint64(0); i < 8; i++ {
		r |= uint64(val[i]) << (8 * i)
	}
	return r
}

func NewCoin(parent protocol.Hash, pHash protocol.Hash, amount uint64) *Coin {
	return &Coin{
		Parent:     parent,
		PuzzleHash: pHash,
		Amount:     amount,
	}
}
