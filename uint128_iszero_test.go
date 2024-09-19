package bits128_test

import (
	"testing"

	"github.com/reiver/go-bits128"
)

func TestUint128_IsZero(t *testing.T) {

	tests := []struct{
		Uint128 bits128.Uint128
		Expected bool
	}{
		{
			Uint128: bits128.Uint128FromUint8(0),
			Expected: true,
		},
		{
			Uint128: bits128.Uint128FromUint8(1),
			Expected: false,
		},
		{
			Uint128: bits128.Uint128FromUint8(2),
			Expected: false,
		},
		{
			Uint128: bits128.Uint128FromUint8(3),
			Expected: false,
		},
		{
			Uint128: bits128.Uint128FromUint8(4),
			Expected: false,
		},
		{
			Uint128: bits128.Uint128FromUint8(5),
			Expected: false,
		},



		{
			Uint128: bits128.MaxUint128(),
			Expected: false,
		},
	}

	for testNumber, test := range tests {

		var actual   bool = test.Uint128.IsZero()
		var expected bool = test.Expected

		if expected != actual {
			t.Errorf("For test #%d, the actual is-zero value is not what was expected.", testNumber)
			t.Logf("EXPECTED: %t", expected)
			t.Logf("ACTUAL:   %t", actual)
			t.Logf("INT128: %s", test.Uint128.HexString())
			continue
		}
	}
}
