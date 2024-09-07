package capability

import "encoding/binary"

const (
	Base           Capability = 1
	BlockHeaders   Capability = 2
	RateLimitsV2   Capability = 3
	NoneResponse   Capability = 4
	MempoolUpdates Capability = 5
)

type Capability uint16

func (c Capability) Encode(enc []byte) ([]byte, error) {
	return binary.BigEndian.AppendUint16(enc, uint16(c)), nil
}
