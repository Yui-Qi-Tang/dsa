package b75

import "testing"

func TestCalculateMinimumHP(t *testing.T) {

	testFuncs := []func([][]int) int{
		calculateMinimumHPv43,
		calculateMinimumHPv42,
		calculateMinimumHPv41,
		calculateMinimumHPv40,
		calculateMinimumHPv39,
		calculateMinimumHPv38,
		calculateMinimumHPv37,
		calculateMinimumHPv36,
		calculateMinimumHPv35,
		calculateMinimumHPv34,
		calculateMinimumHPv33,
		calculateMinimumHPv32,
		calculateMinimumHPv31,
		calculateMinimumHPv30,
		calculateMinimumHPv29,
		calculateMinimumHPv28,
		calculateMinimumHPv27,
		calculateMinimumHPv26,
		calculateMinimumHPv25,
		calculateMinimumHPv24,
		calculateMinimumHPv23,
		calculateMinimumHPv22,
		calculateMinimumHPv21,
		calculateMinimumHPv20,
		calculateMinimumHPv19,
		calculateMinimumHPv18,
		calculateMinimumHPv17,
		calculateMinimumHPv16,
		calculateMinimumHPv15,
		calculateMinimumHPv14,
		calculateMinimumHPv13,
		calculateMinimumHPv12,
		calculateMinimumHPv11,
		calculateMinimumHPv10,
		calculateMinimumHPv9,
		calculateMinimumHPv8,
		calculateMinimumHPv7,
		calculateMinimumHPv6,
		calculateMinimumHPv5,
		calculateMinimumHPv4,
		calculateMinimumHPv3,
		calculateMinimumHPv2,
		//calculateMinimumHPv1,
	}

	tescases := []struct {
		in   [][]int
		want int
	}{
		{
			in: [][]int{
				{-2, -3, 3},
				{-5, -10, 1},
				{10, 30, -5},
			},
			want: 7,
		},
		{
			in:   [][]int{{0}},
			want: 1,
		},
		// { I am not sure this case that is correct
		// 	in:   [][]int{{1}},
		// 	want: 0,
		// },
		{
			in:   [][]int{{-2}},
			want: 3,
		},
		{
			in:   [][]int{{5, 1}, {30, -100}},
			want: 66,
		},
		{
			in:   [][]int{{-2, -3}, {8, -10}},
			want: 5,
		},
	}

	for i, f := range testFuncs {
		t.Logf("test function %d...", i)
		for j, tt := range tescases {
			tmp := make([][]int, len(tt.in))
			for k, v := range tt.in {
				tmp[k] = make([]int, len(v))
				copy(tmp[k], v)
			}
			ans := f(tmp)
			if ans != tt.want {
				t.Fatalf("=> case[%d]: it should be %d, but got %d", j, tt.want, ans)
			}
		}
		t.Logf("function %d is passed...", i)

	}

}
