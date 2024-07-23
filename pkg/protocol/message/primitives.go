package message

type String []byte

func (s String) String() string {
	return string(s)
}

func (s String) Encode(enc *MessageEncoder) (em EncodedMessage, err error) {
	return enc.Encode(string(s))
}

func (s String) Decode(dec *MessageDecoder) (err error) {

	var str string

	if str, err = dec.ParseString(); err != nil {
		return
	}

	copy(s, str)

	return
}
