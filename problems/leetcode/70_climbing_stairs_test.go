package b75

import (
	"fmt"
	"sort"
	"testing"
)

func TestClimbingStairs(t *testing.T) {

	testFuncs := []func(n int) int{
		ClimbingStairsv47,
		ClimbingStairsv46,
		ClimbingStairsv45,
		ClimbingStairsv44,
		ClimbingStairsv43,
		ClimbingStairsv42,
		ClimbingStairsv41,
		ClimbingStairsv40,
		ClimbingStairsv39,
		ClimbingStairsv38,
		ClimbingStairsv37,
		ClimbingStairsv36,
		ClimbingStairsv35,
		ClimbingStairsv34,
		ClimbingStairsv33,
		ClimbingStairsv32,
		ClimbingStairsv31,
		ClimbingStairsv30,
		ClimbingStairsv29,
		ClimbingStairsv28,
		ClimbingStairsv27,
		ClimbingStairsv26,
		ClimbingStairsv25,
		ClimbingStairsv24,
		ClimbingStairsv23,
		ClimbingStairsv22,
		ClimbingStairsv21,
		ClimbingStairsv20,
		ClimbingStairsv19,
		ClimbingStairsv18,
		ClimbingStairsv17,
		ClimbingStairsv16,
		ClimbingStairsv15,
		ClimbingStairsv14,
		ClimbingStairsv13,
		ClimbingStairsv12,
		ClimbingStairsv11,
		ClimbingStairsv10,
		ClimbingStairsv9,
		ClimbingStairsv8,
		ClimbingStairsv7,
		ClimbingStairsv6,
		ClimbingStairsv5,
		ClimbingStairsv4,
		ClimbingStairsv3,
		ClimbingStairsv2,
		ClimbingStairsv1,
		ClimbingStairs,
	}

	testcases := []struct {
		in   int
		want int
	}{
		{
			in:   1,
			want: 1,
		},
		{
			in:   2,
			want: 2,
		},
		{
			in:   3,
			want: 3,
		},
		{
			in:   5,
			want: 8,
		},
		{
			in:   10,
			want: 89,
		},
	}

	for i, f := range testFuncs {
		t.Logf("test function %d...", i)
		for j, tt := range testcases {
			ans := f(tt.in)
			if ans != tt.want {
				t.Fatalf("=> case[%d]it should be %d, but got %d", j, tt.want, ans)
			}
		}
		t.Logf("... function %d is passed", i)

	}

	a := []int{6, 3, 7, 8}
	k := 0
	sort.Ints(a)
	fmt.Println(a)
	for i := 0; i < len(a)-1; i++ {
		if a[i]-a[i+1] == 1 || a[i]-a[i+1] == -1 {
			k++
		} else {
			k = 0
		}
	}
	t.Log(k + 1)

}
