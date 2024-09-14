package full_node

import (
	"github.com/ghia-xch/ghia/pkg/node/protocol"
	"github.com/ghia-xch/ghia/pkg/node/protocol/message"
	"github.com/ghia-xch/ghia/pkg/peer"
)

type RequestPeers struct{}

func (rp *RequestPeers) Type() message.Type { return protocol.RequestPeers }

type RespondPeers struct {
	Peers []peer.TimestamptedPeerInfo
}

func (rp *RespondPeers) Type() message.Type { return protocol.RespondPeers }
