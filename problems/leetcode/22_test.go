package b75

import (
	"reflect"
	"testing"
)

func TestGenerateParenthesis(t *testing.T) {
	testfuncs := []func(int) []string{
		generateParenthesisv34,
		generateParenthesisv33,
		generateParenthesisv32,
		generateParenthesisv31,
		generateParenthesisv30,
		generateParenthesisv29,
		generateParenthesisv28,
		generateParenthesisv27,
		generateParenthesisv26,
		generateParenthesisv25,
		generateParenthesisv24,
		generateParenthesisv23,
		generateParenthesisv22,
		generateParenthesisv21,
		generateParenthesisv20,
		generateParenthesisv19,
		generateParenthesisv18,
		generateParenthesisv17,
		generateParenthesisv16,
		generateParenthesisv15,
		generateParenthesisv14,
		generateParenthesisv13,
		generateParenthesisv12,
		generateParenthesisv11,
		generateParenthesisv10,
		generateParenthesisv9,
		generateParenthesisv8,
		generateParenthesisv7,
		generateParenthesisv6,
		generateParenthesisv5,
		generateParenthesisv4,
		generateParenthesisv3,
		generateParenthesisv2,
		generateParenthesisv1,
	}

	testcases := []struct {
		in   int
		want []string
	}{
		{
			in:   3,
			want: []string{"((()))", "(()())", "(())()", "()(())", "()()()"},
		},
		{
			in:   1,
			want: []string{"()"},
		},
	}

	for i, f := range testfuncs {
		t.Logf("test function %d...", i)
		for j, tt := range testcases {
			ans := f(tt.in)
			if !reflect.DeepEqual(ans, tt.want) {
				t.Fatalf("case[%d]: it should be %v, but got %v", j, tt.want, ans)
			}
		}
		t.Logf("...function[%d] is passed", i)
	}
}
