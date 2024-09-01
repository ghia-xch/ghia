package codec

import (
	"bytes"
	"lukechampine.com/uint128"
	"math"
	"reflect"
	"testing"
)

type encodeTestCase struct {
	value    reflect.Value
	expected []byte
}

type sampleStruct struct {
	a bool
	B bool
}

var (
	valFalse         = false
	valTrue          = true
	valUint8Zero     = uint8(0)
	valUint8Max      = uint8(math.MaxUint8)
	valUint16Zero    = uint16(0)
	valUint16Max     = uint16(math.MaxUint16)
	valUint32Zero    = uint32(0)
	valUint32Max     = uint32(math.MaxUint32)
	valUint64Zero    = uint64(0)
	valUint64Max     = uint64(math.MaxUint64)
	valUint128Zero   = uint128.From64(0)
	valUint128Max, _ = uint128.FromString("340282366920938463463374607431768211455")
	valString        = "hello world"
	valSampleStruct  = sampleStruct{
		true,
		true,
	}
)

var encodeTestCases = []encodeTestCase{
	encodeTestCase{
		value:    reflect.ValueOf(nil),
		expected: []byte{0},
	},
	encodeTestCase{
		value:    reflect.ValueOf(false),
		expected: []byte{0},
	},
	encodeTestCase{
		value:    reflect.ValueOf(true),
		expected: []byte{1},
	},
	encodeTestCase{
		value:    reflect.ValueOf(&valFalse),
		expected: []byte{0},
	},
	encodeTestCase{
		value:    reflect.ValueOf(&valTrue),
		expected: []byte{1},
	},
	encodeTestCase{
		value:    reflect.ValueOf(valUint8Zero),
		expected: []byte{0},
	},
	encodeTestCase{
		value:    reflect.ValueOf(&valUint8Zero),
		expected: []byte{0},
	},
	encodeTestCase{
		value:    reflect.ValueOf(valUint8Max),
		expected: []byte{255},
	},
	encodeTestCase{
		value:    reflect.ValueOf(&valUint8Max),
		expected: []byte{255},
	},
	encodeTestCase{
		value:    reflect.ValueOf(valUint16Zero),
		expected: []byte{0, 0},
	},
	encodeTestCase{
		value:    reflect.ValueOf(&valUint16Zero),
		expected: []byte{0, 0},
	},
	encodeTestCase{
		value:    reflect.ValueOf(valUint16Max),
		expected: []byte{255, 255},
	},
	encodeTestCase{
		value:    reflect.ValueOf(&valUint16Max),
		expected: []byte{255, 255},
	},
	encodeTestCase{
		value:    reflect.ValueOf(valUint32Zero),
		expected: []byte{0, 0, 0, 0},
	},
	encodeTestCase{
		value:    reflect.ValueOf(&valUint32Zero),
		expected: []byte{0, 0, 0, 0},
	},
	encodeTestCase{
		value:    reflect.ValueOf(valUint32Max),
		expected: []byte{255, 255, 255, 255},
	},
	encodeTestCase{
		value:    reflect.ValueOf(&valUint32Max),
		expected: []byte{255, 255, 255, 255},
	},
	encodeTestCase{
		value:    reflect.ValueOf(valUint64Zero),
		expected: []byte{0, 0, 0, 0, 0, 0, 0, 0},
	},
	encodeTestCase{
		value:    reflect.ValueOf(&valUint64Zero),
		expected: []byte{0, 0, 0, 0, 0, 0, 0, 0},
	},
	encodeTestCase{
		value:    reflect.ValueOf(valUint64Max),
		expected: []byte{255, 255, 255, 255, 255, 255, 255, 255},
	},
	encodeTestCase{
		value:    reflect.ValueOf(&valUint64Max),
		expected: []byte{255, 255, 255, 255, 255, 255, 255, 255},
	},
	encodeTestCase{
		value: reflect.ValueOf(valString),
		expected: []byte{
			0, 0, 0, 11, // length 11
			104, 101, 108, 108, 111, 32, 119, 111, 114, 108, 100, // "hello world" in decimal
		},
	},
	encodeTestCase{
		value: reflect.ValueOf(&valString),
		expected: []byte{
			0, 0, 0, 11, // length 11
			104, 101, 108, 108, 111, 32, 119, 111, 114, 108, 100, // "hello world" in decimal
		},
	},
	// structs
	encodeTestCase{
		value:    reflect.ValueOf(valUint128Zero),
		expected: []byte{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
	},
	encodeTestCase{
		value:    reflect.ValueOf(&valUint128Zero),
		expected: []byte{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
	},
	encodeTestCase{
		value:    reflect.ValueOf(valUint128Max),
		expected: []byte{255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255},
	},
	encodeTestCase{
		value:    reflect.ValueOf(&valUint128Max),
		expected: []byte{255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255},
	},
	encodeTestCase{
		value:    reflect.ValueOf(valSampleStruct),
		expected: []byte{1},
	},
}

func TestValueEncode(t *testing.T) {

	var err error

	for _, testCase := range encodeTestCases {

		bArr := make([]byte, 0)

		if bArr, err = encodeValue(testCase.value, bArr); err != nil {
			t.Errorf("Testing %q. Output %q not equal to expected %q. Error: %q", testCase.value, bArr, testCase.expected, err)
		}

		if !bytes.Equal(bArr, testCase.expected) {
			t.Errorf("Testing %q. Output %q not equal to expected %q", testCase.value, bArr, testCase.expected)
		}
	}
}
