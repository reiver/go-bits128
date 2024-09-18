package bits128

import (
	"fmt"
	"math/bits"
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

func Uint128FromUint64(value uint64) Uint128 {
	const lenuint64 = int(unsafe.Sizeof(value))

	switch {
	case lenuint64 <= lenuint:
		return Uint128FromUint(uint(value))
	default:
		const iterations = (lenuint64 / lenuint)

		var val Uint128

		for index:=0; index < iterations; index++ {

			var allUintFs uint
			allUintFs = allUintFs - 1

			var shift = (8*4*index)

			var mask uint64 = uint64(allUintFs) << shift

			val.array[index] = uint((mask & value) >> shift)
		}

		return val
	}
}

func Uint128FromUint32(value uint32) Uint128 {
	return Uint128FromUint64(uint64(value))
}

func Uint128FromUint16(value uint16) Uint128 {
	return Uint128FromUint64(uint64(value))
}

func Uint128FromUint8(value uint8) Uint128 {
	return Uint128FromUint64(uint64(value))
}

func Uint128FromUintsLittleEndian(array [SizeUints]uint) Uint128 {
	return Uint128{
		array: array,
	}
}

func (receiver *Uint128) And(src1 *Uint128, src2 *Uint128) {
	for index:=0; index<SizeUints; index++ {
		receiver.array[index] = src1.array[index] & src2.array[index]
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

// Dec decrements the uint128.
func (receiver *Uint128) Dec() {
	if nil == receiver {
		return
	}

	var borrow uint = 1

	for index:=0; index<SizeUints; index++ {
		receiver.array[index], borrow = bits.Sub(receiver.array[index], 0, borrow)
	}
}

// Inc increments the uint128.
func (receiver *Uint128) Inc() {
	if nil == receiver {
		return
	}

	var carry uint = 1

	for index:=0; index<SizeUints; index++ {
		receiver.array[index], carry = bits.Add(receiver.array[index], 0, carry)
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

func (receiver *Uint128) Xor(src1 *Uint128, src2 *Uint128) {
	for index:=0; index<SizeUints; index++ {
		receiver.array[index] = src1.array[index] ^ src2.array[index]
	}
}
