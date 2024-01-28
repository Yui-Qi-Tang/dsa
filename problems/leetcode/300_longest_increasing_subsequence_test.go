package b75

import "testing"

func TestLengthOfLIS(t *testing.T) {

	testfuncs := []func([]int) int{
		LengthOfLISv40,
		LengthOfLISv39,
		LengthOfLISv38,
		LengthOfLISv37,
		LengthOfLISv36,
		LengthOfLISv35,
		LengthOfLISv34,
		LengthOfLISv33,
		LengthOfLISv32,
		LengthOfLISv31,
		LengthOfLISv30,
		LengthOfLISv29,
		LengthOfLISv28,
		LengthOfLISv27,
		LengthOfLISv26,
		LengthOfLISv25,
		LengthOfLISv24,
		LengthOfLISv23,
		LengthOfLISv22,
		LengthOfLISv21,
		LengthOfLISv20,
		LengthOfLISv19,
		LengthOfLISv18,
		LengthOfLISv17,
		LengthOfLISv16,
		LengthOfLISv15,
		LengthOfLISv14,
		LengthOfLISv13,
		LengthOfLISv12,
		LengthOfLISv11,
		LengthOfLISv10,
		LengthOfLISv9,
		LengthOfLISv8,
		LengthOfLISv7,
		LengthOfLISv6,
		LengthOfLISv5,
		LengthOfLISv4,
		LengthOfLISv3,
		LengthOfLISv2,
		LengthOfLIS,
	}

	testcases := []struct {
		in   []int
		want int
	}{
		{
			in:   []int{10, 9, 2, 5, 3, 7, 101, 18},
			want: 4,
		},
		{
			in:   []int{0, 1, 0, 3, 2, 3},
			want: 4,
		},
		{
			in:   []int{7, 7, 7, 7, 7, 7, 7},
			want: 1,
		},
		{
			in:   []int{0},
			want: 1,
		},
		{in: []int{3, 10, 1}, want: 2},
		{in: []int{4, 10, 4, 3, 8, 9}, want: 3},
	}

	for i, f := range testfuncs {
		t.Logf("start testing function %d...", i)
		for j, tt := range testcases {
			ans := f(tt.in)
			if ans != tt.want {
				t.Fatalf("case[%d] => it should be %d, but got %d", j, tt.want, ans)
			}
		}
		t.Logf("... function[%d] is passed", i)
	}

}
