package bits128_test

import (
	"testing"

	"github.com/reiver/go-bits128"
)

func TestInt128_Cmp(t *testing.T) {

	tests := []struct{
		Value1 bits128.Int128
		Value2 bits128.Int128
		Expected int
	}{
		{
			Value1: bits128.Int128FromInt64(-3),
			Value2: bits128.Int128FromInt64(-3),
			Expected: 0,
		},
		{
			Value1: bits128.Int128FromInt64(-3),
			Value2: bits128.Int128FromInt64(-2),
			Expected: -1,
		},
		{
			Value1: bits128.Int128FromInt64(-3),
			Value2: bits128.Int128FromInt64(-1),
			Expected: -1,
		},
		{
			Value1: bits128.Int128FromInt64(-2),
			Value2: bits128.Int128FromInt64(-3),
			Expected: 1,
		},
		{
			Value1: bits128.Int128FromInt64(-2),
			Value2: bits128.Int128FromInt64(-2),
			Expected: 0,
		},
		{
			Value1: bits128.Int128FromInt64(-2),
			Value2: bits128.Int128FromInt64(-1),
			Expected: -1,
		},
		{
			Value1: bits128.Int128FromInt64(-1),
			Value2: bits128.Int128FromInt64(-3),
			Expected: 1,
		},
		{
			Value1: bits128.Int128FromInt64(-1),
			Value2: bits128.Int128FromInt64(-2),
			Expected: 1,
		},
		{
			Value1: bits128.Int128FromInt64(-1),
			Value2: bits128.Int128FromInt64(-1),
			Expected: 0,
		},









		{
			Value1: bits128.Int128FromInt64(-2),
			Value2: bits128.Int128FromInt64(-2),
			Expected: 0,
		},
		{
			Value1: bits128.Int128FromInt64(-2),
			Value2: bits128.Int128FromInt64(-1),
			Expected: -1,
		},
		{
			Value1: bits128.Int128FromInt64(-2),
			Value2: bits128.Int128FromInt64(0),
			Expected: -1,
		},
		{
			Value1: bits128.Int128FromInt64(-1),
			Value2: bits128.Int128FromInt64(-2),
			Expected: 1,
		},
		{
			Value1: bits128.Int128FromInt64(-1),
			Value2: bits128.Int128FromInt64(-1),
			Expected: 0,
		},
		{
			Value1: bits128.Int128FromInt64(-1),
			Value2: bits128.Int128FromInt64(0),
			Expected: -1,
		},
		{
			Value1: bits128.Int128FromInt64(0),
			Value2: bits128.Int128FromInt64(-2),
			Expected: 1,
		},
		{
			Value1: bits128.Int128FromInt64(0),
			Value2: bits128.Int128FromInt64(-1),
			Expected: 1,
		},
		{
			Value1: bits128.Int128FromInt64(0),
			Value2: bits128.Int128FromInt64(0),
			Expected: 0,
		},









		{
			Value1: bits128.Int128FromInt64(-1),
			Value2: bits128.Int128FromInt64(-1),
			Expected: 0,
		},
		{
			Value1: bits128.Int128FromInt64(-1),
			Value2: bits128.Int128FromInt64(0),
			Expected: -1,
		},
		{
			Value1: bits128.Int128FromInt64(-1),
			Value2: bits128.Int128FromInt64(1),
			Expected: -1,
		},
		{
			Value1: bits128.Int128FromInt64(0),
			Value2: bits128.Int128FromInt64(-1),
			Expected: 1,
		},
		{
			Value1: bits128.Int128FromInt64(0),
			Value2: bits128.Int128FromInt64(0),
			Expected: 0,
		},
		{
			Value1: bits128.Int128FromInt64(0),
			Value2: bits128.Int128FromInt64(1),
			Expected: -1,
		},
		{
			Value1: bits128.Int128FromInt64(1),
			Value2: bits128.Int128FromInt64(-1),
			Expected: 1,
		},
		{
			Value1: bits128.Int128FromInt64(1),
			Value2: bits128.Int128FromInt64(0),
			Expected: 1,
		},
		{
			Value1: bits128.Int128FromInt64(1),
			Value2: bits128.Int128FromInt64(1),
			Expected: 0,
		},









		{
			Value1: bits128.Int128FromInt64(0),
			Value2: bits128.Int128FromInt64(0),
			Expected: 0,
		},
		{
			Value1: bits128.Int128FromInt64(0),
			Value2: bits128.Int128FromInt64(1),
			Expected: -1,
		},
		{
			Value1: bits128.Int128FromInt64(0),
			Value2: bits128.Int128FromInt64(2),
			Expected: -1,
		},
		{
			Value1: bits128.Int128FromInt64(1),
			Value2: bits128.Int128FromInt64(0),
			Expected: 1,
		},
		{
			Value1: bits128.Int128FromInt64(1),
			Value2: bits128.Int128FromInt64(1),
			Expected: 0,
		},
		{
			Value1: bits128.Int128FromInt64(1),
			Value2: bits128.Int128FromInt64(2),
			Expected: -1,
		},
		{
			Value1: bits128.Int128FromInt64(2),
			Value2: bits128.Int128FromInt64(0),
			Expected: 1,
		},
		{
			Value1: bits128.Int128FromInt64(2),
			Value2: bits128.Int128FromInt64(1),
			Expected: 1,
		},
		{
			Value1: bits128.Int128FromInt64(2),
			Value2: bits128.Int128FromInt64(2),
			Expected: 0,
		},









		{
			Value1: bits128.Int128FromInt64(1),
			Value2: bits128.Int128FromInt64(1),
			Expected: 0,
		},
		{
			Value1: bits128.Int128FromInt64(1),
			Value2: bits128.Int128FromInt64(2),
			Expected: -1,
		},
		{
			Value1: bits128.Int128FromInt64(1),
			Value2: bits128.Int128FromInt64(3),
			Expected: -1,
		},
		{
			Value1: bits128.Int128FromInt64(2),
			Value2: bits128.Int128FromInt64(1),
			Expected: 1,
		},
		{
			Value1: bits128.Int128FromInt64(2),
			Value2: bits128.Int128FromInt64(2),
			Expected: 0,
		},
		{
			Value1: bits128.Int128FromInt64(2),
			Value2: bits128.Int128FromInt64(3),
			Expected: -1,
		},
		{
			Value1: bits128.Int128FromInt64(3),
			Value2: bits128.Int128FromInt64(1),
			Expected: 1,
		},
		{
			Value1: bits128.Int128FromInt64(3),
			Value2: bits128.Int128FromInt64(2),
			Expected: 1,
		},
		{
			Value1: bits128.Int128FromInt64(3),
			Value2: bits128.Int128FromInt64(3),
			Expected: 0,
		},









		{
			Value1: bits128.Int128FromUint(0x0),
			Value2: bits128.Int128FromUint(0x5),
			Expected: -1,
		},
		{
			Value1: bits128.Int128FromUint(0x1),
			Value2: bits128.Int128FromUint(0x5),
			Expected: -1,
		},
		{
			Value1: bits128.Int128FromUint(0x2),
			Value2: bits128.Int128FromUint(0x5),
			Expected: -1,
		},
		{
			Value1: bits128.Int128FromUint(0x3),
			Value2: bits128.Int128FromUint(0x5),
			Expected: -1,
		},
		{
			Value1: bits128.Int128FromUint(0x4),
			Value2: bits128.Int128FromUint(0x5),
			Expected: -1,
		},
		{
			Value1: bits128.Int128FromUint(0x5),
			Value2: bits128.Int128FromUint(0x5),
			Expected: 0,
		},
		{
			Value1: bits128.Int128FromUint(0x6),
			Value2: bits128.Int128FromUint(0x5),
			Expected: 1,
		},
		{
			Value1: bits128.Int128FromUint(0x7),
			Value2: bits128.Int128FromUint(0x5),
			Expected: 1,
		},
		{
			Value1: bits128.Int128FromUint(0x8),
			Value2: bits128.Int128FromUint(0x5),
			Expected: 1,
		},
		{
			Value1: bits128.Int128FromUint(0x9),
			Value2: bits128.Int128FromUint(0x5),
			Expected: 1,
		},
		{
			Value1: bits128.Int128FromUint(0xA),
			Value2: bits128.Int128FromUint(0x5),
			Expected: 1,
		},
		{
			Value1: bits128.Int128FromUint(0xB),
			Value2: bits128.Int128FromUint(0x5),
			Expected: 1,
		},
		{
			Value1: bits128.Int128FromUint(0xC),
			Value2: bits128.Int128FromUint(0x5),
			Expected: 1,
		},
		{
			Value1: bits128.Int128FromUint(0xD),
			Value2: bits128.Int128FromUint(0x5),
			Expected: 1,
		},
		{
			Value1: bits128.Int128FromUint(0xE),
			Value2: bits128.Int128FromUint(0x5),
			Expected: 1,
		},
		{
			Value1: bits128.Int128FromUint(0xF),
			Value2: bits128.Int128FromUint(0x5),
			Expected: 1,
		},
		{
			Value1: bits128.Int128FromUint(0x10),
			Value2: bits128.Int128FromUint(0x5),
			Expected: 1,
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
				t.Logf("VALUE-1: %s (%v) (is-zero=%t) (is-neg=%t) (is-pos=%t)", test.Value1.HexString(), test.Value1, test.Value1.IsZero(), test.Value1.IsNegative(), test.Value1.IsPositive())
				t.Logf("VALUE-2: %s (%v) (is-zero=%t) (is-neg=%t) (is-pos=%t)", test.Value2.HexString(), test.Value2, test.Value2.IsZero(), test.Value2.IsNegative(), test.Value2.IsPositive())
				continue
			}
		}
	}
}
