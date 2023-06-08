package b75

import "testing"

func TestLRUCacheLeetcodeCase1(t *testing.T) {

	/* test case from leetcode
	["LRUCache","put","put","put","put","get","get"]
	[[2],[2,1],[1,1],[2,3],[4,1],[1],[2]]
	[null,null,null,null,null,-1,3]
	*/

	// For leetcode
	c := Constructor(2)
	c.Put(1, 1)
	c.Put(2, 2)
	if c.Get(1) != 1 {
		t.Fatalf("it should be 1, but got %d", c.Get(1))
	}
	c.Put(3, 3)
	if c.Get(2) != -1 {
		t.Fatalf("it should be -1, but got %d", c.Get(2))
	}

	c.Put(4, 4)
	if c.Get(1) != -1 {
		t.Fatalf("it should be -1, but got %d", c.Get(1))
	}
	if c.Get(3) != 3 {
		t.Fatalf("it should be 3, but got %d", c.Get(3))
	}
	if c.Get(4) != 4 {
		t.Fatalf("it should be 4, but got %d", c.Get(4))
	}
}

func TestLRUCacheLeetcodeCase2(t *testing.T) {
	/*
		input: ["LRUCache","put","put","put","put","get","get"]
		       [[2],       [2,1], [1,1],[2,3],[4,1],[1],[2]]
		expected: [null,null,null,null,null,-1,3]
	*/

	// For leetcode
	c := Constructor(2)
	c.Put(2, 1)
	c.Put(1, 1)
	c.Put(2, 3)
	c.Put(4, 1)

	if c.Get(1) != -1 {
		t.Fatalf("it should be -1, but got %d", c.Get(1))
	}
	if c.Get(2) != 3 {
		t.Fatalf("it should be 3, but got %d", c.Get(2))
	}
}

type cacher interface {
	Get(key int) int
	Put(key, value int)
}

func TestMyLRUCacheCase1(t *testing.T) {

	size := 2
	testCachers := []cacher{
		constructorv27(size),
		constructorv26(size),
		constructorv25(size),
		constructorv24(size),
		constructorv23(size),
		constructorv22(size),
		constructorv21(size),
		constructorv20(size),
		constructorv19(size),
		constructorv18(size),
		constructorv17(size),
		constructorv16(size),
		constructorv15(size),
		constructorv14(size),
		constructorv13(size),
		constructorv12(size),
		constructorv11(size),
		constructorv10(size),
		constructorv9(size),
		constructorv8(size),
		constructorv7(size),
		constructorv6(size),
		constructorv5(size),
		constructorv4(size),
		constructorv3(size),
		constructorv2(size),
		constructorv1(size),
	}

	testcases := []func(cacher){
		func(c cacher) {
			c.Put(1, 1)
			c.Put(2, 2)
			if c.Get(1) != 1 {
				t.Fatalf("0. it should be 1, but got %d", c.Get(1))
			}
			c.Put(3, 3)
			if c.Get(2) != -1 {
				t.Fatalf("1. it should be -1, but got %d", c.Get(2))
			}

			c.Put(4, 4)
			if c.Get(1) != -1 {
				t.Fatalf("2.it should be -1, but got %d", c.Get(1))
			}
			if c.Get(3) != 3 {
				t.Fatalf("3. it should be 3, but got %d", c.Get(3))
			}
			if c.Get(4) != 4 {
				t.Fatalf("4.it should be 4, but got %d", c.Get(4))
			}
		},
		func(c cacher) {
			c.Put(2, 1)
			c.Put(1, 1)
			c.Put(2, 3)
			c.Put(4, 1)

			if c.Get(1) != -1 {
				t.Fatalf("0. it should be -1, but got %d", c.Get(1))
			}
			if c.Get(2) != 3 {
				t.Fatalf("1. it should be 3, but got %d", c.Get(2))
			}
		},
	}

	for i, c := range testCachers {
		t.Logf("test function %d", i)
		for j, tt := range testcases {
			t.Logf("test case[%d]", j)
			tt(c)
			t.Logf("case[%d] is passed", j)
		}
		t.Logf("function[%d] is passed", i)
	}

}
