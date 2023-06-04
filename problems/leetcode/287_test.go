package b75

import "testing"

func TestFindDuplicate(t *testing.T) {

	testfuncs := []func([]int) int{
		findDuplicatev29,
		findDuplicatev28,
		findDuplicatev27,
		findDuplicatev26,
		findDuplicatev25,
		findDuplicatev24,
		findDuplicatev23,
		findDuplicatev22,
		findDuplicatev21,
		findDuplicatev20,
		findDuplicatev19,
		findDuplicatev18,
		findDuplicatev17,
		findDuplicatev16,
		findDuplicatev15,
		findDuplicatev14,
		findDuplicatev13,
		findDuplicatev12,
		findDuplicatev11,
		findDuplicatev10,
		findDuplicatev9,
		findDuplicatev8,
		findDuplicatev7,
		findDuplicatev6,
		findDuplicatev5,
		findDuplicatev4,
		findDuplicatev3,
		findDuplicatev2,
		findDuplicatev1,
	}

	testcases := []struct {
		in   []int
		want int
	}{
		{
			in:   []int{1, 3, 4, 2, 2},
			want: 2,
		},
		{
			in:   []int{3, 1, 3, 4, 2},
			want: 3,
		},
	}

	for i, f := range testfuncs {
		t.Logf("test function %d", i)
		for j, tt := range testcases {
			ans := f(tt.in)
			if ans != tt.want {
				t.Fatalf("case[%d]: it should be %d, but got %d", j, tt.want, ans)
			}
		}
		t.Logf("function[%d] is passed", i)
	}

}
