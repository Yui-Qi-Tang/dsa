package mathgame

import (
	"math/rand"
	"reflect"
	"testing"
	"time"
)

func TestFindNElementsWithSort(t *testing.T) {
	testcases := []struct {
		n        int
		in, want []int
	}{
		{
			n:    5,
			in:   []int{1, 2, 3},
			want: []int{3, 2, 1},
		},
		{
			n:    5,
			in:   []int{1, 2, 3, 4, 5, 6, 7},
			want: []int{7, 6, 5, 4, 3},
		},
	}

	for _, tt := range testcases {
		ans := sliceWithSort(tt.in).Top(tt.n)
		if !reflect.DeepEqual(ans, tt.want) {
			t.Fatalf("it should be %v, but got %v", tt.want, ans)
		}
	}
	t.Log("passed")
}

func TestFindNElementsWithHeapV2(t *testing.T) {
	testcases := []struct {
		n        int
		in, want []int
	}{
		{
			n:    5,
			in:   []int{1, 2, 3},
			want: []int{3, 2, 1},
		},
		{
			n:    5,
			in:   []int{1, 2, 3, 4, 5, 6, 7},
			want: []int{7, 6, 5, 4, 3},
		},
	}

	for _, tt := range testcases {
		// sliceHeap := make(sliceWithHeap, 0, len(tt.in))
		// heap.Init(&sliceHeap)
		ans := sliceWithHeap(nil).Init(tt.n).Top(tt.in, tt.n)
		if !reflect.DeepEqual(ans, tt.want) {
			t.Fatalf("it should be %v, but got %v", tt.want, ans)
		}
	}
	t.Log("passed")
}
func BenchmarkFindNWithSort(b *testing.B) {
	n := 5

	for i := 0; i < b.N; i++ {
		in := generateSlice(100000)
		_ = sliceWithSort(in).Top(n)
	}
}

func BenchmarkFindNWithHeap(b *testing.B) {
	n := 5
	for i := 0; i < b.N; i++ {
		in := generateSlice(100000)
		_ = sliceWithHeap(nil).Init(n).Top(in, n)
	}
}

func generateSlice(n int) []int {
	return rand.New(rand.NewSource(time.Now().Unix())).Perm(n)
}
