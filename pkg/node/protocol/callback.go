package protocol

type Callback func(dec *MessageDecoder) (err error)

var (
	NullCallback = func(em EncodedMessage) (err error) { return nil }
)
