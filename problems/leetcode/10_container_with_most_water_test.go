package b75

import "testing"

func TestContainerWithMostWater(t *testing.T) {

	testFuncs := []func([]int) int{
		ContainerWithMostWaterv4,
		ContainerWithMostWaterv3,
		ContainerWithMostWaterv2,
		ContainerWithMostWaterv1,
		ContainerWithMostWater,
	}

	testcases := []struct {
		in   []int
		want int
	}{
		{
			in:   []int{1, 8, 6, 2, 5, 4, 8, 3, 7},
			want: 49,
		},
		{
			in:   []int{1, 1},
			want: 1,
		},
		{
			in:   []int{1, 10, 6, 2, 5, 4, 10, 3, 7},
			want: 50,
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
		t.Logf("... function[%d] is passed", i)
	}

}
