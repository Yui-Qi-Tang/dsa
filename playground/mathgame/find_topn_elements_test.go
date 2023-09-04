package mathgame

import (
	"math/rand"
	"reflect"
	"strconv"
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

func TestFindNElementsWithHeapElements(t *testing.T) {
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

	build := func(in []int) []element {
		result := make([]element, 0, len(in))
		for _, v := range in {
			result = append(result, element{
				id:   v + 1,
				name: "user" + strconv.Itoa(v),
				cnt:  v,
			})
		}
		return result
	}

	for _, tt := range testcases {
		ans := topElements(build(tt.in), tt.n)
		if !reflect.DeepEqual(ans, build(tt.want)) {
			t.Fatalf("it should be %v, but got %v", tt.want, ans)
		}
	}
	t.Log("passed")
}

func TestFindNElementsWithSortElements(t *testing.T) {
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

	build := func(in []int) []element {
		result := make([]element, 0, len(in))
		for _, v := range in {
			result = append(result, element{
				id:   v + 1,
				name: "user" + strconv.Itoa(v),
				cnt:  v,
			})
		}
		return result
	}

	for _, tt := range testcases {
		var slice elements = build(tt.in)
		ans := slice.Top(tt.n)
		if !reflect.DeepEqual(ans, build(tt.want)) {
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

func BenchmarkFindTopNFrom1MIntSliceWithSort(b *testing.B) {
	n := 5

	for i := 0; i < b.N; i++ {
		in := generateSlice(1000000)
		_ = sliceWithSort(in).Top(n)
	}
}

func BenchmarkFindTopNFrom1MIntSliceWithMyHeap(b *testing.B) {
	n := 5

	for i := 0; i < b.N; i++ {
		in := generateSlice(1000000)
		_ = top(in, n)

	}
}

func BenchmarkFindTopNFrom1MIntSliceWithBuiltinHeap(b *testing.B) {
	n := 5
	for i := 0; i < b.N; i++ {
		in := generateSlice(1000000)
		swh := make(sliceWithHeap, 0)
		_ = swh.Init(n).Top(in, n)
	}
}

func BenchmarkFindTopNFrom1MStructSliceWithMyHeap(b *testing.B) {
	n := 5
	build := func(in []int) []element {
		result := make([]element, 0, len(in))
		for _, v := range in {
			result = append(result, element{
				id:   v + 1,
				name: "user" + strconv.Itoa(v),
				cnt:  v,
			})
		}
		return result
	}
	in := generateSlice(1000000)
	for i := 0; i < b.N; i++ {
		_ = topElements(build(in), n)
	}
}

func BenchmarkFindTopNFrom1MStructSliceWithSort(b *testing.B) {
	n := 5
	build := func(in []int) []element {
		result := make([]element, 0, len(in))
		for _, v := range in {
			result = append(result, element{
				id:   v + 1,
				name: "user" + strconv.Itoa(v),
				cnt:  v,
			})
		}
		return result
	}
	in := generateSlice(1000000)
	for i := 0; i < b.N; i++ {
		var slice elements = build(in)
		_ = slice.Top(n)
	}
}

func generateSlice(n int) []int {
	return rand.New(rand.NewSource(time.Now().Unix())).Perm(n)
}
