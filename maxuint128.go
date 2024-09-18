package bits128

func MaxUint128() Uint128 {

	var ffffffffffffffffffffffffffffffff uint
	{
		ffffffffffffffffffffffffffffffff--
	}

	var uint128 Uint128

	for index:=0; index<SizeUints; index++ {
		uint128.array[index] = ffffffffffffffffffffffffffffffff
	}

	return uint128
}
