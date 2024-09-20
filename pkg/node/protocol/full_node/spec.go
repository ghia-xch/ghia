package full_node

import (
	"github.com/ghia-xch/ghia/pkg/coin"
	"github.com/ghia-xch/ghia/pkg/coin/program"
	"github.com/ghia-xch/ghia/pkg/node/protocol"
	"github.com/ghia-xch/ghia/pkg/node/protocol/vdf"
	bls12381 "github.com/kilic/bls12-381"
	"lukechampine.com/uint128"
)

//#[streamable]
//pub struct TransactionsInfo {
//// Information that goes along with each transaction block
//generator_root: Bytes32, // sha256 of the block generator in this block
//generator_refs_root: Bytes32, // sha256 of the concatenation of the generator ref list entries
//aggregated_signature: G2Element,
//fees: u64, // This only includes user fees, not block rewards
//cost: u64, // This is the total cost of this block, including CLVM cost, cost of program size and conditions
//reward_claims_incorporated: Vec<Coin>, // These can be in any order
//}

// Information that goes along with each transaction block

type TransactionsInfo struct {
	GeneratorRoot            protocol.Hash
	GeneratorRefsRoot        protocol.Hash
	AggregatedSignature      bls12381.PointG2
	Fees                     uint64
	Cost                     uint64
	RewardClaimsIncorporated []coin.Coin
}

type FoliageTransactionBlock struct {
	PreviousTransactionBlockHash protocol.Hash
	Timestamp                    uint64
	FilterHash                   protocol.Hash
	AdditionsRoot                protocol.Hash
	RemovalsRoot                 protocol.Hash
	TransactionsInfoHash         protocol.Hash
}

type PoolTarget struct {
	PuzzleHash protocol.Hash
	MaxHeight  uint32
}

type FoliageBlockData struct {
	UnfinishedRewardBlockHash protocol.Hash
	PoolTarget                PoolTarget
	PoolSignature             protocol.Optional[bls12381.PointG2]
	FarmerRewardPuzzleHash    protocol.Hash
	ExtensionData             protocol.Hash
}

type Foliage struct {
	PrevBlockHash                    protocol.Hash
	RewardBlockHash                  protocol.Hash
	FoliageBlockData                 FoliageBlockData
	FoliageBlockDataSignature        bls12381.PointG2
	FoliageTransactionBlockHash      protocol.Optional[protocol.Hash]
	FoliageTransactionBlockSignature protocol.Optional[bls12381.PointG2]
}

type RewardChainBlock struct {
	Weight                     uint128.Uint128
	Height                     uint32
	TotalIterations            uint128.Uint128
	SignagePointIndex          uint8
	POSSSCCChallengeHash       protocol.Hash
	ProofOfSpace               ProofOfSpace
	ChallengeChainSPVDF        protocol.Optional[vdf.VDFInfo]
	ChallengeChainSPSignature  bls12381.PointG2
	ChallengeChainIPVDF        vdf.VDFInfo
	RewardChainSPVDF           protocol.Optional[vdf.VDFInfo]
	RewardChainSPSignature     bls12381.PointG2
	RewardChainIPVDF           vdf.VDFInfo
	InfusedChallengeChainIPVDF protocol.Optional[vdf.VDFInfo]
	IsTransactionBlock         bool
}

type FullBlock struct {
	FinishedSubSlots             []EndOfSubSlotBundle
	RewardChainBlock             RewardChainBlock
	ChallengeChainSPProof        protocol.Optional[vdf.VDFProof]
	ChallengeChainIPProof        vdf.VDFProof
	RewardChainSPProof           protocol.Optional[vdf.VDFProof]
	RewardChainIPProof           vdf.VDFProof
	InfusedChallengeChainIPProof protocol.Optional[vdf.VDFProof]
	Foliage                      Foliage
	FoliageTransactionBlock      protocol.Optional[FoliageTransactionBlock]
	TransactionsInfo             protocol.Optional[TransactionsInfo]
	TransactionsGenerator        protocol.Optional[program.Program]
	TransactionsGeneratorRefList []uint32
}
