package b75

import (
	"testing"
)

func TestMaximalSquare(t *testing.T) {

	tfs := []func([][]byte) int{
		maximalSquarev30,
		maximalSquarev29,
		maximalSquarev28,
		maximalSquarev27,
		maximalSquarev26,
		maximalSquarev25,
		maximalSquarev24,
		maximalSquarev23,
		maximalSquarev22,
		maximalSquarev21,
		maximalSquarev20,
		maximalSquarev19,
		maximalSquarev18,
		maximalSquarev17,
		maximalSquarev16,
		maximalSquarev15,
		maximalSquarev14,
		maximalSquarev13,
		maximalSquarev12,
		maximalSquarev11,
		maximalSquarev10,
		maximalSquarev9,
		maximalSquarev8,
		maximalSquarev7,
		maximalSquarev6,
		maximalSquarev5,
		maximalSquarev4,
		maximalSquarev3,
		maximalSquarev2,
		maximalSquarev1,
	}

	testcase := []struct {
		in   [][]byte
		want int
	}{
		{
			in: [][]byte{
				{'1', '0', '1', '0', '0'},
				{'1', '0', '1', '1', '1'},
				{'1', '1', '1', '1', '1'},
				{'1', '0', '0', '1', '0'},
			},
			want: 4,
		},
		{
			in: [][]byte{
				{'1', '0'},
				{'0', '1'},
			},
			want: 1,
		},
		{
			in: [][]byte{
				{'1', '1', '1', '0', '0'},
				{'1', '1', '1', '0', '0'},
				{'1', '1', '1', '1', '1'},
				{'0', '0', '1', '1', '1'},
				{'0', '0', '1', '1', '1'},
			},
			want: 9,
		},
		{
			in: [][]byte{
				{'1'},
			},
			want: 1,
		},
		{
			in: [][]byte{
				{'0'},
			},
			want: 0,
		},
	}

	for i, tf := range tfs {
		t.Logf("test function %d ...", i)
		for j, tt := range testcase {
			ans := tf(tt.in)
			if ans != tt.want {
				t.Fatalf("=> case[%d]: it should be %d, but got %d", j, tt.want, ans)
			}
		}
		t.Logf("... function[%d] is passed", i)
	}

}
