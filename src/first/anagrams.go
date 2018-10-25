// Write a function that reports whether two strings are anagrams of each other, that is,
// they contain the same letters in a different order.
package main

import "fmt"

func main() {
	s := "केदार"
	t := "रदाके"
	fmt.Printf("Are %s and %s anagrams? %v\n", s, t, areAnagrams(s, t))
	s = "abcdeee"
	t = "abecdee"
	fmt.Printf("Are %s and %s anagrams? %v\n", s, t, areAnagrams(s, t))
	s = "कालिदास"
	t = "सालिदाक"
	fmt.Printf("Are %s and %s anagrams? %v\n", s, t, areAnagrams(s, t))
	s = "apoorv"
	t = "vappor"
	fmt.Printf("Are %s and %s anagrams? %v\n", s, t, areAnagrams(s, t))
}

// Strings in Go are UTF-8 encoded, so we can use the standard technique of determining the frequency of
// each Rune in the given strings. This is efficiently done as a map. Equality of the maps determines if the strings
// are anagrams are each other. Another way is to sort both the strings and then equality of sorted strings
// establishes the anagram property; however sorting is more expensive.

func areAnagrams(s, t string) bool {
	if len(s) != len(t) {
		return false
	}
	if len(s) == 0 {
		return true
	}
	return equal(letterFrequency(s), letterFrequency(t))
}

func equal(m1 map[rune]int, m2 map[rune]int) bool {
	if len(m1) != len(m2) {
		return false
	}
	for k1, v1 := range m1 {
		v2, ok := m2[k1]
		if !ok || v1 != v2 {
			return false
		}
	}
	return true
}

func letterFrequency(s string) map[rune]int {
	m := make(map[rune]int, len(s)/3)
	for _, r := range s {
		f := m[r]
		m[r] = f + 1
	}
	return m
}
