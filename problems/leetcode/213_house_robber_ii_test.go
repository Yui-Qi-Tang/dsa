package b75

import "testing"

func TestRobII(t *testing.T) {

	testFuncs := []func([]int) int{
		houseRobberIIv50,
		houseRobberIIv49,
		houseRobberIIv48,
		houseRobberIIv47,
		houseRobberIIv46,
		houseRobberIIv45,
		houseRobberIIv44,
		houseRobberIIv43,
		houseRobberIIv42,
		houseRobberIIv41,
		houseRobberIIv40,
		houseRobberIIv39,
		houseRobberIIv38,
		houseRobberIIv37,
		houseRobberIIv36,
		houseRobberIIv35,
		houseRobberIIv34,
		houseRobberIIv33,
		houseRobberIIv32,
		houseRobberIIv31,
		houseRobberIIv30,
		houseRobberIIv29,
		houseRobberIIv28,
		houseRobberIIv27,
		houseRobberIIv26,
		houseRobberIIv25,
		houseRobberIIv24,
		houseRobberIIv23,
		houseRobberIIv22,
		houseRobberIIv21,
		houseRobberIIv20,
		houseRobberIIv19,
		houseRobberIIv18,
		houseRobberIIv17,
		houseRobberIIv16,
		houseRobberIIv15,
		houseRobberIIv14,
		houseRobberIIv13,
		houseRobberIIv12,
		houseRobberIIv11,
		houseRobberIIv10,
		houseRobberIIv9,
		houseRobberIIv8,
		houseRobberIIv7,
		houseRobberIIv6,
		houseRobberIIv5,
		// houseRobberIIv4,
		// RobII,
		RobIIBetter,
		// houseRobIIv2,
		// houseRobberIIv2,
		// houseRobberIIv3,
	}

	testcases := []struct {
		in   []int
		want int
	}{
		{
			in:   []int{2, 3, 2},
			want: 3,
		},
		{
			in:   []int{1, 2, 3, 1},
			want: 4,
		},
		{
			in:   []int{1, 2, 3},
			want: 3,
		},
		{
			in:   []int{3, 2, 1},
			want: 3,
		},
		{
			in:   []int{1},
			want: 1,
		},
		{
			in:   []int{2, 3},
			want: 3,
		},
		{
			in:   []int{200, 3, 140, 20, 10},
			want: 340,
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
		t.Logf("..function %d is passed", i)
	}

}
