package b75

import "testing"

func TestLongestCommonSubsequence(t *testing.T) {

	testFuncs := []func(string, string) int{
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
