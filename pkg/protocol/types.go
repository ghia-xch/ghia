package protocol

/*
	Contains an incomplete list of message types. See the following url for the full list:

	https://github.com/Chia-Network/chia-blockchain/blob/main/chia/protocols/protocol_message_types.py

*/

const (

	// Universal MessageTypes [Any <-> Any]

	NullType      MessageType = 0
	HandshakeType MessageType = 1
	Error         MessageType = 255

	// FullNode MessageTypes [FullNode <-> FullNode]
	NewPeak                           MessageType = 20
	NewTransaction                    MessageType = 21
	RequestTransaction                MessageType = 22
	RespondTransaction                MessageType = 23
	RequestProofOfWeight              MessageType = 24
	RespondProofOfWeight              MessageType = 25
	RequestBlock                      MessageType = 26
	RespondBlock                      MessageType = 27
	RejectBlock                       MessageType = 28
	RequestBlocks                     MessageType = 29
	RespondBlocks                     MessageType = 30
	RejectBlocks                      MessageType = 31
	NewUnfinishedBlock                MessageType = 32
	RequestUnfinishedBlock            MessageType = 33
	RespondUnfinishedBlock            MessageType = 34
	NewSignagePointOrEndOfSubSlot     MessageType = 35
	RequestSignagePointOrEndOfSubSlot MessageType = 36
	RespondSignagePoint               MessageType = 37
	RespondEndOfSubSlot               MessageType = 38
	RequestMempoolTransactions        MessageType = 39
	RequestCompactVDF                 MessageType = 40
	RespondCompactVDF                 MessageType = 41
	NewCompactVDF                     MessageType = 42
	RequestPeers                      MessageType = 43
	RespondPeers                      MessageType = 44

	// Wallet MessageTypes [Wallet <-> FullNode]
	RequestPuzzleSolution  MessageType = 45
	RespondPuzzleSolution  MessageType = 46
	RejectPuzzleSolution   MessageType = 47
	SendTransaction        MessageType = 48
	RespondSendTransaction MessageType = 49
	NewPeakWallet          MessageType = 50
	RequestBlockHeader     MessageType = 51
	RespondBlockHeader     MessageType = 52
	RejectBlockHeader      MessageType = 53
	RequestRemovals        MessageType = 54
	RespondRemovals        MessageType = 55
	RejectRemovals         MessageType = 56
	RequestAdditions       MessageType = 57
	RespondAdditions       MessageType = 58
	RejectAdditions        MessageType = 59
	RequestHeaderBlocks    MessageType = 60
	RejectHeaderBlocks     MessageType = 61
	RespondHeaderBlocks    MessageType = 62

	// Introducer MessageTypes [Introducer <-> FullNode]
	RequestPeersIntroducer MessageType = 63
	RespondPeersIntroducer MessageType = 64

	// Wallet Updates
	CoinStateUpdate              MessageType = 69
	RegisterInterestInPuzzleHash MessageType = 70
	RespondInterestInPuzzleHash  MessageType = 71
	RegisterInterestInCoin       MessageType = 72
	RespondInterestInCoin        MessageType = 73
	RequestChildren              MessageType = 74
	RespondChildren              MessageType = 75
	RequestSESHashes             MessageType = 76
	RespondSESHashes             MessageType = 77

	RequestBlockHeaders MessageType = 86
	RejectBlockHeaders  MessageType = 87
	RespondBlockHeaders MessageType = 88

	RequestFeeEstimates MessageType = 89
	RespondFeeEstimates MessageType = 90

	// FullNode Updates
	NoneResponse            MessageType = 91
	NewUnfinishedBlock2     MessageType = 92
	RequestUnfinishedBlock2 MessageType = 93

	// Wallet Sync
	RequestRemovePuzzleSubscriptions MessageType = 94
	RespondRemovePuzzleSubscriptions MessageType = 95
	RequestRemoveCoinSubscriptions   MessageType = 96
	RespondRemoveCoinSubscriptions   MessageType = 97
	RequestPuzzleState               MessageType = 98
	RespondPuzzleState               MessageType = 99
	RejectPuzzleState                MessageType = 100
	RequestCoinState                 MessageType = 101
	RespondCoinState                 MessageType = 102
	RejectCoinState                  MessageType = 103

	// Wallet Mempool Updates
	MempoolItemsAdded   MessageType = 104
	MempoolItemsRemoved MessageType = 105
	RequestCostInfo     MessageType = 106
	RespondCostInfo     MessageType = 107
)

var (

	/*
		Singletons are messages that require no response from the peer.
	*/

	singletons = map[MessageType]bool{

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

	validResponses = map[MessageType][]MessageType{

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

func HasNoExpectedResponse(mt MessageType) (isNotExpected bool) {

	_, isNotExpected = singletons[mt]

	return
}

func HasExpectedResponse(mt MessageType) (isExpected bool) {

	_, isExpected = validResponses[mt]

	return
}

func IsValidResponse(sent MessageType, recv MessageType) (isValid bool) {

	var responses []MessageType

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
