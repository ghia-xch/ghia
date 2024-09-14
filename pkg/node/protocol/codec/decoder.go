package codec

import (
	"encoding/binary"
	"errors"
	"github.com/ghia-xch/ghia/pkg/node/protocol/message"
	"reflect"
)

func Decode(in Encodable, em message.EncodedMessage) error {

	inType := reflect.ValueOf(in)

	if in == nil {
		return errors.New("decodable type in is empty.")
	}

	if len(em) == 0 {
		return errors.New("decodable message is empty.")
	}

	if message.Type(em[0]) != in.Type() {
		return errors.New("message types do not match.")
	}

	var err error

	if _, err = decodeValue(inType, em.Data()); err != nil {
		return err
	}

	return nil
}

func decodeValue(in reflect.Value, b []byte) ([]byte, error) {

	switch in.Kind() {
	case reflect.Pointer:
		return decodeValue(in.Elem(), b)
	case reflect.Struct:
		return decodeStruct(in, b)
	case reflect.Slice:

		var err error

		var sliceLen = reflect.New(reflect.TypeOf(uint32(0))).Elem()

		if b, err = DecodeElement(sliceLen, b); err != nil {
			return nil, err
		}

		for i := uint32(0); i < sliceLen.Interface().(uint32); i++ {

			var sliceVal = reflect.New(in.Type().Elem()).Elem()

			if b, err = decodeValue(sliceVal, b); err != nil {
				return nil, err
			}

			in.Set(reflect.Append(in, sliceVal))
		}

	case reflect.Map:

		var err error

		var mapLen = reflect.New(reflect.TypeOf(uint32(0))).Elem()

		if b, err = DecodeElement(mapLen, b); err != nil {
			return nil, err
		}

		for i := uint32(0); i < mapLen.Interface().(uint32); i++ {

			var mapKey = reflect.New(in.Type().Key()).Elem()

			if b, err = DecodeElement(mapKey, b); err != nil {
				return nil, err
			}

			var mapVal = reflect.New(in.Type().Elem()).Elem()

			if b, err = DecodeElement(mapVal, b); err != nil {
				return nil, err
			}

			in.SetMapIndex(mapKey, mapVal)
		}

		return b, nil

	default:
		return DecodeElement(in, b)
	}

	return nil, nil
}

func decodeStruct(in reflect.Value, b []byte) ([]byte, error) {

	if in.Kind() != reflect.Struct {
		return nil, errors.New("expected struct")
	}

	var err error

	for i := 0; i < in.NumField(); i++ {

		if !in.Type().Field(i).IsExported() {
			continue
		}

		if b, err = decodeValue(in.Field(i), b); err != nil {
			return nil, err
		}
	}

	return b, nil
}

type DecodableElement interface {
	Decode(b []byte) ([]byte, error)
}

func DecodeElement(elem reflect.Value, b []byte) ([]byte, error) {

	decoderInterface := reflect.TypeOf((*DecodableElement)(nil)).Elem()

	if elem.Addr().Type().Implements(decoderInterface) {

		res := elem.Addr().MethodByName("Decode").Call([]reflect.Value{reflect.ValueOf(b)})

		if res[1].IsValid() {
			return res[0].Interface().([]byte), nil
		}

		return nil, res[1].Interface().(error)
	}

	switch elem.Kind() {
	case reflect.Bool:
		elem.SetBool(b[0] == 1)
		return b[1:], nil
	case reflect.Uint8:
		elem.SetUint(uint64(b[0]))
		return b[1:], nil
	case reflect.Uint16:
		elem.SetUint(uint64(binary.BigEndian.Uint16(b)))
		return b[2:], nil
	case reflect.Uint32:
		elem.SetUint(uint64(binary.BigEndian.Uint32(b)))
		return b[4:], nil
	case reflect.Uint64:
		elem.SetUint(uint64(binary.BigEndian.Uint64(b)))
		return b[8:], nil
	case reflect.String:

		var strLen uint32 = binary.BigEndian.Uint32(b)

		b = b[4:]

		elem.SetString(string(b[:strLen]))

		return b[strLen:], nil

	}

	return b, errors.New("couldn't find element")
}
