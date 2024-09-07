package introducer

import (
	"github.com/ghia-xch/ghia/pkg/node/protocol"
	"github.com/ghia-xch/ghia/pkg/node/protocol/message"
	"github.com/ghia-xch/ghia/pkg/peer"
)

type RequestPeersIntroducer byte

func (r *RequestPeersIntroducer) Encode(enc *message.MessageEncoder) (em message.EncodedMessage, err error) {

	enc.Reset(protocol.RequestPeersIntroducer, nil)

	return enc.Encode()
}

func (r *RequestPeersIntroducer) Decode(dec *message.MessageDecoder, em message.EncodedMessage) (err error) {
	return nil
}

type RespondPeersIntroducer []*peer.TimestamptedPeerInfo

func (r *RespondPeersIntroducer) Encode(enc *message.MessageEncoder) (em message.EncodedMessage, err error) {

	enc.Reset(protocol.RespondPeersIntroducer, nil)

	return enc.Encode(r)
}

func (r *RespondPeersIntroducer) Decode(dec *message.MessageDecoder, em message.EncodedMessage) (err error) {

	if err = dec.Reset(em); err != nil {
		return
	}

	return nil
}
