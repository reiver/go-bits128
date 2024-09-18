package bits128

import (
	"testing"
)

func TestUint128FromUint(t *testing.T) {

	tests := []struct{
		Value uint
	}{
		{
			Value: 0x0,
		},
		{
			Value: 0x1,
		},
		{
			Value: 0x2,
		},
		{
			Value: 0x3,
		},
		{
			Value: 0x4,
		},
		{
			Value: 0x5,
		},
		{
			Value: 0x6,
		},
		{
			Value: 0x7,
		},
		{
			Value: 0x8,
		},
		{
			Value: 0x9,
		},
		{
			Value: 0xA,
		},
		{
			Value: 0xB,
		},
		{
			Value: 0xC,
		},
		{
			Value: 0xD,
		},
		{
			Value: 0xE,
		},
		{
			Value: 0xF,
		},
		{
			Value: 0x10,
		},



		{
			Value: 0x1F,
		},
		{
			Value: 0x20,
		},



		{
			Value: 0x2F,
		},
		{
			Value: 0x30,
		},



		{
			Value: 0x3F,
		},
		{
			Value: 0x40,
		},



		{
			Value: 0xEF,
		},
		{
			Value: 0xF0,
		},



		{
			Value: 0xFF,
		},
		{
			Value: 0x100,
		},



		{
			Value: 0x1FF,
		},
		{
			Value: 0x200,
		},



		{
			Value: 0x2FF,
		},
		{
			Value: 0x300,
		},



		{
			Value: 0x3FF,
		},
		{
			Value: 0x400,
		},



		{
			Value: 0x4FF,
		},
		{
			Value: 0x500,
		},



		{
			Value: 0xFFF,
		},
		{
			Value: 0x1000,
		},



		{
			Value: 0x1FFF,
		},
		{
			Value: 0x2000,
		},



		{
			Value: 0xFFFF,
		},
		{
			Value: 0x10000,
		},



		{
			Value: 0xFFFFF,
		},
		{
			Value: 0x100000,
		},



		{
			Value: 0xFFFFFF,
		},
		{
			Value: 0x1000000,
		},



		{
			Value: 0xFFFFFFF,
		},
		{
			Value: 0x10000000,
		},



		{
			Value: 0xFFFFFFF0,
		},
		{
			Value: 0xFFFFFFF1,
		},
		{
			Value: 0xFFFFFFF2,
		},
		{
			Value: 0xFFFFFFF3,
		},
		{
			Value: 0xFFFFFFF4,
		},
		{
			Value: 0xFFFFFFF5,
		},
		{
			Value: 0xFFFFFFF6,
		},
		{
			Value: 0xFFFFFFF7,
		},
		{
			Value: 0xFFFFFFF8,
		},
		{
			Value: 0xFFFFFFF9,
		},
		{
			Value: 0xFFFFFFFA,
		},
		{
			Value: 0xFFFFFFFB,
		},
		{
			Value: 0xFFFFFFFC,
		},
		{
			Value: 0xFFFFFFFD,
		},
		{
			Value: 0xFFFFFFFE,
		},
		{
			Value: 0xFFFFFFFF,
		},









		{
			Value: 0xFEDCBA987654321F,
		},



		{
			Value: 0xFFFFFFFFFFFFFFEF,
		},
		{
			Value: 0xFFFFFFFFFFFFFFF0,
		},
		{
			Value: 0xFFFFFFFFFFFFFFF1,
		},
		{
			Value: 0xFFFFFFFFFFFFFFF2,
		},
		{
			Value: 0xFFFFFFFFFFFFFFF3,
		},
		{
			Value: 0xFFFFFFFFFFFFFFF4,
		},
		{
			Value: 0xFFFFFFFFFFFFFFF5,
		},
		{
			Value: 0xFFFFFFFFFFFFFFF6,
		},
		{
			Value: 0xFFFFFFFFFFFFFFF7,
		},
		{
			Value: 0xFFFFFFFFFFFFFFF8,
		},
		{
			Value: 0xFFFFFFFFFFFFFFF9,
		},
		{
			Value: 0xFFFFFFFFFFFFFFFA,
		},
		{
			Value: 0xFFFFFFFFFFFFFFFB,
		},
		{
			Value: 0xFFFFFFFFFFFFFFFC,
		},
		{
			Value: 0xFFFFFFFFFFFFFFFD,
		},
		{
			Value: 0xFFFFFFFFFFFFFFFE,
		},
		{
			Value: 0xFFFFFFFFFFFFFFFF,
		},
	}

	for testNumber, test := range tests {
		var val Uint128 = Uint128FromUint(test.Value)

		{
			expected := test.Value
			actual   := val.array[0]

			if expected != actual {
				t.Errorf("For test #%d, the actual first element of the (internal) array is not what was expected." , testNumber)
				t.Logf("EXPECTED: %d", expected)
				t.Logf("ACTUAL:   %d", actual)
				continue
			}
		}

		{
			for i:=1; i < len(val.array); i++ {
				const expected = 0
				actual := val.array[i]

				if expected != actual {
					t.Errorf("For test #%d, and array index %d, the actual value is not whas what expected.", testNumber, i)
					t.Logf("EXPECTED: [%d] %d", i, expected)
					t.Logf("ACTUAL:   [%d] %d", i, actual)
				}
			}
		}
	}
}
