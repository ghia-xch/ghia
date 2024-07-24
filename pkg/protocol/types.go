package protocol

import (
	"github.com/ghia-xch/ghia/pkg/protocol/primitive"
)

const (
	HandshakeType                     primitive.MessageType = 1
	NewPeak                           primitive.MessageType = 20
	NewTransaction                    primitive.MessageType = 21
	RequestTransaction                primitive.MessageType = 22
	RespondTransaction                primitive.MessageType = 23
	RequestProofOfWeight              primitive.MessageType = 24
	RespondProofOfWeight              primitive.MessageType = 25
	RequestBlock                      primitive.MessageType = 26
	RespondBlock                      primitive.MessageType = 27
	RejectBlock                       primitive.MessageType = 28
	RequestBlocks                     primitive.MessageType = 29
	RespondBlocks                     primitive.MessageType = 30
	RejectBlocks                      primitive.MessageType = 31
	NewUnfinishedBlock                primitive.MessageType = 32
	RequestUnfinishedBlock            primitive.MessageType = 33
	RespondUnfinishedBlock            primitive.MessageType = 34
	NewSignagePointOrEndOfSubSlot     primitive.MessageType = 35
	RequestSignagePointOrEndOfSubSlot primitive.MessageType = 36
	RespondSignagePoint               primitive.MessageType = 37
	RespondEndOfSubSlot               primitive.MessageType = 38
	RequestMempoolTransactions        primitive.MessageType = 39
	RequestCompactVDF                 primitive.MessageType = 40
	RespondCompactVDF                 primitive.MessageType = 41
	NewCompactVDF                     primitive.MessageType = 42
	RequestPeers                      primitive.MessageType = 43
	RespondPeers                      primitive.MessageType = 44
	NoneResponse                      primitive.MessageType = 91
)

var (
	singletonMessages = []primitive.MessageType{
		NewPeak,
		NewTransaction,
		NewUnfinishedBlock,
		// NewUnfinishedBlock2,
		NewSignagePointOrEndOfSubSlot,
		RequestMempoolTransactions,
		NewCompactVDF,
		// CoinStateUpdate,
		// MempoolItemsAdded,
		// MempoolItemsRemoved,
	}

	validResponseMessages = map[primitive.MessageType][]primitive.MessageType{
		RequestTransaction:     {RespondTransaction},
		RequestProofOfWeight:   {RespondProofOfWeight},
		RequestBlock:           {RespondBlock, RejectBlock},
		RequestBlocks:          {RespondBlocks, RejectBlocks},
		RequestUnfinishedBlock: {RespondUnfinishedBlock},
		// RequestUnfinishedBlock2: {RespondUnfinishedBlock},
		RequestSignagePointOrEndOfSubSlot: {RespondSignagePoint, RespondEndOfSubSlot},
		RequestCompactVDF:                 {RespondCompactVDF},
		RequestPeers:                      {RespondPeers},
	}
)

func HasExpectedResponse(mt primitive.MessageType) (isExpected bool) {

	_, isExpected = validResponseMessages[mt]

	return
}

func IsValidResponse(sent primitive.MessageType, recv primitive.MessageType) (isValid bool) {

	var responses []primitive.MessageType

	responses, isValid = validResponseMessages[sent]

	if !isValid {
		return
	}

	for r := range responses {
		if responses[r] == recv {
			return true
		}
	}

	return false
}
