package b75

import (
	"reflect"
	"testing"
)

func TestTwoSum(t *testing.T) {

	testFunc := []func(nums []int, target int) []int{
		TwoSumv50,
		TwoSumv49,
		TwoSumv48,
		TwoSumv47,
		TwoSumv46,
		TwoSumv45,
		TwoSumv44,
		TwoSumv43,
		TwoSumv42,
		TwoSumv41,
		TwoSumv40,
		TwoSumv39,
		TwoSumv38,
		TwoSumv37,
		TwoSumv36,
		TwoSumv35,
		TwoSumv34,
		TwoSumv33,
		TwoSumv32,
		TwoSumv31,
		TwoSumv30,
		TwoSumv29,
		TwoSumv28,
		TwoSumv27,
		TwoSumv26,
		TwoSumv25,
		TwoSumv24,
		TwoSumv23,
		TwoSumv22,
		TwoSumv21,
		TwoSumv20,
		TwoSumv19,
		TwoSumv18,
		TwoSumv17,
		TwoSumv16,
		TwoSumv15,
		TwoSumv14,
		TwoSumv13,
		TwoSumv12,
		TwoSumv11,
		TwoSumv10,
		TwoSumv9,
		TwoSumv8,
		TwoSumv7,
		TwoSumv6,
		TwoSumv5,
		TwoSumv4,
		TwoSumv3,
		TwoSumv2,
		TwoSumv1,
		TwoSum,
	}

	testcases := []struct {
		in     []int
		target int
		want   []int
	}{

		{
			in:     []int{2, 7, 11, 15},
			target: 9,
			want:   []int{0, 1},
		},
		{
			in:     []int{3, 2, 4},
			target: 6,
			want:   []int{1, 2},
		},
		{
			in:     []int{3, 3},
			target: 6,
			want:   []int{0, 1},
		},
	}

	for i, f := range testFunc {
		t.Logf("test function %d...", i)
		for j, tt := range testcases {
			ans := f(tt.in, tt.target)
			if !reflect.DeepEqual(tt.want, ans) {
				t.Fatalf("=> case[%d]: it should be: %v, but got: %v", j, tt.want, ans)
			}
		}
		t.Logf("... test function %d is passed", i)
	}

}
