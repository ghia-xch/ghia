package protocol

import "github.com/ghia-xch/ghia/pkg/node/protocol/message"

type MessageHandler struct {
	Type     message.Type
	Callback Callback
}

func Handler(t message.Type, callback Callback) MessageHandler {
	return MessageHandler{
		Type:     t,
		Callback: callback,
	}
}
