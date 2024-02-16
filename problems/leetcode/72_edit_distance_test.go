package b75

import (
	"testing"
)

func TestEditDistance(t *testing.T) {

	testFuncs := []func(string, string) int{
		editDistancev40,
		editDistancev39,
		editDistancev38,
		editDistancev37,
		editDistancev36,
		editDistancev35,
		editDistancev34,
		editDistancev33,
		editDistancev32,
		editDistancev31,
		editDistancev30,
		editDistancev29,
		editDistancev28,
		editDistancev27,
		editDistancev26,
		editDistancev25,
		editDistancev24,
		editDistancev23,
		editDistancev22,
		editDistancev21,
		editDistancev20,
		editDistancev19,
		editDistancev18,
		editDistancev17,
		editDistancev16,
		editDistancev15,
		editDistancev14,
		editDistancev13,
		editDistancev12,
		editDistancev11,
		editDistancev10,
		editDistancev9,
		editDistancev8,
		editDistancev7,
		editDistancev6,
		editDistancev5,
		editDistancev4,
		editDistancev3,
		editDistancev2,
		editDistancev1,
		EditDistance,
	}

	testcases := []struct {
		word1, word2 string
		want         int
	}{
		{
			word1: "horse", word2: "ros",
			want: 3,
		},
		{
			word1: "acd", word2: "abd",
			want: 1,
		},
		{
			word1: "abc", word2: "bc",
			want: 1,
		},
		{
			word1: "", word2: "bc",
			want: 2,
		},
		{
			word1: "bc", word2: "",
			want: 2,
		},
		{
			word1: "bc", word2: "bc",
			want: 0,
		},
		{
			word1: "zoologicoarchaeologist",
			word2: "zoogeologist",
			want:  10,
		},
	}

	for i, f := range testFuncs {
		t.Logf("test function %d...", i)
		for j, tt := range testcases {
			ans := f(tt.word1, tt.word2)
			if ans != tt.want {
				t.Fatalf("=> case[%d]: it should be %d, but got %d", j, tt.want, ans)
			}
		}
		t.Logf("... function %d is passed", i)
	}

}
