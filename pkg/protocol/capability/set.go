package capability

import "github.com/ghia-xch/ghia/pkg/protocol/message"

var DefaultSet = map[Capability]string{
	Base:         "1",
	BlockHeaders: "1",
	RateLimitsV2: "1",
	//NoneResponse:   "0",
	//MempoolUpdates: "1",
}

type Set map[Capability]string

func (s Set) IsEnabled(capability Capability) bool {
	return false
}

// Encode expects a List[Tuple[uint16,string]]
func (s Set) Encode(enc *message.MessageEncoder) (em message.EncodedMessage, err error) {

	if em, err = enc.Encode(uint32(len(s))); err != nil {
		return
	}

	for key, value := range s {
		if em, err = enc.Encode(uint16(key), value); err != nil {
			return
		}
	}

	return enc.Bytes(), nil
}

func (s Set) Decode(dec *message.MessageDecoder) (err error) {
	return nil
}
