package b75

import "testing"

func TestBestTimeToBuAndSellStock(t *testing.T) {

	testfuncs := []func([]int) int{
		maxProfitV65,
		maxProfitV64,
		maxProfitV63,
		maxProfitV62,
		maxProfitV61,
		maxProfitV60,
		maxProfitV59,
		maxProfitV58,
		maxProfitV57,
		maxProfitV56,
		maxProfitV55,
		maxProfitV54,
		maxProfitV53,
		maxProfitV52,
		maxProfitV51,
		maxProfitV50,
		maxProfitV49,
		maxProfitV48,
		maxProfitV47,
		maxProfitV46,
		maxProfitV45,
		maxProfitV44,
		maxProfitV43,
		maxProfitV42,
		maxProfitV41,
		maxProfitV40,
		maxProfitV39,
		maxProfitV38,
		maxProfitV37,
		maxProfitV36,
		maxProfitV35,
		maxProfitV34,
		maxProfitV33,
		maxProfitV32,
		maxProfitV31,
		maxProfitV30,
		maxProfitV29,
		maxProfitV28,
		maxProfitV27,
		maxProfitV26,
		maxProfitV25,
		maxProfitV24,
		maxProfitV23,
		maxProfitV22,
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
