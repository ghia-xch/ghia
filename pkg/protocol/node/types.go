package node

import (
	"github.com/ghia-xch/ghia/pkg/protocol/primitive/message"
)

type Type uint8

func (t Type) Encode(enc *message.MessageEncoder) (em message.EncodedMessage, err error) {
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
