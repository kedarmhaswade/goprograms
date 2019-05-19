package ch4

func x(b bool)  {
	if b {
		println(b)
	}
	x := 5
	panic(x)
}
