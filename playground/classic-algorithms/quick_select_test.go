package classicalgorithms

import (
	"testing"
)

func TestQuickSelect(t *testing.T) {
	nums := []int{1, 5, 8, 3, 6, 2, 7, 4, 9, 10}
	k := 3
	ans := 8
	kthLargest := quickSelect(nums, k)

	if kthLargest != ans {
		t.Fatalf("it should be %d, but got %d", ans, kthLargest)
	}

}
