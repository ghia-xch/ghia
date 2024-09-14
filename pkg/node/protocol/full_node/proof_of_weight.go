package full_node

import (
	"github.com/ghia-xch/ghia/pkg/node/protocol"
)

type RequestProofOfWeight struct {
	TotalNumberOfBlocks uint32
	Tip                 protocol.Hash
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
