package full_node

import (
	"github.com/ghia-xch/ghia/pkg/node/protocol"
	"github.com/ghia-xch/ghia/pkg/node/protocol/message"
	"github.com/ghia-xch/ghia/pkg/peer"
)

//@streamable
//@dataclass(frozen=True)
//class RequestPeers(Streamable):
//"""
//Return full list of peers
//"""

type RequestPeers struct{}

func (rp *RequestPeers) Type() message.Type { return protocol.RequestPeers }

//@streamable
//@dataclass(frozen=True)
//class RespondPeers(Streamable):
//peer_list: List[TimestampedPeerInfo]

type RespondPeers struct {
	Peers []peer.TimestamptedPeerInfo
}

func (rp *RespondPeers) Type() message.Type { return protocol.RespondPeers }
