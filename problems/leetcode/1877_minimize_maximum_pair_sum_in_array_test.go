package b75

import (
	"fmt"
	"testing"
)

func TestMinPairSum(t *testing.T) {

	testfuncs := []func([]int) int{
		minPairSumv55,
		minPairSumv54,
		minPairSumv53,
		minPairSumv52,
		minPairSumv51,
		minPairSumv50,
		minPairSumv49,
		minPairSumv48,
		minPairSumv47,
		minPairSumv46,
		minPairSumv45,
		minPairSumv44,
		minPairSumv43,
		minPairSumv42,
		minPairSumv41,
		minPairSumv40,
		minPairSumv39,
		minPairSumv38,
		minPairSumv37,
		minPairSumv36,
		minPairSumv35,
		minPairSumv34,
		minPairSumv33,
		minPairSumv32,
		minPairSumv31,
		minPairSumv30,
		minPairSumv29,
		minPairSumv28,
		minPairSumv27,
		minPairSumv26,
		minPairSumv25,
		minPairSumv24,
		minPairSumv23,
		minPairSumv22,
		minPairSumv21,
		minPairSumv20,
		minPairSumv19,
		minPairSumv18,
		minPairSumv17,
		minPairSumv16,
		minPairSumv15,
		minPairSumv14,
		minPairSumv13,
		minPairSumv12,
		minPairSumv11,
		minPairSumv10,
		minPairSumv9,
		minPairSumv8,
		minPairSumv7,
		minPairSumv6,
		minPairSumv5,
		minPairSumv4,
		minPairSumv3,
		minPairSumv2,
		minPairSumv1,
	}

	testcases := []struct {
		in   []int
		want int
	}{
		{
			in:   []int{3, 5, 2, 3},
			want: 7,
		},
		{
			in:   []int{3, 5, 4, 2, 4, 6},
			want: 8,
		},
	}

	for i, f := range testfuncs {
		t.Run(fmt.Sprintf("test function %d", i), func(t *testing.T) {
			for j, tt := range testcases {
				ans := f(tt.in)
				if ans != tt.want {
					t.Fatalf("case[%d]: it should be %d, but got %d", j, tt.want, ans)
				}
			}
		})
	}

}
