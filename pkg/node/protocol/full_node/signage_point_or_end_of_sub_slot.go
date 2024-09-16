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

type EndOfSubSlotBundle struct{}
