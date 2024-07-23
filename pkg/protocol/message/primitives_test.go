package message

import (
	"testing"
)

func TestStringCodec(t *testing.T) {

	var hello = "Hello World."
	var longer = "A much longer string that should be truncated on decode."

	var helloStr = String{&hello}
	var longerStr = String{&longer}

	var em EncodedMessage
	var err error

	var encoder = NewMessageEncoder(1024)

	encoder.Reset(NullType, &NullId)

	if em, err = helloStr.Encode(encoder); err != nil {
		t.Error("String encode failed.")
	}

	dec := NewMessageDecoder()

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
