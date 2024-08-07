package protocol

type MessageHandler struct {
	Type     MessageType
	Callback Callback
}

func Handler(t MessageType, callback Callback) MessageHandler {
	return MessageHandler{
		Type:     t,
		Callback: callback,
	}
}
