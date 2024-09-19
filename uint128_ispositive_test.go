package bits128_test

import (
	"testing"

	"github.com/reiver/go-bits128"
)

func TestUint128_IsPositive(t *testing.T) {

	tests := []struct{
		Uint128 bits128.Uint128
		Expected bool
	}{
		{
			Uint128: bits128.Uint128FromUint8(0),
			Expected: false,
		},
		{
			Uint128: bits128.Uint128FromUint8(1),
			Expected: true,
		},
		{
			Uint128: bits128.Uint128FromUint8(2),
			Expected: true,
		},
		{
			Uint128: bits128.Uint128FromUint8(3),
			Expected: true,
		},
		{
			Uint128: bits128.Uint128FromUint8(4),
			Expected: true,
		},
		{
			Uint128: bits128.Uint128FromUint8(5),
			Expected: true,
		},



		{
			Uint128: bits128.MaxUint128(),
			Expected: true,
		},
	}

	for testNumber, test := range tests {

		var actual   bool = test.Uint128.IsPositive()
		var expected bool = test.Expected

		if expected != actual {
			t.Errorf("For test #%d, the actual is-positive value is not what was expected.", testNumber)
			t.Logf("EXPECTED: %t", expected)
			t.Logf("ACTUAL:   %t", actual)
			t.Logf("INT128: %s", test.Uint128.HexString())
			continue
		}
	}
}
