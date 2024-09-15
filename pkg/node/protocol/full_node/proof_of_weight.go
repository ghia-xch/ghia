package full_node

import (
	"github.com/ghia-xch/ghia/pkg/node/protocol"
	"github.com/ghia-xch/ghia/pkg/node/protocol/message"
	"github.com/ghia-xch/ghia/pkg/node/protocol/vdf"
)

type RequestProofOfWeight struct {
	TotalNumberOfBlocks uint32
	Tip                 protocol.Hash
}

func (r *RequestProofOfWeight) Type() message.Type { return protocol.RequestProofOfWeight }

type RespondProofOfWeight struct {
	WeightProof WeightProof
	Tip         [32]byte
}

func (r *RespondProofOfWeight) Type() message.Type { return protocol.RespondProofOfWeight }

///
///
///

type HeaderBlock struct {
}

type SubSlotData struct {
}

type SubEpochChallengeSegment struct {
	SubEpochN     uint32
	SubSlots      []SubSlotData
	RcSlotEndInfo *vdf.VDFInfo
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
