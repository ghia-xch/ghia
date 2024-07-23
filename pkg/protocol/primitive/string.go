package primitive

type String struct{ *string }

func (s *String) String() string {
	return *s.string
}

func (s *String) Encode(enc *MessageEncoder) (em EncodedMessage, err error) {
	return enc.Encode(*s.string)
}

func (s *String) Decode(dec *MessageDecoder) (err error) {

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
