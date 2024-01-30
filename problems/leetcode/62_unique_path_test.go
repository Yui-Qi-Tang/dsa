package b75

import "testing"

func TestUniquePaths(t *testing.T) {

	testFuncs := []func(m, n int) int{
		uniquePathsv28,
		uniquePathsv27,
		uniquePathsv26,
		uniquePathsv25,
		uniquePathsv24,
		uniquePathsv23,
		uniquePathsv22,
		uniquePathsv21,
		uniquePathsv20,
		uniquePathsv19,
		uniquePathsv18,
		uniquePathsv17,
		uniquePathsv16,
		uniquePathsv15,
		uniquePathsv14,
		uniquePathsv13,
		uniquePathsv12,
		uniquePathsv11,
		uniquePathsv10,
		uniquePathsv9,
		uniquePathsv8,
		uniquePathsv7,
		uniquePathsv6,
		uniquePathsv5,
		uniquePathsv4,
		uniquePathsv3,
		uniquePathsv2,
		uniquePathsv1,
		UniquePaths,
	}

	testcases := []struct {
		m, n, want int
	}{
		{m: 3, n: 7, want: 28},
		{m: 3, n: 2, want: 3},
	}

	for i, f := range testFuncs {
		t.Logf("test function %d...", i)
		for j, tt := range testcases {
			ans := f(tt.m, tt.n)
			if ans != tt.want {
				t.Fatalf("=> case[%d]: it should be %d, but got %d", j, tt.want, ans)
			}
		}
		t.Logf("... function %d is passed", i)
	}

}

func TestUnipathFormula(t *testing.T) {
	for i := 1; i <= 4; i++ {
		for j := 1; j <= 9; j++ {
			t.Logf("=> unipath(%d, %d)=%d", i, j, uniquePathsv3(i, j))
		}
	}
}
