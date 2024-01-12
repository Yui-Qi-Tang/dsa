package b75

import "testing"

func TestCoinChangeII(t *testing.T) {

	// 1-9 too slow to use, except 2
	testFuncs := []func(amount int, coins []int) int{
		// 1-D dp, for now, I can't understand what's going on...
		changev24,
		changev23,
		changev21,
		changev12,
		changev2, // the fastest, this can be passed by leetcode

		// dfs way
		changev25,
		changev22,
		changev20,
		changev19,
		changev18,
		changev17,
		changev16,
		changev15,
		changev14,
		changev13,
		changev11,
		changev10,
		changev9,
		changev8,
		changev7,
		changev6,
		changev5,
		changev4,
		changev3,
		changev1,
	}

	testcases := []struct {
		amount int
		coins  []int
		want   int
	}{
		// edge case, the can't be processed by dfs way
		{
			amount: 500,
			coins:  []int{3, 5, 7, 8, 9, 10, 11},
			want:   35502874,
		},
		{
			amount: 5,
			coins:  []int{1, 2, 5},
			want:   4,
		},
		{
			amount: 3,
			coins:  []int{2},
			want:   0,
		},
		{
			amount: 10,
			coins:  []int{10},
			want:   1,
		},
	}

	for i, f := range testFuncs {
		t.Logf("test function %d...", i)
		for j, tt := range testcases {
			if j == 0 && i != 0 {
				// skip case[0] if function isn't 0
				continue
			}
			ans := f(tt.amount, tt.coins)
			if ans != tt.want {
				t.Fatalf("=> case[%d]: it should be %d, but got %d", j, tt.want, ans)
			}
		}
		t.Logf("... function %d is passed ...", i)
	}

}
