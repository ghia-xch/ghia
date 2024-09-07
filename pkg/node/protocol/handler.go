package protocol

import "github.com/ghia-xch/ghia/pkg/node/protocol/message"

type MessageHandler struct {
	Type     message.MessageType
	Callback Callback
}

func Handler(t message.MessageType, callback Callback) MessageHandler {
	return MessageHandler{
		Type:     t,
		Callback: callback,
	}
}
