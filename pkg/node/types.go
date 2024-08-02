package node

import (
	"github.com/ghia-xch/ghia/pkg/protocol/primitive"
)

type Type uint8

func (t Type) Encode(enc *primitive.MessageEncoder) (em primitive.EncodedMessage, err error) {
	return enc.Encode(uint8(t))
}

const (
	FullNode   Type = 1
	Harvester  Type = 2
	Farmer     Type = 3
	Timelord   Type = 4
	Introducer Type = 5
	Wallet     Type = 6
	DataLayer  Type = 7
)
