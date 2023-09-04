package mathgame

import (
	"math/rand"
	"strconv"
	"testing"
	"time"
)

func TestHeap(t *testing.T) {
	heap := make(arr, 0)

	rand.New(rand.NewSource(time.Now().Unix()))
	nums := rand.Perm(10)

	for _, num := range nums {
		heap.push(num)
	}

	t.Log(nums)
	for heap.len() > 0 {
		v := heap.pop()
		t.Log(v)
	}

	t.Log(top(nums, 5))
}

func TestHeapElements(t *testing.T) {
	heap := make(elements, 0)

	rand.New(rand.NewSource(time.Now().Unix()))
	nums := rand.Perm(10)

	for _, num := range nums {
		heap.push(element{
			id:   num + 1,
			name: "user" + strconv.Itoa(num),
			cnt:  num,
		})
	}

	t.Log(nums)
	for heap.len() > 0 {
		v := heap.pop()
		t.Log(v)
	}

}
