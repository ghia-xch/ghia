package network

import "github.com/ghia-xch/ghia/pkg/protocol/message"

var (
	Mainnet = &Network{message.NewString("mainnet")}
	Testnet = &Network{message.NewString("testnet")}
	Simnet  = &Network{message.NewString("simnet")}
)

type Network struct{ message.String }

var DefaultNetwork *Network = Mainnet

func NewNetwork(str string) *Network {
	return &Network{message.NewString(str)}
}
