package primitive

import (
	"github.com/ghia-xch/ghia/pkg/protocol/primitive/message"
	"testing"
)

func TestStringCodec(t *testing.T) {

	var hello = "Hello World."
	var longer = "A much longer string that should be truncated on decode."

	var helloStr = String{&hello}
	var longerStr = String{&longer}

	var em message.EncodedMessage
	var err error

	var encoder = message.NewMessageEncoder(1024)

	encoder.Reset(message.NullType, &message.NullId)

	if em, err = helloStr.Encode(encoder); err != nil {
		t.Error("String encode failed.")
	}

	dec := message.NewMessageDecoder()

	dec.Reset(em)

	if err = longerStr.Decode(dec); err != nil {
		t.Error("String decode failed.")
	}

	if len(*helloStr.string) != len(*longerStr.string) {
		t.Errorf("String length mismatch. Expected \"%d\", got \"%d.\"", len(*helloStr.string), len(*longerStr.string))
	}

	if longerStr.String() != helloStr.String() {
		t.Errorf("String mismatch. Expected \"%v\", got \"%v\"", helloStr, longerStr)
	}
}
