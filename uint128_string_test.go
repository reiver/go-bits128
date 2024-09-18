package bits128

import (
	"testing"
)

func TestUint128_String_uint(t *testing.T) {

	tests := []struct{
		Value uint
		Expected string
	}{
		{
			Value:                                    0x0,
			Expected: "0x00000000000000000000000000000000",
		},
		{
			Value:                                    0x1,
			Expected: "0x00000000000000000000000000000001",
		},
		{
			Value:                                    0x2,
			Expected: "0x00000000000000000000000000000002",
		},
		{
			Value:                                    0x3,
			Expected: "0x00000000000000000000000000000003",
		},
		{
			Value:                                    0x4,
			Expected: "0x00000000000000000000000000000004",
		},
		{
			Value:                                    0x5,
			Expected: "0x00000000000000000000000000000005",
		},
		{
			Value:                                    0x6,
			Expected: "0x00000000000000000000000000000006",
		},
		{
			Value:                                    0x7,
			Expected: "0x00000000000000000000000000000007",
		},
		{
			Value:                                    0x8,
			Expected: "0x00000000000000000000000000000008",
		},
		{
			Value:                                    0x9,
			Expected: "0x00000000000000000000000000000009",
		},
		{
			Value:                                    0xA,
			Expected: "0x0000000000000000000000000000000A",
		},
		{
			Value:                                    0xB,
			Expected: "0x0000000000000000000000000000000B",
		},
		{
			Value:                                    0xC,
			Expected: "0x0000000000000000000000000000000C",
		},
		{
			Value:                                    0xD,
			Expected: "0x0000000000000000000000000000000D",
		},
		{
			Value:                                    0xE,
			Expected: "0x0000000000000000000000000000000E",
		},
		{
			Value:                                    0xF,
			Expected: "0x0000000000000000000000000000000F",
		},
		{
			Value:                                    0x10,
			Expected: "0x00000000000000000000000000000010",
		},



		{
			Value:                                   0x1F,
			Expected: "0x0000000000000000000000000000001F",
		},
		{
			Value:                                   0x20,
			Expected: "0x00000000000000000000000000000020",
		},



		{
			Value:                                   0x2F,
			Expected: "0x0000000000000000000000000000002F",
		},
		{
			Value:                                   0x30,
			Expected: "0x00000000000000000000000000000030",
		},



		{
			Value:                                   0x3F,
			Expected: "0x0000000000000000000000000000003F",
		},
		{
			Value:                                   0x40,
			Expected: "0x00000000000000000000000000000040",
		},



		{
			Value:                                   0xEF,
			Expected: "0x000000000000000000000000000000EF",
		},
		{
			Value:                                   0xF0,
			Expected: "0x000000000000000000000000000000F0",
		},



		{
			Value:                                   0xFF,
			Expected: "0x000000000000000000000000000000FF",
		},
		{
			Value:                                  0x100,
			Expected: "0x00000000000000000000000000000100",
		},



		{
			Value:                                  0x1FF,
			Expected: "0x000000000000000000000000000001FF",
		},
		{
			Value:                                  0x200,
			Expected: "0x00000000000000000000000000000200",
		},



		{
			Value:                                  0x2FF,
			Expected: "0x000000000000000000000000000002FF",
		},
		{
			Value:                                  0x300,
			Expected: "0x00000000000000000000000000000300",
		},



		{
			Value:                                  0x3FF,
			Expected: "0x000000000000000000000000000003FF",
		},
		{
			Value:                                  0x400,
			Expected: "0x00000000000000000000000000000400",
		},



		{
			Value:                                  0x4FF,
			Expected: "0x000000000000000000000000000004FF",
		},
		{
			Value:                                  0x500,
			Expected: "0x00000000000000000000000000000500",
		},



		{
			Value:                                  0xFFF,
			Expected: "0x00000000000000000000000000000FFF",
		},
		{
			Value:                                 0x1000,
			Expected: "0x00000000000000000000000000001000",
		},



		{
			Value:                                 0x1FFF,
			Expected: "0x00000000000000000000000000001FFF",
		},
		{
			Value:                                 0x2000,
			Expected: "0x00000000000000000000000000002000",
		},



		{
			Value:                                 0xFFFF,
			Expected: "0x0000000000000000000000000000FFFF",
		},
		{
			Value:                                0x10000,
			Expected: "0x00000000000000000000000000010000",
		},



		{
			Value:                                0xFFFFF,
			Expected: "0x000000000000000000000000000FFFFF",
		},
		{
			Value:                               0x100000,
			Expected: "0x00000000000000000000000000100000",
		},



		{
			Value:                               0xFFFFFF,
			Expected: "0x00000000000000000000000000FFFFFF",
		},
		{
			Value                             : 0x1000000,
			Expected: "0x00000000000000000000000001000000",
		},



		{
			Value:                              0xFFFFFFF,
			Expected: "0x0000000000000000000000000FFFFFFF",
		},
		{
			Value:                             0x10000000,
			Expected: "0x00000000000000000000000010000000",
		},



		{
			Value:                             0xFFFFFFF0,
			Expected: "0x000000000000000000000000FFFFFFF0",
		},
		{
			Value:                             0xFFFFFFFB,
			Expected: "0x000000000000000000000000FFFFFFFB",
		},
		{
			Value:                             0xFFFFFFFC,
			Expected: "0x000000000000000000000000FFFFFFFC",
		},
		{
			Value:                             0xFFFFFFFD,
			Expected: "0x000000000000000000000000FFFFFFFD",
		},
		{
			Value:                             0xFFFFFFF4,
			Expected: "0x000000000000000000000000FFFFFFF4",
		},
		{
			Value:                             0xFFFFFFF5,
			Expected: "0x000000000000000000000000FFFFFFF5",
		},
		{
			Value:                             0xFFFFFFF6,
			Expected: "0x000000000000000000000000FFFFFFF6",
		},
		{
			Value:                             0xFFFFFFF7,
			Expected: "0x000000000000000000000000FFFFFFF7",
		},
		{
			Value:                             0xFFFFFFF8,
			Expected: "0x000000000000000000000000FFFFFFF8",
		},
		{
			Value:                             0xFFFFFFF9,
			Expected: "0x000000000000000000000000FFFFFFF9",
		},
		{
			Value:                             0xFFFFFFFA,
			Expected: "0x000000000000000000000000FFFFFFFA",
		},
		{
			Value:                             0xFFFFFFFB,
			Expected: "0x000000000000000000000000FFFFFFFB",
		},
		{
			Value:                             0xFFFFFFFC,
			Expected: "0x000000000000000000000000FFFFFFFC",
		},
		{
			Value:                             0xFFFFFFFD,
			Expected: "0x000000000000000000000000FFFFFFFD",
		},
		{
			Value:                             0xFFFFFFFE,
			Expected: "0x000000000000000000000000FFFFFFFE",
		},
		{
			Value:                             0xFFFFFFFF,
			Expected: "0x000000000000000000000000FFFFFFFF",
		},









		{
			Value:                     0xFFFFFFFFFFFFFFFF,
			Expected: "0x0000000000000000FFFFFFFFFFFFFFFF",
		},



		{
			Value:                     0xFEDCBA987654321F,
			Expected: "0x0000000000000000FEDCBA987654321F",
		},
	}

	for testNumber, test := range tests {
		var val Uint128 = Uint128FromUint(test.Value)

		{
			expected := test.Expected
			actual   := val.String()

			if expected != actual {
				t.Errorf("For test #%d, the actual stringer-value is not what was expected." , testNumber)
				t.Logf("EXPECTED: %q", expected)
				t.Logf("ACTUAL:   %q", actual)
				t.Logf("VALUE:    0x%X", test.Value)
				continue
			}
		}
	}
}
