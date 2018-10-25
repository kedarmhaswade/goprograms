// Slices are not "comparable", then how can they be used as keys of a map, or members of a set?
// 	m := make(map[[]int]string) => key type must not be a function, map, or slice
// Consider the possibility of a slice being a key of a map.
// The map keys must be immutable, otherwise a whole lot can go wrong and since slices are mutable
// structures, they shouldn't be allowed to be used as keys of maps. Can we "work around" this
// limitation?
// If a slice is unmodifiable, then it is possible that this can be simulated. But how can we enforce immutability?
package main

func main() {
	//m := make(map[[]int]string) // not allowed!
}
