package b75

import "testing"

func TestMinCostClimbingStairs(t *testing.T) {

	testFuncs := []func([]int) int{
		minCostClimbingStairsv34,
		minCostClimbingStairsv33,
		minCostClimbingStairsv32,
		minCostClimbingStairsv31,
		minCostClimbingStairsv30,
		minCostClimbingStairsv29,
		minCostClimbingStairsv28,
		minCostClimbingStairsv27,
		minCostClimbingStairsv26,
		minCostClimbingStairsv25,
		minCostClimbingStairsv24,
		minCostClimbingStairsv23,
		minCostClimbingStairsv22,
		minCostClimbingStairsv21,
		minCostClimbingStairsv20,
		minCostClimbingStairsv19,
		minCostClimbingStairsv18,
		minCostClimbingStairsv17,
		minCostClimbingStairsv16,
		minCostClimbingStairsv15,
		minCostClimbingStairsv14,
		minCostClimbingStairsv13,
		minCostClimbingStairsv12,
		minCostClimbingStairsv11,
		minCostClimbingStairsv10,
		minCostClimbingStairsv9,
		minCostClimbingStairsv8,
		minCostClimbingStairsv7,
		minCostClimbingStairsv6,
		minCostClimbingStairsv5,
		minCostClimbingStairsv4,
		minCostClimbingStairsv3,
		MinCostClimbingStairsV2,
		MinCostClimbingStairs,
	}

	testcases := []struct {
		in   []int
		want int
	}{
		{
			in:   []int{10, 15, 20},
			want: 15,
		},
		{
			in:   []int{1, 100, 1, 1, 1, 100, 1, 1, 100, 1},
			want: 6,
		},
	}

	for i, f := range testFuncs {
		t.Logf("test function %d...", i)
		for j, tt := range testcases {
			in := make([]int, len(tt.in))
			copy(in, tt.in)
			ans := f(in)
			if ans != tt.want {
				t.Fatalf("=> case[%d]: it should be %d, but got %d", j, tt.want, ans)
			}
		}
		t.Logf("... test function %d is passed", i)
	}

}
