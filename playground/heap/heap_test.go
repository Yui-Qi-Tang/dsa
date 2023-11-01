package heap

import (
	"math/rand"
	"testing"
	"time"
)

func TestHeap(t *testing.T) {

	rand.NewSource(time.Now().Unix())

	var arr []int
	input := rand.Perm(10)
	push(&arr, input...)
	t.Log("input=>", input)
	for i := 0; i < len(input); i++ {
		t.Log(pop(&arr))
	}
}
