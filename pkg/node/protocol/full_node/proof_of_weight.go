package full_node

import (
	"github.com/ghia-xch/ghia/pkg/node/protocol"
	"github.com/ghia-xch/ghia/pkg/node/protocol/message"
)

type RequestProofOfWeight struct {
	TotalNumberOfBlocks uint32
	Tip                 protocol.Hash
}

func (r RequestProofOfWeight) Encode(enc *message.MessageEncoder) (em message.EncodedMessage, err error) {

	return enc.Encode(
		r.TotalNumberOfBlocks,
		r.Tip,
	)
}

func (r *RequestProofOfWeight) Decode(dec *message.MessageDecoder) (err error) {

	if r.TotalNumberOfBlocks, err = dec.ParseUint32(); err != nil {
		return err
	}

	//if r.Tip, err = dec.ParseHash(); err != nil {
	//	return err
	//}

	return nil
}

func CreateRequestProofOfWeight(nb uint32, t protocol.Hash) (em message.EncodedMessage) {
	em, _ = RequestProofOfWeight{
		TotalNumberOfBlocks: nb,
		Tip:                 t,
	}.Encode(message.NewMessageEncoder(42))
	return
}

type HeaderBlock struct {
}

type SubSlotData struct {
}

type ClassGroupElement [100]byte

type VDFInfo struct {
	Challenge     protocol.Hash
	NumIterations uint64
	Output        ClassGroupElement
}

type SubEpochChallengeSegment struct {
	SubEpochN     uint32
	SubSlots      []SubSlotData
	RcSlotEndInfo *VDFInfo
}

type SubEpochData struct {
	RewardChainHash      protocol.Hash
	NumBlocksOverflow    uint8
	NewSubSlotIterations *uint64
	NewDifficulty        *uint64
}

type WeightProof struct {
	SubEpochs        []SubEpochData
	SubEpochSegments []SubEpochChallengeSegment
	RecentChainData  []HeaderBlock
}

type RespondProofOfWeight struct {
	WeightProof WeightProof
	Tip         [32]byte
}
