package full_node

import (
	"github.com/ghia-xch/ghia/pkg/node/protocol"
	"github.com/ghia-xch/ghia/pkg/node/protocol/message"
	"github.com/ghia-xch/ghia/pkg/node/protocol/vdf"
)

type NewSignagePointOrEndOfSubSlot struct {
	PreviousChallengeHash protocol.Hash `cenc:"optional"`
	ChallengeHash         protocol.Hash
	IndexFromChallenge    uint8
	LastRCInfusion        protocol.Hash
}

func (n *NewSignagePointOrEndOfSubSlot) Type() message.Type {
	return protocol.NewSignagePointOrEndOfSubSlot
}

type RequestSignagePointOrEndOfSubSlot struct {
	ChallengeHash      protocol.Hash
	IndexFromChallenge uint8
	LastRCInfusion     protocol.Hash
}

func (n *RequestSignagePointOrEndOfSubSlot) Type() message.Type {
	return protocol.RequestSignagePointOrEndOfSubSlot
}

type RespondSignagePoint struct {
	IndexFromChallenge  uint8
	ChallengeChainVDF   vdf.VDFInfo
	ChallengeChainProof vdf.VDFProof
	RewardChainVDF      vdf.VDFInfo
	RewardChainProof    vdf.VDFProof
}

func (n *RespondSignagePoint) Type() message.Type {
	return protocol.RespondSignagePoint
}

type RespondEndOfSubSlot struct {
	EndOfSubSlot EndOfSubSlotBundle
}

func (n *RespondEndOfSubSlot) Type() message.Type {
	return protocol.RespondEndOfSubSlot
}

///
///
///

type SubSlotProofs struct {
	CCSlotProof  vdf.VDFProof
	ICCSlotProof *vdf.VDFProof `cenc:"optional"`
	RCSlotProof  vdf.VDFProof
}

type RewardChainSubSlot struct {
	EndOfSlotVDF   vdf.VDFInfo
	CCSubSlotHash  protocol.Hash
	ICCSubSlotHash *protocol.Hash `cenc:"optional"`
	Deficit        uint8
}

type InfusedChallengeChainSubSlot struct {
	ICCEndOfSlotVDF vdf.VDFInfo
}

type ChallengeChainSubSlot struct {
	CCEndOfSlotVDF       vdf.VDFInfo
	InfusedCCSubSlotHash *protocol.Hash `cenc:"optional"`
	SubEpochSummaryHash  *protocol.Hash `cenc:"optional"`
	NewSubSlotIterations *uint64        `cenc:"optional"`
	NewDifficulty        *uint64        `cenc:"optional"`
}

type EndOfSubSlotBundle struct {
	ChallengeChain        ChallengeChainSubSlot
	InfusedChallengeChain InfusedChallengeChainSubSlot `cenc:"optional"`
	RewardChain           RewardChainSubSlot
	Proofs                SubSlotProofs
}
