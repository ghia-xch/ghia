package node

type Type uint8

const (
	FullNode   Type = 1
	Harvester  Type = 2
	Farmer     Type = 3
	Timelord   Type = 4
	Introducer Type = 5
	Wallet     Type = 6
	DataLayer  Type = 7
)
