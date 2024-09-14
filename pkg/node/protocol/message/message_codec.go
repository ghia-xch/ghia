package message

import (
	"encoding/binary"
	"lukechampine.com/uint128"
)

var MessageAttributeNotEncodableError error

type Encodable interface {
	Encode(enc *MessageEncoder) (em EncodedMessage, err error)
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

		case Encodable:

			if em, err = a.Encode(m); err != nil {
				return
			}

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

		case uint64:

			buf = m.raw[m.pos : m.pos+8]
			m.pos += 8

			binary.BigEndian.PutUint64(buf, a)

		case uint128.Uint128:

			buf = m.raw[m.pos : m.pos+8]
			m.pos += 8

			binary.BigEndian.PutUint64(buf, a.Hi)

			buf = m.raw[m.pos : m.pos+8]
			m.pos += 8

			binary.BigEndian.PutUint64(buf, a.Lo)

		case string:

			buf = m.raw[m.pos : m.pos+4]
			m.pos += 4

			binary.BigEndian.PutUint32(buf, uint32(len(a)))

			buf = m.raw[m.pos : m.pos+len(a)]

			copy(buf, a)

			m.pos += len(a)

		case []byte:

			buf = m.raw[m.pos : m.pos+4]
			m.pos += 4

			binary.BigEndian.PutUint32(buf, uint32(len(a)))

			buf = m.raw[m.pos : m.pos+len(a)]

			copy(buf, a)

			m.pos += len(a)

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
	msgType Type
	id      Id
	dataLen uint32
	em      EncodedMessage
	pos     uint32
}

const MinHeaderLength = 6
const MaxHeaderLength = 8
const HasEncodedIdFlag = 1

func (md *MessageDecoder) parseHeader() (err error) {

	// Parse the message type.

	if md.em == nil || len(md.em) < MinHeaderLength {
		return MessageAttributeNotDecodableError
	}

	md.msgType = Type(md.em[0])
	md.pos++

	// Parse the message optional message id.

	var id uint16

	if md.em[md.pos] == HasEncodedIdFlag {

		if len(md.em) < MaxHeaderLength {
			return MessageAttributeNotDecodableError
		}

		if id, err = md.ParseUint16(); err != nil {
			return MessageAttributeNotDecodableError
		}

		md.id = Id(id)
	}

	md.pos++

	// Parse the data frame

	var dataLen uint32

	if dataLen, err = md.ParseUint32(); err != nil {
		return MessageAttributeNotDecodableError
	}

	if uint32(len(md.em[md.pos:])) != dataLen {
		return MessageAttributeNotDecodableError
	}

	return nil
}

func (md *MessageDecoder) Reset(em EncodedMessage) (err error) {

	md.em = em
	md.pos = 0

	return md.parseHeader()
}

func (md *MessageDecoder) Bytes() []byte {
	return md.em
}

func (md *MessageDecoder) Type() Type {
	return md.msgType
}

func (md *MessageDecoder) HasId() bool {
	return md.id != 0
}

func (md *MessageDecoder) Id() Id {
	return md.id
}

func (md *MessageDecoder) ParseUint8() (value uint8, err error) {

	if len(md.em) == 0 {
		return 0, MessageAttributeNotDecodableError
	}

	md.pos++

	return md.em[0], nil
}

const (
	Uint8Len     = 1
	Uint16Len    = 2
	Uint32Len    = 4
	Uint64Len    = 8
	StringMinLen = 5
)

func (md *MessageDecoder) ParseUint16() (value uint16, err error) {

	if len(md.em[md.pos:]) < Uint16Len {
		return 0, MessageAttributeNotDecodableError
	}

	value = binary.BigEndian.Uint16(md.em[md.pos:])

	md.pos += Uint16Len

	return
}

func (md *MessageDecoder) ParseUint32() (value uint32, err error) {

	if len(md.em[md.pos:]) < Uint32Len {
		return 0, MessageAttributeNotDecodableError
	}

	value = binary.BigEndian.Uint32(md.em[md.pos:])

	md.pos += Uint32Len

	return
}

func (md *MessageDecoder) ParseUint64() (value uint64, err error) {

	if len(md.em[md.pos:]) < Uint64Len {
		return 0, MessageAttributeNotDecodableError
	}

	value = binary.BigEndian.Uint64(md.em[md.pos:])

	md.pos += Uint64Len

	return
}

func (md *MessageDecoder) ParseUint128() (value uint128.Uint128, err error) {

	if value.Hi, err = md.ParseUint64(); err != nil {
		return value, err
	}

	if value.Lo, err = md.ParseUint64(); err != nil {
		return value, err
	}

	return
}

func (md *MessageDecoder) ParseString() (value string, err error) {

	var strLen uint32

	if strLen, err = md.ParseUint32(); err != nil {
		return "", err
	}

	if uint32(len(md.em[md.pos:])) < strLen {
		return "", MessageAttributeNotDecodableError
	}

	value = string(md.em[md.pos : md.pos+strLen])

	md.pos += strLen

	return
}

func (md *MessageDecoder) ParseBytes() (value []byte, err error) {

	var valLen uint32

	if valLen, err = md.ParseUint32(); err != nil {
		return nil, err
	}

	if len(md.em[md.pos:]) < int(valLen) {
		return nil, MessageAttributeNotDecodableError
	}

	value = md.em[md.pos : md.pos+valLen]

	md.pos += valLen

	return
}

//func (md *MessageDecoder) ParseHash() (value protocol.Hash, err error) {
//
//	if len(md.em[md.pos:]) < 32 {
//		return protocol.Hash([]byte{}), MessageAttributeNotDecodableError
//	}
//
//	value = protocol.Hash(md.em[md.pos : md.pos+32])
//
//	md.pos += 32
//
//	return
//}

func NewMessageDecoder() *MessageDecoder {
	return &MessageDecoder{em: nil, pos: 0}
}
