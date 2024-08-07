package network

import (
	"errors"
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

func Select(str string) (*Network, error) {

	switch str {
	case "mainnet":
		return Mainnet, nil
	case "testnet":
		return Testnet, nil
	case "simnet":
		return Simnet, nil
	default:
		return nil, errors.New("'" + str + "' is not a valid network.")
	}
}
