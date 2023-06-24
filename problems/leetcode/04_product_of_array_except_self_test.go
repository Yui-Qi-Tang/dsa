package b75

import (
	"reflect"
	"testing"
)

func TestProductOfArrayExceptSelfv1(t *testing.T) {

	testfuncs := []func([]int) []int{
		ProductOfArrayExceptSelf40,
		ProductOfArrayExceptSelf39,
		ProductOfArrayExceptSelf38,
		ProductOfArrayExceptSelf37,
		ProductOfArrayExceptSelf36,
		ProductOfArrayExceptSelf35,
		ProductOfArrayExceptSelf34,
		ProductOfArrayExceptSelf33,
		ProductOfArrayExceptSelf32,
		ProductOfArrayExceptSelf31,
		ProductOfArrayExceptSelf30,
		ProductOfArrayExceptSelf29,
		ProductOfArrayExceptSelf28,
		ProductOfArrayExceptSelf27,
		ProductOfArrayExceptSelf26,
		ProductOfArrayExceptSelf25,
		ProductOfArrayExceptSelf24,
		ProductOfArrayExceptSelf23,
		ProductOfArrayExceptSelf22,
		ProductOfArrayExceptSelf21,
		ProductOfArrayExceptSelf20,
		ProductOfArrayExceptSelf19,
		ProductOfArrayExceptSelf18,
		ProductOfArrayExceptSelf17,
		ProductOfArrayExceptSelf16,
		ProductOfArrayExceptSelf15,
		ProductOfArrayExceptSelf14,
		ProductOfArrayExceptSelf13,
		ProductOfArrayExceptSelf12,
		ProductOfArrayExceptSelf11,
		ProductOfArrayExceptSelf10,
		ProductOfArrayExceptSelf9,
		ProductOfArrayExceptSelf8,
		ProductOfArrayExceptSelf7,
		ProductOfArrayExceptSelf6,
		ProductOfArrayExceptSelf5,
		ProductOfArrayExceptSelf4,
		ProductOfArrayExceptSelf3,
		ProductOfArrayExceptSelf2,
		ProductOfArrayExceptSelf,
	}

	testcases := []struct {
		in   []int
		want []int
	}{
		{in: []int{1, 2, 3, 4}, want: []int{24, 12, 8, 6}},
		{in: []int{-1, 1, 0, -3, 3}, want: []int{0, 0, 9, 0, 0}},
		{in: []int{0, 2, 3, 4}, want: []int{24, 0, 0, 0}},
		{in: []int{3, 2, 1}, want: []int{2, 3, 6}},
	}

	for i, f := range testfuncs {
		t.Logf("test function %d...", i)
		for j, tt := range testcases {
			ans := f(tt.in)
			if !reflect.DeepEqual(ans, tt.want) {
				t.Fatalf("case[%d]: it should be %v, but got %v", j, tt.want, ans)
			}
		}
		t.Logf("... function[%d] is passed", i)
	}

}
