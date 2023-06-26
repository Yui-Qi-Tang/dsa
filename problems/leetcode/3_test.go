package b75

import "testing"

func TestLengthOfLongestSubstring(t *testing.T) {

	testfuncs := []func(string) int{
		lengthOfLongestSubstringv25,
		lengthOfLongestSubstringv24,
		lengthOfLongestSubstringv23,
		lengthOfLongestSubstringv22,
		lengthOfLongestSubstringv21,
		lengthOfLongestSubstringv20,
		lengthOfLongestSubstringv19,
		lengthOfLongestSubstringv18,
		lengthOfLongestSubstringv17,
		lengthOfLongestSubstringv16,
		lengthOfLongestSubstringv15,
		lengthOfLongestSubstringv14,
		lengthOfLongestSubstringv13,
		lengthOfLongestSubstringv12,
		lengthOfLongestSubstringv11,
		lengthOfLongestSubstringv10,
		lengthOfLongestSubstringv9,
		lengthOfLongestSubstringv8,
		lengthOfLongestSubstringv7,
		lengthOfLongestSubstringv6,
		lengthOfLongestSubstringv5,
		lengthOfLongestSubstringv4,
		lengthOfLongestSubstringv3,
		lengthOfLongestSubstringv2,
		lengthOfLongestSubstringv1,
	}

	testcases := []struct {
		in   string
		want int
	}{
		{
			in:   "abcabcbb",
			want: 3,
		},
		{
			in:   "bbbbb",
			want: 1,
		},
		{
			in:   "pwwkew",
			want: 3,
		},
		{
			in:   "abcebcbb",
			want: 4,
		},
		{
			in:   "abce",
			want: 4,
		},
		{
			in:   "bcbb",
			want: 2,
		},
		{
			in:   "dvdf",
			want: 3,
		},
	}

	for i, f := range testfuncs {
		t.Logf("test function %d...", i)

		for j, tt := range testcases {
			ans := f(tt.in)
			if ans != tt.want {
				t.Fatalf("case[%d]: it should be: %d, but got %d", j, tt.want, ans)
			}
		}

		t.Logf("... function[%d] is passed", i)
	}

}
