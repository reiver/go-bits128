package bits128

type Int128 struct {
	uint128 Uint128
}

func Int128FromUint(value uint) Int128 {
	return Int128{
		uint128: Uint128FromUint(value),
	}
}

func Int128FromInt64(value int64) Int128 {

	if 0 <= value {
		return Int128FromUint64(uint64(value))
	}

	{
		const minint64 int64 = -0x8000000000000000

		if minint64 == value {
			var int128 Int128 = Int128FromInt64(value + 1)

			int128.Dec()
			return int128
		}
	}

	{
		var uint128 Uint128 = Uint128FromUint64(uint64(-1 * value))

		var mask Uint128 = MaxUint128()
		uint128.Xor(&uint128, &mask)

		uint128.Inc()

		return Int128{
			uint128: uint128,
		}
	}
}

func Int128FromInt32(value int32) Int128 {
	return Int128FromInt64(int64(value))
}

func Int128FromInt16(value int16) Int128 {
	return Int128FromInt64(int64(value))
}

func Int128FromInt8(value int8) Int128 {
	return Int128FromInt64(int64(value))
}

func Int128FromUint64(value uint64) Int128 {
	return Int128{
		uint128: Uint128FromUint64(value),
	}
}

func Int128FromUint32(value uint32) Int128 {
	return Int128{
		uint128: Uint128FromUint32(value),
	}
}

func Int128FromUint16(value uint16) Int128 {
	return Int128{
		uint128: Uint128FromUint16(value),
	}
}

func Int128FromUint8(value uint8) Int128 {
	return Int128{
		uint128: Uint128FromUint8(value),
	}
}

func (src1 *Int128) Cmp(src2 *Int128) int {

	if src1.uint128 == src2.uint128 {
		return 0
	}

	{
		var src1Zero bool = src1.IsZero()
		var src2Zero bool = src2.IsZero()

		var src1Negative bool = src1.IsNegative()
		var src2Negative bool = src2.IsNegative()

		var src1Positive bool = !src1Zero && !src1Negative
		var src2Positive bool = !src2Zero && !src2Negative

		switch {
		case (src1Positive || src1Zero) && src2Negative:
			return 1
		case src1Negative && (src2Positive || src2Zero):
			return -1
		default:
			return src1.uint128.Cmp(&src2.uint128)
		}
	}
}

func (receiver *Int128) Dec() {
	receiver.uint128.Dec()
}

func (receiver *Int128) Inc() {
	receiver.uint128.Inc()
}

func (receiver *Int128) IsNegative() bool {
	var zero uint = 0
	var xffffffffffffffff uint = (zero - 1)
	var x7fffffffffffffff uint = (xffffffffffffffff >> 1)
	var x8000000000000000 uint = x7fffffffffffffff ^ xffffffffffffffff

	var last uint = receiver.uint128.array[SizeUints-1]

	return x8000000000000000 == (last & x8000000000000000)
}

func (receiver *Int128) IsPositive() bool {
	return !receiver.IsNegative() && !receiver.IsZero()
}

func (receiver *Int128) IsZero() bool {
	return receiver.uint128.IsZero()
}

func (receiver Int128) HexString() string {
	return receiver.uint128.HexString()
}
