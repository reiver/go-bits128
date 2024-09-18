package bits128_test

import (
	"testing"

	"github.com/reiver/go-bits128"
)

func TestUint128_And(t *testing.T) {

	tests := []struct{
		Value1 bits128.Uint128
		Value2 bits128.Uint128
		Expected bits128.Uint128
	}{
		{
			Value1:   bits128.Uint128FromUintsLittleEndian([2]uint{0xFFFFFFFFFFFFFFFF,0xFFFFFFFFFFFFFFFF}),
			Value2:   bits128.Uint128FromUintsLittleEndian([2]uint{0x0000000000000000,0x0000000000000000}),
			Expected: bits128.Uint128FromUintsLittleEndian([2]uint{0x0000000000000000,0x0000000000000000}),
		},
		{
			Value1:   bits128.Uint128FromUintsLittleEndian([2]uint{0x0000000000000000,0x0000000000000000}),
			Value2:   bits128.Uint128FromUintsLittleEndian([2]uint{0xFFFFFFFFFFFFFFFF,0xFFFFFFFFFFFFFFFF}),
			Expected: bits128.Uint128FromUintsLittleEndian([2]uint{0x0000000000000000,0x0000000000000000}),
		},



		{
			Value1:   bits128.Uint128FromUintsLittleEndian([2]uint{0xFFFFFFFFFFFFFFFF,0xFFFFFFFFFFFFFFFF}),
			Value2:   bits128.Uint128FromUintsLittleEndian([2]uint{0x0000000000000001,0x0000000000000000}),
			Expected: bits128.Uint128FromUintsLittleEndian([2]uint{0x0000000000000001,0x0000000000000000}),
		},
		{
			Value1:   bits128.Uint128FromUintsLittleEndian([2]uint{0x0000000000000001,0x0000000000000000}),
			Value2:   bits128.Uint128FromUintsLittleEndian([2]uint{0xFFFFFFFFFFFFFFFF,0xFFFFFFFFFFFFFFFF}),
			Expected: bits128.Uint128FromUintsLittleEndian([2]uint{0x0000000000000001,0x0000000000000000}),
		},
		{
			Value1:   bits128.Uint128FromUintsLittleEndian([2]uint{0xFFFFFFFFFFFFFFFF,0xFFFFFFFFFFFFFFFF}),
			Value2:   bits128.Uint128FromUintsLittleEndian([2]uint{0x0000000000000002,0x0000000000000000}),
			Expected: bits128.Uint128FromUintsLittleEndian([2]uint{0x0000000000000002,0x0000000000000000}),
		},
		{
			Value1:   bits128.Uint128FromUintsLittleEndian([2]uint{0x0000000000000002,0x0000000000000000}),
			Value2:   bits128.Uint128FromUintsLittleEndian([2]uint{0xFFFFFFFFFFFFFFFF,0xFFFFFFFFFFFFFFFF}),
			Expected: bits128.Uint128FromUintsLittleEndian([2]uint{0x0000000000000002,0x0000000000000000}),
		},
		{
			Value1:   bits128.Uint128FromUintsLittleEndian([2]uint{0xFFFFFFFFFFFFFFFF,0xFFFFFFFFFFFFFFFF}),
			Value2:   bits128.Uint128FromUintsLittleEndian([2]uint{0x0000000000000004,0x0000000000000000}),
			Expected: bits128.Uint128FromUintsLittleEndian([2]uint{0x0000000000000004,0x0000000000000000}),
		},
		{
			Value1:   bits128.Uint128FromUintsLittleEndian([2]uint{0x0000000000000004,0x0000000000000000}),
			Value2:   bits128.Uint128FromUintsLittleEndian([2]uint{0xFFFFFFFFFFFFFFFF,0xFFFFFFFFFFFFFFFF}),
			Expected: bits128.Uint128FromUintsLittleEndian([2]uint{0x0000000000000004,0x0000000000000000}),
		},
		{
			Value1:   bits128.Uint128FromUintsLittleEndian([2]uint{0xFFFFFFFFFFFFFFFF,0xFFFFFFFFFFFFFFFF}),
			Value2:   bits128.Uint128FromUintsLittleEndian([2]uint{0x0000000000000008,0x0000000000000000}),
			Expected: bits128.Uint128FromUintsLittleEndian([2]uint{0x0000000000000008,0x0000000000000000}),
		},
		{
			Value1:   bits128.Uint128FromUintsLittleEndian([2]uint{0x0000000000000008,0x0000000000000000}),
			Value2:   bits128.Uint128FromUintsLittleEndian([2]uint{0xFFFFFFFFFFFFFFFF,0xFFFFFFFFFFFFFFFF}),
			Expected: bits128.Uint128FromUintsLittleEndian([2]uint{0x0000000000000008,0x0000000000000000}),
		},



		{
			Value1:   bits128.Uint128FromUintsLittleEndian([2]uint{0x60754ec013a1bddf,0xcd3226e094b432b6}),
			Value2:   bits128.Uint128FromUintsLittleEndian([2]uint{0x0b73602cd34bce78,0x45afb17b4e3bdc8d}),
			Expected: bits128.Uint128FromUintsLittleEndian([2]uint{0x0071400013018C58,0x4522206004301084}),
		},



		{
			Value1:   bits128.Uint128FromUintsLittleEndian([2]uint{0xFFFFFFFFFFFFFFFF,0xFFFFFFFFFFFFFFFF}),
			Value2:   bits128.Uint128FromUintsLittleEndian([2]uint{0x0000000000000010,0x0000000000000000}),
			Expected: bits128.Uint128FromUintsLittleEndian([2]uint{0x0000000000000010,0x0000000000000000}),
		},
		{
			Value1:   bits128.Uint128FromUintsLittleEndian([2]uint{0x0000000000000010,0x0000000000000000}),
			Value2:   bits128.Uint128FromUintsLittleEndian([2]uint{0xFFFFFFFFFFFFFFFF,0xFFFFFFFFFFFFFFFF}),
			Expected: bits128.Uint128FromUintsLittleEndian([2]uint{0x0000000000000010,0x0000000000000000}),
		},
		{
			Value1:   bits128.Uint128FromUintsLittleEndian([2]uint{0xFFFFFFFFFFFFFFFF,0xFFFFFFFFFFFFFFFF}),
			Value2:   bits128.Uint128FromUintsLittleEndian([2]uint{0x0000000000000020,0x0000000000000000}),
			Expected: bits128.Uint128FromUintsLittleEndian([2]uint{0x0000000000000020,0x0000000000000000}),
		},
		{
			Value1:   bits128.Uint128FromUintsLittleEndian([2]uint{0x0000000000000020,0x0000000000000000}),
			Value2:   bits128.Uint128FromUintsLittleEndian([2]uint{0xFFFFFFFFFFFFFFFF,0xFFFFFFFFFFFFFFFF}),
			Expected: bits128.Uint128FromUintsLittleEndian([2]uint{0x0000000000000020,0x0000000000000000}),
		},
		{
			Value1:   bits128.Uint128FromUintsLittleEndian([2]uint{0xFFFFFFFFFFFFFFFF,0xFFFFFFFFFFFFFFFF}),
			Value2:   bits128.Uint128FromUintsLittleEndian([2]uint{0x0000000000000040,0x0000000000000000}),
			Expected: bits128.Uint128FromUintsLittleEndian([2]uint{0x0000000000000040,0x0000000000000000}),
		},
		{
			Value1:   bits128.Uint128FromUintsLittleEndian([2]uint{0x0000000000000040,0x0000000000000000}),
			Value2:   bits128.Uint128FromUintsLittleEndian([2]uint{0xFFFFFFFFFFFFFFFF,0xFFFFFFFFFFFFFFFF}),
			Expected: bits128.Uint128FromUintsLittleEndian([2]uint{0x0000000000000040,0x0000000000000000}),
		},
		{
			Value1:   bits128.Uint128FromUintsLittleEndian([2]uint{0xFFFFFFFFFFFFFFFF,0xFFFFFFFFFFFFFFFF}),
			Value2:   bits128.Uint128FromUintsLittleEndian([2]uint{0x0000000000000080,0x0000000000000000}),
			Expected: bits128.Uint128FromUintsLittleEndian([2]uint{0x0000000000000080,0x0000000000000000}),
		},
		{
			Value1:   bits128.Uint128FromUintsLittleEndian([2]uint{0x0000000000000080,0x0000000000000000}),
			Value2:   bits128.Uint128FromUintsLittleEndian([2]uint{0xFFFFFFFFFFFFFFFF,0xFFFFFFFFFFFFFFFF}),
			Expected: bits128.Uint128FromUintsLittleEndian([2]uint{0x0000000000000080,0x0000000000000000}),
		},



		{
			Value1:   bits128.Uint128FromUintsLittleEndian([2]uint{0xFFFFFFFFFFFFFFFF,0xFFFFFFFFFFFFFFFF}),
			Value2:   bits128.Uint128FromUintsLittleEndian([2]uint{0xFFFFFFFFFFFFFFFF,0xFFFFFFFFFFFFFFFF}),
			Expected: bits128.Uint128FromUintsLittleEndian([2]uint{0xFFFFFFFFFFFFFFFF,0xFFFFFFFFFFFFFFFF}),
		},
	}

	for testNumber, test := range tests {
		var actual bits128.Uint128

		actual.And(&test.Value1, &test.Value2)

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
