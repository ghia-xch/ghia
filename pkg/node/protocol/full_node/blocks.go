package full_node

import (
	"github.com/ghia-xch/ghia/pkg/node/protocol"
	"github.com/ghia-xch/ghia/pkg/node/protocol/message"
)

type RequestBlocks struct {
	StartHeight             uint32
	EndHeight               uint32
	IncludeTransactionBlock bool
}

func (r *RequestBlocks) Type() message.Type { return protocol.RequestBlocks }

type RespondBlocks struct {
	StartHeight uint32
	EndHeight   uint32
	Blocks      []FullBlock
}

func (r *RespondBlocks) Type() message.Type { return protocol.RespondBlocks }

type RejectBlocks struct {
	StartHeight uint32
	EndHeight   uint32
}

func (r *RejectBlocks) Type() message.Type { return protocol.RejectBlocks }
