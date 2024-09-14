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

func (np *NewPeak) Type() message.Type { return protocol.NewPeak }
