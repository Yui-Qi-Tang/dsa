package disjointsets

import (
	"reflect"
	"testing"
)

func TestDisjointSets(t *testing.T) {

	testcases := []struct {
		vertexs []byte
		edges   [][]byte
		want    []string
	}{
		{
			vertexs: []byte{'a', 'b', 'c', 'd', 'e', 'f', 'g', 'h', 'i', 'j'},
			edges: [][]byte{
				{'b', 'd'},
				{'e', 'g'},
				{'a', 'c'},
				{'h', 'i'},
				{'a', 'b'},
				{'e', 'f'},
				{'b', 'c'},
			},
			want: []string{"acbd", "egf", "hi", "j"},
		},
		{
			vertexs: []byte{'a', 'b', 'c', 'd', 'e', 'f', 'g', 'h', 'i', 'j', 'k'},
			edges: [][]byte{
				{'d', 'i'},
				{'f', 'k'},
				{'g', 'i'},
				{'b', 'g'},
				{'a', 'h'},
				{'i', 'j'},
				{'d', 'k'},
				{'b', 'j'},
				{'d', 'f'},
				{'g', 'j'},
				{'a', 'e'},
			},
			want: []string{"ahe", "bdigjfk", "c"},
		},
	}

	for _, tt := range testcases {
		ds := New()
		ds.MakeSets(tt.vertexs...)
		for _, edge := range tt.edges {
			ds.Union(edge[0], edge[1])
		}

		ans := ds.Collections()
		if len(ans) != len(tt.want) {
			t.Fatalf("the length of collection should be %d, but got %d", len(tt.want), len(ans))
		}

		if !reflect.DeepEqual(ans, tt.want) {
			t.Fatalf("the collection should be %v, but got %v", tt.want, ans)
		}
	}
}
