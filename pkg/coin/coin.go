package coin

import (
	"github.com/ghia-xch/ghia/pkg/node/protocol"
)

type Coin struct {
	Parent     protocol.Hash
	PuzzleHash protocol.Hash
	Amount     uint64
}
