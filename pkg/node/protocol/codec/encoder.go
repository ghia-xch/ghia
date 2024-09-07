package codec

import (
	"encoding/binary"
	"errors"
	"github.com/ghia-xch/ghia/pkg/node/protocol"
	log "github.com/sirupsen/logrus"
	"lukechampine.com/uint128"
	"reflect"
)

type Encodable interface {
	Type() protocol.MessageType
}

type EncodableElement interface {
	Encode(enc []byte) ([]byte, error)
}

const DefaultEncodableCapacity = 8192

func Encode(id *protocol.Id, in Encodable) (em protocol.EncodedMessage, err error) {

	var b, headerDataLen []byte
	var headerLen int

	b, headerLen, headerDataLen = prepareHeader(id, in)

	inType := reflect.ValueOf(in)

	if b, err = encodeValue(inType, b); err != nil {
		return nil, err
	}

	binary.BigEndian.PutUint32(headerDataLen, uint32(len(b)-headerLen))

	return b, nil
}

const headerSizeWithoutId = 6
const headerSizeWithId = 8

func prepareHeader(id *protocol.Id, in Encodable) (b []byte, headerLen int, headerDataLen []byte) {

	if id != nil {

		b = make([]byte, headerSizeWithId, DefaultEncodableCapacity)

		b[0] = byte(in.Type())
		b[1] = 1

		binary.BigEndian.PutUint16(b[2:4], uint16(*id))

		return b, headerSizeWithId, b[headerSizeWithId-4 : headerSizeWithId]
	}

	b = make([]byte, headerSizeWithoutId, DefaultEncodableCapacity)

	b[0] = byte(in.Type())
	b[1] = 0

	return b, headerSizeWithoutId, b[headerSizeWithoutId-4 : headerSizeWithoutId]
}

func encodeValue(v reflect.Value, b []byte) ([]byte, error) {

	if !v.IsValid() {
		return []byte{0}, nil
	}

	var err error

	switch v.Kind() {
	case reflect.Pointer:
		return encodeValue(v.Elem(), b)
	case reflect.Struct:
		return encodeStruct(v, b)
	case reflect.Slice:

		// encode the length of the slice as uint32
		if b, err = EncodeElement(uint32(v.Len()), b); err != nil {
			return nil, err
		}

		// iterate and encode the values
		for i := 0; i < v.Len(); i++ {
			if b, err = encodeValue(v.Index(i), b); err != nil {
				return nil, err
			}
		}

	case reflect.Map:

		keys := v.MapKeys()

		// encode the length of the map as uint32
		if b, err = EncodeElement(uint32(len(keys)), b); err != nil {
			return nil, err
		}

		for _, key := range keys {

			// Encode the key
			if b, err = encodeValue(key, b); err != nil {
				return nil, err
			}

			// Encode the value
			if b, err = encodeValue(v.MapIndex(key), b); err != nil {
				return nil, err
			}
		}

		return b, nil

	default:

		if v.CanInterface() {
			return EncodeElement(v.Interface(), b)
		}

		return nil, errors.New("unsupported type")
	}

	return nil, nil
}

func encodeStruct(in reflect.Value, b []byte) ([]byte, error) {

	if in.Kind() != reflect.Struct {
		return nil, errors.New("expected struct")
	}

	var err error

	for i := 0; i < in.NumField(); i++ {

		if !in.Type().Field(i).IsExported() {
			continue
		}

		if b, err = encodeValue(in.Field(i), b); err != nil {
			return nil, err
		}
	}

	return b, nil
}

func EncodeElement(in any, b []byte) ([]byte, error) {

	switch v := in.(type) {
	case nil:
		return append(b, byte(0)), nil
	case bool:

		if v {
			return append(b, byte(1)), nil
		}

		return append(b, byte(0)), nil

	case uint8:
		return append(b, v), nil
	case uint16:
		return binary.BigEndian.AppendUint16(b, v), nil
	case uint32:
		return binary.BigEndian.AppendUint32(b, v), nil
	case uint64:
		return binary.BigEndian.AppendUint64(b, v), nil
	case string:
		b = binary.BigEndian.AppendUint32(b, uint32(len(v)))
		return append(b, []byte(v)...), nil
	case uint128.Uint128:
		b = binary.BigEndian.AppendUint64(b, v.Hi)
		return binary.BigEndian.AppendUint64(b, v.Lo), nil
	case protocol.Hash:
		return append(b, v.Bytes()...), nil
	case []byte:
		b = binary.BigEndian.AppendUint32(b, uint32(len(v)))
		return append(b, v...), nil
	case EncodableElement:
		return v.Encode(b)
	}

	log.Errorf("invalid element type %T", in)

	return nil, errors.New("invalid element type")
}
