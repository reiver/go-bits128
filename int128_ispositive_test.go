package bits128_test

import (
	"testing"

	"github.com/reiver/go-bits128"
)

func TestInt128_IsPositive(t *testing.T) {

	tests := []struct{
		Int128 bits128.Int128
		Expected bool
	}{
		{
			Int128: bits128.MinInt128(),
			Expected: false,
		},



		{
			Int128: bits128.Int128FromInt8(-5),
			Expected: false,
		},
		{
			Int128: bits128.Int128FromInt8(-4),
			Expected: false,
		},
		{
			Int128: bits128.Int128FromInt8(-3),
			Expected: false,
		},
		{
			Int128: bits128.Int128FromInt8(-2),
			Expected: false,
		},
		{
			Int128: bits128.Int128FromInt8(-1),
			Expected: false,
		},
		{
			Int128: bits128.Int128FromInt8(0),
			Expected: false,
		},
		{
			Int128: bits128.Int128FromInt8(1),
			Expected: true,
		},
		{
			Int128: bits128.Int128FromInt8(2),
			Expected: true,
		},
		{
			Int128: bits128.Int128FromInt8(3),
			Expected: true,
		},
		{
			Int128: bits128.Int128FromInt8(4),
			Expected: true,
		},
		{
			Int128: bits128.Int128FromInt8(5),
			Expected: true,
		},



		{
			Int128: bits128.MaxInt128(),
			Expected: true,
		},
	}

	for testNumber, test := range tests {

		var actual   bool = test.Int128.IsPositive()
		var expected bool = test.Expected

		if expected != actual {
			t.Errorf("For test #%d, the actual is-positive value is not what was expected.", testNumber)
			t.Logf("EXPECTED: %t", expected)
			t.Logf("ACTUAL:   %t", actual)
			t.Logf("INT128: %s", test.Int128.HexString())
			continue
		}
	}
}
