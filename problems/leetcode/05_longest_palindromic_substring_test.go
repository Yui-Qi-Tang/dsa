package b75

import "testing"

func TestLongestPalindrome(t *testing.T) {

	testFuncs := []func(s string) string{
		longestPalindromev40,
		longestPalindromev39,
		longestPalindromev38,
		longestPalindromev37,
		longestPalindromev36,
		longestPalindromev35,
		longestPalindromev34,
		longestPalindromev33,
		longestPalindromev32,
		longestPalindromev31,
		longestPalindromev30,
		longestPalindromev29,
		longestPalindromev28,
		longestPalindromev27,
		longestPalindromev26,
		longestPalindromev25,
		longestPalindromev24,
		longestPalindromev23,
		longestPalindromev22,
		longestPalindromev21,
		longestPalindromev20,
		longestPalindromev19,
		longestPalindromev18,
		longestPalindromev17,
		longestPalindromev16,
		longestPalindromev15,
		longestPalindromev14,
		longestPalindromev13,
		longestPalindromev12,
		longestPalindromev11,
		longestPalindromev10,
		longestPalindromev9,
		longestPalindromev8,
		longestPalindromev7,
		longestPalindromev6,
		longestPalindromev5,
		longestPalindromev4,
		longestPalindromev3,
		longestPalindromev2,
		longestPalindrome,
		longestPalindromev1,
	}

	testcases := []struct {
		in, want string
	}{
		{
			in:   "babad",
			want: "bab",
		},
		{
			in:   "cbbd",
			want: "bb",
		},
	}

	for i, f := range testFuncs {
		t.Logf("test function %d...", i)
		for j, tt := range testcases {
			ans := f(tt.in)
			if ans != tt.want {
				t.Fatalf("=> case[%d]: it should be %s, but got %s", j, tt.want, ans)
			}
		}
		t.Logf("... test function %d is passed", i)
	}
}
