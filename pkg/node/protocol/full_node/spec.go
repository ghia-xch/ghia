package full_node

///

type FullBlock struct{}

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
