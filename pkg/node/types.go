package node

import (
	"github.com/ghia-xch/ghia/pkg/node/protocol"
)

type Type uint8

func (t Type) Encode(enc *protocol.MessageEncoder) (em protocol.EncodedMessage, err error) {
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
