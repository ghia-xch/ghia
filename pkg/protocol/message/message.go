package message

const NullType Type = 0

var NullId Id = Id(0)

type Type uint8
type Id uint16

type Message interface {
	Type() Type
	Id() Id
	Data() []byte
}
