package primitive

import "github.com/ghia-xch/ghia/pkg/protocol"

type String struct{ *string }

func (s *String) String() string {
	return *s.string
}

func (s *String) Encode(enc *protocol.MessageEncoder) (em protocol.EncodedMessage, err error) {
	return enc.Encode(*s.string)
}

func (s *String) Decode(dec *protocol.MessageDecoder) (err error) {

	var str string

	if str, err = dec.ParseString(); err != nil {
		return
	}

	s.string = &str

	return
}

func NewString(str string) String {
	return String{&str}
}
