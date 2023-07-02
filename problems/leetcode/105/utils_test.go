package b75

import "testing"

func TestBuild(t *testing.T) {
	ins := [][]int{
		{1, 2, 5, 3, 4, 6, 7},
		{1, 2, 3, null, null, 4},
	}

	for i := range ins {
		root := build(ins[i])
		t.Log("origin=>", ins[i])
		t.Log("preorder =>", printPreorder(root))
		t.Log("inorder =>", printInorder(root))
		t.Log("postorder =>", printPostorder(root))
	}
}
