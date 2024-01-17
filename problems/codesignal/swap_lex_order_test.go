package codesignal

import (
	"fmt"
	"testing"
)

func TestSwapLexOrder(t *testing.T) {
	testfuncs := []func(string, [][]int) string{
		swapLexOrderv44,
		swapLexOrderv43,
		swapLexOrderv42,
		swapLexOrderv41,
		swapLexOrderv40,
		swapLexOrderv39,
		swapLexOrderv38,
		swapLexOrderv37,
		swapLexOrderv36,
		swapLexOrderv35,
		swapLexOrderv34,
		swapLexOrderv33,
		swapLexOrderv32,
		swapLexOrderv31,
		swapLexOrderv30,
		swapLexOrderv29,
		swapLexOrderv28,
		swapLexOrderv27,
		swapLexOrderv26,
		swapLexOrderv25,
		swapLexOrderv24,
		swapLexOrderv23,
		swapLexOrderv22,
		swapLexOrderv21,
		swapLexOrderv20,
		swapLexOrderv19,
		swapLexOrderv18,
		swapLexOrderv17,
		swapLexOrderv16,
		swapLexOrderv15,
		swapLexOrderv14,
		swapLexOrderv13,
		swapLexOrderv12,
		swapLexOrderv11,
		swapLexOrderv10,
		swapLexOrderv9,
		swapLexOrderv8,
		swapLexOrderv7,
		swapLexOrderv6,
		swapLexOrderv5,
		swapLexOrderv4,
		swapLexOrderv3,
		swapLexOrderv2,
		swapLexOrderv1,
	}

	testcases := []struct {
		in    string
		pairs [][]int
		want  string
	}{
		{
			in: "abdc",
			pairs: [][]int{
				{1, 4},
				{3, 4},
			},
			want: "dbca",
		},
		{
			in: "abcdefgh",
			pairs: [][]int{
				{1, 4},
				{7, 8},
			},
			want: "dbcaefhg",
		},
		{
			in: "acxrabdz",
			pairs: [][]int{
				{1, 3},
				{6, 8},
				{3, 8},
				{2, 7},
			},
			want: "zdxrabca",
		},
		{
			in:    "z",
			pairs: [][]int{},
			want:  "z",
		},
		{
			in: "dznsxamwoj",
			pairs: [][]int{
				{1, 2},
				{3, 4},
				{6, 5},
				{8, 10},
			},
			want: "zdsnxamwoj",
		},
		{
			in: "fixmfbhyutghwbyezkveyameoamqoi",
			pairs: [][]int{
				{8, 5},
				{10, 8},
				{4, 18},
				{20, 12},
				{5, 2},
				{17, 2},
				{13, 25},
				{29, 12},
				{22, 2},
				{17, 11},
			},
			want: "fzxmybhtuigowbyefkvhyameoamqei",
		},
		{
			in: "lvvyfrbhgiyexoirhunnuejzhesylojwbyatfkrv",
			pairs: [][]int{
				{13, 23},
				{13, 28},
				{15, 20},
				{24, 29},
				{6, 7},
				{3, 4},
				{21, 30},
				{2, 13},
				{12, 15},
				{19, 23},
				{10, 19},
				{13, 14},
				{6, 16},
				{17, 25},
				{6, 21},
				{17, 26},
				{5, 6},
				{12, 24},
			},
			want: "lyyvurrhgxyzvonohunlfejihesiebjwbyatfkrv",
		},
		{
			in:    "a",
			pairs: [][]int{},
			want:  "a",
		},
	}

	for i, f := range testfuncs {
		t.Run(fmt.Sprintf("test function %d", i), func(t *testing.T) {
			for j, tt := range testcases {
				ans := f(tt.in, tt.pairs)
				if ans != tt.want {
					t.Fatalf("case[%d]: it should be %s, but got %s", j, tt.want, ans)
				}
			}
		})
	}
}
