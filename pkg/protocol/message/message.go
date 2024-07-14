package message

type Type uint8
type Id uint16

type Message interface {
	Type() Type
	Id() Id
	Data() []byte
}
