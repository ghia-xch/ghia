package codec

import (
	"encoding/binary"
	"errors"
	"github.com/davecgh/go-spew/spew"
	"github.com/ghia-xch/ghia/pkg/node/protocol"
	"lukechampine.com/uint128"
	"reflect"
)

type Codeable interface {
	Type() protocol.MessageType
}

const DefaultEncodableSize = 8192

func Encode(in Codeable) (em protocol.EncodedMessage, err error) {

	b := make([]byte, 7, DefaultEncodableSize)

	b[0] = byte(in.Type())

	// We should add future support for setting an id here.
	// At this time no message types support them.

	inType := reflect.ValueOf(in)

	if b, err = encodeValue(inType, b); err != nil {
		return nil, err
	}

	binary.BigEndian.PutUint32(b[3:7], uint32(len(b)))

	return b, nil
}

func encodeValue(in reflect.Value, b []byte) ([]byte, error) {

	if !in.IsValid() {
		return []byte{0}, nil
	}

	switch in.Kind() {
	case reflect.Struct:
		return encodeStruct(in, b)
	case reflect.Ptr:
		return encodeValue(in.Elem(), b)
	default:

		if in.CanInterface() {
			return encodeElem(in.Interface(), b)
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

		f := in.Field(i)

		spew.Dump(f.Kind())

		switch f.Kind() {

		case reflect.Slice:

			//

		case reflect.Struct:

			if b, err = encodeStruct(f, b); err != nil {
				return nil, err
			}

		case reflect.Pointer:

			if b, err = encodeValue(f.Elem(), b); err != nil {
				return nil, err
			}

		default:

			if b, err = encodeElem(f.Interface(), b); err != nil {
				return nil, err
			}

		}
	}

	return b, nil
}

func encodeElem(in any, b []byte) ([]byte, error) {

	switch v := in.(type) {
	case nil:
		return append(b, byte(0)), nil
	case bool:

		if v {
			return []byte{1}, nil
		}

		return []byte{0}, nil

	case uint8:
		return append(b, v), nil
	case uint16:
		return binary.BigEndian.AppendUint16(b, v), nil
	case uint32:
		return binary.BigEndian.AppendUint32(b, v), nil
	case uint64:
		return binary.BigEndian.AppendUint64(b, v), nil
	case uint128.Uint128:
		b = binary.BigEndian.AppendUint64(b, v.Hi)
		return binary.BigEndian.AppendUint64(b, v.Lo), nil
	case string:
		b = binary.BigEndian.AppendUint32(b, uint32(len(v)))
		return append(b, []byte(v)...), nil
	case protocol.Hash:
		return append(b, v.Bytes()...), nil
	case []byte:
		b = binary.BigEndian.AppendUint32(b, uint32(len(v)))
		return append(b, v...), nil
	}

	return nil, errors.New("invalid element type")
}

func Decode(in Codeable, em protocol.EncodedMessage) error {

	inType := reflect.ValueOf(in)

	if inType.Kind() != reflect.Ptr {
		return errors.New("expected pointer to struct")
	}

	if protocol.MessageType(em[0]) != in.Type() {
		return errors.New("message types to not match")
	}

	// We should add future support for setting an id here.
	// At this time no message types support them.

	return nil
}
