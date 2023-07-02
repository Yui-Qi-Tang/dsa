package graph

type node struct {
	val       string
	vistied   bool
	neighbors []*node
}
