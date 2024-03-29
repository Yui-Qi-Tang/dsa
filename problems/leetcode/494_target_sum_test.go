package b75

import "testing"

func TestFindTargetSumWays(t *testing.T) {
	testFuncs := []func(nums []int, target int) int{
		findTargetSumWaysv50,
		findTargetSumWaysv49,
		findTargetSumWaysv48,
		findTargetSumWaysv47,
		findTargetSumWaysv46,
		findTargetSumWaysv45,
		findTargetSumWaysv44,
		findTargetSumWaysv43,
		findTargetSumWaysv42,
		findTargetSumWaysv41,
		findTargetSumWaysv40,
		findTargetSumWaysv39,
		findTargetSumWaysv38,
		findTargetSumWaysv37,
		findTargetSumWaysv36,
		findTargetSumWaysv35,
		findTargetSumWaysv34,
		findTargetSumWaysv33,
		findTargetSumWaysv32,
		findTargetSumWaysv31,
		findTargetSumWaysv30,
		findTargetSumWaysv29,
		findTargetSumWaysv28,
		findTargetSumWaysv27,
		findTargetSumWaysv26,
		findTargetSumWaysv25,
		findTargetSumWaysv24,
		findTargetSumWaysv23,
		findTargetSumWaysv22,
		findTargetSumWaysv21,
		findTargetSumWaysv20,
		findTargetSumWaysv19,
		findTargetSumWaysv18,
		findTargetSumWaysv17,
		findTargetSumWaysv16,
		findTargetSumWaysv15,
		findTargetSumWaysv14,
		findTargetSumWaysv13,
		findTargetSumWaysv12,
		findTargetSumWaysv11,
		findTargetSumWaysv10,
		findTargetSumWaysv9,
		findTargetSumWaysv8,
		findTargetSumWaysv7, // I don't why leetcode say my code is exeecded the time?
		findTargetSumWaysv6,
		findTargetSumWaysv5,
		findTargetSumWaysv4,
		findTargetSumWaysv3,
		findTargetSumWaysv2,
		findTargetSumWaysv1,
	}

	testcases := []struct {
		in     []int
		target int
		want   int
	}{
		{
			in:     []int{1, 1, 1, 1, 1},
			target: 3,
			want:   5,
		},
		{
			in:     []int{1},
			target: 1,
			want:   1,
		},

		{
			in:     []int{7, 46, 36, 49, 5, 34, 25, 39, 41, 38, 49, 47, 17, 11, 1, 41, 7, 16, 23, 13},
			target: 3,
			want:   5756,
		},
	}

	for i, f := range testFuncs {
		t.Logf("test function %d...", i)
		for j, tt := range testcases {
			ans := f(tt.in, tt.target)
			if ans != tt.want {
				t.Fatalf("=> case[%d]: it should be %d, but got %d", j, tt.want, ans)
			}
		}
		t.Logf("...function %d is passed", i)
	}

}
