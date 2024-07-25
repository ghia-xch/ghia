package protocol

import (
	"github.com/ghia-xch/ghia/pkg/protocol/primitive"
)

/*
	Contains an incomplete list of message types. See the following url for the full list:

	https://github.com/Chia-Network/chia-blockchain/blob/main/chia/protocols/protocol_message_types.py

*/

const (

	// Universal MessageTypes [Any <-> Any]

	NullType      primitive.MessageType = 0
	HandshakeType primitive.MessageType = 1
	Error         primitive.MessageType = 255

	// FullNode MessageTypes [FullNode <-> FullNode]
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

	// Wallet MessageTypes [Wallet <-> FullNode]
	RequestPuzzleSolution  primitive.MessageType = 45
	RespondPuzzleSolution  primitive.MessageType = 46
	RejectPuzzleSolution   primitive.MessageType = 47
	SendTransaction        primitive.MessageType = 48
	RespondSendTransaction primitive.MessageType = 49
	NewPeakWallet          primitive.MessageType = 50
	RequestBlockHeader     primitive.MessageType = 51
	RespondBlockHeader     primitive.MessageType = 52
	RejectBlockHeader      primitive.MessageType = 53
	RequestRemovals        primitive.MessageType = 54
	RespondRemovals        primitive.MessageType = 55
	RejectRemovals         primitive.MessageType = 56
	RequestAdditions       primitive.MessageType = 57
	RespondAdditions       primitive.MessageType = 58
	RejectAdditions        primitive.MessageType = 59
	RequestHeaderBlocks    primitive.MessageType = 60
	RejectHeaderBlocks     primitive.MessageType = 61
	RespondHeaderBlocks    primitive.MessageType = 62

	// Introducer MessageTypes [Introducer <-> FullNode]
	RequestPeersIntroducer primitive.MessageType = 63
	RespondPeersIntroducer primitive.MessageType = 64

	// Wallet Updates
	CoinStateUpdate              primitive.MessageType = 69
	RegisterInterestInPuzzleHash primitive.MessageType = 70
	RespondInterestInPuzzleHash  primitive.MessageType = 71
	RegisterInterestInCoin       primitive.MessageType = 72
	RespondInterestInCoin        primitive.MessageType = 73
	RequestChildren              primitive.MessageType = 74
	RespondChildren              primitive.MessageType = 75
	RequestSESHashes             primitive.MessageType = 76
	RespondSESHashes             primitive.MessageType = 77

	RequestBlockHeaders primitive.MessageType = 86
	RejectBlockHeaders  primitive.MessageType = 87
	RespondBlockHeaders primitive.MessageType = 88

	RequestFeeEstimates primitive.MessageType = 89
	RespondFeeEstimates primitive.MessageType = 90

	// FullNode Updates
	NoneResponse            primitive.MessageType = 91
	NewUnfinishedBlock2     primitive.MessageType = 92
	RequestUnfinishedBlock2 primitive.MessageType = 93

	// Wallet Sync
	RequestRemovePuzzleSubscriptions primitive.MessageType = 94
	RespondRemovePuzzleSubscriptions primitive.MessageType = 95
	RequestRemoveCoinSubscriptions   primitive.MessageType = 96
	RespondRemoveCoinSubscriptions   primitive.MessageType = 97
	RequestPuzzleState               primitive.MessageType = 98
	RespondPuzzleState               primitive.MessageType = 99
	RejectPuzzleState                primitive.MessageType = 100
	RequestCoinState                 primitive.MessageType = 101
	RespondCoinState                 primitive.MessageType = 102
	RejectCoinState                  primitive.MessageType = 103

	// Wallet Mempool Updates
	MempoolItemsAdded   primitive.MessageType = 104
	MempoolItemsRemoved primitive.MessageType = 105
	RequestCostInfo     primitive.MessageType = 106
	RespondCostInfo     primitive.MessageType = 107
)

var (

	/*
		Singletons are messages that require no response from the peer.
	*/

	singletons = map[primitive.MessageType]bool{

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

	validResponses = map[primitive.MessageType][]primitive.MessageType{

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

func HasNoExpectedResponse(mt primitive.MessageType) (isNotExpected bool) {

	_, isNotExpected = singletons[mt]

	return
}

func HasExpectedResponse(mt primitive.MessageType) (isExpected bool) {

	_, isExpected = validResponses[mt]

	return
}

func IsValidResponse(sent primitive.MessageType, recv primitive.MessageType) (isValid bool) {

	var responses []primitive.MessageType

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
