package full_node

import (
	"github.com/ghia-xch/ghia/pkg/node/protocol"
	"github.com/ghia-xch/ghia/pkg/node/protocol/message"
)

type NewCompactVDF struct {
	Height     uint32
	HeaderHash protocol.Hash
	FieldVDF   uint8
	VDFInfo    VDFInfo
}

func (c *NewCompactVDF) Type() message.Type { return protocol.NewCompactVDF }

type RequestCompactVDF struct {
	Height     uint32
	HeaderHash protocol.Hash
	FieldVDF   uint8
	VDFInfo    VDFInfo
}

func (c *RequestCompactVDF) Type() message.Type { return protocol.RequestCompactVDF }

//@streamable
//@dataclass(frozen=True)
//class RespondCompactVDF(Streamable):
//height: uint32
//header_hash: bytes32
//field_vdf: uint8
//vdf_info: VDFInfo
//vdf_proof: VDFProof
//
