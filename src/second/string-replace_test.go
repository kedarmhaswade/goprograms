package second
import (
	"fmt"
	"testing"
)

func BenchmarkRemoveNonNumeric1(b *testing.B) {
	n1 := "1-(408) 318-4424"
	for i := 0; i < b.N; i++ {
		RemoveNonNumeric1(n1)
	}
}
func BenchmarkRemoveNonNumeric2(b *testing.B) {
	n1 := "1-(408) 318-4424"
	for i := 0; i < b.N; i++ {
		RemoveNonNumeric2(n1)
	}
}

func TestEqual(t *testing.T) {
	n1 := "1-(408) 318-4424"
	fmt.Printf("%v\n", RemoveNonNumeric1(n1))
	fmt.Printf("%v\n", RemoveNonNumeric2(n1))
}
