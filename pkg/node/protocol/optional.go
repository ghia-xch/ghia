package protocol

import (
	"errors"
	"github.com/ghia-xch/ghia/pkg/node/protocol/codec"
	"reflect"
)

type Optional[T any] struct {
	Value *T
}

func (o Optional[T]) Encode(b []byte) (res []byte, err error) {

	if o.Value == nil {
		return codec.EncodeRaw([]byte{0}, b)
	}

	if res, err = codec.EncodeRaw([]byte{1}, b); err != nil {
		return
	}

	return codec.EncodeElement(o.Value, res)
}

func (o *Optional[T]) Decode(b []byte) (res []byte, err error) {

	if len(b) == 0 {
		return nil, errors.New("invalid hash length")
	}

	if b[0] == 0 {
		return b[1:], nil
	}

	return codec.DecodeElement(reflect.ValueOf(*o.Value), b)
}
