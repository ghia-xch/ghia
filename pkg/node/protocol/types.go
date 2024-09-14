package protocol

import "github.com/ghia-xch/ghia/pkg/node/protocol/message"

/*
	Contains an incomplete list of message types. See the following url for the full list:

	https://github.com/Chia-Network/chia-blockchain/blob/main/chia/protocols/protocol_message_types.py

*/

const (

	// Universal MessageTypes [Any <-> Any]

	NullType      message.Type = 0
	HandshakeType message.Type = 1
	Error         message.Type = 255

	// FullNode MessageTypes [FullNode <-> FullNode]
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

	// Wallet MessageTypes [Wallet <-> FullNode]
	RequestPuzzleSolution  message.Type = 45
	RespondPuzzleSolution  message.Type = 46
	RejectPuzzleSolution   message.Type = 47
	SendTransaction        message.Type = 48
	RespondSendTransaction message.Type = 49
	NewPeakWallet          message.Type = 50
	RequestBlockHeader     message.Type = 51
	RespondBlockHeader     message.Type = 52
	RejectBlockHeader      message.Type = 53
	RequestRemovals        message.Type = 54
	RespondRemovals        message.Type = 55
	RejectRemovals         message.Type = 56
	RequestAdditions       message.Type = 57
	RespondAdditions       message.Type = 58
	RejectAdditions        message.Type = 59
	RequestHeaderBlocks    message.Type = 60
	RejectHeaderBlocks     message.Type = 61
	RespondHeaderBlocks    message.Type = 62

	// Introducer MessageTypes [Introducer <-> FullNode]
	RequestPeersIntroducer message.Type = 63
	RespondPeersIntroducer message.Type = 64

	// Wallet Updates
	CoinStateUpdate              message.Type = 69
	RegisterInterestInPuzzleHash message.Type = 70
	RespondInterestInPuzzleHash  message.Type = 71
	RegisterInterestInCoin       message.Type = 72
	RespondInterestInCoin        message.Type = 73
	RequestChildren              message.Type = 74
	RespondChildren              message.Type = 75
	RequestSESHashes             message.Type = 76
	RespondSESHashes             message.Type = 77

	RequestBlockHeaders message.Type = 86
	RejectBlockHeaders  message.Type = 87
	RespondBlockHeaders message.Type = 88

	RequestFeeEstimates message.Type = 89
	RespondFeeEstimates message.Type = 90

	// FullNode Updates
	NoneResponse            message.Type = 91
	NewUnfinishedBlock2     message.Type = 92
	RequestUnfinishedBlock2 message.Type = 93

	// Wallet Sync
	RequestRemovePuzzleSubscriptions message.Type = 94
	RespondRemovePuzzleSubscriptions message.Type = 95
	RequestRemoveCoinSubscriptions   message.Type = 96
	RespondRemoveCoinSubscriptions   message.Type = 97
	RequestPuzzleState               message.Type = 98
	RespondPuzzleState               message.Type = 99
	RejectPuzzleState                message.Type = 100
	RequestCoinState                 message.Type = 101
	RespondCoinState                 message.Type = 102
	RejectCoinState                  message.Type = 103

	// Wallet Mempool Updates
	MempoolItemsAdded   message.Type = 104
	MempoolItemsRemoved message.Type = 105
	RequestCostInfo     message.Type = 106
	RespondCostInfo     message.Type = 107
)

var (

	/*
		Singletons are messages that require no response from the peer.
	*/

	singletons = map[message.Type]bool{

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

	validResponses = map[message.Type][]message.Type{

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

func HasNoExpectedResponse(mt message.Type) (isNotExpected bool) {

	_, isNotExpected = singletons[mt]

	return
}

func HasExpectedResponse(mt message.Type) (isExpected bool) {

	_, isExpected = validResponses[mt]

	return
}

func IsValidResponse(sent message.Type, recv message.Type) (isValid bool) {

	var responses []message.Type

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

func TypeAsString(mt message.Type) string {

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
