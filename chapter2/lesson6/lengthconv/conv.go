package lengthconv

func MToF(m Meter) Foot {
	return Foot(m * lengthCoefficient)
}

func FToM(f Foot) Meter {
	return Meter(f / lengthCoefficient)
}
