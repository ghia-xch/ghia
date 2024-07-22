package network

import "github.com/ghia-xch/ghia/pkg/protocol/message"

type Network string

func (n Network) Encode(enc *message.MessageEncoder) (em message.EncodedMessage, err error) {
	return enc.Encode(string(n))
}

const (
	Mainnet Network = "mainnet"
	Testnet Network = "testnet"
	Simnet  Network = "simnet"
)

const DefaultNetwork Network = Mainnet
