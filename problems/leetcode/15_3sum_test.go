package b75

import (
	"testing"
)

func TestThreeSum(t *testing.T) {

	// run one test function at once
	testfuncs := []func([]int) [][]int{
		ThreeSumv23,
		ThreeSumv22,
		//ThreeSumv21,
		//ThreeSumv20,
		//ThreeSumv19,
		//ThreeSumv18,
		//ThreeSumv17,
		//ThreeSumv16,
		//ThreeSumv15,
		//ThreeSumv14,
		//ThreeSumv13,
		//ThreeSumv12,
		//ThreeSumv11,
		//ThreeSumv10,
		//ThreeSumv9,
		//ThreeSumv8,
		//ThreeSumv7,
		//ThreeSumv6,
		//ThreeSumv5,
		//ThreeSumv4,
		//ThreeSumv3,
		//ThreeSumv2,
		//ThreeSumv1,
	}

	testcases := []struct {
		in   []int
		want [][]int
	}{
		{
			in: []int{-1, -1, 2, 5},
			want: [][]int{
				{2, -1, -1}, // {-1, -1, 2}, {-1, 0, 1},
			},
		},
		{
			in: []int{-1, 0, 1, 2, -1, -4},
			want: [][]int{
				{-1, 0, 1}, {2, -1, -1}, // {-1, -1, 2}, {-1, 0, 1},
			},
		},
		{
			in:   []int{0, 1, 1},
			want: [][]int{},
		},
		{
			in: []int{0, 0, 0},
			want: [][]int{
				{0, 0, 0},
			},
		},
	}

	equal := func(a, b [][]int) bool {

		m1, m2 := make(map[int]bool, len(a)), make(map[int]bool, len(b))

		for _, v := range a {
			for _, x := range v {
				m1[x] = true
			}
		}

		for _, v := range b {
			for _, x := range v {
				m2[x] = true
			}
		}

		for k := range m1 {
			if !m2[k] {
				return false
			}
		}

		return true
	}

	for j, f := range testfuncs {
		t.Logf("test function %d...", j)
		for i := 0; i < len(testcases); i++ {
			ans := f(testcases[i].in)
			if !equal(ans, testcases[i].want) {
				t.Fatalf("case[%d]: it should be %v, but got %v", i, testcases[i].want, ans)
			}
		}
		t.Logf("... function[%d] is passed", j)
	}

}
