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

func TestFindNElementsWithHeap(t *testing.T) {
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
		ans := sliceWithHeap(nil).Init(tt.n).Top(tt.in, tt.n)
		if !reflect.DeepEqual(ans, tt.want) {
			t.Fatalf("it should be %v, but got %v", tt.want, ans)
		}
	}
	t.Log("passed")
}

func TestFindNElementsWithMyHeap(t *testing.T) {
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
		ans := top(tt.in, tt.n)
		if !reflect.DeepEqual(ans, tt.want) {
			t.Fatalf("it should be %v, but got %v", tt.want, ans)
		}
	}
	t.Log("passed")
}

func BenchmarkFindNWithSort(b *testing.B) {
	n := 5

	for i := 0; i < b.N; i++ {
		in := generateSlice(1000000)
		_ = sliceWithSort(in).Top(n)
	}
}

func BenchmarkFindNWithMyHeap(b *testing.B) {
	n := 5

	for i := 0; i < b.N; i++ {
		in := generateSlice(1000000)
		_ = top(in, n)

	}
}

func BenchmarkFindNWithHeap(b *testing.B) {
	n := 5
	for i := 0; i < b.N; i++ {
		in := generateSlice(1000000)
		swh := make(sliceWithHeap, 0)
		_ = swh.Init(n).Top(in, n)
	}
}

func generateSlice(n int) []int {
	return rand.New(rand.NewSource(time.Now().Unix())).Perm(n)
}
