package codec

import (
	"bytes"
	"reflect"
	"testing"
)

type decodeTestCase struct {
	encodedValue           []byte
	targetValue            any
	expectedType           reflect.Kind
	expectedValue          any
	expectedRemainingBytes []byte
}

var decodeTestCases = []decodeTestCase{
	{
		encodedValue:           []byte{0, 1},
		targetValue:            true,
		expectedType:           reflect.Bool,
		expectedValue:          false,
		expectedRemainingBytes: []byte{1},
	},
	{
		encodedValue:           []byte{1, 0},
		targetValue:            false,
		expectedType:           reflect.Bool,
		expectedValue:          true,
		expectedRemainingBytes: []byte{0},
	},
}

func TestValueDecode(t *testing.T) {

	var err error

	for _, testCase := range decodeTestCases {

		var remainingBytes []byte

		if remainingBytes, err = decodeValue(reflect.ValueOf(testCase.targetValue), testCase.encodedValue); err != nil {
			t.Errorf("Testing %q. Output %q not equal to expected %q. Error: %q", testCase.encodedValue, testCase.targetValue, testCase.expectedValue, err)
		}

		if bytes.Compare(remainingBytes, testCase.expectedRemainingBytes) != 0 {
			t.Errorf("expected remaining bytes to be %v, got %v", testCase.expectedRemainingBytes, remainingBytes)
		}

		if testCase.expectedType != reflect.TypeOf(testCase.targetValue).Kind() {
			t.Errorf("expected type %v but got %v", testCase.expectedType, reflect.TypeOf(testCase.targetValue).Kind())
		}
	}
}
