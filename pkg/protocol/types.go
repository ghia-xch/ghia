package protocol

import (
	"github.com/ghia-xch/ghia/pkg/protocol/primitive"
)

const (
	HandshakeType                     primitive.Type = 1
	NewPeak                           primitive.Type = 20
	NewTransaction                    primitive.Type = 21
	RequestTransaction                primitive.Type = 22
	RespondTransaction                primitive.Type = 23
	RequestProofOfWeight              primitive.Type = 24
	RespondProofOfWeight              primitive.Type = 25
	RequestBlock                      primitive.Type = 26
	RespondBlock                      primitive.Type = 27
	RejectBlock                       primitive.Type = 28
	RequestBlocks                     primitive.Type = 29
	RespondBlocks                     primitive.Type = 30
	RejectBlocks                      primitive.Type = 31
	NewUnfinishedBlock                primitive.Type = 32
	RequestUnfinishedBlock            primitive.Type = 33
	RespondUnfinishedBlock            primitive.Type = 34
	NewSignagePointOrEndOfSubSlot     primitive.Type = 35
	RequestSignagePointOrEndOfSubSlot primitive.Type = 36
	RespondSignagePoint               primitive.Type = 37
	RespondEndOfSubSlot               primitive.Type = 38
	RequestMempoolTransactions        primitive.Type = 39
	RequestCompactVDF                 primitive.Type = 40
	RespondCompactVDF                 primitive.Type = 41
	NewCompactVDF                     primitive.Type = 42
	RequestPeers                      primitive.Type = 43
	RespondPeers                      primitive.Type = 44
	NoneResponse                      primitive.Type = 91
)
