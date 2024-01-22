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
		{in: []int{0}, want: false},
	}

	testfuncs := []func([]int) bool{
		ContainDuplicatev41,
		ContainDuplicatev40,
		ContainDuplicatev39,
		ContainDuplicatev38,
		ContainDuplicatev37,
		ContainDuplicatev36,
		ContainDuplicatev35,
		ContainDuplicatev34,
		ContainDuplicatev33,
		ContainDuplicatev32,
		ContainDuplicatev31,
		ContainDuplicatev30,
		ContainDuplicatev29,
		ContainDuplicatev28,
		ContainDuplicatev27,
		ContainDuplicatev26,
		ContainDuplicatev25,
		ContainDuplicatev24,
		ContainDuplicatev23,
		ContainDuplicatev22,
		ContainDuplicatev21,
		ContainDuplicatev20,
		ContainDuplicatev19,
		ContainDuplicatev18,
		ContainDuplicatev17,
		ContainDuplicatev16,
		ContainDuplicatev15,
		ContainDuplicatev14,
		ContainDuplicatev13,
		ContainDuplicatev12,
		ContainDuplicatev11,
		ContainDuplicatev10,
		ContainDuplicatev9,
		ContainDuplicatev8,
		ContainDuplicatev7,
		ContainDuplicatev6,
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
