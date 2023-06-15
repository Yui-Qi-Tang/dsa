package b75

import "testing"

func TestCodes(t *testing.T) {

	type codecer interface {
		serialize(root *TreeNode) string
		deserialize(data string) *TreeNode
	}

	/*
		        We should to consider the nagative number, so for keeping simple, we can choice
				the format of the serilaized string as "num1,n,num2,..." where 'num_i' is the number which contains postive and nagetive and
				'n' is denoted the null value

				the test will ignore the null values of the tree, so just checkout the

				codec := constructor()
				tree := codec.deserialize(input)
				serString := codec.serialize(tree)
				tree2 := codec.deserialize(serString)

				pass if codec.serialize(tree2) == serString is true otherwise failed

	*/

	leetcodeVerCodec := Constructor()
	testfuncs := []codecer{
		constructorv17(),
		constructorv16(),
		constructorv15(),
		constructorv14(),
		constructorv13(),
		constructorv12(),
		constructorv11(),
		constructorv10(),
		constructorv9(),
		constructorv8(),
		constructorv7(),
		constructorv6(),
		constructorv5(),
		constructorv4(),
		constructorv3(),
		constructorv2(),
		Constructorv1(),
		&leetcodeVerCodec,
	}

	testcases := []struct {
		in string
	}{
		{
			in: "4,-7,-3,n,n,-9,-3,9,-7,-4,n,6,n,-6,-6,n,n,0,6,5,n,9,n,n,-1,-4,n,n,n,-2",
		},
		{
			in: "1,2,3,n,n,4,5",
		},
		{
			in: "",
		},
	}

	var sameTree func(p, q *TreeNode) bool

	sameTree = func(p, q *TreeNode) bool {
		if p == nil || q == nil {
			return p == q
		}

		if p.Val != q.Val {
			return false
		}

		return sameTree(p.Left, q.Left) && sameTree(p.Right, q.Right)
	}

	for i, f := range testfuncs {
		t.Logf("test function %d", i)

		for j, tt := range testcases {
			tree1 := f.deserialize(tt.in)
			serialize1 := f.serialize(tree1)
			tree2 := f.deserialize(serialize1)
			serialize2 := f.serialize(tree2)

			if !sameTree(tree1, tree2) {
				t.Fatalf("case[%d]: the tree is generated from %s is not equal to the tree is generated from %s", j, tt.in, serialize1)
			}

			if serialize1 != serialize2 {
				t.Fatalf("case[%d]: it should be %s , but got %s", j, serialize1, serialize2)
			}
		}

		t.Logf("function[%d] is passed", i)
	}

}
