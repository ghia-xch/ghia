package node

type Type uint8

func (t Type) Encode(enc []byte) ([]byte, error) {
	return append(enc, byte(t)), nil
}

const (
	FullNode   Type = 1
	Harvester  Type = 2
	Farmer     Type = 3
	Timelord   Type = 4
	Introducer Type = 5
	Wallet     Type = 6
	DataLayer  Type = 7
)
