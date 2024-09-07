package network

import (
	"errors"
	"github.com/ghia-xch/ghia/pkg/node/protocol/codec"
)

var (
	Mainnet Network = "mainnet"
	Testnet Network = "testnet"
	Simnet  Network = "simnet"
)

type Network string

func (n Network) String() string {
	return string(n)
}

func (n Network) Encode(enc []byte) ([]byte, error) {
	return codec.EncodeElement(string(n), enc)
}

var DefaultNetwork Network = Mainnet

func NewNetwork(str string) Network {
	return Network(str)
}

func Select(str string) (Network, error) {

	switch str {
	case "mainnet":
		return Mainnet, nil
	case "testnet":
		return Testnet, nil
	case "simnet":
		return Simnet, nil
	default:
		return Network(""), errors.New("'" + str + "' is not a valid network.")
	}
}
