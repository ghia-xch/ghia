package full_node

import (
	"github.com/ghia-xch/ghia/pkg/node/protocol"
	"github.com/ghia-xch/ghia/pkg/node/protocol/message"
	"github.com/ghia-xch/ghia/pkg/node/protocol/vdf"
	bls12381 "github.com/kilic/bls12-381"
	"lukechampine.com/uint128"
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

//#[streamable]
//pub struct ProofOfSpace {
//challenge: Bytes32,
//pool_public_key: Option<G1Element>,
//pool_contract_puzzle_hash: Option<Bytes32>,
//plot_public_key: G1Element,
//size: u8,
//proof: Bytes,
//}

type HeaderBlock struct {
}

type ProofOfSpace struct {
	Challenge              protocol.Hash
	PoolPublicKey          *bls12381.PointG1 `cenc:"optional"`
	PoolContractPuzzleHash *protocol.Hash    `cenc:"optional"`
	PlotPublicKey          bls12381.PointG1
	Size                   uint8
	Proof                  []byte
}

type SubSlotData struct {
	ProofOfSpace      *ProofOfSpace    `cenc:"optional"`
	CCSignagePoint    *vdf.VDFProof    `cenc:"optional"`
	CCInfusionPoint   *vdf.VDFProof    `cenc:"optional"`
	ICCInfusionPoint  *vdf.VDFProof    `cenc:"optional"`
	CCSPVDFInfo       *vdf.VDFInfo     `cenc:"optional"`
	SignagePointIndex *uint8           `cenc:"optional"`
	CCSlotEnd         *vdf.VDFProof    `cenc:"optional"`
	ICCSlotEnd        *vdf.VDFProof    `cenc:"optional"`
	CCSlotEndInfo     *vdf.VDFInfo     `cenc:"optional"`
	ICCSlotEndInfo    *vdf.VDFInfo     `cenc:"optional"`
	CCIPVDFInfo       *vdf.VDFInfo     `cenc:"optional"`
	ICCIPVDFInfo      *vdf.VDFInfo     `cenc:"optional"`
	TotalIterations   *uint128.Uint128 `cenc:"optional"`
}

type SubEpochChallengeSegment struct {
	SubEpochN     uint32
	SubSlots      []SubSlotData
	RcSlotEndInfo *vdf.VDFInfo `cenc:"optional"`
}

type SubEpochData struct {
	RewardChainHash      protocol.Hash
	NumBlocksOverflow    uint8
	NewSubSlotIterations *uint64 `cenc:"optional"`
	NewDifficulty        *uint64 `cenc:"optional"`
}

type WeightProof struct {
	SubEpochs        []SubEpochData
	SubEpochSegments []SubEpochChallengeSegment
	RecentChainData  []HeaderBlock
}
