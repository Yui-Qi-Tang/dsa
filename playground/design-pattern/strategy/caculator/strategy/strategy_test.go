package strategy

import "testing"

func TestAdditionImplementation(t *testing.T) {
	testcases := []struct {
		var1, var2, want int
	}{
		{
			var1: 1,
			var2: 2,
			want: 3,
		},
		{
			var1: 1,
			var2: -1,
			want: 0,
		},
	}

	for i, tt := range testcases {
		ans := Add{InputTwoValues{Value1: tt.var1, Value2: tt.var2}}.Caculate()
		if ans != tt.want {
			t.Fatalf("case[%d]: it should be %d, but got %d", i, tt.want, ans)
		}
	}

}

func TestMultiplicationImplementation(t *testing.T) {
	testcases := []struct {
		var1, var2, want int
	}{
		{
			var1: 1,
			var2: 2,
			want: 2,
		},
		{
			var1: 1,
			var2: -1,
			want: -1,
		},
	}

	for i, tt := range testcases {
		ans := Multiplication{InputTwoValues{Value1: tt.var1, Value2: tt.var2}}.Caculate()
		if ans != tt.want {
			t.Fatalf("case[%d]: it should be %d, but got %d", i, tt.want, ans)
		}
	}
}
