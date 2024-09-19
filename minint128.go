package bits128

func MinInt128() Int128 {
	var x Int128 = MaxInt128()
	x.Inc()
	return x
}
