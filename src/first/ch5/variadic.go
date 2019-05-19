// Covers variadic functions from gopl.io
package ch5

import "bytes"

// Exercise 5.15: Write variadic functions max and min, analogous to sum. What should these functions do when called
// with no arguments? Write variants that require at least one argument.

// In general, it is rather hard to decide on a good API with variadic functions when the number of arguments
// provided is less than that required for the proper operation of the function, if the API is to be usable.
// It also varies according to the function itself. For example, a variadic function sum that sums all the int arguments
// passed to it can return 0 when no arguments are passed. But for a function like max which returns the maximum of
// given arguments, there is no appropriate value to return when no arguments are passed. Similarly for min.
// So, one should only provide max and min (variadic) functions with at least one required argument.
// Doing so sidesteps the issues associated with providing variadic functions for which _all_ the arguments are optional.

func max(first int, rest ...int) int {
	m := first
	for _, n := range rest {
		if n > m {
			m = n
		}
	}
	return m
}

func min(first int, rest ...int) int {
	m := first
	for _, n := range rest {
		if n < m {
			m = n
		}
	}
	return m
}

// Parameterizing the "behavior" from above functions

func linear(comparator func(x int, y int) bool, first int, rest ...int) int {
	result := first
	for _, n := range rest {
		if (comparator(n, result)) {
			result = n
		}
	}
	return result
}

func max1(first int, rest ...int) int {
	return linear(func(x int, y int) bool {
		if x > y {
			return true
		}
		return false
	}, first, rest...)
}

func min1(first int, rest ...int) int {
	return linear(func(x int, y int) bool {
		if x < y {
			return true
		}
		return false
	}, first, rest...)
}


// 5.16 Write a variadic version of strings.Join
// Here's a description from strings.Join:
// Join concatenates the elements of a to create a single string. The separator string
// sep is placed between elements in the resulting string.
// It would be ideal to test it out with strings.Join
func Join(sep string, strings ...string) string {
	n := len(strings)
	if n == 0 {
		return ""
	}
	var buf bytes.Buffer
	for i := 0; i < n-1; i++ {
		buf.WriteString(strings[i])
		buf.WriteString(sep)
	}
	buf.WriteString(strings[n-1])
	return buf.String()
}
