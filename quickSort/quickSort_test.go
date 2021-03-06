package quickSort

import (
	"fmt"
	"sort"
	"sort_algorithm/sequence"
	"testing"
)

func TestQuickSort(t *testing.T) {
	seq := quickSort(sequence.Seq)
	if isSorted := sort.IsSorted(seq); !isSorted {
		t.Errorf("quick sort failed")
	}
	fmt.Printf("%v\n", seq)
}
