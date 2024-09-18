package bits128

import (
	"fmt"
//	"math/bits"
	"unsafe"
)


var unused uint

// The length of a uint128, measured in number-of-bytes.
const lenuint128 = (128 / 8) // == 16

// The length of a uint, measured in number-of-bytes.
const lenuint = int(unsafe.Sizeof(unused))

// The length of the array in the bits128.Uint128 struct, measured in number-of-uints.
// NOTE that this is measured in uints (rather than number-of-bytes).
const lenarray = lenuint128 / lenuint

// We expected the length of a uint128 to be divisble by the length of a uint.
// If it isn't, we panic().
func init() {
	if 0 != lenuint128 % lenuint {
		panic(fmt.Sprintf("bits128: expected the length (in bytes) of a 'uint128' (i.e., %d) to be divisible by the length (in bytes) of a 'uint' (i.e., %d), but it isn't", lenuint128, lenuint))
	}
}

// Uint128 is an implementation of a uint128 type.
// I.e., an unsigned 128-bit integer.
type Uint128 struct {
	// We use 'uint' rather than 'uint64' (or something else) to make it so we are using the optimal word length for the CPU architecture.
	// So, 'uint' MIGHT be 'uint64' or MIGHT be 'uint32' or MIGHT be something else.
	array [lenarray]uint
}

func Uint128FromUint(value uint) Uint128 {
	return Uint128{
		array: [lenarray]uint{
			value,
		},
	}
}

func (receiver Uint128) String() string {
	var buffer [34]byte
	var p []byte = buffer[0:0]

	p = append(p, "0x"...)

	{
		var format string = "%0"+fmt.Sprintf("%d", 2*lenuint)+"X"

		for index := (len(receiver.array)-1); index >= 0; index-- {
			p = append(p, fmt.Sprintf(format, receiver.array[index])...)
		}
	}

	return string(p)
}
