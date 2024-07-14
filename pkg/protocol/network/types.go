package network

type Network string

const (
	Mainnet Network = "mainnet"
	Testnet Network = "testnet"
	Simnet  Network = "simnet"
)

const DefaultNetwork Network = Mainnet
