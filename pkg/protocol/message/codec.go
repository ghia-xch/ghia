package message

import (
	"encoding/binary"
)

type EncodedMessage []byte

var MessageAttributeNotEncodableError error

type Encodable interface {
	Encode(enc *MessageEncoder) (em EncodedMessage, err error)
	Decode(dec *MessageDecoder) (err error)
}

type MessageEncoder struct {
	header []byte
	rawLen []byte
	raw    []byte
	pos    int
}

func (m *MessageEncoder) Reset(t Type, id *Id) {

	// The default header is 6 bytes.
	// One byte for the message type,
	// one nil byte for the optional id,
	// and four bytes for the length of the data frame.
	m.header = m.raw[0:6]

	// Assign the message type.
	m.header[0] = uint8(t)

	// Locate the raw data length.
	m.rawLen = m.header[2:6]

	/*
		if id != nil {

			// If an optional id is present, an additional 2 bytes is allocated to the header.
			m.header = m.raw[0:7]

			binary.BigEndian.PutUint16(m.header[1:2], uint16(*id))

			// Locate the raw data length.
			m.rawLen = m.header[3:7]
		}
	*/

	m.pos = len(m.header)
}

func (m *MessageEncoder) Bytes() []byte {
	return m.raw[0:m.pos]
}

func (m *MessageEncoder) Encode(attrs ...any) (em EncodedMessage, err error) {

	var buf []byte

	for _, attr := range attrs {

		switch a := attr.(type) {

		case nil:

			m.raw[m.pos] = 0
			m.pos++

		case uint8:

			m.raw[m.pos] = a
			m.pos++

		case uint16:

			buf = m.raw[m.pos : m.pos+2]
			m.pos += 2

			binary.BigEndian.PutUint16(buf, a)

		case uint32:

			buf = m.raw[m.pos : m.pos+4]
			m.pos += 4

			binary.BigEndian.PutUint32(buf, a)

		case string:

			buf = m.raw[m.pos : m.pos+4]
			m.pos += 4

			binary.BigEndian.PutUint32(buf, uint32(len(a)))

			buf = m.raw[m.pos : m.pos+len(a)]

			copy(buf, []byte(a))

			m.pos += len(a)

		case Encodable:

			if em, err = a.Encode(m); err != nil {
				return
			}

		default:
			return nil, MessageAttributeNotEncodableError
		}
	}

	binary.BigEndian.PutUint32(m.rawLen, uint32(len(m.Bytes())-len(m.header)))

	return m.Bytes(), nil
}

func NewMessageEncoder(size int) *MessageEncoder {

	raw := make([]byte, size)

	return &MessageEncoder{
		raw[0:4],
		raw[1:4],
		raw,
		0,
	}
}

var MessageAttributeNotDecodableError error

type MessageDecoder struct {
	em  EncodedMessage
	pos uint32
}

func (md *MessageDecoder) Reset(em EncodedMessage) {
	md.em = em
	md.pos = 0
}

func (md *MessageDecoder) ParseUint8() (value uint8, err error) {

	if len(md.em) == 0 {
		return 0, MessageAttributeNotDecodableError
	}

	md.pos++

	return md.em[0], nil
}

func (md *MessageDecoder) ParseUint16(em EncodedMessage) (value uint16, err error) {

	if len(em) < 2 {
		return 0, MessageAttributeNotDecodableError
	}

	return binary.BigEndian.Uint16(em), nil
}

func (md *MessageDecoder) ParseUint32(em EncodedMessage) (value uint32, err error) {

	if len(em) < 4 {
		return 0, MessageAttributeNotDecodableError
	}

	return binary.BigEndian.Uint32(em), nil
}

func (md *MessageDecoder) ParseString(em EncodedMessage) (value string, err error) {

	if len(em) < 5 || binary.BigEndian.Uint32(em) < uint32(len(em)-4) {
		return "", MessageAttributeNotDecodableError
	}

	return string(em[4 : 4+binary.BigEndian.Uint32(em)]), err
}
