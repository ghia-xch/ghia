package protocol

import (
	"github.com/ghia-xch/ghia/pkg/protocol/primitive/message"
)

const (
	HandshakeType                     message.Type = 1
	NewPeak                           message.Type = 20
	NewTransaction                    message.Type = 21
	RequestTransaction                message.Type = 22
	RespondTransaction                message.Type = 23
	RequestProofOfWeight              message.Type = 24
	RespondProofOfWeight              message.Type = 25
	RequestBlock                      message.Type = 26
	RespondBlock                      message.Type = 27
	RejectBlock                       message.Type = 28
	RequestBlocks                     message.Type = 29
	RespondBlocks                     message.Type = 30
	RejectBlocks                      message.Type = 31
	NewUnfinishedBlock                message.Type = 32
	RequestUnfinishedBlock            message.Type = 33
	RespondUnfinishedBlock            message.Type = 34
	NewSignagePointOrEndOfSubSlot     message.Type = 35
	RequestSignagePointOrEndOfSubSlot message.Type = 36
	RespondSignagePoint               message.Type = 37
	RespondEndOfSubSlot               message.Type = 38
	RequestMempoolTransactions        message.Type = 39
	RequestCompactVDF                 message.Type = 40
	RespondCompactVDF                 message.Type = 41
	NewCompactVDF                     message.Type = 42
	RequestPeers                      message.Type = 43
	RespondPeers                      message.Type = 44
	NoneResponse                      message.Type = 91
)
