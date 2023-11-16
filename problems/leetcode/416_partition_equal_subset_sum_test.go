package b75

import "testing"

func TestCanPartition(t *testing.T) {

	testfuncs := []func([]int) bool{
		canPartitionv12,
		canPartitionv11,
		canPartitionv10,
		canPartitionv9,
		canPartitionv8,
		canPartitionv7,
		canPartitionv6,
		canPartitionv5,
		canPartitionv4,
		canPartitionv3,
		canPartitionv2,
		canPartitionv1,
	}

	testcases := []struct {
		in   []int
		want bool
	}{
		{
			in:   []int{1, 5, 11, 5},
			want: true,
		},
		{
			in:   []int{1, 2, 3, 5},
			want: false,
		},
		{
			in:   []int{3, 3, 3, 4, 5},
			want: true,
		},
		{
			in:   []int{14, 9, 8, 4, 3, 2},
			want: true,
		},
		{
			in:   []int{1, 2, 5},
			want: false,
		},
	}

	for i, f := range testfuncs {
		t.Logf("test function %d...", i)
		for j, tt := range testcases {
			ans := f(tt.in)
			if ans != tt.want {
				t.Fatalf("case[%d]: it should be %v, but got %v", j, tt.want, ans)
			}
		}
		t.Logf("...function[%d] is passed", i)
	}
}
