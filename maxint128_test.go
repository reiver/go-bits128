package bits128_test

import (
	"testing"

	"strings"

	"github.com/reiver/go-bits128"
)

func TestMaxuint128(t *testing.T) {

	var expected string = "0x7" + strings.Repeat("F", 2*bits128.SizeBytes - 1)

	var actual string = bits128.MaxInt128().HexString()

	if expected != actual {
		t.Errorf("The actual max-int128 is not what was expected.")
		t.Logf("EXPECTED: %q", expected)
		t.Logf("ACTUAL  : %q", actual)
		return
	}
}
