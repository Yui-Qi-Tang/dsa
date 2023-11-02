package heap

import (
	"math/rand"
	"testing"
	"time"
)

func TestHeap(t *testing.T) {

	rand.NewSource(time.Now().Unix())
	input := rand.Perm(10)

	var arr []int
	push(&arr, input...)
	t.Log("start test heap on []int...")
	first := pop(&arr)
	for i := 0; i < len(input)-1; i++ {
		second := pop(&arr)
		if first > second {
			t.Fatal("can not statify min-heapify")
		}
		first = second
	}
	t.Log("... Passed")
}

func TestInts(t *testing.T) {
	rand.NewSource(time.Now().Unix())
	input := rand.Perm(10)
	intsHeap := make(ints, 0, len(input))
	t.Log("start test heap on ints...")
	intsHeap.push(input...)
	first := intsHeap.pop()
	for i := 0; i < len(input)-1; i++ {
		second := intsHeap.pop()
		if first > second {
			t.Fatal("can not statify min-heapify")
		}
		first = second
	}
	t.Log("... Passed")
}
