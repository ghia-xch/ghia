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

type RespondCompactVDF struct {
	Height     uint32
	HeaderHash protocol.Hash
	FieldVDF   uint8
	VDFInfo    VDFInfo
	VDFProof   VDFProof
}

func (c *RespondCompactVDF) Type() message.Type { return protocol.RespondCompactVDF }

///
///
///

type VDFProof struct{}
