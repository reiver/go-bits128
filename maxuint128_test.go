package bits128_test

import (
	"testing"

	"github.com/reiver/go-bits128"
)

func TestMaxuint128(t *testing.T) {

	var expected string = "0xFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFF"

	var actual string = bits128.MaxUint128().String()

	if expected != actual {
		t.Errorf("The actual max-uint128 is not what was expected.")
		t.Logf("EXPECTED: %q", expected)
		t.Logf("ACTUAL  : %q", actual)
		return
	}
}
