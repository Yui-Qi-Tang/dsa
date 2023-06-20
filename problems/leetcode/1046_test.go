package b75

import "testing"

/*

i, i*2+1. i
0, 1, 2
1, 3, 4
2, 5, 6
*/

func TestLastStoneWeight(t *testing.T) {

	testfuncs := []func([]int) int{
		lastStoneWeightv21,
		lastStoneWeightv20,
		lastStoneWeightv19,
		lastStoneWeightv18,
		lastStoneWeightv17,
		lastStoneWeightv16,
		lastStoneWeightv15,
		lastStoneWeightv14,
		lastStoneWeightv13,
		lastStoneWeightv12,
		lastStoneWeightv11,
		lastStoneWeightv10,
		lastStoneWeightv9,
		lastStoneWeightv8,
		lastStoneWeightv7,
		lastStoneWeightv6,
		lastStoneWeightv5,
		lastStoneWeightv4,
		lastStoneWeightv3,
		lastStoneWeightv2,
		lastStoneWeightv1,
		lastStoneWeight,
	}

	testcases := []struct {
		in   []int
		want int
	}{
		{in: []int{2, 7, 4, 1, 8, 1}, want: 1},
		{in: []int{1}, want: 1},
		{in: []int{2, 2}, want: 0},
		{in: []int{4, 3, 4, 3, 2}, want: 2},
	}

	for i, f := range testfuncs {
		t.Logf("test function %d", i)
		for j, tt := range testcases {
			ans := f(tt.in)
			if ans != tt.want {
				t.Fatalf("case[%d]: it should be %d, but got %d", j, tt.want, ans)
			}
		}
		t.Logf("function[%d] is passed", i)
	}

}
