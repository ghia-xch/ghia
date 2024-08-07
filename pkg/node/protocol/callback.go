package protocol

type Callback func(em EncodedMessage) (err error)

var (
	NullCallback = func(em EncodedMessage) (err error) { return nil }
)
