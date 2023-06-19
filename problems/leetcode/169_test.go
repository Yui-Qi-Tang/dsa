package b75

import "testing"

func TestMajorityElement(t *testing.T) {

	testfuncs := []func([]int) int{
		majorityElementv20,
		majorityElementv19,
		majorityElementv18,
		majorityElementv17,
		majorityElementv16,
		majorityElementv16,
		majorityElementv15,
		majorityElementv14,
		majorityElementv13,
		majorityElementv12,
		majorityElementv11,
		majorityElementv10,
		majorityElementv9,
		majorityElementv8,
		majorityElementv7,
		majorityElementv6,
		majorityElementv5,
		majorityElementv4,
		majorityElementv3,
		majorityElementv2,
		majorityElementv1,
		majorityElement,
	}

	testcases := []struct {
		in   []int
		want int
	}{
		{
			in:   []int{3, 2, 3},
			want: 3,
		},
		{
			in:   []int{2, 2, 1, 1, 1, 2, 2},
			want: 2,
		},
	}

	for i, f := range testfuncs {
		t.Logf("test function %d...", i)

		for j, tt := range testcases {
			ans := f(tt.in)
			if ans != tt.want {
				t.Fatalf("case[%d]: it should be %d, but got %d", j, tt.want, ans)
			}
		}

		t.Logf("function[%d] is passed", i)
	}

}
