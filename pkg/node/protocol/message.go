package protocol

var (
	NullMessageType    = MessageType(0)
	NullId          Id = Id(0)
)

type MessageType uint8
type Id uint16

type Message interface {
	Type() MessageType
	Id() Id
	Data() []byte
}
