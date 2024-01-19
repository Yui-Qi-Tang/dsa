package b75

import "testing"

func TestLongestCommonSubsequence(t *testing.T) {

	testFuncs := []func(string, string) int{
		longestCommonSubsequencev39,
		longestCommonSubsequencev38,
		longestCommonSubsequencev37,
		longestCommonSubsequencev36,
		longestCommonSubsequencev35,
		longestCommonSubsequencev34,
		longestCommonSubsequencev33,
		longestCommonSubsequencev32,
		longestCommonSubsequencev31,
		longestCommonSubsequencev30,
		longestCommonSubsequencev29,
		longestCommonSubsequencev28,
		longestCommonSubsequencev27,
		longestCommonSubsequencev26,
		longestCommonSubsequencev25,
		longestCommonSubsequencev24,
		longestCommonSubsequencev23,
		longestCommonSubsequencev22,
		longestCommonSubsequencev21,
		longestCommonSubsequencev20,
		longestCommonSubsequencev19,
		longestCommonSubsequencev18,
		longestCommonSubsequencev17,
		longestCommonSubsequencev16,
		longestCommonSubsequencev15,
		longestCommonSubsequencev14,
		longestCommonSubsequencev13,
		longestCommonSubsequencev12,
		longestCommonSubsequencev11,
		longestCommonSubsequencev10,
		longestCommonSubsequencev9,
		longestCommonSubsequencev8,
		longestCommonSubsequencev7,
		longestCommonSubsequencev6,
		longestCommonSubsequencev5,
		longestCommonSubsequencev4,
		longestCommonSubsequencev3,
		longestCommonSubsequencev2,
		longestCommonSubsequencev1,
		LongestCommonSubsequence,
		longestCommonSubsequence,
	}

	testcases := []struct {
		text1, text2 string
		want         int
	}{
		{
			text1: "abc",
			text2: "abc",
			want:  3,
		},
		{
			text1: "abc",
			text2: "def",
			want:  0,
		},
		{
			text1: "a",
			text2: "ab",
			want:  1,
		},
		{
			text1: "ab",
			text2: "a",
			want:  1,
		},
		{
			text1: "abc",
			text2: "defa",
			want:  1,
		},
	}

	for i, f := range testFuncs {
		t.Logf("test function %d...", i)
		for j, tt := range testcases {
			ans := f(tt.text1, tt.text2)
			if ans != tt.want {
				t.Fatalf("=> case[%d]: it should be %d, but got %d", j, tt.want, ans)
			}
		}
		t.Logf("... test function %d is passed", i)

	}
}
