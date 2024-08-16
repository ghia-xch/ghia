package full_node

import (
	"github.com/ghia-xch/ghia/pkg/node/protocol"
	"lukechampine.com/uint128"
)

type NewPeak struct {
	HeaderHash                protocol.Hash
	Height                    uint32
	Weight                    uint128.Uint128
	ForkPointWithPreviousPeak uint32
	UnfinishedRewardBlockHash protocol.Hash
}

func (n *NewPeak) Encode(enc *protocol.MessageEncoder) (em protocol.EncodedMessage, err error) {
	return enc.Encode(
		n.HeaderHash,
		n.Height,
		n.Weight,
		n.ForkPointWithPreviousPeak,
		n.UnfinishedRewardBlockHash,
	)
}

func (n *NewPeak) Decode(dec *protocol.MessageDecoder) (err error) {

	if n.HeaderHash, err = dec.ParseHash(); err != nil {
		return
	}

	if n.Height, err = dec.ParseUint32(); err != nil {
		return
	}

	if n.Weight, err = dec.ParseUint128(); err != nil {
		return
	}

	if n.ForkPointWithPreviousPeak, err = dec.ParseUint32(); err != nil {
		return
	}

	if n.UnfinishedRewardBlockHash, err = dec.ParseHash(); err != nil {
		return
	}

	return
}
