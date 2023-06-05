package b75

import (
	"container/heap"
)

/*
1046. Last Stone Weight
Easy
5.2K
91
Companies
You are given an array of integers stones where stones[i] is the weight of the ith stone.

We are playing a game with the stones. On each turn, we choose the heaviest two stones and smash them together. Suppose the heaviest two stones have weights x and y with x <= y. The result of this smash is:

If x == y, both stones are destroyed, and
If x != y, the stone of weight x is destroyed, and the stone of weight y has new weight y - x.
At the end of the game, there is at most one stone left.

Return the weight of the last remaining stone. If there are no stones left, return 0.



Example 1:

Input: stones = [2,7,4,1,8,1]
Output: 1
Explanation:
We combine 7 and 8 to get 1 so the array converts to [2,4,1,1,1] then,
we combine 2 and 4 to get 2 so the array converts to [2,1,1,1] then,
we combine 2 and 1 to get 1 so the array converts to [1,1,1] then,
we combine 1 and 1 to get 0 so the array converts to [1] then that's the value of the last stone.
Example 2:

Input: stones = [1]
Output: 1

Constraints:

1 <= stones.length <= 30
1 <= stones[i] <= 1000

*/

type priorityQueuev8 []int

func (p priorityQueuev8) Len() int { return len(p) }

func (p priorityQueuev8) Less(i, j int) bool {
	return p[i] > p[j]
}

func (p *priorityQueuev8) Push(x any) {
	*p = append(*p, x.(int))
}

func (p priorityQueuev8) Swap(i, j int) {
	p[i], p[j] = p[j], p[i]
}

func (p *priorityQueuev8) Pop() any {
	snap := *p
	n := snap[snap.Len()-1]
	snap = snap[:snap.Len()-1]
	*p = snap
	return n
}

func lastStoneWeightv8(stones []int) int {
	if len(stones) == 0 {
		return 0
	}

	in := priorityQueuev8(stones)
	heap.Init(&in)
	abs := func(x int) int {
		if x < 0 {
			return -x
		}
		return x
	}

	for in.Len() > 1 {
		first := heap.Pop(&in).(int)
		second := heap.Pop(&in).(int)

		if first != second {
			heap.Push(&in, abs(first-second))
		}

	}

	if in.Len() == 0 {
		return 0
	}

	return in[0]
}

type priorityQueuev7 []int

func (p priorityQueuev7) Len() int { return len(p) }

func (p priorityQueuev7) Less(i, j int) bool {
	return p[i] > p[j]
}

func (p *priorityQueuev7) Push(x any) {
	*p = append(*p, x.(int))
}

func (p priorityQueuev7) Swap(i, j int) {
	p[i], p[j] = p[j], p[i]
}

func (p *priorityQueuev7) Pop() any {
	snap := *p
	v := snap[len(snap)-1]
	snap = snap[:len(snap)-1]
	*p = snap
	return v
}

func lastStoneWeightv7(stones []int) int {
	if len(stones) == 0 {
		return 0
	}

	in := priorityQueuev7(stones)
	heap.Init(&in)

	abs := func(x int) int {
		if x < 0 {
			return -x
		}
		return x
	}

	for len(in) > 1 {
		first := heap.Pop(&in).(int)
		second := heap.Pop(&in).(int)
		if first != second {
			heap.Push(&in, abs(first-second))
		}
	}

	if len(in) == 0 {
		return 0
	}

	return in[0]
}

type priorityQueuev6 []int

func (p priorityQueuev6) Len() int { return len(p) }

func (p priorityQueuev6) Less(i, j int) bool {
	return p[i] > p[j]
}

func (p *priorityQueuev6) Push(x any) {
	*p = append(*p, x.(int))
}

func (p priorityQueuev6) Swap(i, j int) {
	p[i], p[j] = p[j], p[i]
}

func (p *priorityQueuev6) Pop() any {
	snap := *p
	n := snap[len(snap)-1]
	snap = snap[:len(snap)-1]
	*p = snap
	return n
}

func lastStoneWeightv6(stones []int) int {
	if len(stones) == 0 {
		return 0
	}

	in := priorityQueuev6(stones)
	heap.Init(&in)

	abs := func(x int) int {
		if x < 0 {
			return -x
		}

		return x
	}

	for len(in) > 1 {
		first := heap.Pop(&in).(int)
		second := heap.Pop(&in).(int)
		if first != second {
			heap.Push(&in, abs(first-second))
		}
	}

	if len(in) == 0 {
		return 0
	}
	return in[0]
}

type priorityQueuev5 []int

func (p priorityQueuev5) Len() int { return len(p) }

func (p priorityQueuev5) Less(i, j int) bool {
	return p[i] > p[j]
}

func (p *priorityQueuev5) Push(x any) {
	*p = append(*p, x.(int))
}

func (p priorityQueuev5) Swap(i, j int) {
	p[i], p[j] = p[j], p[i]
}

func (p *priorityQueuev5) Pop() any {

	snap := *p
	e := snap[len(snap)-1]
	snap = snap[:len(snap)-1]
	*p = snap

	return e
}

func lastStoneWeightv5(stones []int) int {
	if len(stones) == 0 {
		return 0
	}

	queue := priorityQueuev5(stones)
	heap.Init(&queue)
	abs := func(x int) int {
		if x < 0 {
			return -x
		}
		return x
	}

	for len(queue) > 1 {
		n1 := heap.Pop(&queue).(int)
		n2 := heap.Pop(&queue).(int)

		if n1 != n2 {
			heap.Push(&queue, abs(n1-n2))
		}
	}

	if len(queue) == 0 {
		return 0
	}

	return queue[0]
}

type priorityQueuev4 []int

func (p priorityQueuev4) Len() int { return len(p) }

func (p priorityQueuev4) Less(i, j int) bool {
	return p[i] > p[j]
}

func (p *priorityQueuev4) Push(x any) {
	*p = append(*p, x.(int))
}

func (p priorityQueuev4) Swap(i, j int) {
	p[i], p[j] = p[j], p[i]
}

func (p *priorityQueuev4) Pop() any {
	snap := *p
	n := snap[len(snap)-1]
	snap = snap[:len(snap)-1]
	*p = snap
	return n
}

func lastStoneWeightv4(stones []int) int {
	in := priorityQueuev4(stones)
	heap.Init(&in)

	abs := func(x int) int {
		if x < 0 {
			return -x
		}
		return x
	}

	for len(in) > 1 {
		first := heap.Pop(&in).(int)
		second := heap.Pop(&in).(int)

		if first != second {
			heap.Push(&in, abs(first-second))
		}
	}

	if len(in) == 0 {
		return 0
	}

	return in[0]
}

type priorityQueuev3 []int

func (p priorityQueuev3) Len() int {
	return len(p)
}

func (p priorityQueuev3) Less(i, j int) bool {
	return p[i] > p[j]
}

func (p *priorityQueuev3) Push(x any) {
	*p = append(*p, x.(int))
}

func (p priorityQueuev3) Swap(i, j int) {
	p[i], p[j] = p[j], p[i]
}

func (p *priorityQueuev3) Pop() any {
	clone := *p
	v := clone[len(clone)-1]
	clone = clone[:len(clone)-1]
	*p = clone
	return v
}

func lastStoneWeightv3(stones []int) int {
	in := priorityQueuev3(stones)
	heap.Init(&in)

	abs := func(x int) int {
		if x < 0 {
			return -x
		}
		return x
	}

	for len(in) > 1 {
		first := heap.Pop(&in).(int)
		second := heap.Pop(&in).(int)
		if first != second {
			heap.Push(&in, abs(first-second))
		}
	}

	if len(in) == 0 {
		return 0
	}

	return in[0]
}

type priorityQueuev2 []int

func (p priorityQueuev2) Len() int { return len(p) }
func (p priorityQueuev2) Less(i, j int) bool {
	return p[i] > p[j]
}

func (p *priorityQueuev2) Push(x any) {
	*p = append(*p, x.(int))
}

func (p priorityQueuev2) Swap(i, j int) {
	p[i], p[j] = p[j], p[i]
}

func (p *priorityQueuev2) Pop() any {
	old := *p
	v := old[len(old)-1]

	old = old[:len(old)-1]
	*p = old
	return v
}

func lastStoneWeightv2(stones []int) int {
	if len(stones) == 0 {
		return 0
	}

	in := priorityQueuev2(stones)
	heap.Init(&in)

	abs := func(x int) int {
		if x < 0 {
			return -x
		}

		return x
	}

	for len(in) > 1 {
		first := heap.Pop(&in).(int)
		second := heap.Pop(&in).(int)
		if first != second {
			heap.Push(&in, abs(first-second))
		}
	}

	if len(in) == 0 {
		return 0
	}

	return in[0]
}

type priorityQueuev1 []int

func (p priorityQueuev1) Len() int { return len(p) }
func (p priorityQueuev1) Less(i, j int) bool {
	return p[i] > p[j]
}

func (p *priorityQueuev1) Push(x any) {
	*p = append(*p, x.(int))
}

func (p priorityQueuev1) Swap(i, j int) {
	p[i], p[j] = p[j], p[i]
}

func (p *priorityQueuev1) Pop() any {
	q := *p
	v := q[len(q)-1]
	*p = q[:len(q)-1]
	return v
}

func lastStoneWeightv1(stones []int) int {
	if len(stones) == 0 {
		return 0
	}

	in := priorityQueuev1(stones)
	heap.Init(&in)

	abs := func(x int) int {
		if x < 0 {
			return -x
		}
		return x
	}

	for len(in) > 1 {
		first := heap.Pop(&in).(int)
		second := heap.Pop(&in).(int)
		if abs(first-second) > 0 {
			heap.Push(&in, abs(first-second))
		}
	}
	if len(in) == 0 {
		return 0
	}

	return in[0]
}

type priorityQueue []int

func (p priorityQueue) Len() int { return len(p) }
func (p priorityQueue) Less(i, j int) bool {
	return p[i] > p[j]
}
func (p *priorityQueue) Push(x any) {
	*p = append(*p, x.(int))
}
func (p priorityQueue) Swap(i, j int) {
	p[i], p[j] = p[j], p[i]
}
func (p *priorityQueue) Pop() any {
	old := *p
	oldLen := len(old)
	v := old[oldLen-1]
	*p = old[:oldLen-1]
	return v
}

func lastStoneWeight(stones []int) int {

	if len(stones) == 1 {
		return stones[0]
	}

	abs := func(a int) int {
		if a < 0 {
			return -a
		}
		return a
	}

	in := priorityQueue(stones)
	heap.Init(&in)

	for in.Len() > 1 {
		first := heap.Pop(&in).(int)
		second := heap.Pop(&in).(int)
		diff := abs(first - second)
		if diff != 0 {
			heap.Push(&in, diff)
		}
	}

	if in.Len() == 0 {
		return 0
	}

	return in[0]
}
