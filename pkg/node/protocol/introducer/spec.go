package introducer

import (
	"github.com/ghia-xch/ghia/pkg/node/protocol"
	"github.com/ghia-xch/ghia/pkg/node/protocol/message"
	"github.com/ghia-xch/ghia/pkg/peer"
)

type RequestPeersIntroducer struct{}

func (r *RequestPeersIntroducer) Type() message.Type {
	return protocol.RequestPeersIntroducer
}

type RespondPeersIntroducer struct {
	Peers []*peer.TimestamptedPeerInfo
}

func (r *RespondPeersIntroducer) Type() message.Type {
	return protocol.RespondPeersIntroducer
}
