package vdf

import (
	"github.com/ghia-xch/ghia/pkg/node/protocol"
	"github.com/ghia-xch/ghia/pkg/node/protocol/codec"
)

type ClassGroupElement [100]byte

func (c ClassGroupElement) Encode(enc []byte) ([]byte, error) {
	return codec.EncodeRaw(c[:], enc)
}

type VDFInfo struct {
	Challenge     protocol.Hash
	NumIterations uint64
	Output        ClassGroupElement
}
