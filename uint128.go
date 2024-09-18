package bits128

import (
	"fmt"
//	"math/bits"
	"unsafe"
)


var unused uint

// SizeBytes equals the size of a uint128, measured in bytes.
// I.e., 16 (== 128 / 8).
const SizeBytes = (128 / 8) // == 16

// The length of a uint, measured in number-of-bytes.
const lenuint = int(unsafe.Sizeof(unused))

// SizeUints equals the size of a uint128, as measured in the number of 'uint'.
//
// So, for example, if a 'uint' is 32-bits in size, then SizeUints is 4 (== 128 / 32).
//
// And, for example, if a 'uint' is 64-bits in size, then SizeUints is 2 (== 128 / 64).
const SizeUints = SizeBytes / lenuint

// We expected the length of a uint128 to be divisble by the length of a uint.
// If it isn't, we panic().
func init() {
	if 0 != SizeBytes % lenuint {
		panic(fmt.Sprintf("bits128: expected the length (in bytes) of a 'uint128' (i.e., %d) to be divisible by the length (in bytes) of a 'uint' (i.e., %d), but it isn't", SizeBytes, lenuint))
	}
}

// Uint128 is an implementation of a uint128 type.
// I.e., an unsigned 128-bit integer.
type Uint128 struct {
	// We use 'uint' rather than 'uint64' (or something else) to make it so we are using the optimal word length for the CPU architecture.
	// So, 'uint' MIGHT be 'uint64' or MIGHT be 'uint32' or MIGHT be something else.
	array [SizeUints]uint
}

func Uint128FromUint(value uint) Uint128 {
	return Uint128{
		array: [SizeUints]uint{
			value,
		},
	}
}

func Uint128FromUintsLittleEndian(array [SizeUints]uint) Uint128 {
	return Uint128{
		array: array,
	}
}

func (src1 *Uint128) Cmp(src2 *Uint128) int {
	switch {
	case src1 == src2:
		return 0
	case src1.array == src2.array:
		return 0
	default:
		for i:=(SizeUints-1); i>=0; i-- {
			val1 := src1.array[i]
			val2 := src2.array[i]

			switch {
			case val1 < val2:
				return -1
			case val1 > val2:
				return 1
			}
		}
		return 0
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
