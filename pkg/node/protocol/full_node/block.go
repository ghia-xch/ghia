package full_node

import (
	"github.com/ghia-xch/ghia/pkg/node/protocol"
	"github.com/ghia-xch/ghia/pkg/node/protocol/message"
)

type RequestBlock struct {
	Height                  uint32
	IncludeTransactionBlock bool
}

func (r *RequestBlock) Type() message.Type { return protocol.RequestBlock }

type RespondBlock struct {
	Block FullBlock
}

func (r *RespondBlock) Type() message.Type { return protocol.RespondBlock }

type RejectBlock struct {
	Height uint32
}

func (r *RejectBlock) Type() message.Type { return protocol.RejectBlock }
