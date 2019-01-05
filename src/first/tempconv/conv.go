package tempconv

func FToC(f Fahrenheit) Celsius {
	return 5 * Celsius(f-32) / 9
}

func CToF(c Celsius) Fahrenheit {
	return Fahrenheit(c*9)/5 + 32
}

func CToK(c Celsius) Kelvin {
	return Kelvin(c - AbsoluteZeroC) // i.e. c + 273.15
}

func KToC(k Kelvin) Celsius {
	return Celsius(k + Kelvin(AbsoluteZeroC)) // i.e. k - 273.15
}

func KToF(k Kelvin) Fahrenheit {
	return CToF(KToC(k))
}

func FToK(f Fahrenheit) Kelvin {
	return CToK(FToC(f))
}
