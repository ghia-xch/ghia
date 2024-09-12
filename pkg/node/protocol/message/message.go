package message

var (
	NullMessageType    = MessageType(0)
	NullId          Id = Id(0)
)

type MessageType uint8

func (m MessageType) Equals(other MessageType) bool {
	return m == other
}

type Id uint16

type Message interface {
	Type() MessageType
	Id() Id
	Data() []byte
}
