package b75

import (
	"fmt"
	"testing"
)

func TestWordBreak(t *testing.T) {

	testfuncs := []func(string, []string) bool{
		wordBreakv41,
		wordBreakv40,
		wordBreakv39,
		wordBreakv38,
		wordBreakv37,
		wordBreakv36,
		wordBreakv35,
		wordBreakv34,
		wordBreakv33,
		wordBreakv32,
		wordBreakv31,
		wordBreakv30,
		wordBreakv29,
		wordBreakv28,
		wordBreakv27,
		wordBreakv26,
		wordBreakv25,
		wordBreakv24,
		wordBreakv23,
		wordBreakv22,
		wordBreakv21,
		wordBreakv20,
		wordBreakv19,
		wordBreakv18,
		wordBreakv17,
		wordBreakv16,
		wordBreakv15,
		wordBreakv14,
		wordBreakv13,
		wordBreakv12,
		wordBreakv11,
		wordBreakv10,
		wordBreakv9,
		wordBreakv8,
		wordBreakv7,
		wordBreakv6,
		wordBreakv5,
		wordBreakv4,
		wordBreakv3,
		wordBreakv2,
		wordBreakv1,
	}

	testcases := []struct {
		s    string
		dict []string
		want bool
	}{
		{
			s:    "leetcode",
			dict: []string{"leet", "code"},
			want: true,
		},
		{
			s:    "applepenapple",
			dict: []string{"apple", "pen"},
			want: true,
		},
		{
			s:    "catsandog",
			dict: []string{"cats", "dog", "sand", "and", "cat"},
			want: false,
		},
	}

	for i, f := range testfuncs {
		t.Run(fmt.Sprintf("test function[%d]", i), func(t *testing.T) {
			for j, tt := range testcases {
				ans := f(tt.s, tt.dict)
				if ans != tt.want {
					t.Fatalf("case[%d]: it should be %v, but got %v", j, tt.want, ans)
				}
			}
		})
	}
}
