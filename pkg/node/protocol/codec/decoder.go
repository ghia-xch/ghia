package codec

import (
	"errors"
	"github.com/ghia-xch/ghia/pkg/node/protocol/message"
	"reflect"
)

func Decode(in Encodable, em message.EncodedMessage) error {

	inType := reflect.ValueOf(in)

	if inType.Kind() != reflect.Ptr {
		return errors.New("expected pointer to struct")
	}

	if message.MessageType(em[0]) != in.Type() {
		return errors.New("message types to not match")
	}

	// We should add future support for setting an id here.
	// At this time no message types support them.

	return nil
}
