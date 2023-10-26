package b75

import "testing"

func TestLongestCommonPrefix(t *testing.T) {

	testfuncs := []func([]string) string{
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
