package mathgame

import (
	"container/heap"
	"sort"
)

// Finder represents the interface for finding the elements
type Finder interface {
	// Top returns elements with n
	Top(n int) []int
}

type sliceWithSort []int

func (s sliceWithSort) Len() int { return len(s) }

func (s sliceWithSort) Top(n int) []int {
	if s.Len() < n {
		n = s.Len()
	}
	sort.Slice(s, func(i, j int) bool {
		return s[i] > s[j]
	})
	return s[:n]
}

type sliceWithHeap []int

func (s sliceWithHeap) Len() int           { return len(s) }
func (s sliceWithHeap) Less(i, j int) bool { return s[i] < s[j] }
func (s *sliceWithHeap) Push(x any)        { *s = append(*s, x.(int)) }
func (s sliceWithHeap) Swap(i, j int)      { s[i], s[j] = s[j], s[i] }
func (s *sliceWithHeap) Pop() any {
	old := *s
	n := old[old.Len()-1]
	*s = old[:s.Len()-1]
	return n
}

func (sliceWithHeap) Init(size int) sliceWithHeap {
	s := make(sliceWithHeap, 0, size)
	heap.Init(&s)
	return s
}

func (s sliceWithHeap) Top(in []int, n int) []int {
	for _, v := range in {
		heap.Push(&s, v)
		if s.Len() > n {
			heap.Pop(&s)
		}
	}
	result := make([]int, min(s.Len(), n))
	for s.Len() > 0 {
		v := heap.Pop(&s).(int)
		result[s.Len()] = v
	}

	return result
}

func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}
