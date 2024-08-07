package introducer

import (
	protocol2 "github.com/ghia-xch/ghia/pkg/node/protocol"
	"github.com/ghia-xch/ghia/pkg/peer"
)

type RequestPeersIntroducer byte

func (r *RequestPeersIntroducer) Encode(enc *protocol2.MessageEncoder) (em protocol2.EncodedMessage, err error) {

	enc.Reset(protocol2.RequestPeersIntroducer, nil)

	return enc.Encode()
}

func (r *RequestPeersIntroducer) Decode(dec *protocol2.MessageDecoder, em protocol2.EncodedMessage) (err error) {
	return nil
}

type RespondPeersIntroducer []*peer.TimestamptedPeerInfo

func (r *RespondPeersIntroducer) Encode(enc *protocol2.MessageEncoder) (em protocol2.EncodedMessage, err error) {

	enc.Reset(protocol2.RespondPeersIntroducer, nil)

	return enc.Encode(r)
}

func (r *RespondPeersIntroducer) Decode(dec *protocol2.MessageDecoder, em protocol2.EncodedMessage) (err error) {

	if err = dec.Reset(em); err != nil {
		return
	}

	return nil
}
