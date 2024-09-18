package bits128_test

import (
	"testing"

	"github.com/reiver/go-bits128"
)

func TestUint128_String(t *testing.T) {

	tests := []struct{
		Value bits128.Uint128
		Expected string
	}{
		{
			Value:            bits128.Uint128FromUint(0x0),
			Expected: "0x00000000000000000000000000000000",
		},
		{
			Value:            bits128.Uint128FromUint64(0x0),
			Expected:   "0x00000000000000000000000000000000",
		},



		{
			Value:            bits128.Uint128FromUint(0x1),
			Expected: "0x00000000000000000000000000000001",
		},
		{
			Value:            bits128.Uint128FromUint64(0x1),
			Expected:   "0x00000000000000000000000000000001",
		},



		{
			Value:            bits128.Uint128FromUint(0x2),
			Expected: "0x00000000000000000000000000000002",
		},
		{
			Value:            bits128.Uint128FromUint64(0x2),
			Expected:   "0x00000000000000000000000000000002",
		},



		{
			Value:            bits128.Uint128FromUint(0x3),
			Expected: "0x00000000000000000000000000000003",
		},
		{
			Value:            bits128.Uint128FromUint(0x4),
			Expected: "0x00000000000000000000000000000004",
		},
		{
			Value:            bits128.Uint128FromUint(0x5),
			Expected: "0x00000000000000000000000000000005",
		},
		{
			Value:            bits128.Uint128FromUint(0x6),
			Expected: "0x00000000000000000000000000000006",
		},
		{
			Value:            bits128.Uint128FromUint(0x7),
			Expected: "0x00000000000000000000000000000007",
		},
		{
			Value:            bits128.Uint128FromUint(0x8),
			Expected: "0x00000000000000000000000000000008",
		},
		{
			Value:            bits128.Uint128FromUint(0x9),
			Expected: "0x00000000000000000000000000000009",
		},
		{
			Value:            bits128.Uint128FromUint(0xA),
			Expected: "0x0000000000000000000000000000000A",
		},
		{
			Value:            bits128.Uint128FromUint(0xB),
			Expected: "0x0000000000000000000000000000000B",
		},
		{
			Value:            bits128.Uint128FromUint(0xC),
			Expected: "0x0000000000000000000000000000000C",
		},
		{
			Value:            bits128.Uint128FromUint(0xD),
			Expected: "0x0000000000000000000000000000000D",
		},
		{
			Value:            bits128.Uint128FromUint(0xE),
			Expected: "0x0000000000000000000000000000000E",
		},
		{
			Value:            bits128.Uint128FromUint(0xF),
			Expected: "0x0000000000000000000000000000000F",
		},
		{
			Value:            bits128.Uint128FromUint(0x10),
			Expected: "0x00000000000000000000000000000010",
		},



		{
			Value:           bits128.Uint128FromUint(0x1F),
			Expected: "0x0000000000000000000000000000001F",
		},
		{
			Value:           bits128.Uint128FromUint(0x20),
			Expected: "0x00000000000000000000000000000020",
		},



		{
			Value:           bits128.Uint128FromUint(0x2F),
			Expected: "0x0000000000000000000000000000002F",
		},
		{
			Value:           bits128.Uint128FromUint(0x30),
			Expected: "0x00000000000000000000000000000030",
		},



		{
			Value:           bits128.Uint128FromUint(0x3F),
			Expected: "0x0000000000000000000000000000003F",
		},
		{
			Value:           bits128.Uint128FromUint(0x40),
			Expected: "0x00000000000000000000000000000040",
		},



		{
			Value:           bits128.Uint128FromUint(0xEF),
			Expected: "0x000000000000000000000000000000EF",
		},
		{
			Value:           bits128.Uint128FromUint(0xF0),
			Expected: "0x000000000000000000000000000000F0",
		},



		{
			Value:           bits128.Uint128FromUint(0xFF),
			Expected: "0x000000000000000000000000000000FF",
		},
		{
			Value:          bits128.Uint128FromUint(0x100),
			Expected: "0x00000000000000000000000000000100",
		},



		{
			Value:          bits128.Uint128FromUint(0x1FF),
			Expected: "0x000000000000000000000000000001FF",
		},
		{
			Value:          bits128.Uint128FromUint(0x200),
			Expected: "0x00000000000000000000000000000200",
		},



		{
			Value:          bits128.Uint128FromUint(0x2FF),
			Expected: "0x000000000000000000000000000002FF",
		},
		{
			Value:          bits128.Uint128FromUint(0x300),
			Expected: "0x00000000000000000000000000000300",
		},



		{
			Value:          bits128.Uint128FromUint(0x3FF),
			Expected: "0x000000000000000000000000000003FF",
		},
		{
			Value:          bits128.Uint128FromUint(0x400),
			Expected: "0x00000000000000000000000000000400",
		},



		{
			Value:          bits128.Uint128FromUint(0x4FF),
			Expected: "0x000000000000000000000000000004FF",
		},
		{
			Value:          bits128.Uint128FromUint(0x500),
			Expected: "0x00000000000000000000000000000500",
		},



		{
			Value:          bits128.Uint128FromUint(0xFFF),
			Expected: "0x00000000000000000000000000000FFF",
		},
		{
			Value:         bits128.Uint128FromUint(0x1000),
			Expected: "0x00000000000000000000000000001000",
		},



		{
			Value:         bits128.Uint128FromUint(0x1FFF),
			Expected: "0x00000000000000000000000000001FFF",
		},
		{
			Value:         bits128.Uint128FromUint(0x2000),
			Expected: "0x00000000000000000000000000002000",
		},



		{
			Value:         bits128.Uint128FromUint(0xFFFF),
			Expected: "0x0000000000000000000000000000FFFF",
		},
		{
			Value:        bits128.Uint128FromUint(0x10000),
			Expected: "0x00000000000000000000000000010000",
		},



		{
			Value:        bits128.Uint128FromUint(0xFFFFF),
			Expected: "0x000000000000000000000000000FFFFF",
		},
		{
			Value:       bits128.Uint128FromUint(0x100000),
			Expected: "0x00000000000000000000000000100000",
		},



		{
			Value:       bits128.Uint128FromUint(0xFFFFFF),
			Expected: "0x00000000000000000000000000FFFFFF",
		},
		{
			Value:      bits128.Uint128FromUint(0x1000000),
			Expected: "0x00000000000000000000000001000000",
		},



		{
			Value:      bits128.Uint128FromUint(0xFFFFFFF),
			Expected: "0x0000000000000000000000000FFFFFFF",
		},
		{
			Value:     bits128.Uint128FromUint(0x10000000),
			Expected: "0x00000000000000000000000010000000",
		},



		{
			Value:     bits128.Uint128FromUint(0xFFFFFFF0),
			Expected: "0x000000000000000000000000FFFFFFF0",
		},
		{
			Value:     bits128.Uint128FromUint(0xFFFFFFFB),
			Expected: "0x000000000000000000000000FFFFFFFB",
		},
		{
			Value:     bits128.Uint128FromUint(0xFFFFFFFC),
			Expected: "0x000000000000000000000000FFFFFFFC",
		},
		{
			Value:     bits128.Uint128FromUint(0xFFFFFFFD),
			Expected: "0x000000000000000000000000FFFFFFFD",
		},
		{
			Value:     bits128.Uint128FromUint(0xFFFFFFF4),
			Expected: "0x000000000000000000000000FFFFFFF4",
		},
		{
			Value:     bits128.Uint128FromUint(0xFFFFFFF5),
			Expected: "0x000000000000000000000000FFFFFFF5",
		},
		{
			Value:     bits128.Uint128FromUint(0xFFFFFFF6),
			Expected: "0x000000000000000000000000FFFFFFF6",
		},
		{
			Value:     bits128.Uint128FromUint(0xFFFFFFF7),
			Expected: "0x000000000000000000000000FFFFFFF7",
		},
		{
			Value:     bits128.Uint128FromUint(0xFFFFFFF8),
			Expected: "0x000000000000000000000000FFFFFFF8",
		},
		{
			Value:     bits128.Uint128FromUint(0xFFFFFFF9),
			Expected: "0x000000000000000000000000FFFFFFF9",
		},
		{
			Value:     bits128.Uint128FromUint(0xFFFFFFFA),
			Expected: "0x000000000000000000000000FFFFFFFA",
		},
		{
			Value:     bits128.Uint128FromUint(0xFFFFFFFB),
			Expected: "0x000000000000000000000000FFFFFFFB",
		},
		{
			Value:     bits128.Uint128FromUint(0xFFFFFFFC),
			Expected: "0x000000000000000000000000FFFFFFFC",
		},
		{
			Value:     bits128.Uint128FromUint(0xFFFFFFFD),
			Expected: "0x000000000000000000000000FFFFFFFD",
		},



		{
			Value:     bits128.Uint128FromUint(0xFFFFFFFE),
			Expected: "0x000000000000000000000000FFFFFFFE",
		},
		{
			Value:     bits128.Uint128FromUint64(0xFFFFFFFE),
			Expected:   "0x000000000000000000000000FFFFFFFE",
		},



		{
			Value:     bits128.Uint128FromUint(0xFFFFFFFF),
			Expected: "0x000000000000000000000000FFFFFFFF",
		},
		{
			Value:     bits128.Uint128FromUint64(0xFFFFFFFF),
			Expected:   "0x000000000000000000000000FFFFFFFF",
		},









		{
			Value: bits128.Uint128FromUint(0xFEDCBA987654321F),
			Expected:     "0x0000000000000000FEDCBA987654321F",
		},
		{
			Value: bits128.Uint128FromUint64(0xFEDCBA987654321F),
			Expected:       "0x0000000000000000FEDCBA987654321F",
		},



		{
			Value: bits128.Uint128FromUint(0xFFFFFFFFFFFFFFFF),
			Expected:     "0x0000000000000000FFFFFFFFFFFFFFFF",
		},
		{
			Value: bits128.Uint128FromUint64(0xFFFFFFFFFFFFFFFF),
			Expected:       "0x0000000000000000FFFFFFFFFFFFFFFF",
		},









		{
			Value: bits128.Uint128FromUintsLittleEndian([2]uint{0xac43eff54de6477c,0xa551f311cf94c7d7}),
			Expected:                          "0xA551F311CF94C7D7AC43EFF54DE6477C",
		},



		{
			Value: bits128.Uint128FromUintsLittleEndian([2]uint{0xe4f7f439a272aa52,0xfa96edfa7f21f7fe}),
			Expected:                          "0xFA96EDFA7F21F7FEE4F7F439A272AA52",
		},
	}

	for testNumber, test := range tests {
		var val bits128.Uint128 = test.Value

		{
			expected := test.Expected
			actual   := val.String()

			if expected != actual {
				t.Errorf("For test #%d, the actual stringer-value is not what was expected." , testNumber)
				t.Logf("EXPECTED: %q", expected)
				t.Logf("ACTUAL:   %q", actual)
				continue
			}
		}
	}
}
