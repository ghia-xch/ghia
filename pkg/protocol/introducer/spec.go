package introducer

import (
	"github.com/ghia-xch/ghia/pkg/peer"
	"github.com/ghia-xch/ghia/pkg/protocol"
)

type RequestPeersIntroducer byte

func (r *RequestPeersIntroducer) Encode(enc *protocol.MessageEncoder) (em protocol.EncodedMessage, err error) {

	enc.Reset(protocol.RequestPeersIntroducer, nil)

	return enc.Encode()
}

func (r *RequestPeersIntroducer) Decode(dec *protocol.MessageDecoder, em protocol.EncodedMessage) (err error) {
	return nil
}

type RespondPeersIntroducer []*peer.TimestamptedPeerInfo

func (r *RespondPeersIntroducer) Encode(enc *protocol.MessageEncoder) (em protocol.EncodedMessage, err error) {

	enc.Reset(protocol.RespondPeersIntroducer, nil)

	return enc.Encode(r)
}

func (r *RespondPeersIntroducer) Decode(dec *protocol.MessageDecoder, em protocol.EncodedMessage) (err error) {

	if err = dec.Reset(em); err != nil {
		return
	}

	return nil
}
