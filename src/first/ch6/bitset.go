package ch6

import "bytes"

// A bit vector uses a slice of unsigned integer values or “words,”
// each bit of which represents a possible element of the set.
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

// UnionWith sets this to the union of this and that and returns this set
func (this *IntSet) UnionWith(that *IntSet) *IntSet {
	thisLen := len(this.words)
	thatLen := len(that.words)
	var i, j int
	for ; i < thisLen && j < thatLen; i, j = i+1, j+1 {
		this.words[i] |= that.words[j]
	}
	if i == thisLen {
		this.words = append(this.words, that.words[j:]...)
	}
	this.words = append(this.words, this.words[i:]...)
	return this
}

// AddAll adds all the given integers to this set. The caller is guaranteed that the given integers
// are in the set after this call returns as long as they are non negative.
func (this *IntSet) AddAll(a ... int) {
	for _, n := range a {
		this.Add(n) // return value is ignored
	}
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
	panic("nyi")
}

// Returns the size (cardinality) of this set.
func (this *IntSet) Len() int {
	size := 0
	if this == nil {
		return size
	}
	for _, w := range this.words {
		size += popCount(w)
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

//Equals returns true if this and
func (this *IntSet) Equals(that *IntSet) bool {
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

///////////// PRIVATE ///////
func popCount(u uint64) int {
	// a much more efficient function is possible
	pc := 0
	for i := 0; i < 64; i++ {
		if u>>uint(i)&1 == 1 {
			pc += 1
		}
	}
	return pc
}

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
