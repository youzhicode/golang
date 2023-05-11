package tempconv

func CToF(c Celsius) Fahrenheit {
	return Fahrenheit(c * 9 / 5 + 32)
}

func FToC(f Fahrenheit) Celsius {
	return Celsius((f - 32) * 5 / 9)
}

func FToM(foot LengthForFoot) LengthForMeter {
	return LengthForMeter(foot * 0.3048)
}

func MToF(meter LengthForMeter) LengthForFoot {
	return LengthForFoot(meter / 0.3048)
}

func PToK(pound WeightForPound) WeightForKG {
	return WeightForKG(pound * 0.4535)
}

func KToP(kg WeightForKG) WeightForPound {
	return WeightForPound(kg / 0.4535)
}
