package protocol

import (
	"encoding/binary"
	"errors"
	"lukechampine.com/uint128"
	"reflect"
)

type Codeable interface {
	Type() MessageType
}

const DefaultEncodableSize = 8192

func Encode(in Codeable) (em EncodedMessage, err error) {

	b := make([]byte, 7, DefaultEncodableSize)

	b[0] = byte(in.Type())

	// We should add future support for setting an id here.
	// At this time no message types support them.

	inType := reflect.ValueOf(in)

	if inType.Kind() != reflect.Ptr {
		return nil, errors.New("expected pointer to struct")
	}

	if b, err = encodeStruct(inType.Elem(), b); err != nil {
		return nil, err
	}

	binary.BigEndian.PutUint32(b[3:7], uint32(len(b)))

	return b, nil
}

func encodeStruct(in reflect.Value, b []byte) ([]byte, error) {

	if in.Kind() != reflect.Struct {
		return nil, errors.New("expected struct")
	}

	var err error

	for i := 0; i < in.NumField(); i++ {

		f := in.Field(i)

		switch f.Kind() {

		case reflect.Slice:

			//

		case reflect.Struct:

			if b, err = encodeStruct(f.Elem(), b); err != nil {
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
	case Hash:
		return append(b, v.Bytes()...), nil
	case []byte:
		b = binary.BigEndian.AppendUint32(b, uint32(len(v)))
		return append(b, v...), nil
	}

	return nil, errors.New("invalid element type")
}

func Decode(in Codeable, em EncodedMessage) error {

	inType := reflect.ValueOf(in)

	if inType.Kind() != reflect.Ptr {
		return errors.New("expected pointer to struct")
	}

	if MessageType(em[0]) != in.Type() {
		return errors.New("message types to not match")
	}

	// We should add future support for setting an id here.
	// At this time no message types support them.

	return nil
}
