package graph

import (
	"testing"
)

func TestDFS(t *testing.T) {
	// create graph
	a := &node{val: "A", neighbors: make([]*node, 0)}
	b := &node{val: "B", neighbors: make([]*node, 0)}
	c := &node{val: "C", neighbors: make([]*node, 0)}
	d := &node{val: "D", neighbors: make([]*node, 0)}
	e := &node{val: "E", neighbors: make([]*node, 0)}
	f := &node{val: "F", neighbors: make([]*node, 0)}

	a.neighbors = append(a.neighbors, b, d, e)
	b.neighbors = append(b.neighbors, a, c)
	c.neighbors = append(c.neighbors, b, e)
	d.neighbors = append(d.neighbors, a, e)
	e.neighbors = append(e.neighbors, a, c, d, f)
	f.neighbors = append(f.neighbors, e)

	// DFS
	stack := []*node{a} // get started from A
	a.vistied = true
	result := ""
	for len(stack) > 0 {
		n := stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		result += n.val
		for _, nb := range n.neighbors {
			if !nb.vistied {
				nb.vistied = true
				stack = append(stack, nb)
			}
		}
	}
	if result != "AEFCDB" {
		t.Fatalf("it should be 'AEFCDB', but got '%s'", result)
	}
	t.Log("passed")
}

func TestBFS(t *testing.T) {
	// create graph
	a := &node{val: "A", neighbors: make([]*node, 0)}
	b := &node{val: "B", neighbors: make([]*node, 0)}
	c := &node{val: "C", neighbors: make([]*node, 0)}
	d := &node{val: "D", neighbors: make([]*node, 0)}
	e := &node{val: "E", neighbors: make([]*node, 0)}
	f := &node{val: "F", neighbors: make([]*node, 0)}

	a.neighbors = append(a.neighbors, b, d, e)
	b.neighbors = append(b.neighbors, a, c)
	c.neighbors = append(c.neighbors, b, e)
	d.neighbors = append(d.neighbors, a, e)
	e.neighbors = append(e.neighbors, a, c, d, f)
	f.neighbors = append(f.neighbors, e)

	// BFS
	queue := []*node{a} // get started from A
	a.vistied = true
	result := ""
	for len(queue) > 0 {
		n := queue[0]
		queue = queue[1:]
		result += n.val
		for _, nb := range n.neighbors {
			if !nb.vistied {
				nb.vistied = true
				queue = append(queue, nb)
			}
		}
	}
	if result != "ABDECF" {
		t.Fatalf("it should be 'AEFCDB', but got '%s'", result)
	}
	t.Log("passed")
}
