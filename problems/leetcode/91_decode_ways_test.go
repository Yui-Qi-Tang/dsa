package b75

import "testing"

func TestNumDecodings(t *testing.T) {

	testFuncs := []func(s string) int{
		numDecodingsv36,
		numDecodingsv35,
		numDecodingsv34,
		numDecodingsv33,
		numDecodingsv32,
		numDecodingsv31,
		numDecodingsv30,
		numDecodingsv29,
		numDecodingsv28,
		numDecodingsv27,
		numDecodingsv26,
		numDecodingsv25,
		numDecodingsv24,
		numDecodingsv23,
		numDecodingsv22,
		numDecodingsv21,
		numDecodingsv20,
		numDecodingsv19,
		numDecodingsv18,
		numDecodingsv17,
		numDecodingsv16,
		numDecodingsv15,
		numDecodingsv14,
		numDecodingsv13,
		numDecodingsv12,
		numDecodingsv11,
		numDecodingsv10,
		numDecodingsv9,
		numDecodingsv8,
		numDecodingsv7,
		numDecodingsv6,
		numDecodingsv5,
		numDecodingsv4,
		numDecodingsv3,
		numDecodingsv2,
		numDecodingsv1,
		NumDecodings,
		numDecodings,
	}

	testcases := []struct {
		in   string
		want int
	}{
		{
			in:   "10",
			want: 1,
		},
		{
			in:   "12",
			want: 2,
		},
		{
			in:   "226",
			want: 3,
		},
		{
			in:   "06",
			want: 0,
		},
		{
			in:   "610",
			want: 1,
		},
		{
			in:   "1201234",
			want: 3,
		},
	}

	for i, f := range testFuncs {
		t.Logf("test function %d...", i)
		for j, tt := range testcases {
			ans := f(tt.in)
			if ans != tt.want {
				t.Fatalf("=> case[%d]: it should be %d, but got %d", j, tt.want, ans)
			}
		}
		t.Logf("... function %d is passed", i)
	}
}
