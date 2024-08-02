package capability

const (
	Base           Capability = 1
	BlockHeaders   Capability = 2
	RateLimitsV2   Capability = 3
	NoneResponse   Capability = 4
	MempoolUpdates Capability = 5
)

type Capability uint16
