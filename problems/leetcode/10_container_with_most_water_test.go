package b75

import "testing"

func TestContainerWithMostWater(t *testing.T) {

	testFuncs := []func([]int) int{
		ContainerWithMostWaterv24,
		ContainerWithMostWaterv23,
		ContainerWithMostWaterv22,
		ContainerWithMostWaterv21,
		ContainerWithMostWaterv20,
		ContainerWithMostWaterv19,
		ContainerWithMostWaterv18,
		ContainerWithMostWaterv17,
		ContainerWithMostWaterv16,
		ContainerWithMostWaterv15,
		ContainerWithMostWaterv14,
		ContainerWithMostWaterv13,
		ContainerWithMostWaterv12,
		ContainerWithMostWaterv11,
		ContainerWithMostWaterv10,
		ContainerWithMostWaterv9,
		ContainerWithMostWaterv8,
		ContainerWithMostWaterv7,
		ContainerWithMostWaterv6,
		ContainerWithMostWaterv5,
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
