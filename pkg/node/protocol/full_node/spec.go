package full_node

import (
	"encoding/binary"
	"github.com/ghia-xch/ghia/pkg/node/protocol/message"
)

////

type RequestBlock struct {
	Height                  uint32
	IncludeTransactionBlock bool
}

type RejectBlock struct {
	Height uint32
}

// type RequestBlocks struct {
// 	StartHeight             uint32
// 	EndHeight               uint32
// 	IncludeTransactionBlock bool
//}

var RequestBlocksType message.MessageType = 29

type RequestBlocks [10]byte

func RequestBlocksMessage(start uint32, end uint32, includeTxBlock bool) (r RequestBlocks) {

	r[0] = byte(RequestBlocksType)

	binary.BigEndian.PutUint32(r[1:5], start)
	binary.BigEndian.PutUint32(r[6:10], end)

	if includeTxBlock {
		r[9] = 1
	}

	return
}

func (r RequestBlocks) Encode() (em message.EncodedMessage, err error) {
	return r[:], nil
}

/*func (r RequestBlocks) Decode(em message.EncodedMessage) error {

	if len(em) != 10 {

	}

	r = em[0:10]
}*/

var RespondBlocksType message.MessageType = 38

type FullBlock struct{}

type RespondBlocks struct {
	startHeight uint32
	endHeight   uint32
	blocks      []FullBlock
}

//@streamable
//@dataclass(frozen=True)
//class RespondBlocks(Streamable):
//start_height: uint32
//end_height: uint32
//blocks: List[FullBlock]
//
//
//@streamable
//@dataclass(frozen=True)
//class RejectBlocks(Streamable):
//start_height: uint32
//end_height: uint32
//
//
//@streamable
//@dataclass(frozen=True)
//class RespondBlock(Streamable):
//block: FullBlock
//
//
//@streamable
//@dataclass(frozen=True)
//class NewUnfinishedBlock(Streamable):
//unfinished_reward_hash: bytes32
//
//
//@streamable
//@dataclass(frozen=True)
//class RequestUnfinishedBlock(Streamable):
//unfinished_reward_hash: bytes32
//
//
//@streamable
//@dataclass(frozen=True)
//class NewUnfinishedBlock2(Streamable):
//unfinished_reward_hash: bytes32
//foliage_hash: Optional[bytes32]
//
//
//@streamable
//@dataclass(frozen=True)
//class RequestUnfinishedBlock2(Streamable):
//unfinished_reward_hash: bytes32
//foliage_hash: Optional[bytes32]
//
//
//@streamable
//@dataclass(frozen=True)
//class RespondUnfinishedBlock(Streamable):
//unfinished_block: UnfinishedBlock
//
//

//@streamable
//@dataclass(frozen=True)
//class NewSignagePointOrEndOfSubSlot(Streamable):
//prev_challenge_hash: Optional[bytes32]
//challenge_hash: bytes32
//index_from_challenge: uint8
//last_rc_infusion: bytes32
//
//
//@streamable
//@dataclass(frozen=True)
//class RequestSignagePointOrEndOfSubSlot(Streamable):
//challenge_hash: bytes32
//index_from_challenge: uint8
//last_rc_infusion: bytes32
//
//
//@streamable
//@dataclass(frozen=True)
//class RespondSignagePoint(Streamable):
//index_from_challenge: uint8
//challenge_chain_vdf: VDFInfo
//challenge_chain_proof: VDFProof
//reward_chain_vdf: VDFInfo
//reward_chain_proof: VDFProof
//
//
//@streamable
//@dataclass(frozen=True)
//class RespondEndOfSubSlot(Streamable):
//end_of_slot_bundle: EndOfSubSlotBundle
//
//

//@streamable
//@dataclass(frozen=True)
//class RequestMempoolTransactions(Streamable):
//filter: bytes
//
//
//@streamable
//@dataclass(frozen=True)
//class NewCompactVDF(Streamable):
//height: uint32
//header_hash: bytes32
//field_vdf: uint8
//vdf_info: VDFInfo
//
//
//@streamable
//@dataclass(frozen=True)
//class RequestCompactVDF(Streamable):
//height: uint32
//header_hash: bytes32
//field_vdf: uint8
//vdf_info: VDFInfo
//
//
//@streamable
//@dataclass(frozen=True)
//class RespondCompactVDF(Streamable):
//height: uint32
//header_hash: bytes32
//field_vdf: uint8
//vdf_info: VDFInfo
//vdf_proof: VDFProof
//
//
//@streamable
//@dataclass(frozen=True)
//class RequestPeers(Streamable):
//"""
//Return full list of peers
//"""
//
//
//@streamable
//@dataclass(frozen=True)
//class RespondPeers(Streamable):
//peer_list: List[TimestampedPeerInfo]
