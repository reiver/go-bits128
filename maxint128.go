package bits128

func MaxInt128() Int128 {
	var mask uint
	{
		mask--

		mask = mask >> 1
	}

	var uint128 Uint128 = MaxUint128()

	uint128.array[SizeUints-1] =uint128.array[SizeUints-1] & mask

	return Int128{
		uint128:uint128,
	}
}
