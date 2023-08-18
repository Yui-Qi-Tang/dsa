package b75

import "testing"

func TestLongestConsecutive(t *testing.T) {

	testFuncs := []func(nums []int) int{
		longestConsecutivev35,
		longestConsecutivev34,
		longestConsecutivev33,
		longestConsecutivev32,
		longestConsecutivev31,
		longestConsecutivev30,
		longestConsecutivev29,
		longestConsecutivev28,
		longestConsecutivev27,
		longestConsecutivev26,
		longestConsecutivev25,
		longestConsecutivev24,
		longestConsecutivev23,
		longestConsecutivev22,
		longestConsecutivev21,
		longestConsecutivev20,
		longestConsecutivev19,
		longestConsecutivev18,
		longestConsecutivev17,
		longestConsecutivev16,
		longestConsecutivev15,
		longestConsecutivev14,
		longestConsecutivev13,
		longestConsecutivev12,
		longestConsecutivev11,
		longestConsecutivev10,
		longestConsecutivev9,
		longestConsecutivev8,
		longestConsecutivev7,
		longestConsecutivev6,
		longestConsecutivev5,
		longestConsecutivev4,
		longestConsecutivev3,
		longestConsecutivev2,
		longestConsecutivev1,
	}

	testcases := []struct {
		in   []int
		want int
	}{
		{
			in:   []int{1, 2, 1, 2},
			want: 2,
		},
		{
			in:   []int{2, 2, 2, 2, 2, 2, 2, 2},
			want: 1,
		},
		{
			in:   []int{4, 4, 4, 4, 2, 2, 2, 2, 2, 3, 2, 1},
			want: 4,
		},
		{
			in:   []int{2, 3, 4},
			want: 3,
		},
		{
			in:   []int{100, 4},
			want: 1,
		},
		{
			in:   []int{100, 4, 200, 1, 3, 2},
			want: 4,
		},
		{
			in:   []int{100, 2, 4, 200, 1, 3, 2},
			want: 4,
		},
		{
			in:   []int{8, 7, 1, 2, 3, 4, 200, 1, 3, 2},
			want: 4,
		},
		{
			in:   []int{0, 3, 7, 2, 5, 8, 4, 6, 0, 1},
			want: 9,
		},
	}

	for i, f := range testFuncs {
		t.Logf("test function %d...", i)
		for j, tt := range testcases {
			ans := f(tt.in)
			if ans != tt.want {
				t.Fatalf("=> case[%d]: it should be %d, but got %d", j, tt.want, ans)
			}
		}
		t.Logf("... function %d is passed", i)
	}

}
