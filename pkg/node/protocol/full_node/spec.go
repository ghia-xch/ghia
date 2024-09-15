package full_node

///

type FullBlock struct{}

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
