package bits128_test

import (
	"testing"

	"github.com/reiver/go-bits128"
)

func TestUint128_Xor(t *testing.T) {

	tests := []struct{
		Value1 bits128.Uint128
		Value2 bits128.Uint128
		Expected bits128.Uint128
	}{
		{
			Value1:   bits128.Uint128FromUintsLittleEndian([2]uint{0xFFFFFFFFFFFFFFFF,0xFFFFFFFFFFFFFFFF}),
			Value2:   bits128.Uint128FromUintsLittleEndian([2]uint{0x0000000000000000,0x0000000000000000}),
			Expected: bits128.Uint128FromUintsLittleEndian([2]uint{0xFFFFFFFFFFFFFFFF,0xFFFFFFFFFFFFFFFF}),
		},
		{
			Value1:   bits128.Uint128FromUintsLittleEndian([2]uint{0x0000000000000000,0x0000000000000000}),
			Value2:   bits128.Uint128FromUintsLittleEndian([2]uint{0xFFFFFFFFFFFFFFFF,0xFFFFFFFFFFFFFFFF}),
			Expected: bits128.Uint128FromUintsLittleEndian([2]uint{0xFFFFFFFFFFFFFFFF,0xFFFFFFFFFFFFFFFF}),
		},



		{
			Value1:   bits128.Uint128FromUintsLittleEndian([2]uint{0xFFFFFFFFFFFFFFFF,0xFFFFFFFFFFFFFFFF}),
			Value2:   bits128.Uint128FromUintsLittleEndian([2]uint{0x0000000000000001,0x0000000000000000}),
			Expected: bits128.Uint128FromUintsLittleEndian([2]uint{0xFFFFFFFFFFFFFFFE,0xFFFFFFFFFFFFFFFF}),
		},
		{
			Value1:   bits128.Uint128FromUintsLittleEndian([2]uint{0x0000000000000001,0x0000000000000000}),
			Value2:   bits128.Uint128FromUintsLittleEndian([2]uint{0xFFFFFFFFFFFFFFFF,0xFFFFFFFFFFFFFFFF}),
			Expected: bits128.Uint128FromUintsLittleEndian([2]uint{0xFFFFFFFFFFFFFFFE,0xFFFFFFFFFFFFFFFF}),
		},
		{
			Value1:   bits128.Uint128FromUintsLittleEndian([2]uint{0xFFFFFFFFFFFFFFFF,0xFFFFFFFFFFFFFFFF}),
			Value2:   bits128.Uint128FromUintsLittleEndian([2]uint{0x0000000000000002,0x0000000000000000}),
			Expected: bits128.Uint128FromUintsLittleEndian([2]uint{0xFFFFFFFFFFFFFFFD,0xFFFFFFFFFFFFFFFF}),
		},
		{
			Value1:   bits128.Uint128FromUintsLittleEndian([2]uint{0x0000000000000002,0x0000000000000000}),
			Value2:   bits128.Uint128FromUintsLittleEndian([2]uint{0xFFFFFFFFFFFFFFFF,0xFFFFFFFFFFFFFFFF}),
			Expected: bits128.Uint128FromUintsLittleEndian([2]uint{0xFFFFFFFFFFFFFFFD,0xFFFFFFFFFFFFFFFF}),
		},
		{
			Value1:   bits128.Uint128FromUintsLittleEndian([2]uint{0xFFFFFFFFFFFFFFFF,0xFFFFFFFFFFFFFFFF}),
			Value2:   bits128.Uint128FromUintsLittleEndian([2]uint{0x0000000000000004,0x0000000000000000}),
			Expected: bits128.Uint128FromUintsLittleEndian([2]uint{0xFFFFFFFFFFFFFFFB,0xFFFFFFFFFFFFFFFF}),
		},
		{
			Value1:   bits128.Uint128FromUintsLittleEndian([2]uint{0x0000000000000004,0x0000000000000000}),
			Value2:   bits128.Uint128FromUintsLittleEndian([2]uint{0xFFFFFFFFFFFFFFFF,0xFFFFFFFFFFFFFFFF}),
			Expected: bits128.Uint128FromUintsLittleEndian([2]uint{0xFFFFFFFFFFFFFFFB,0xFFFFFFFFFFFFFFFF}),
		},
		{
			Value1:   bits128.Uint128FromUintsLittleEndian([2]uint{0xFFFFFFFFFFFFFFFF,0xFFFFFFFFFFFFFFFF}),
			Value2:   bits128.Uint128FromUintsLittleEndian([2]uint{0x0000000000000008,0x0000000000000000}),
			Expected: bits128.Uint128FromUintsLittleEndian([2]uint{0xFFFFFFFFFFFFFFF7,0xFFFFFFFFFFFFFFFF}),
		},
		{
			Value1:   bits128.Uint128FromUintsLittleEndian([2]uint{0x0000000000000008,0x0000000000000000}),
			Value2:   bits128.Uint128FromUintsLittleEndian([2]uint{0xFFFFFFFFFFFFFFFF,0xFFFFFFFFFFFFFFFF}),
			Expected: bits128.Uint128FromUintsLittleEndian([2]uint{0xFFFFFFFFFFFFFFF7,0xFFFFFFFFFFFFFFFF}),
		},



		{
			Value1:   bits128.Uint128FromUintsLittleEndian([2]uint{0x60754ec013a1bddf,0xcd3226e094b432b6}),
			Value2:   bits128.Uint128FromUintsLittleEndian([2]uint{0x0b73602cd34bce78,0x45afb17b4e3bdc8d}),
			Expected: bits128.Uint128FromUintsLittleEndian([2]uint{0x6B062EECC0EA73A7,0x889D979BDA8FEE3B}),
		},



		{
			Value1:   bits128.Uint128FromUintsLittleEndian([2]uint{0xFFFFFFFFFFFFFFFF,0xFFFFFFFFFFFFFFFF}),
			Value2:   bits128.Uint128FromUintsLittleEndian([2]uint{0x0000000000000010,0x0000000000000000}),
			Expected: bits128.Uint128FromUintsLittleEndian([2]uint{0xFFFFFFFFFFFFFFEF,0xFFFFFFFFFFFFFFFF}),
		},
		{
			Value1:   bits128.Uint128FromUintsLittleEndian([2]uint{0x0000000000000010,0x0000000000000000}),
			Value2:   bits128.Uint128FromUintsLittleEndian([2]uint{0xFFFFFFFFFFFFFFFF,0xFFFFFFFFFFFFFFFF}),
			Expected: bits128.Uint128FromUintsLittleEndian([2]uint{0xFFFFFFFFFFFFFFEF,0xFFFFFFFFFFFFFFFF}),
		},
		{
			Value1:   bits128.Uint128FromUintsLittleEndian([2]uint{0xFFFFFFFFFFFFFFFF,0xFFFFFFFFFFFFFFFF}),
			Value2:   bits128.Uint128FromUintsLittleEndian([2]uint{0x0000000000000020,0x0000000000000000}),
			Expected: bits128.Uint128FromUintsLittleEndian([2]uint{0xFFFFFFFFFFFFFFDF,0xFFFFFFFFFFFFFFFF}),
		},
		{
			Value1:   bits128.Uint128FromUintsLittleEndian([2]uint{0x0000000000000020,0x0000000000000000}),
			Value2:   bits128.Uint128FromUintsLittleEndian([2]uint{0xFFFFFFFFFFFFFFFF,0xFFFFFFFFFFFFFFFF}),
			Expected: bits128.Uint128FromUintsLittleEndian([2]uint{0xFFFFFFFFFFFFFFDF,0xFFFFFFFFFFFFFFFF}),
		},
		{
			Value1:   bits128.Uint128FromUintsLittleEndian([2]uint{0xFFFFFFFFFFFFFFFF,0xFFFFFFFFFFFFFFFF}),
			Value2:   bits128.Uint128FromUintsLittleEndian([2]uint{0x0000000000000040,0x0000000000000000}),
			Expected: bits128.Uint128FromUintsLittleEndian([2]uint{0xFFFFFFFFFFFFFFBF,0xFFFFFFFFFFFFFFFF}),
		},
		{
			Value1:   bits128.Uint128FromUintsLittleEndian([2]uint{0x0000000000000040,0x0000000000000000}),
			Value2:   bits128.Uint128FromUintsLittleEndian([2]uint{0xFFFFFFFFFFFFFFFF,0xFFFFFFFFFFFFFFFF}),
			Expected: bits128.Uint128FromUintsLittleEndian([2]uint{0xFFFFFFFFFFFFFFBF,0xFFFFFFFFFFFFFFFF}),
		},
		{
			Value1:   bits128.Uint128FromUintsLittleEndian([2]uint{0xFFFFFFFFFFFFFFFF,0xFFFFFFFFFFFFFFFF}),
			Value2:   bits128.Uint128FromUintsLittleEndian([2]uint{0x0000000000000080,0x0000000000000000}),
			Expected: bits128.Uint128FromUintsLittleEndian([2]uint{0xFFFFFFFFFFFFFF7F,0xFFFFFFFFFFFFFFFF}),
		},
		{
			Value1:   bits128.Uint128FromUintsLittleEndian([2]uint{0x0000000000000080,0x0000000000000000}),
			Value2:   bits128.Uint128FromUintsLittleEndian([2]uint{0xFFFFFFFFFFFFFFFF,0xFFFFFFFFFFFFFFFF}),
			Expected: bits128.Uint128FromUintsLittleEndian([2]uint{0xFFFFFFFFFFFFFF7F,0xFFFFFFFFFFFFFFFF}),
		},



		{
			Value1:   bits128.Uint128FromUintsLittleEndian([2]uint{0xFFFFFFFFFFFFFFFF,0xFFFFFFFFFFFFFFFF}),
			Value2:   bits128.Uint128FromUintsLittleEndian([2]uint{0xFFFFFFFFFFFFFFFF,0xFFFFFFFFFFFFFFFF}),
			Expected: bits128.Uint128FromUintsLittleEndian([2]uint{0x0000000000000000,0x0000000000000000}),
		},
	}

	for testNumber, test := range tests {
		var actual bits128.Uint128

		actual.Xor(&test.Value1, &test.Value2)

		{
			expected := test.Expected

			if expected != actual {
				t.Errorf("For test #%d, the actual cmp-value is not what was expected." , testNumber)
				t.Logf("EXPECTED: %v", expected)
				t.Logf("ACTUAL:   %v", actual)
				t.Logf("VALUE-1:  %v", test.Value1)
				t.Logf("VALUE-2:  %v", test.Value2)
				continue
			}
		}
	}
}
