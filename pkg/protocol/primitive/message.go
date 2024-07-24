package primitive

const NullType MessageType = 0

var NullId Id = Id(0)

type MessageType uint8
type Id uint16

type Message interface {
	Type() MessageType
	Id() Id
	Data() []byte
}
