package ch6

import "bytes"

// A bit vector uses a slice of unsigned integer values or “words,”
// each bit of which represents an element of the set.
// The set contains i if the i-th bit is set.
// The following program demonstrates a simple bit vector type with these methods.

// An IntSet is a set of small non-negative integers.
// Its zero value represents the empty set.
type IntSet struct {
	words []uint64 //uint64 is important because we need control over number and value of bits
}

// Has returns true if the given integer is in the set, false otherwise
func (this *IntSet) Has(m int) bool {
	if m < 0 {
		return false
	}
	return this.hasAt(this.coordinatesFor(m))
}

// Add adds the given integer m if it is non negative; returns false if m is negative.
// Expands the backing array as needed and returns true; sets the relevant bit in the word.
func (this *IntSet) Add(m int) bool {
	if m < 0 {
		return false
	}
	idx, off := this.coordinatesFor(m)
	n := len(this.words)
	for i := n; i <= idx; i++ {
		this.words = append(this.words, 0)
	}
	this.words[idx] |= 1 << off
	return true
}

// AddAll adds all the given integers to this set. The caller is guaranteed that the given integers
// are in the set after this call returns as long as they are non negative.
func (this *IntSet) AddAll(a ... int) *IntSet {
	for _, n := range a {
		this.Add(n) // return value is ignored
	}
	return this
}

// String returns a textual representation of the members of this set.
// The returned string is sorted. If the set contains 100 or more elements,
// then the first 99 elements, followed by an ellipsis, and the last element
// are present in the returned string, otherwise all the elements are present.
// The string starts with a '{' and ends with a '}'.
func (this *IntSet) String() string {
	return this.StringN(100)
}

// Generalization of the above method.
func (this *IntSet) StringN(n int) string {
	var buf bytes.Buffer
	buf.WriteString("\n")
	return buf.String()
}

// Returns the size (cardinality) of this set.
func (this *IntSet) Size() int {
	if this == nil {
		return 0
	}
	size := 0
	for _, w := range this.words {
		size += PopCount(w)
	}
	return size
}

// Remove removes x from the set, returns true if the element was removed, false otherwise
func (this *IntSet) Remove(x int) bool {
	idx, off := this.coordinatesFor(x)
	if idx >= len(this.words) {
		return false // not removed!
	}
	this.words[idx] &^= 1 << off // clear the bit for x => remove the element x
	return true                  // removed!
}

// Clear removes all elements from the set.
// If the argument release is true, the memory currently taken is released too, not otherwise
func (this *IntSet) Clear(release bool) {
	if release {
		this.words = nil // make gc work
	} else {
		for i := 0; i < len(this.words); i++ {
			this.words[i] = 0 // retain memory, but clear all bits => remove all elements
		}
	}
}

//Copy returns a copy of the set
func (this *IntSet) Copy() *IntSet {
	words2 := make([]uint64, len(this.words))
	copy(words2, this.words)
	return &IntSet{words2}
}

//Equals returns true if this and that if they are either the same set or contain the same elements
func (this *IntSet) Equals(that *IntSet) bool {
	if this == that {
		return true
	}
	thisLen := len(this.words)
	if thisLen != len(that.words) {
		return false
	}
	for i := 0; i < thisLen; i++ {
		if this.words[i] != that.words[i] {
			return false
		}
	}
	return true
}

//////// set operations

// UnionWith creates a new set that is a "union" of this set and that set; does not mutate this set
func (this *IntSet) UnionWith(that *IntSet) *IntSet {
	thisLen, thatLen, result := getResult(this.words, that.words)
	var i int
	for ; i < thisLen && i < thatLen; i++ {
		result.words[i] = this.words[i] | that.words[i]
	}
	if i == thisLen {
		result.words = append(result.words, that.words[i:]...)
	}
	result.words = append(result.words, this.words[i:]...)
	return result
}

// IntersectWith returns a new set that contains elements present in both this set and that set
func (this *IntSet) IntersectWith(that *IntSet) *IntSet {
	if this == nil || that == nil {
		return &IntSet{} // every null set is different, can't use a constant
	}
	thisLen := len(this.words)
	thatLen := len(that.words)
	result := &IntSet{} // start empty
	n := 0
	if thisLen < thatLen {
		result.words = make([]uint64, thisLen)
		n = thisLen
	} else {
		result.words = make([]uint64, thatLen)
		n = thatLen
	}
	for i := 0; i < n; i++ {
		result.words[i] = this.words[i] & that.words[i]
	}
	return result
}

// DifferenceWith returns a new set that represents the difference of this set and that set -- a set that
// contains every element in 'this' set, but not 'that' set.
func (this *IntSet) DifferenceWith(that *IntSet) *IntSet {
	if this == nil {
		return &IntSet{}
	}
	if that == nil {
		return this.Copy()
	}
	thisLen := len(this.words)
	thatLen := len(that.words)
	result := &IntSet{}
	var i int
	for ; i < thisLen && i < thatLen; i++ {
		result.words = append(result.words, this.words[i] &^ that.words[i])
	}
	if i == thatLen {
		result.words = append(result.words, this.words[i:]...)
	}
	return result
}

//PopCount counts the population, aka the set bits in the given number
func PopCount(u uint64) int {
	// a much more efficient function is possible
	pc := 0
	for i := 0; i < 64; i++ {
		if u>>uint(i)&1 == 1 {
			pc += 1
		}
	}
	return pc
}
///////////// PRIVATE ///////
// coordinatesFor gives the word and bit indexes of the given integer in this set
func (this *IntSet) coordinatesFor(m int) (int, uint64) {
	idx := m / 64         // index of the word which may contain this integer
	off := uint64(m % 64) // the bit (0 being lsb) representing this integer
	return idx, off
}

func (this *IntSet) hasAt(idx int, off uint64) bool {
	n := len(this.words)
	if idx >= n {
		return false
	}
	return (this.words[idx]>>off)&1 == 1

}

// getResult returns an empty IntSet with a big-enough capacity that can hold elements
// in either IntSet's internal representation (slice).
func getResult(thisWords []uint64, thatWords []uint64) (int, int, *IntSet) {
	thisLen := len(thisWords)
	thatLen := len(thatWords)
	var result = &IntSet{}
	if thisLen > thatLen {
		result.words = make([]uint64, thisLen)
	} else {
		result.words = make([]uint64, thatLen)
	}
	return thisLen, thatLen, result
}
