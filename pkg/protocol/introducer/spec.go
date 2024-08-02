package introducer

import (
	"github.com/ghia-xch/ghia/pkg/peer"
	"github.com/ghia-xch/ghia/pkg/protocol"
	"github.com/ghia-xch/ghia/pkg/protocol/primitive"
)

type RequestPeersIntroducer byte

func (r *RequestPeersIntroducer) Encode(enc *primitive.MessageEncoder) (em primitive.EncodedMessage, err error) {

	enc.Reset(protocol.RequestPeersIntroducer, nil)

	return enc.Encode()
}

func (r *RequestPeersIntroducer) Decode(dec *primitive.MessageDecoder, em primitive.EncodedMessage) (err error) {
	return nil
}

type RespondPeersIntroducer []*peer.TimestamptedPeerInfo

func (r *RespondPeersIntroducer) Encode(enc *primitive.MessageEncoder) (em primitive.EncodedMessage, err error) {

	enc.Reset(protocol.RespondPeersIntroducer, nil)

	return enc.Encode(r)
}

func (r *RespondPeersIntroducer) Decode(dec *primitive.MessageDecoder, em primitive.EncodedMessage) (err error) {

	if err = dec.Reset(em); err != nil {
		return
	}

	return nil
}
