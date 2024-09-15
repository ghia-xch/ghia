package full_node

import (
	"github.com/ghia-xch/ghia/pkg/node/protocol"
	"github.com/ghia-xch/ghia/pkg/node/protocol/message"
)

type RequestMempoolTransactions struct {
	Filter []byte
}

func (r *RequestMempoolTransactions) Type() message.Type { return protocol.RequestMempoolTransactions }
