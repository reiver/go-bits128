package bits128

func MaxUint128() Uint128 {

	// If uint is 64-bit, then this will end up being: 0xFFFFFFFFFFFFFFFF.
	// If uint is 32-bit, then this will end up being: 0xFFFFFFFF.
	//
	// I.e., in 64-bit: 0x0000000000000000 - 1 = 0xFFFFFFFFFFFFFFFF
	// I.e., in 32-bit: 0x00000000         - 1 = 0xFFFFFFFF
	var ffffffffffffffff uint
	{
		ffffffffffffffff--
	}

	var uint128 Uint128
	{
		for index:=0; index<SizeUints; index++ {
			uint128.array[index] = ffffffffffffffff
		}
	}

	return uint128
}
