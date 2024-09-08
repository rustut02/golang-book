package weightconv

func KToP(k Kilos) Pounds {
	return Pounds(k / weightCoefficient)
}

func PToK(p Pounds) Kilos {
	return Kilos(p * weightCoefficient)
}
