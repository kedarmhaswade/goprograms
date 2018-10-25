package nick

// Write a recursive function that, given two strings, returns whether the first string is a
// subsequence of the second. For example, given "hac" and "cathartic", you should return true,
// but given "bat" and "table", you should return false.
// We treat strings as sequences of bytes, without any regards to encoding.
// Empty string is a subsequence of all strings including an empty string.
func IsSubsequence(subseq string, seq string) bool {
	if len(subseq) > len(seq) {
		return false
	}
	if len(subseq) == 0 {
		return true
	}
	if len(seq) == 0 {
		return false
	}
	f := indexOf(subseq[0], seq)
	if f < 0 {
		return false
	}
	return IsSubsequence(subseq[1:], seq[f+1:])
}

func IsSubsequenceIter(subseq string, seq string) bool {
	matchedIndex := 0 // index of the character from subsequence that matched
	sslen := len(subseq)
	slen := len(seq)
	for i, j := 0, 0; i < sslen && j < slen; { // i: subseq, j: seq
		if subseq[i] != seq[j] {
			j += 1
		} else {
			i += 1
			j += 1
			matchedIndex += 1
		}
	}
	return sslen == matchedIndex
}

///////////////////// Unexported Functions ////////////////

// indexOf returns the index of the given byte in a given string, returns -1 if the byte is not in string.
func indexOf(b byte, s string) int {
	n := len(s)
	for i := 0; i < n; i++ {
		if b == s[i] {
			return i
		}
	}
	return -1
}