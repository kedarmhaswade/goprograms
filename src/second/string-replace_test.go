package second

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func BenchmarkRemoveNonNumericReplacer(b *testing.B) {
	n1 := "1-(408) 318-4424"
	for i := 0; i < b.N; i++ {
		RemoveNonNumeric1(n1)
	}
}
func BenchmarkRemoveNonNumericSimple(b *testing.B) {
	n1 := "1-(408) 318-4424"
	for i := 0; i < b.N; i++ {
		RemoveNonNumeric2(n1)
	}
}

func TestEqual(t *testing.T) {
	for _, n1 := range []string{"+०९१ १२३४५-५४३२१", "+1(408) 318-4649"} {
		useReplacer := RemoveNonNumeric1(n1)
		excludeNonNumeric := RemoveNonNumeric2(n1)
		if ! assert.Equal(t, true, useReplacer == excludeNonNumeric) {
			t.Errorf("%v and %v are not equal, but they should be!", useReplacer, excludeNonNumeric)
		}
	}
}
