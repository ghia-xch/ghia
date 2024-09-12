package full_node

import (
	"github.com/ghia-xch/ghia/pkg/node/protocol"
	"github.com/ghia-xch/ghia/pkg/node/protocol/message"
	"lukechampine.com/uint128"
)

type NewPeak struct {
	HeaderHash                protocol.Hash
	Height                    uint32
	Weight                    uint128.Uint128
	ForkPointWithPreviousPeak uint32
	UnfinishedRewardBlockHash protocol.Hash
}

func (n *NewPeak) Type() message.MessageType {
	return protocol.NewPeak
}

func CreateNewPeak() *NewPeak {
	return &NewPeak{
		HeaderHash:                protocol.Hash{},
		Height:                    0,
		Weight:                    uint128.Uint128{},
		ForkPointWithPreviousPeak: 0,
		UnfinishedRewardBlockHash: protocol.Hash{},
	}
}
