package message

import (
	"testing"
)

type encodedMessageTypeTestCase struct {
	em           EncodedMessage
	expectedType Type
}

var encodedMessageTypeTestCases = []encodedMessageTypeTestCase{
	{
		em:           []byte{},
		expectedType: NilMessageType,
	},
	{
		em:           []byte{0},
		expectedType: NilMessageType,
	},
	{
		em:           []byte{255},
		expectedType: Type(255),
	},
}

func TestEncodedMessageType(t *testing.T) {

	for _, testCase := range encodedMessageTypeTestCases {

		if !testCase.em.Type().Equals(testCase.expectedType) {
			t.Errorf("Testing %q. Output %q not equal to expected %q", testCase.em, testCase.em.Type(), testCase.expectedType)
		}
	}
}

const nilId = Id(0)

type encodedMessageIdTestCase struct {
	em         EncodedMessage
	expectedId *Id
}

var encodedMessageIdTestCases = []encodedMessageIdTestCase{
	{
		em:         []byte{},
		expectedId: nil,
	},
	{
		em:         []byte{0},
		expectedId: nil,
	},
	{
		em:         []byte{0, 0, 0, 0, 0, 0},
		expectedId: nil,
	},
	{
		em:         []byte{0, 1, 0, 0, 0, 0},
		expectedId: &nilId,
	},
}
