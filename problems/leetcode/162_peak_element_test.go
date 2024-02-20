package b75

import "testing"

func TestFindPeakElement(t *testing.T) {

	testFuncs := []func(nums []int) int{
		findPeakElementv45,
		findPeakElementv44,
		findPeakElementv43,
		findPeakElementv42,
		findPeakElementv41,
		findPeakElementv40,
		findPeakElementv39,
		findPeakElementv38,
		findPeakElementv37,
		findPeakElementv36,
		findPeakElementv35,
		findPeakElementv34,
		findPeakElementv33,
		findPeakElementv32,
		findPeakElementv31,
		findPeakElementv30,
		findPeakElementv29,
		findPeakElementv28,
		findPeakElementv27,
		findPeakElementv26,
		findPeakElementv25,
		findPeakElementv24,
		findPeakElementv23,
		findPeakElementv22,
		findPeakElementv21,
		findPeakElementv20,
		findPeakElementv19,
		findPeakElementv18,
		findPeakElementv17,
		findPeakElementv16,
		findPeakElementv15,
		findPeakElementv14,
		findPeakElementv13,
		findPeakElementv12,
		findPeakElementv11,
		findPeakElementv10,
		findPeakElementv9,
		findPeakElementv8,
		findPeakElementv7,
		findPeakElementv6,
		findPeakElementv5,
		findPeakElementv4,
		findPeakElementv3,
		findPeakElementv2,
		findPeakElementv1,
		FindPeakElement,
		findPeakElement,
	}

	testcases := []struct {
		in      []int
		wantIdx int
	}{
		{
			in:      []int{1, 2, 3, 1},
			wantIdx: 2, // nums[2] = 3
		},
		{
			in:      []int{1, 2, 1, 3, 5, 6, 4},
			wantIdx: 5, // nums[5] = 6
		},
		{
			in:      []int{1, 2, 1, 3, 5, 6, 4, 7},
			wantIdx: 5, // nums[5] = 6
		},
	}

	for i, f := range testFuncs {
		t.Logf("test function %d ...", i)
		for j, tt := range testcases {
			ans := f(tt.in)
			if ans != tt.wantIdx {
				t.Fatalf("=> case [%d]: it should be %d, but got %d", j, tt.wantIdx, ans)
			}
		}
		t.Logf("... function %d is passed", i)
	}
}
