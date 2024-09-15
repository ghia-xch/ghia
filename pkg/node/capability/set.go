package capability

var DefaultSet = map[Capability]string{
	Base:         "1",
	BlockHeaders: "1",
	RateLimitsV2: "1",
	//NoneResponse:   "0",
	//MempoolUpdates: "1",
}

type Set map[Capability]string

func (s Set) IsEnabled(capability Capability) bool {

	val, ok := s[capability]

	if !ok || val == "0" {
		return false
	}

	return true
}
