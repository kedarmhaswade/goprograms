package nick

// Given a list of n distinct elements, write a function that lists all subsets of those elements.

// There is an iterative algorithm, and, of course, a recursive one. The idea is that since we know that
// the number of subsets is exponential (2^n) in the size of the set, we can use the binary vector of length
// n to give a decent iterative algorithm.

//GetSubsets returns a set of all subsets of the given set.
// Each element of the returned slice is a slice representing a subset.
func GetSubsets(set []int) [][]int {
	n := len(set)
	if n == 0 {
		var ss [][]int
		ss = append(ss, []int{}) // the null set
		return ss
	}
	first := set[0]
	rem := GetSubsets(set[1:])
	var ss [][]int
	for _, subset := range rem {
		ss = append(ss, subset)
		ss = append(ss, append(subset, first))
	}
	return ss
}

// Given a non negative integer n, write a function that lists all strings formed from exactly n
// pairs of balanced parentheses. For example, given n = 3, you'd list these five strings:
// ((())) (()()) (())() ()(()) ()()()`