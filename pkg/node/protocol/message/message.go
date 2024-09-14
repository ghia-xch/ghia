package message

var (
	NullMessageType    = Type(0)
	NullId          Id = Id(0)
)

type Type uint8

func (m Type) Equals(other Type) bool {
	return m == other
}

type Id uint16

type Message interface {
	Type() Type
	Id() Id
	Data() []byte
}
