package b75

import "testing"

func TestContainDuplicate(t *testing.T) {

	testcases := []struct {
		in   []int
		want bool
	}{
		{in: []int{1, 2, 3, 1}, want: true},
		{in: []int{1, 2, 3, 4}, want: false},
		{in: []int{1, 1, 1, 3, 3, 4, 3, 2, 4, 2}, want: true},
	}

	testfuncs := []func([]int) bool{
		ContainDuplicatev5,
		ContainDuplicatev4,
		ContainDuplicatev3,
		ContainDuplicatev2,
		ContainDuplicatev1,
		ContainDuplicate,
	}

	for i, f := range testfuncs {
		t.Logf("start test function: %d", i)
		for j, tt := range testcases {
			ans := f(tt.in)
			if ans != tt.want {
				t.Fatalf("case[%d]: it should be: %v, but got: %v", j, tt.want, ans)
			}
		}
		t.Logf("function[%d] is passed", i)
	}

}
