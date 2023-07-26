package b75

import "testing"

func TestCoinChange(t *testing.T) {

	// test functions
	testFuncs := []func(coins []int, amount int) int{
		coinChangev24,
		coinChangev23,
		coinChangev22,
		coinChangev21,
		coinChangev20,
		coinChangev19,
		coinChangev18,
		coinChangev17,
		coinChangev16,
		coinChangev15,
		coinChangev14,
		coinChangev13,
		coinChangev12,
		coinChangev11,
		cofinChangev10,
		coinChangev9,
		coinChangev8,
		coinChangev7,
		coinChangev6,
		coinChangev5,
		coinChangev4,
		coinChangev3,
		coinChangev2,
		CoinChange,
		coinChange,
	}
	// testcases
	testcases := []struct {
		coins  []int
		amount int
		want   int
	}{
		{
			coins:  []int{1, 2, 5},
			amount: 11,
			want:   3,
		},
		{
			coins:  []int{2},
			amount: 3,
			want:   -1,
		},
		{
			coins:  []int{1},
			amount: 0,
			want:   0,
		},
		{
			coins:  []int{20, 50},
			amount: 60,
			want:   3,
		},
	}

	for i, f := range testFuncs {
		t.Logf("test function %d...", i)

		for j, tt := range testcases {
			ans := f(tt.coins, tt.amount)
			if ans != tt.want {
				t.Fatalf("error case[%d] on function[%d]: it should be %d, but got %d", j, i, tt.want, ans)
			}
		}
		t.Logf("... function %d is passed", i)
	}

	t.Log("done")
}
