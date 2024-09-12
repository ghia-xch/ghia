package protocol

import "github.com/ghia-xch/ghia/pkg/node/protocol/message"

type Callback func(em message.EncodedMessage) (err error)

var (
	NullCallback = func(em message.EncodedMessage) (err error) { return nil }
)
