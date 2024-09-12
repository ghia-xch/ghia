package message

import "encoding/binary"

const (
	minEncodedMessageLen   = 6
	minIdEncodedMessageLen = 8

	NilMessageType = MessageType(0)
)

type EncodedMessage []byte

func (em EncodedMessage) Type() MessageType {

	if len(em) == 0 {
		return NilMessageType
	}

	return MessageType(em[0])
}

func (em EncodedMessage) HasId() bool {

	if len(em) < minEncodedMessageLen {
		return false
	}

	return em[1] == 1
}

func (em EncodedMessage) Id() *Id {

	if !em.HasId() {
		return nil
	}

	id := Id(binary.BigEndian.Uint16(em[2:3]))

	return &id
}

func (em EncodedMessage) Data() []byte {

	if em == nil || len(em) <= minEncodedMessageLen {
		return nil
	}

	if !em.HasId() {
		return em[minEncodedMessageLen:]
	}

	return em[minIdEncodedMessageLen:]
}
