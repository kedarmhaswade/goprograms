// Demonstrates the grade-school multiplication algorithm by representing numbers as slices of digits
package main

import "fmt"

func main() {
	a := []int8{3, 2, 1, 5, 0, 6, 9, 7, 8}
	b := []int8{5, 2, 1, 4, 0, 6, 9, 7, 8, 0, 1}
	checkAdditionWith(a[:], b[:], 879605123, 10879604125)
	a = []int8{3, 2, 1}
	b = []int8{5, 9}
	checkAdditionWith(a, b, 123, 95)
	checkDigitMultiplicationWith(a, 9, 123*9)
	var big int64 = 12243675623
	checkMultiply(4233467, big)
}
func checkMultiply(first, second int64) {
	// only non-negative numbers for now
	f := toSlice(first)
	s := toSlice(second)
	r := multiply(f, s)
	p := first * second
	printBigEndian(r)
	fmt.Printf("%d\n", p)
	fmt.Printf("%v\n", isEqual(r, int64(p)))
}
func checkDigitMultiplicationWith(a []int8, n int8, p int64) {
	res := multiplyByDigit(a, n)
	fmt.Printf("%v\n", isEqual(res, p)) // note: we write in a big-endian manner;)
	printBigEndian(res)
	fmt.Printf("%d\n", p)
}
func checkAdditionWith(first []int8, second []int8, f, s int64) {
	n1 := int64(f)
	fmt.Printf("%v\n", isEqual(first, n1)) // note: we write in a big-endian manner;)
	n2 := int64(s)
	fmt.Printf("%v\n", isEqual(second, n2))
	fmt.Printf("%v\n", isEqual(addDecimal(first, second), n1+n2))
}

func addDecimal(first, second []int8) []int8 {
	n1 := len(first)
	n2 := len(second)
	if n1 > n2 {
		t := first
		first = second
		second = t
		x := n1
		n1 = n2
		n2 = x
	}
	// len(first) <= len(second)
	result := make([]int8, n2+1) // just in case there is a carry
	var c int8 = 0               // carry
	for i := 0; i < n1; i++ {
		var s2 = first[i] + second[i] + c
		result[i] = s2 % 10
		c = s2 / 10
	}
	for i := n1; i < n2; i++ {
		var s2 = second[i] + c
		result[i] = s2 % 10
		c = s2 / 10
	}
	if c > 0 {
		result[n2] = c
	} else {
		result = result[0:n2]
	}
	return result
}
func multiplyByDigit(num []int8, dig int8) []int8 {
	n := len(num)
	result := make([]int8, 0)
	var c int8
	for i := 0; i < n; i++ {
		p := num[i]*dig + c
		result = append(result, p%10)
		c = p / 10
	}
	if c != 0 {
		result = append(result, c)
	}
	return result
}

func multiply(first, second []int8) []int8 {
	var r []int8
	for i := 0; i < len(second); i++ {
		m := multiplyByDigit(first, second[i])
		a := nzeros(i)
		a = appendIt(a, m)
		r = addDecimal(r, a)
	}
	return r
}
func appendIt(z []int8, nz []int8) []int8 {
	for i := 0; i < len(nz); i++ {
		z = append(z, nz[i])
	}
	return z
}
func nzeros(i int) []int8 {
	var r []int8
	for j := 0; j < i; j++ {
		r = append(r, 0)
	}
	return r
}

func isEqual(digits []int8, value int64) bool {
	// this function is really of limited use, recursion is not that onerous
	if len(digits) == 1 {
		if int8(value) == digits[0] { // explicit conversion is safe
			return true
		}
		return false
	}
	return int8(value%10) == digits[0] && isEqual(digits[1:], value/10)
}
func printBigEndian(num []int8) {
	// not using recursion here :)
	for i := len(num) - 1; i >= 0; i-- {
		fmt.Printf("%d", num[i])
	}
	fmt.Println()
}
func toSlice(x int64) []int8 {
	if x < 0 {
		x = -x
	}
	if x < 10 {
		return []int8{int8(x)} // safe conversion
	}
	var r []int8
	for x >= 10 {
		r = append(r, int8(x%10))
		x /= 10
	}
	return append(r, int8(x))
}
