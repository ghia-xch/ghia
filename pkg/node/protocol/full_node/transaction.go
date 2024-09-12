package full_node

import (
	"github.com/ghia-xch/ghia/pkg/coin"
	"github.com/ghia-xch/ghia/pkg/node/protocol"
	"github.com/ghia-xch/ghia/pkg/node/protocol/message"
)

type NewTransaction struct {
	TransactionId protocol.Hash
	Cost          uint64
	Fees          uint64
}

func (r *NewTransaction) Type() message.MessageType { return protocol.NewTransaction }

type RequestTransaction struct {
	TransactionId protocol.Hash
}

func (r *RequestTransaction) Type() message.MessageType { return protocol.RequestTransaction }

type RespondTransaction struct {
	Transaction coin.SpendBundle
}

func (r *RespondTransaction) Type() message.MessageType { return protocol.RespondTransaction }
