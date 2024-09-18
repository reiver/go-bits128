package bits128_test

import (
	"testing"

	"github.com/reiver/go-bits128"
)

func TestUint128_Cmp(t *testing.T) {

	tests := []struct{
		Value1 bits128.Uint128
		Value2 bits128.Uint128
		Expected int
	}{
		{
			Value1: bits128.Uint128FromUint(0x0),
			Value2: bits128.Uint128FromUint(0x5),
			Expected: -1,
		},
		{
			Value1: bits128.Uint128FromUint(0x1),
			Value2: bits128.Uint128FromUint(0x5),
			Expected: -1,
		},
		{
			Value1: bits128.Uint128FromUint(0x2),
			Value2: bits128.Uint128FromUint(0x5),
			Expected: -1,
		},
		{
			Value1: bits128.Uint128FromUint(0x3),
			Value2: bits128.Uint128FromUint(0x5),
			Expected: -1,
		},
		{
			Value1: bits128.Uint128FromUint(0x4),
			Value2: bits128.Uint128FromUint(0x5),
			Expected: -1,
		},
		{
			Value1: bits128.Uint128FromUint(0x5),
			Value2: bits128.Uint128FromUint(0x5),
			Expected: 0,
		},
		{
			Value1: bits128.Uint128FromUint(0x6),
			Value2: bits128.Uint128FromUint(0x5),
			Expected: 1,
		},
		{
			Value1: bits128.Uint128FromUint(0x7),
			Value2: bits128.Uint128FromUint(0x5),
			Expected: 1,
		},
		{
			Value1: bits128.Uint128FromUint(0x8),
			Value2: bits128.Uint128FromUint(0x5),
			Expected: 1,
		},
		{
			Value1: bits128.Uint128FromUint(0x9),
			Value2: bits128.Uint128FromUint(0x5),
			Expected: 1,
		},
		{
			Value1: bits128.Uint128FromUint(0xA),
			Value2: bits128.Uint128FromUint(0x5),
			Expected: 1,
		},
		{
			Value1: bits128.Uint128FromUint(0xB),
			Value2: bits128.Uint128FromUint(0x5),
			Expected: 1,
		},
		{
			Value1: bits128.Uint128FromUint(0xC),
			Value2: bits128.Uint128FromUint(0x5),
			Expected: 1,
		},
		{
			Value1: bits128.Uint128FromUint(0xD),
			Value2: bits128.Uint128FromUint(0x5),
			Expected: 1,
		},
		{
			Value1: bits128.Uint128FromUint(0xE),
			Value2: bits128.Uint128FromUint(0x5),
			Expected: 1,
		},
		{
			Value1: bits128.Uint128FromUint(0xF),
			Value2: bits128.Uint128FromUint(0x5),
			Expected: 1,
		},
		{
			Value1: bits128.Uint128FromUint(0x10),
			Value2: bits128.Uint128FromUint(0x5),
			Expected: 1,
		},









		{
			Value1: bits128.Uint128FromUintsLittleEndian([2]uint{0xac43eff54de6477c,0xa551f311cf94c7d7}),
			Value2: bits128.Uint128FromUintsLittleEndian([2]uint{0xe4f7f439a272aa52,0xfa96edfa7f21f7fe}),
			Expected: -1,
		},
		{
			Value1: bits128.Uint128FromUintsLittleEndian([2]uint{0xe4f7f439a272aa52,0xfa96edfa7f21f7fe}),
			Value2: bits128.Uint128FromUintsLittleEndian([2]uint{0xac43eff54de6477c,0xa551f311cf94c7d7}),
			Expected: 1,
		},



		{
			Value1: bits128.Uint128FromUintsLittleEndian([2]uint{0xFFFFFFFFFFFFFFFE,0xFFFFFFFFFFFFFFFF}),
			Value2: bits128.Uint128FromUintsLittleEndian([2]uint{0xFFFFFFFFFFFFFFFF,0xFFFFFFFFFFFFFFFF}),
			Expected: -1,
		},
		{
			Value1: bits128.Uint128FromUintsLittleEndian([2]uint{0xFFFFFFFFFFFFFFFF,0xFFFFFFFFFFFFFFFF}),
			Value2: bits128.Uint128FromUintsLittleEndian([2]uint{0xFFFFFFFFFFFFFFFE,0xFFFFFFFFFFFFFFFF}),
			Expected: 1,
		},
		{
			Value1: bits128.Uint128FromUintsLittleEndian([2]uint{0xFFFFFFFFFFFFFFFF,0xFFFFFFFFFFFFFFFF}),
			Value2: bits128.Uint128FromUintsLittleEndian([2]uint{0xFFFFFFFFFFFFFFFF,0xFFFFFFFFFFFFFFFF}),
			Expected: 0,
		},
	}

	for testNumber, test := range tests {
		actual := test.Value1.Cmp(&test.Value2)

		{
			expected := test.Expected

			if expected != actual {
				t.Errorf("For test #%d, the actual cmp-value is not what was expected." , testNumber)
				t.Logf("EXPECTED: %d", expected)
				t.Logf("ACTUAL:   %d", actual)
				t.Logf("VALUE-1: %v", test.Value1)
				t.Logf("VALUE-2: %v", test.Value2)
				continue
			}
		}
	}
}
