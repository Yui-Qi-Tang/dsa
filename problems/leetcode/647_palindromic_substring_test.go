package b75

import "testing"

func TestCountSubstrings(t *testing.T) {
	testFuncs := []func(s string) int{
		countSubstringsv37,
		countSubstringsv36,
		countSubstringsv35,
		countSubstringsv34,
		countSubstringsv33,
		countSubstringsv32,
		countSubstringsv31,
		countSubstringsv30,
		countSubstringsv29,
		countSubstringsv28,
		countSubstringsv27,
		countSubstringsv26,
		countSubstringsv25,
		countSubstringsv24,
		countSubstringsv23,
		countSubstringsv22,
		countSubstringsv21,
		countSubstringsv20,
		countSubstringsv19,
		countSubstringsv18,
		countSubstringsv17,
		countSubstringsv16,
		countSubstringsv15,
		countSubstringsv14,
		countSubstringsv13,
		countSubstringsv12,
		countSubstringsv11,
		countSubstringsv10,
		countSubstringsv9,
		countSubstringsv8,
		countSubstringsv7,
		countSubstringsv6,
		countSubstringsv5,
		countSubstringsv4,
		countSubstringsv3,
		countSubstrings,
		countSubstringsv1,
		countSubstringsv2,
	}

	testcases := []struct {
		in   string
		want int
	}{
		{
			in:   "abc",
			want: 3,
		},
		{
			in:   "aaa",
			want: 6,
		},
		{
			in:   "fdsklf",
			want: 6,
		},
	}

	for i, f := range testFuncs {
		t.Logf("test function %d...", i)

		for j, tt := range testcases {

			ans := f(tt.in)
			if ans != tt.want {
				t.Fatalf("=> case [%d]: it should be %d, but got %d", j, tt.want, ans)
			}

		}

		t.Logf("... function %d is passed...", i)
	}

}
