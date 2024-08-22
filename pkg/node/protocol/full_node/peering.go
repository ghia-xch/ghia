package full_node

import (
	"github.com/ghia-xch/ghia/pkg/node/protocol"
	"github.com/ghia-xch/ghia/pkg/peer"
)

//@streamable
//@dataclass(frozen=True)
//class RequestPeers(Streamable):
//"""
//Return full list of peers
//"""
//
//
//@streamable
//@dataclass(frozen=True)
//class RespondPeers(Streamable):
//peer_list: List[TimestampedPeerInfo]

type RequestPeers [7]byte

func CreateRequestPeers() (em protocol.EncodedMessage) {

	var rp RequestPeers

	rp[0] = byte(protocol.RequestPeers)

	return rp[:]
}

type RespondPeers struct {
	peers []peer.TimestamptedPeerInfo
}

func (rp *RespondPeers) Encode(enc *protocol.MessageEncoder) (em protocol.EncodedMessage, err error) {

	if _, err = enc.Encode(uint32(len(rp.peers))); err != nil {
		return
	}

	for _, peer := range rp.peers {
		if _, err = peer.Encode(enc); err != nil {
			return
		}
	}

	return em, nil
}

func (rp *RespondPeers) Decode(dec *protocol.MessageDecoder) (err error) {

	var length uint32

	if length, err = dec.ParseUint32(); err != nil {
		return
	}

	rp.peers = make([]peer.TimestamptedPeerInfo, length)

	for i := 0; i < int(length); i++ {

		var peerInfo peer.TimestamptedPeerInfo

		if err = peerInfo.Decode(dec); err != nil {
			return
		}

		rp.peers[i] = peerInfo
	}

	return nil
}
