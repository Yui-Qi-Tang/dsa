package b75

import "testing"

func TestCoinChangeIIWithDPTopDown(t *testing.T) {
	t.Log("start testing the dp top down...")
	testFuncs := []func(amount int, coins []int) int{
		changev71,
		changev69,
		changev67,
		changev66,
		changev65,
		changev64,
		changev61,
		changev60,
		changev57,
		changev55,
		changev54,
		changev52,
		changev49,
		changev47,
		changev45,
		changev43,
		changev41,
		changev40,
		changev37,
		changev35,
		changev33,
		changev31,
		changev28,
		changev26,
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
		// edge case, the can't be processed by top down way
		// {
		// 	amount: 500,
		// 	coins:  []int{3, 5, 7, 8, 9, 10, 11},
		// 	want:   35502874,
		// },
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
		{
			amount: 5,
			coins:  []int{1, 5},
			want:   2,
		},
	}

	for i, f := range testFuncs {
		for j, tt := range testcases {
			ans := f(tt.amount, tt.coins)
			if ans != tt.want {
				t.Fatalf("=> case[%d]: it should be %d, but got %d", j, tt.want, ans)
			}
		}
		t.Logf("... function %d is passed ...", i)
	}
	t.Log("... Passed")
}

func TestCoinChangeIIWithBottomUp(t *testing.T) {
	t.Log("start testing the dp bottom up...")

	testFuncs := []func(amount int, coins []int) int{
		changev72,
		changev70,
		changev68,
		changev63,
		changev62,
		changev59,
		changev58,
		changev56,
		changev53,
		changev51,
		changev50,
		changev48,
		changev46,
		changev44,
		changev42,
		changev39,
		changev38,
		changev36,
		changev34,
		changev32,
		changev30,
		changev29,
		changev27,
		changev24,
		changev23,
		changev21,
		changev12,
		changev2,
	}

	testcases := []struct {
		amount int
		coins  []int
		want   int
	}{
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
		{
			amount: 5,
			coins:  []int{1, 5},
			want:   2,
		},
	}

	for i, f := range testFuncs {
		for j, tt := range testcases {
			ans := f(tt.amount, tt.coins)
			if ans != tt.want {
				t.Fatalf("=> case[%d]: it should be %d, but got %d", j, tt.want, ans)
			}
		}
		t.Logf("... function %d is passed ...", i)
	}
	t.Log("... Passed")
}
