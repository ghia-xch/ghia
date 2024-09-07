package protocol

import "github.com/ghia-xch/ghia/pkg/node/protocol/message"

/*
	Contains an incomplete list of message types. See the following url for the full list:

	https://github.com/Chia-Network/chia-blockchain/blob/main/chia/protocols/protocol_message_types.py

*/

const (

	// Universal MessageTypes [Any <-> Any]

	NullType      message.MessageType = 0
	HandshakeType message.MessageType = 1
	Error         message.MessageType = 255

	// FullNode MessageTypes [FullNode <-> FullNode]
	NewPeak                           message.MessageType = 20
	NewTransaction                    message.MessageType = 21
	RequestTransaction                message.MessageType = 22
	RespondTransaction                message.MessageType = 23
	RequestProofOfWeight              message.MessageType = 24
	RespondProofOfWeight              message.MessageType = 25
	RequestBlock                      message.MessageType = 26
	RespondBlock                      message.MessageType = 27
	RejectBlock                       message.MessageType = 28
	RequestBlocks                     message.MessageType = 29
	RespondBlocks                     message.MessageType = 30
	RejectBlocks                      message.MessageType = 31
	NewUnfinishedBlock                message.MessageType = 32
	RequestUnfinishedBlock            message.MessageType = 33
	RespondUnfinishedBlock            message.MessageType = 34
	NewSignagePointOrEndOfSubSlot     message.MessageType = 35
	RequestSignagePointOrEndOfSubSlot message.MessageType = 36
	RespondSignagePoint               message.MessageType = 37
	RespondEndOfSubSlot               message.MessageType = 38
	RequestMempoolTransactions        message.MessageType = 39
	RequestCompactVDF                 message.MessageType = 40
	RespondCompactVDF                 message.MessageType = 41
	NewCompactVDF                     message.MessageType = 42
	RequestPeers                      message.MessageType = 43
	RespondPeers                      message.MessageType = 44

	// Wallet MessageTypes [Wallet <-> FullNode]
	RequestPuzzleSolution  message.MessageType = 45
	RespondPuzzleSolution  message.MessageType = 46
	RejectPuzzleSolution   message.MessageType = 47
	SendTransaction        message.MessageType = 48
	RespondSendTransaction message.MessageType = 49
	NewPeakWallet          message.MessageType = 50
	RequestBlockHeader     message.MessageType = 51
	RespondBlockHeader     message.MessageType = 52
	RejectBlockHeader      message.MessageType = 53
	RequestRemovals        message.MessageType = 54
	RespondRemovals        message.MessageType = 55
	RejectRemovals         message.MessageType = 56
	RequestAdditions       message.MessageType = 57
	RespondAdditions       message.MessageType = 58
	RejectAdditions        message.MessageType = 59
	RequestHeaderBlocks    message.MessageType = 60
	RejectHeaderBlocks     message.MessageType = 61
	RespondHeaderBlocks    message.MessageType = 62

	// Introducer MessageTypes [Introducer <-> FullNode]
	RequestPeersIntroducer message.MessageType = 63
	RespondPeersIntroducer message.MessageType = 64

	// Wallet Updates
	CoinStateUpdate              message.MessageType = 69
	RegisterInterestInPuzzleHash message.MessageType = 70
	RespondInterestInPuzzleHash  message.MessageType = 71
	RegisterInterestInCoin       message.MessageType = 72
	RespondInterestInCoin        message.MessageType = 73
	RequestChildren              message.MessageType = 74
	RespondChildren              message.MessageType = 75
	RequestSESHashes             message.MessageType = 76
	RespondSESHashes             message.MessageType = 77

	RequestBlockHeaders message.MessageType = 86
	RejectBlockHeaders  message.MessageType = 87
	RespondBlockHeaders message.MessageType = 88

	RequestFeeEstimates message.MessageType = 89
	RespondFeeEstimates message.MessageType = 90

	// FullNode Updates
	NoneResponse            message.MessageType = 91
	NewUnfinishedBlock2     message.MessageType = 92
	RequestUnfinishedBlock2 message.MessageType = 93

	// Wallet Sync
	RequestRemovePuzzleSubscriptions message.MessageType = 94
	RespondRemovePuzzleSubscriptions message.MessageType = 95
	RequestRemoveCoinSubscriptions   message.MessageType = 96
	RespondRemoveCoinSubscriptions   message.MessageType = 97
	RequestPuzzleState               message.MessageType = 98
	RespondPuzzleState               message.MessageType = 99
	RejectPuzzleState                message.MessageType = 100
	RequestCoinState                 message.MessageType = 101
	RespondCoinState                 message.MessageType = 102
	RejectCoinState                  message.MessageType = 103

	// Wallet Mempool Updates
	MempoolItemsAdded   message.MessageType = 104
	MempoolItemsRemoved message.MessageType = 105
	RequestCostInfo     message.MessageType = 106
	RespondCostInfo     message.MessageType = 107
)

var (

	/*
		Singletons are messages that require no response from the peer.
	*/

	singletons = map[message.MessageType]bool{

		// FullNode
		NewPeak:                       true,
		NewTransaction:                true,
		NewUnfinishedBlock:            true,
		NewUnfinishedBlock2:           true,
		NewSignagePointOrEndOfSubSlot: true,
		RequestMempoolTransactions:    true,
		NewCompactVDF:                 true,

		// Wallet
		CoinStateUpdate:     true,
		MempoolItemsAdded:   true,
		MempoolItemsRemoved: true,

		// Not in spec, but makes sense to be here.
		NewPeakWallet: true,
	}

	validResponses = map[message.MessageType][]message.MessageType{

		// FullNode
		RequestTransaction:                {RespondTransaction},
		RequestProofOfWeight:              {RespondProofOfWeight},
		RequestBlock:                      {RespondBlock, RejectBlock},
		RequestBlocks:                     {RespondBlocks, RejectBlocks},
		RequestUnfinishedBlock:            {RespondUnfinishedBlock},
		RequestUnfinishedBlock2:           {RespondUnfinishedBlock},
		RequestBlockHeader:                {RespondBlockHeader, RejectBlockHeader},
		RequestRemovals:                   {RespondRemovals, RejectRemovals},
		RequestAdditions:                  {RespondAdditions, RejectAdditions},
		RequestSignagePointOrEndOfSubSlot: {RespondSignagePoint, RespondEndOfSubSlot},
		RequestCompactVDF:                 {RespondCompactVDF},
		RequestPeers:                      {RespondPeers},

		// Wallet
		RequestHeaderBlocks:          {RespondHeaderBlocks, RejectHeaderBlocks, RejectBlockHeaders},
		RegisterInterestInPuzzleHash: {RespondInterestInPuzzleHash},
		RegisterInterestInCoin:       {RespondInterestInCoin},
		RequestChildren:              {RespondChildren},
		RequestSESHashes:             {RespondSESHashes},
		RequestBlockHeaders:          {RespondBlockHeaders, RejectBlockHeaders, RejectHeaderBlocks},
		RequestFeeEstimates:          {RespondFeeEstimates},

		// Introducer
		RequestPeersIntroducer: {RespondPeersIntroducer},

		// Wallet Sync
		RequestPuzzleSolution:            {RespondPuzzleSolution, RejectPuzzleSolution},
		SendTransaction:                  {RespondSendTransaction},
		RequestRemovePuzzleSubscriptions: {RespondRemovePuzzleSubscriptions},
		RequestRemoveCoinSubscriptions:   {RespondRemoveCoinSubscriptions},
		RequestPuzzleState:               {RespondPuzzleState, RejectPuzzleState},
		RequestCoinState:                 {RespondCoinState, RejectCoinState},
		RequestCostInfo:                  {RespondCostInfo},
	}
)

func HasNoExpectedResponse(mt message.MessageType) (isNotExpected bool) {

	_, isNotExpected = singletons[mt]

	return
}

func HasExpectedResponse(mt message.MessageType) (isExpected bool) {

	_, isExpected = validResponses[mt]

	return
}

func IsValidResponse(sent message.MessageType, recv message.MessageType) (isValid bool) {

	var responses []message.MessageType

	responses, isValid = validResponses[sent]

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

func TypeAsString(mt message.MessageType) string {

	switch mt {
	case NullType:
		return "null"
	case HandshakeType:
		return "handshake"
	case Error:
		return "error"
	case NewPeak:
		return "new peak"
	case NewTransaction:
		return "new transaction"
	case RequestTransaction:
		return "request transaction"
	case RespondTransaction:
		return "respond transaction"
	case RequestProofOfWeight:
		return "request proof of weight"
	case RespondProofOfWeight:
		return "respond proof of weight"
	case RequestBlock:
		return "request block"
	case RespondBlock:
		return "respond block"
	case RejectBlock:
		return "reject block"
	case RequestBlocks:
		return "request blocks"
	case RespondBlocks:
		return "respond blocks"
	case RejectBlocks:
		return "reject blocks"
	case NewUnfinishedBlock:
		return "new unfinished block"
	case RequestUnfinishedBlock:
		return "request unfinished block"
	case RespondUnfinishedBlock:
		return "respond unfinished block"
	case NewSignagePointOrEndOfSubSlot:
		return "new signage point or end of sub slot"
	case RequestSignagePointOrEndOfSubSlot:
		return "request signage point or end of sub slot"
	case RespondSignagePoint:
		return "respond signage point"
	case RespondEndOfSubSlot:
		return "respond end of sub slot"
	case RequestMempoolTransactions:
		return "request mempool transactions"
	case RequestCompactVDF:
		return "request compact vdf"
	case RespondCompactVDF:
		return "respond compact vdf"
	case NewCompactVDF:
		return "new compact vdf"
	case RequestPeers:
		return "request peers"
	case RespondPeers:
		return "respond peers"

	default:
		return "unknown"
	}
}
