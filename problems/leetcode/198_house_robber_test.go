package b75

import "testing"

func TestRob(t *testing.T) {

	testFuncs := []func([]int) int{
		houseRobber47,
		houseRobber46,
		houseRobber45,
		houseRobber44,
		houseRobber43,
		houseRobber42,
		houseRobber41,
		houseRobber40,
		houseRobber39,
		houseRobber38,
		houseRobber37,
		houseRobber36,
		houseRobber35,
		houseRobber34,
		houseRobber33,
		houseRobber32,
		houseRobber31,
		houseRobber30,
		houseRobber29,
		houseRobber28,
		houseRobber27,
		houseRobber26,
		houseRobber25,
		houseRobber24,
		houseRobber23,
		houseRobber22,
		houseRobber21,
		houseRobber20,
		houseRobber19,
		houseRobber18,
		houseRobber17,
		houseRobber16,
		houseRobber15,
		houseRobberv14,
		houseRobberv13,
		houseRobberv12,
		houseRobberv11,
		houseRobberv10,
		houseRobberv9,
		houseRobberv8,
		houseRobberv7,
		houseRobberv6,
		houseRobberv5,
		RobBetterDPv2,

		// bad implementation bellow. Keep them and learn
		// houseRobberv4,
		// Rob,
		// RobBetterDP,
		// hourseRobber,
		// houseRobberv1,
		// houseRobberv2,
		// houseRobberv3,
	}

	testcases := []struct {
		in   []int
		want int
	}{
		{
			in:   []int{1, 2, 3, 1},
			want: 4,
		},
		{
			in:   []int{2, 7, 9, 3, 1},
			want: 12,
		},
		{
			in:   []int{2, -1, 9, -2, 1},
			want: 12,
		},
		{
			in:   []int{-2, 1, -9, 2, -1},
			want: 3,
		},
		{
			in:   []int{2, 1, 1, 2},
			want: 4,
		},
		{
			in:   []int{1},
			want: 1,
		},
	}

	for i, f := range testFuncs {
		t.Logf("test function %d...", i)
		for j, tt := range testcases {
			in := make([]int, len(tt.in))
			copy(in, tt.in)
			ans := f(in)
			if ans != tt.want {
				t.Fatalf("=> case[%d] it should be %d, but got %d", j, tt.want, ans)
			}
		}
		t.Logf("... function %d is passed", i)
	}

}
