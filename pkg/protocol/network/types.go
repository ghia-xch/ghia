package network

import "github.com/ghia-xch/ghia/pkg/protocol/message"

type Network struct {
	message.String
}

var (
	Mainnet Network = Network{message.String("mainnet")}
	Testnet Network = Network{message.String("testnet")}
	Simnet  Network = Network{message.String("simnet")}
)

var DefaultNetwork Network = Mainnet
