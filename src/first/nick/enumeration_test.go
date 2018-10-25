package nick

import (
	"fmt"
	"testing"
)

func TestSubsets(t *testing.T) {
	ss := GetSubsets([]int{1, 2, 3, 4})
	fmt.Printf("%+v\n", ss)
}

