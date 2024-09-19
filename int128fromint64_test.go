package bits128_test

import (
	"testing"

	"github.com/reiver/go-bits128"
)

func TestInt128FromInt64(t *testing.T) {

	tests := []struct{
		Int64 int64
		Expected string
	}{
		{
			Int64:                    -0x8000000000000000,
			Expected: "0xFFFFFFFFFFFFFFFF8000000000000000",
		},
		{
			Int64:                    -0x7FFFFFFFFFFFFFFF,
			Expected: "0xFFFFFFFFFFFFFFFF8000000000000001",
		},
		{
			Int64:                    -0x7FFFFFFFFFFFFFFE,
			Expected: "0xFFFFFFFFFFFFFFFF8000000000000002",
		},
		{
			Int64:                    -0x7FFFFFFFFFFFFFFD,
			Expected: "0xFFFFFFFFFFFFFFFF8000000000000003",
		},
		{
			Int64:                    -0x7FFFFFFFFFFFFFFC,
			Expected: "0xFFFFFFFFFFFFFFFF8000000000000004",
		},
		{
			Int64:                    -0x7FFFFFFFFFFFFFFB,
			Expected: "0xFFFFFFFFFFFFFFFF8000000000000005",
		},
		{
			Int64:                    -0x7FFFFFFFFFFFFFFA,
			Expected: "0xFFFFFFFFFFFFFFFF8000000000000006",
		},
		{
			Int64:                    -0x7FFFFFFFFFFFFFF9,
			Expected: "0xFFFFFFFFFFFFFFFF8000000000000007",
		},
		{
			Int64:                    -0x7FFFFFFFFFFFFFF8,
			Expected: "0xFFFFFFFFFFFFFFFF8000000000000008",
		},
		{
			Int64:                    -0x7FFFFFFFFFFFFFF7,
			Expected: "0xFFFFFFFFFFFFFFFF8000000000000009",
		},
		{
			Int64:                    -0x7FFFFFFFFFFFFFF6,
			Expected: "0xFFFFFFFFFFFFFFFF800000000000000A",
		},
		{
			Int64:                    -0x7FFFFFFFFFFFFFF5,
			Expected: "0xFFFFFFFFFFFFFFFF800000000000000B",
		},
		{
			Int64:                    -0x7FFFFFFFFFFFFFF4,
			Expected: "0xFFFFFFFFFFFFFFFF800000000000000C",
		},
		{
			Int64:                    -0x7FFFFFFFFFFFFFF3,
			Expected: "0xFFFFFFFFFFFFFFFF800000000000000D",
		},
		{
			Int64:                    -0x7FFFFFFFFFFFFFF2,
			Expected: "0xFFFFFFFFFFFFFFFF800000000000000E",
		},
		{
			Int64:                    -0x7FFFFFFFFFFFFFF1,
			Expected: "0xFFFFFFFFFFFFFFFF800000000000000F",
		},
		{
			Int64:                    -0x7FFFFFFFFFFFFFF0,
			Expected: "0xFFFFFFFFFFFFFFFF8000000000000010",
		},
		{
			Int64:                    -0x7FFFFFFFFFFFFFEF,
			Expected: "0xFFFFFFFFFFFFFFFF8000000000000011",
		},
		{
			Int64:                    -0x7FFFFFFFFFFFFFEE,
			Expected: "0xFFFFFFFFFFFFFFFF8000000000000012",
		},
		{
			Int64:                    -0x7FFFFFFFFFFFFFED,
			Expected: "0xFFFFFFFFFFFFFFFF8000000000000013",
		},
		{
			Int64:                    -0x7FFFFFFFFFFFFFEC,
			Expected: "0xFFFFFFFFFFFFFFFF8000000000000014",
		},
		{
			Int64:                    -0x7FFFFFFFFFFFFFEB,
			Expected: "0xFFFFFFFFFFFFFFFF8000000000000015",
		},
		{
			Int64:                    -0x7FFFFFFFFFFFFFEA,
			Expected: "0xFFFFFFFFFFFFFFFF8000000000000016",
		},
		{
			Int64:                    -0x7FFFFFFFFFFFFFE9,
			Expected: "0xFFFFFFFFFFFFFFFF8000000000000017",
		},
		{
			Int64:                    -0x7FFFFFFFFFFFFFE8,
			Expected: "0xFFFFFFFFFFFFFFFF8000000000000018",
		},
		{
			Int64:                    -0x7FFFFFFFFFFFFFE7,
			Expected: "0xFFFFFFFFFFFFFFFF8000000000000019",
		},
		{
			Int64:                    -0x7FFFFFFFFFFFFFE6,
			Expected: "0xFFFFFFFFFFFFFFFF800000000000001A",
		},
		{
			Int64:                    -0x7FFFFFFFFFFFFFE5,
			Expected: "0xFFFFFFFFFFFFFFFF800000000000001B",
		},
		{
			Int64:                    -0x7FFFFFFFFFFFFFE4,
			Expected: "0xFFFFFFFFFFFFFFFF800000000000001C",
		},
		{
			Int64:                    -0x7FFFFFFFFFFFFFE3,
			Expected: "0xFFFFFFFFFFFFFFFF800000000000001D",
		},
		{
			Int64:                    -0x7FFFFFFFFFFFFFE2,
			Expected: "0xFFFFFFFFFFFFFFFF800000000000001E",
		},
		{
			Int64:                    -0x7FFFFFFFFFFFFFE1,
			Expected: "0xFFFFFFFFFFFFFFFF800000000000001F",
		},
		{
			Int64:                    -0x7FFFFFFFFFFFFFE0,
			Expected: "0xFFFFFFFFFFFFFFFF8000000000000020",
		},









		{
			Int64:                                    -60,
			Expected: "0xFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFC4",
		},
		{
			Int64:                                    -59,
			Expected: "0xFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFC5",
		},
		{
			Int64:                                    -58,
			Expected: "0xFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFC6",
		},
		{
			Int64:                                    -57,
			Expected: "0xFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFC7",
		},
		{
			Int64:                                    -56,
			Expected: "0xFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFC8",
		},
		{
			Int64:                                    -55,
			Expected: "0xFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFC9",
		},
		{
			Int64:                                    -54,
			Expected: "0xFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFCA",
		},
		{
			Int64:                                    -53,
			Expected: "0xFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFCB",
		},
		{
			Int64:                                    -52,
			Expected: "0xFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFCC",
		},
		{
			Int64:                                    -51,
			Expected: "0xFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFCD",
		},
		{
			Int64:                                    -50,
			Expected: "0xFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFCE",
		},
		{
			Int64:                                    -49,
			Expected: "0xFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFCF",
		},
		{
			Int64:                                    -48,
			Expected: "0xFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFD0",
		},
		{
			Int64:                                    -47,
			Expected: "0xFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFD1",
		},
		{
			Int64:                                    -46,
			Expected: "0xFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFD2",
		},
		{
			Int64:                                    -45,
			Expected: "0xFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFD3",
		},
		{
			Int64:                                    -44,
			Expected: "0xFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFD4",
		},
		{
			Int64:                                    -43,
			Expected: "0xFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFD5",
		},
		{
			Int64:                                    -42,
			Expected: "0xFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFD6",
		},
		{
			Int64:                                    -41,
			Expected: "0xFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFD7",
		},
		{
			Int64:                                    -40,
			Expected: "0xFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFD8",
		},
		{
			Int64:                                    -39,
			Expected: "0xFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFD9",
		},
		{
			Int64:                                    -38,
			Expected: "0xFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFDA",
		},
		{
			Int64:                                    -37,
			Expected: "0xFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFDB",
		},
		{
			Int64:                                    -36,
			Expected: "0xFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFDC",
		},
		{
			Int64:                                    -35,
			Expected: "0xFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFDD",
		},
		{
			Int64:                                    -34,
			Expected: "0xFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFDE",
		},
		{
			Int64:                                    -33,
			Expected: "0xFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFDF",
		},
		{
			Int64:                                    -32,
			Expected: "0xFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFE0",
		},
		{
			Int64:                                    -31,
			Expected: "0xFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFE1",
		},
		{
			Int64:                                    -30,
			Expected: "0xFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFE2",
		},
		{
			Int64:                                    -29,
			Expected: "0xFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFE3",
		},
		{
			Int64:                                    -28,
			Expected: "0xFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFE4",
		},
		{
			Int64:                                    -27,
			Expected: "0xFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFE5",
		},
		{
			Int64:                                    -26,
			Expected: "0xFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFE6",
		},
		{
			Int64:                                    -25,
			Expected: "0xFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFE7",
		},
		{
			Int64:                                    -24,
			Expected: "0xFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFE8",
		},
		{
			Int64:                                    -23,
			Expected: "0xFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFE9",
		},
		{
			Int64:                                    -22,
			Expected: "0xFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFEA",
		},
		{
			Int64:                                    -21,
			Expected: "0xFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFEB",
		},
		{
			Int64:                                    -20,
			Expected: "0xFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFEC",
		},
		{
			Int64:                                    -19,
			Expected: "0xFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFED",
		},
		{
			Int64:                                    -18,
			Expected: "0xFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFEE",
		},
		{
			Int64:                                    -17,
			Expected: "0xFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFEF",
		},
		{
			Int64:                                    -16,
			Expected: "0xFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFF0",
		},
		{
			Int64:                                    -15,
			Expected: "0xFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFF1",
		},
		{
			Int64:                                    -14,
			Expected: "0xFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFF2",
		},
		{
			Int64:                                    -13,
			Expected: "0xFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFF3",
		},
		{
			Int64:                                    -12,
			Expected: "0xFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFF4",
		},
		{
			Int64:                                    -11,
			Expected: "0xFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFF5",
		},
		{
			Int64:                                    -10,
			Expected: "0xFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFF6",
		},
		{
			Int64:                                     -9,
			Expected: "0xFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFF7",
		},
		{
			Int64:                                     -8,
			Expected: "0xFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFF8",
		},
		{
			Int64:                                     -7,
			Expected: "0xFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFF9",
		},
		{
			Int64:                                     -6,
			Expected: "0xFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFA",
		},
		{
			Int64:                                     -5,
			Expected: "0xFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFB",
		},
		{
			Int64:                                     -4,
			Expected: "0xFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFC",
		},
		{
			Int64:                                     -3,
			Expected: "0xFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFD",
		},
		{
			Int64:                                     -2,
			Expected: "0xFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFE",
		},
		{
			Int64:                                     -1,
			Expected: "0xFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFF",
		},
		{
			Int64:                                      0,
			Expected: "0x00000000000000000000000000000000",
		},
		{
			Int64:                                      1,
			Expected: "0x00000000000000000000000000000001",
		},
		{
			Int64:                                      2,
			Expected: "0x00000000000000000000000000000002",
		},
		{
			Int64:                                      3,
			Expected: "0x00000000000000000000000000000003",
		},
		{
			Int64:                                      4,
			Expected: "0x00000000000000000000000000000004",
		},
		{
			Int64:                                      5,
			Expected: "0x00000000000000000000000000000005",
		},
		{
			Int64:                                      6,
			Expected: "0x00000000000000000000000000000006",
		},
		{
			Int64:                                      7,
			Expected: "0x00000000000000000000000000000007",
		},
		{
			Int64:                                      8,
			Expected: "0x00000000000000000000000000000008",
		},
		{
			Int64:                                      9,
			Expected: "0x00000000000000000000000000000009",
		},
		{
			Int64:                                     10,
			Expected: "0x0000000000000000000000000000000A",
		},
		{
			Int64:                                     11,
			Expected: "0x0000000000000000000000000000000B",
		},
		{
			Int64:                                     12,
			Expected: "0x0000000000000000000000000000000C",
		},
		{
			Int64:                                     13,
			Expected: "0x0000000000000000000000000000000D",
		},
		{
			Int64:                                     14,
			Expected: "0x0000000000000000000000000000000E",
		},
		{
			Int64:                                     15,
			Expected: "0x0000000000000000000000000000000F",
		},
		{
			Int64:                                     16,
			Expected: "0x00000000000000000000000000000010",
		},
		{
			Int64:                                     17,
			Expected: "0x00000000000000000000000000000011",
		},
		{
			Int64:                                     18,
			Expected: "0x00000000000000000000000000000012",
		},
		{
			Int64:                                     19,
			Expected: "0x00000000000000000000000000000013",
		},
		{
			Int64:                                     20,
			Expected: "0x00000000000000000000000000000014",
		},
		{
			Int64:                                     21,
			Expected: "0x00000000000000000000000000000015",
		},
		{
			Int64:                                     22,
			Expected: "0x00000000000000000000000000000016",
		},
		{
			Int64:                                     23,
			Expected: "0x00000000000000000000000000000017",
		},
		{
			Int64:                                     24,
			Expected: "0x00000000000000000000000000000018",
		},
		{
			Int64:                                     25,
			Expected: "0x00000000000000000000000000000019",
		},
		{
			Int64:                                     26,
			Expected: "0x0000000000000000000000000000001A",
		},
		{
			Int64:                                     27,
			Expected: "0x0000000000000000000000000000001B",
		},
		{
			Int64:                                     28,
			Expected: "0x0000000000000000000000000000001C",
		},
		{
			Int64:                                     29,
			Expected: "0x0000000000000000000000000000001D",
		},
		{
			Int64:                                     30,
			Expected: "0x0000000000000000000000000000001E",
		},
		{
			Int64:                                     31,
			Expected: "0x0000000000000000000000000000001F",
		},
		{
			Int64:                                     32,
			Expected: "0x00000000000000000000000000000020",
		},



		{
			Int64:                    9223372036854775792,
			Expected: "0x00000000000000007FFFFFFFFFFFFFF0",
		},
		{
			Int64:                    9223372036854775793,
			Expected: "0x00000000000000007FFFFFFFFFFFFFF1",
		},
		{
			Int64:                    9223372036854775794,
			Expected: "0x00000000000000007FFFFFFFFFFFFFF2",
		},
		{
			Int64:                    9223372036854775795,
			Expected: "0x00000000000000007FFFFFFFFFFFFFF3",
		},
		{
			Int64:                    9223372036854775796,
			Expected: "0x00000000000000007FFFFFFFFFFFFFF4",
		},
		{
			Int64:                    9223372036854775797,
			Expected: "0x00000000000000007FFFFFFFFFFFFFF5",
		},
		{
			Int64:                    9223372036854775798,
			Expected: "0x00000000000000007FFFFFFFFFFFFFF6",
		},
		{
			Int64:                    9223372036854775799,
			Expected: "0x00000000000000007FFFFFFFFFFFFFF7",
		},
		{
			Int64:                    9223372036854775800,
			Expected: "0x00000000000000007FFFFFFFFFFFFFF8",
		},
		{
			Int64:                    9223372036854775801,
			Expected: "0x00000000000000007FFFFFFFFFFFFFF9",
		},
		{
			Int64:                    9223372036854775802,
			Expected: "0x00000000000000007FFFFFFFFFFFFFFA",
		},
		{
			Int64:                    9223372036854775803,
			Expected: "0x00000000000000007FFFFFFFFFFFFFFB",
		},
		{
			Int64:                    9223372036854775804,
			Expected: "0x00000000000000007FFFFFFFFFFFFFFC",
		},
		{
			Int64:                    9223372036854775805,
			Expected: "0x00000000000000007FFFFFFFFFFFFFFD",
		},
		{
			Int64:                    9223372036854775806,
			Expected: "0x00000000000000007FFFFFFFFFFFFFFE",
		},
		{
			Int64:                    9223372036854775807,
			Expected: "0x00000000000000007FFFFFFFFFFFFFFF",
		},









		{
			Int64:                     0x7FFFFFFFFFFFFFDF,
			Expected: "0x00000000000000007FFFFFFFFFFFFFDF",
		},
		{
			Int64:                     0x7FFFFFFFFFFFFFE0,
			Expected: "0x00000000000000007FFFFFFFFFFFFFE0",
		},
		{
			Int64:                     0x7FFFFFFFFFFFFFE1,
			Expected: "0x00000000000000007FFFFFFFFFFFFFE1",
		},
		{
			Int64:                     0x7FFFFFFFFFFFFFE2,
			Expected: "0x00000000000000007FFFFFFFFFFFFFE2",
		},
		{
			Int64:                     0x7FFFFFFFFFFFFFE3,
			Expected: "0x00000000000000007FFFFFFFFFFFFFE3",
		},
		{
			Int64:                     0x7FFFFFFFFFFFFFE4,
			Expected: "0x00000000000000007FFFFFFFFFFFFFE4",
		},
		{
			Int64:                     0x7FFFFFFFFFFFFFE5,
			Expected: "0x00000000000000007FFFFFFFFFFFFFE5",
		},
		{
			Int64:                     0x7FFFFFFFFFFFFFE6,
			Expected: "0x00000000000000007FFFFFFFFFFFFFE6",
		},
		{
			Int64:                     0x7FFFFFFFFFFFFFE7,
			Expected: "0x00000000000000007FFFFFFFFFFFFFE7",
		},
		{
			Int64:                     0x7FFFFFFFFFFFFFE8,
			Expected: "0x00000000000000007FFFFFFFFFFFFFE8",
		},
		{
			Int64:                     0x7FFFFFFFFFFFFFE9,
			Expected: "0x00000000000000007FFFFFFFFFFFFFE9",
		},
		{
			Int64:                     0x7FFFFFFFFFFFFFEA,
			Expected: "0x00000000000000007FFFFFFFFFFFFFEA",
		},
		{
			Int64:                     0x7FFFFFFFFFFFFFEB,
			Expected: "0x00000000000000007FFFFFFFFFFFFFEB",
		},
		{
			Int64:                     0x7FFFFFFFFFFFFFEC,
			Expected: "0x00000000000000007FFFFFFFFFFFFFEC",
		},
		{
			Int64:                     0x7FFFFFFFFFFFFFED,
			Expected: "0x00000000000000007FFFFFFFFFFFFFED",
		},
		{
			Int64:                     0x7FFFFFFFFFFFFFEE,
			Expected: "0x00000000000000007FFFFFFFFFFFFFEE",
		},
		{
			Int64:                     0x7FFFFFFFFFFFFFEF,
			Expected: "0x00000000000000007FFFFFFFFFFFFFEF",
		},
		{
			Int64:                     0x7FFFFFFFFFFFFFF0,
			Expected: "0x00000000000000007FFFFFFFFFFFFFF0",
		},
		{
			Int64:                     0x7FFFFFFFFFFFFFF1,
			Expected: "0x00000000000000007FFFFFFFFFFFFFF1",
		},
		{
			Int64:                     0x7FFFFFFFFFFFFFF2,
			Expected: "0x00000000000000007FFFFFFFFFFFFFF2",
		},
		{
			Int64:                     0x7FFFFFFFFFFFFFF3,
			Expected: "0x00000000000000007FFFFFFFFFFFFFF3",
		},
		{
			Int64:                     0x7FFFFFFFFFFFFFF4,
			Expected: "0x00000000000000007FFFFFFFFFFFFFF4",
		},
		{
			Int64:                     0x7FFFFFFFFFFFFFF5,
			Expected: "0x00000000000000007FFFFFFFFFFFFFF5",
		},
		{
			Int64:                     0x7FFFFFFFFFFFFFF6,
			Expected: "0x00000000000000007FFFFFFFFFFFFFF6",
		},
		{
			Int64:                     0x7FFFFFFFFFFFFFF7,
			Expected: "0x00000000000000007FFFFFFFFFFFFFF7",
		},
		{
			Int64:                     0x7FFFFFFFFFFFFFF8,
			Expected: "0x00000000000000007FFFFFFFFFFFFFF8",
		},
		{
			Int64:                     0x7FFFFFFFFFFFFFF9,
			Expected: "0x00000000000000007FFFFFFFFFFFFFF9",
		},
		{
			Int64:                     0x7FFFFFFFFFFFFFFA,
			Expected: "0x00000000000000007FFFFFFFFFFFFFFA",
		},
		{
			Int64:                     0x7FFFFFFFFFFFFFFB,
			Expected: "0x00000000000000007FFFFFFFFFFFFFFB",
		},
		{
			Int64:                     0x7FFFFFFFFFFFFFFC,
			Expected: "0x00000000000000007FFFFFFFFFFFFFFC",
		},
		{
			Int64:                     0x7FFFFFFFFFFFFFFD,
			Expected: "0x00000000000000007FFFFFFFFFFFFFFD",
		},
		{
			Int64:                     0x7FFFFFFFFFFFFFFE,
			Expected: "0x00000000000000007FFFFFFFFFFFFFFE",
		},
		{
			Int64:                     0x7FFFFFFFFFFFFFFF,
			Expected: "0x00000000000000007FFFFFFFFFFFFFFF",
		},
	}

	for testNumber, test := range tests {

		var actual string = bits128.Int128FromInt64(test.Int64).HexString()

		var expected string = test.Expected

		if expected != actual {
			t.Errorf("For test #%d, the actual int128 is not what was expected.", testNumber)
			t.Logf("EXPECTED: %q", expected)
			t.Logf("ACTUAL:   %q", actual)
			continue
		}
	}
}
