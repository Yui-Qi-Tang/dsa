package b75

import "testing"

func TestBestTimeToBuAndSellStock(t *testing.T) {

	testfuncs := []func([]int) int{
		maxProfitV21,
		maxProfitV20,
		maxProfitV19,
		maxProfitV18,
		maxProfitV17,
		maxProfitV16,
		maxProfitV15,
		maxProfitV14,
		maxProfitV13,
		maxProfitV12,
		maxProfitV11,
		maxProfitV10,
		maxProfitV9,
		maxProfitV8,
		maxProfitV7,
		maxProfitV6,
		maxProfitV5,
		maxProfitV4,
		maxProfitV3,
		maxProfitV2,
		maxProfitV1,
	}

	testcases := []struct {
		in   []int
		want int
	}{
		{
			in:   []int{7, 1, 5, 3, 6, 4},
			want: 5,
		},
		{
			in:   []int{7, 6, 5, 4, 3, 2, 1},
			want: 0,
		},
		{
			in:   []int{7, 2, 5, 1, 6, 4},
			want: 5,
		},
		{
			in:   []int{1, 2, 3, 4, 5, 6, 7},
			want: 6,
		},
		{
			in:   []int{2, 4, 1},
			want: 2,
		},
	}

	for i, f := range testfuncs {
		t.Logf("test function %d", i)
		for j, tt := range testcases {
			ans := f(tt.in)
			if ans != tt.want {
				t.Fatalf("case[%d]:it should be %d, but got %d", j, tt.want, ans)
			}
		}
		t.Logf("function[%d] is passed", i)
	}

}
