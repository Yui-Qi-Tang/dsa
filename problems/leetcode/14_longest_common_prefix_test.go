package b75

import "testing"

func TestLongestCommonPrefix(t *testing.T) {

	testfuncs := []func([]string) string{
		longestCommonPrefixv24,
		longestCommonPrefixv23,
		longestCommonPrefixv22,
		longestCommonPrefixv21,
		longestCommonPrefixv20,
		longestCommonPrefixv19,
		longestCommonPrefixv18,
		longestCommonPrefixv17,
		longestCommonPrefixv16,
		longestCommonPrefixv15,
		longestCommonPrefixv14,
		longestCommonPrefixv13,
		longestCommonPrefixv12,
		longestCommonPrefixv11,
		longestCommonPrefixv10,
		longestCommonPrefixv9,
		longestCommonPrefixv8,
		longestCommonPrefixv7,
		longestCommonPrefixv6,
		longestCommonPrefixv5,
		longestCommonPrefixv4,
		longestCommonPrefixv3,
		longestCommonPrefixv2,
		longestCommonPrefixv1,
	}

	testcases := []struct {
		in   []string
		want string
	}{
		{
			in:   []string{"flower", "flow", "flight"},
			want: "fl",
		},
		{
			in:   []string{"dog", "racecar", "car"},
			want: "",
		},
		{
			in:   []string{"flower", "flow", "flight", "dog", "racecar", "car"},
			want: "",
		},
		{
			in:   []string{"folwer", "flow", "flight"},
			want: "f",
		},
		{
			in:   []string{"a"},
			want: "a",
		},
		{
			in:   []string{"cir", "car"},
			want: "c",
		},
		{
			in:   []string{"aaa", "aa", "aaa"},
			want: "aa",
		},
	}

	for i, f := range testfuncs {
		t.Logf("test function %d...", i)
		for j, tt := range testcases {
			ans := f(tt.in)
			if ans != tt.want {
				t.Fatalf("case[%d]: it should be %s, but got %s", j, tt.want, ans)
			}
		}
		t.Logf("... function[%d] is passed", i)
	}

}
