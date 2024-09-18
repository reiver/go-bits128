package bits128_test

import (
	"testing"

	"strings"

	"github.com/reiver/go-bits128"
)

func TestMaxUint128(t *testing.T) {

	var expected string = "0x" + strings.Repeat("F", 2*bits128.SizeBytes)

	var actual string = bits128.MaxUint128().HexString()

	if expected != actual {
		t.Errorf("The actual max-uint128 is not what was expected.")
		t.Logf("EXPECTED: %q", expected)
		t.Logf("ACTUAL  : %q", actual)
		return
	}
}
