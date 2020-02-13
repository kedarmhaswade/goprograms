package main

//import "golang.org/x/tour/pic"
//
////Implement Pic.
////It should return a slice of length dy, each element of which
////is a slice of dx 8-bit unsigned integers. When you run the program,
////it will display your picture, interpreting the integers as gray scale (well, blue scale) values.
////
////The choice of image is up to you. Interesting functions include (x+y)/2, x*y, and x^y.
////
////(You need to use a loop to allocate each []uint8 inside the [][]uint8.)
////
////(Use uint8(intValue) to convert between types.)
//func Pic(dx, dy int) [][]uint8 {
//	a := make([][]uint8, dy)
//	for i := range a {
//		a[i] = make([]uint8, dx)
//		for j := range a[i] {
//			a[i][j] = uint8((i + j) >> 1)
//		}
//	}
//	return a
//}
//
//func main() {
//	pic.Show(Pic)
//}
