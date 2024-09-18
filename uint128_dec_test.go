package bits128_test

import (
	"testing"

	"github.com/reiver/go-bits128"
)

func TestUint128_Dec(t *testing.T) {

	tests := []struct{
		Value bits128.Uint128
		Expected bits128.Uint128
	}{
		{
			Value:    bits128.Uint128FromUintsLittleEndian([2]uint{0x0000000000000001,0x0000000000000000}),
			Expected: bits128.Uint128FromUintsLittleEndian([2]uint{0x0000000000000000,0x0000000000000000}),
		},
		{
			Value:    bits128.Uint128FromUintsLittleEndian([2]uint{0x0000000000000002,0x0000000000000000}),
			Expected: bits128.Uint128FromUintsLittleEndian([2]uint{0x0000000000000001,0x0000000000000000}),
		},
		{
			Value:    bits128.Uint128FromUintsLittleEndian([2]uint{0x0000000000000003,0x0000000000000000}),
			Expected: bits128.Uint128FromUintsLittleEndian([2]uint{0x0000000000000002,0x0000000000000000}),
		},
		{
			Value:    bits128.Uint128FromUintsLittleEndian([2]uint{0x0000000000000004,0x0000000000000000}),
			Expected: bits128.Uint128FromUintsLittleEndian([2]uint{0x0000000000000003,0x0000000000000000}),
		},
		{
			Value:    bits128.Uint128FromUintsLittleEndian([2]uint{0x0000000000000005,0x0000000000000000}),
			Expected: bits128.Uint128FromUintsLittleEndian([2]uint{0x0000000000000004,0x0000000000000000}),
		},
		{
			Value:    bits128.Uint128FromUintsLittleEndian([2]uint{0x0000000000000007,0x0000000000000000}),
			Expected: bits128.Uint128FromUintsLittleEndian([2]uint{0x0000000000000006,0x0000000000000000}),
		},
		{
			Value:    bits128.Uint128FromUintsLittleEndian([2]uint{0x0000000000000009,0x0000000000000000}),
			Expected: bits128.Uint128FromUintsLittleEndian([2]uint{0x0000000000000008,0x0000000000000000}),
		},
		{
			Value:    bits128.Uint128FromUintsLittleEndian([2]uint{0x000000000000000A,0x0000000000000000}),
			Expected: bits128.Uint128FromUintsLittleEndian([2]uint{0x0000000000000009,0x0000000000000000}),
		},
		{
			Value:    bits128.Uint128FromUintsLittleEndian([2]uint{0x000000000000000B,0x0000000000000000}),
			Expected: bits128.Uint128FromUintsLittleEndian([2]uint{0x000000000000000A,0x0000000000000000}),
		},
		{
			Value:    bits128.Uint128FromUintsLittleEndian([2]uint{0x000000000000000C,0x0000000000000000}),
			Expected: bits128.Uint128FromUintsLittleEndian([2]uint{0x000000000000000B,0x0000000000000000}),
		},
		{
			Value:    bits128.Uint128FromUintsLittleEndian([2]uint{0x000000000000000D,0x0000000000000000}),
			Expected: bits128.Uint128FromUintsLittleEndian([2]uint{0x000000000000000C,0x0000000000000000}),
		},
		{
			Value:    bits128.Uint128FromUintsLittleEndian([2]uint{0x000000000000000E,0x0000000000000000}),
			Expected: bits128.Uint128FromUintsLittleEndian([2]uint{0x000000000000000D,0x0000000000000000}),
		},
		{
			Value:    bits128.Uint128FromUintsLittleEndian([2]uint{0x000000000000000F,0x0000000000000000}),
			Expected: bits128.Uint128FromUintsLittleEndian([2]uint{0x000000000000000E,0x0000000000000000}),
		},
		{
			Value:    bits128.Uint128FromUintsLittleEndian([2]uint{0x0000000000000010,0x0000000000000000}),
			Expected: bits128.Uint128FromUintsLittleEndian([2]uint{0x000000000000000F,0x0000000000000000}),
		},
		{
			Value:    bits128.Uint128FromUintsLittleEndian([2]uint{0x0000000000000011,0x0000000000000000}),
			Expected: bits128.Uint128FromUintsLittleEndian([2]uint{0x0000000000000010,0x0000000000000000}),
		},
		{
			Value:    bits128.Uint128FromUintsLittleEndian([2]uint{0x0000000000000012,0x0000000000000000}),
			Expected: bits128.Uint128FromUintsLittleEndian([2]uint{0x0000000000000011,0x0000000000000000}),
		},
		{
			Value:    bits128.Uint128FromUintsLittleEndian([2]uint{0x0000000000000013,0x0000000000000000}),
			Expected: bits128.Uint128FromUintsLittleEndian([2]uint{0x0000000000000012,0x0000000000000000}),
		},
		{
			Value:    bits128.Uint128FromUintsLittleEndian([2]uint{0x0000000000000014,0x0000000000000000}),
			Expected: bits128.Uint128FromUintsLittleEndian([2]uint{0x0000000000000013,0x0000000000000000}),
		},
		{
			Value:    bits128.Uint128FromUintsLittleEndian([2]uint{0x0000000000000015,0x0000000000000000}),
			Expected: bits128.Uint128FromUintsLittleEndian([2]uint{0x0000000000000014,0x0000000000000000}),
		},
		{
			Value:    bits128.Uint128FromUintsLittleEndian([2]uint{0x0000000000000016,0x0000000000000000}),
			Expected: bits128.Uint128FromUintsLittleEndian([2]uint{0x0000000000000015,0x0000000000000000}),
		},
		{
			Value:    bits128.Uint128FromUintsLittleEndian([2]uint{0x0000000000000017,0x0000000000000000}),
			Expected: bits128.Uint128FromUintsLittleEndian([2]uint{0x0000000000000016,0x0000000000000000}),
		},
		{
			Value:    bits128.Uint128FromUintsLittleEndian([2]uint{0x0000000000000018,0x0000000000000000}),
			Expected: bits128.Uint128FromUintsLittleEndian([2]uint{0x0000000000000017,0x0000000000000000}),
		},
		{
			Value:    bits128.Uint128FromUintsLittleEndian([2]uint{0x0000000000000019,0x0000000000000000}),
			Expected: bits128.Uint128FromUintsLittleEndian([2]uint{0x0000000000000018,0x0000000000000000}),
		},
		{
			Value:    bits128.Uint128FromUintsLittleEndian([2]uint{0x000000000000001A,0x0000000000000000}),
			Expected: bits128.Uint128FromUintsLittleEndian([2]uint{0x0000000000000019,0x0000000000000000}),
		},
		{
			Value:    bits128.Uint128FromUintsLittleEndian([2]uint{0x000000000000001B,0x0000000000000000}),
			Expected: bits128.Uint128FromUintsLittleEndian([2]uint{0x000000000000001A,0x0000000000000000}),
		},
		{
			Value:    bits128.Uint128FromUintsLittleEndian([2]uint{0x000000000000001C,0x0000000000000000}),
			Expected: bits128.Uint128FromUintsLittleEndian([2]uint{0x000000000000001B,0x0000000000000000}),
		},
		{
			Value:    bits128.Uint128FromUintsLittleEndian([2]uint{0x000000000000001D,0x0000000000000000}),
			Expected: bits128.Uint128FromUintsLittleEndian([2]uint{0x000000000000001C,0x0000000000000000}),
		},
		{
			Value:    bits128.Uint128FromUintsLittleEndian([2]uint{0x000000000000001E,0x0000000000000000}),
			Expected: bits128.Uint128FromUintsLittleEndian([2]uint{0x000000000000001D,0x0000000000000000}),
		},
		{
			Value:    bits128.Uint128FromUintsLittleEndian([2]uint{0x000000000000001F,0x0000000000000000}),
			Expected: bits128.Uint128FromUintsLittleEndian([2]uint{0x000000000000001E,0x0000000000000000}),
		},
		{
			Value:    bits128.Uint128FromUintsLittleEndian([2]uint{0x0000000000000020,0x0000000000000000}),
			Expected: bits128.Uint128FromUintsLittleEndian([2]uint{0x000000000000001F,0x0000000000000000}),
		},



		{
			Value:    bits128.Uint128FromUintsLittleEndian([2]uint{0x825C99712043E01D,0xED7BA4708E54465E}),
			Expected: bits128.Uint128FromUintsLittleEndian([2]uint{0x825C99712043E01C,0xED7BA4708E54465E}),
		},



		{
			Value:    bits128.Uint128FromUintsLittleEndian([2]uint{0xFFFFFFFFFFFFFFFB,0xFFFFFFFFFFFFFFFF}),
			Expected: bits128.Uint128FromUintsLittleEndian([2]uint{0xFFFFFFFFFFFFFFFA,0xFFFFFFFFFFFFFFFF}),
		},
		{
			Value:    bits128.Uint128FromUintsLittleEndian([2]uint{0xFFFFFFFFFFFFFFFC,0xFFFFFFFFFFFFFFFF}),
			Expected: bits128.Uint128FromUintsLittleEndian([2]uint{0xFFFFFFFFFFFFFFFB,0xFFFFFFFFFFFFFFFF}),
		},
		{
			Value:    bits128.Uint128FromUintsLittleEndian([2]uint{0xFFFFFFFFFFFFFFFD,0xFFFFFFFFFFFFFFFF}),
			Expected: bits128.Uint128FromUintsLittleEndian([2]uint{0xFFFFFFFFFFFFFFFC,0xFFFFFFFFFFFFFFFF}),
		},
		{
			Value:    bits128.Uint128FromUintsLittleEndian([2]uint{0xFFFFFFFFFFFFFFFE,0xFFFFFFFFFFFFFFFF}),
			Expected: bits128.Uint128FromUintsLittleEndian([2]uint{0xFFFFFFFFFFFFFFFD,0xFFFFFFFFFFFFFFFF}),
		},
		{
			Value:    bits128.Uint128FromUintsLittleEndian([2]uint{0xFFFFFFFFFFFFFFFF,0xFFFFFFFFFFFFFFFF}),
			Expected: bits128.Uint128FromUintsLittleEndian([2]uint{0xFFFFFFFFFFFFFFFE,0xFFFFFFFFFFFFFFFF}),
		},
		{
			Value:    bits128.Uint128FromUintsLittleEndian([2]uint{0x0000000000000000,0x0000000000000000}),
			Expected: bits128.Uint128FromUintsLittleEndian([2]uint{0xFFFFFFFFFFFFFFFF,0xFFFFFFFFFFFFFFFF}),
		},
	}

	for testNumber, test := range tests {
		var actual bits128.Uint128 = test.Value
		actual.Dec()

		{
			expected := test.Expected

			if expected != actual {
				t.Errorf("For test #%d, the actual dec-value is not what was expected." , testNumber)
				t.Logf("EXPECTED: %v", expected)
				t.Logf("ACTUAL:   %v", actual)
				t.Logf("VALUE:    %v", test.Value)
				continue
			}
		}
	}
}
