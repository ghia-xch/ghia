package network

import (
	"github.com/ghia-xch/ghia/pkg/protocol/primitive"
)

var (
	Mainnet = &Network{primitive.NewString("mainnet")}
	Testnet = &Network{primitive.NewString("testnet")}
	Simnet  = &Network{primitive.NewString("simnet")}
)

type Network struct{ primitive.String }

var DefaultNetwork *Network = Mainnet

func NewNetwork(str string) *Network {
	return &Network{primitive.NewString(str)}
}
