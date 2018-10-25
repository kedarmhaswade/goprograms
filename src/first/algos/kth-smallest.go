// Given a sorted array of integers, find the kth smallest square of an integer.
// k starts from 0
// Example:
// array: {-9, -8, -4, -3, -1, 2, 5, 6, 10}, k = 2
// answer (-3 * -3) = 9
package algos

func KthSmallestSquareInSortedArray(a []int, k int) int {
	i := findSmallestAbsoluteIndex(a) // linear: finds the index of the element with smallest absolute value in a sorted slice
	//i := findSmallestAbsoluteIndexBS(a, 0, len(a)) // bs: finds the index of the element with smallest absolute value in a sorted slice
	// the smallest absolute value gives the smallest square
	// pi is the index of the smallest non negative number
	// ni is the index of the biggest negative number
	var ni, pi int
	if a[i] < 0 {
		ni = i
		pi = i + 1
	} else {
		pi = i
		ni = i - 1
	}
	var n, sq int
	// we only increment pi and we only decrement ni
	for ni >= 0 && pi < len(a) && n < k {
		if abs(a[ni]) <= abs(a[pi]) {
			sq = a[ni] * a[ni]
			ni -= 1
		} else {
			sq = a[pi] * a[pi]
			pi += 1
		}
		n += 1
	}
	for ni >= 0 && n < k {
		sq = a[ni] * a[ni]
		ni -= 1
		n += 1
	}
	for pi < len(a) && n < k {
		sq = a[pi] * a[pi]
		pi += 1
		n += 1
	}
	if n < k {
		return -1 // there are not enough elements!
	}
	return sq
}

func findSmallestAbsoluteIndex(a []int) int {
	var i int
	// for now, let me use linear search, which is, of course, suboptimal
	for i = 1; i < len(a); {
		if a[i] > a[i-1] && abs(a[i]) < abs(a[i-1]) {
			i += 1
		} else {
			break
		}
	}
	return i - 1
}

// findSmallestAbsoluteIndexBS uses binary search to find the index of the element withe smallest absolute value
func findSmallestAbsoluteIndexBS(a []int, lo, hi int) int {
	// assumption: lo and hi are valid indices, lo in inclusive, hi is exclusive
	for {
		mi := (lo + hi) >> 1
		if mi-1 >= 0 && mi+1 < hi && a[mi-1] <= 0 && a[mi+1] > 0 { // we are at the lowest point
			return mi
		} else if mi-1 >= 0 && mi+1 < hi && a[mi-1] <= 0 && a[mi+1] < 0 {
			lo = mi
		} else if mi-1 >= 0 && mi+1 < hi && a[mi-1] >= 0 && a[mi+1] > 0 {
			hi = mi + 1
		} else {
			return mi
		}
	}
}
func abs(n int) int {
	if n < 0 {
		return -n
	}
	return n
}
