package b75

import "testing"

func TestIsInterleave(t *testing.T) {

	testfuncs := []func(s1 string, s2 string, s3 string) bool{
		isInterleavev5,
		isInterleavev4,
		isInterleavev3,
		isInterleavev2,
		isInterleavev1,
	}

	testcases := []struct {
		s1, s2, s3 string
		want       bool
	}{
		{
			s1:   "aabcc",
			s2:   "dbbca",
			s3:   "aadbbcbcac",
			want: true,
		},
		{
			s1:   "aabcc",
			s2:   "dbbca",
			s3:   "aadbbbaccc",
			want: false,
		},
		{
			s1:   "",
			s2:   "",
			s3:   "",
			want: true,
		},
		{
			s1:   "",
			s2:   "",
			s3:   "a",
			want: false,
		},
		{
			s1:   "abababababababababababababababababababababababababababababababababababababababababababababababababbb",
			s2:   "babababababababababababababababababababababababababababababababababababababababababababababababaaaba",
			s3:   "abababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababbb",
			want: false,
		},
	}

	for i, f := range testfuncs {
		t.Logf("start testing function %d...", i)
		for j, tt := range testcases {
			ans := f(tt.s1, tt.s2, tt.s3)
			if ans != tt.want {
				t.Fatalf("case[%d]: it shoud be %v but got %v", j, tt.want, ans)
			}
		}
		t.Logf("... passed")
	}
}
