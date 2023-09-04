package mathgame

import "sort"

type arr []int

func (a *arr) push(value int) {
	*a = append(*a, value)
	a.up(len(*a) - 1)
}

func (a arr) up(i int) {
	for i > 0 {
		parent := (i - 1) / 2
		if a.less(i, parent) {
			return
		}
		a.swap(i, parent)
		i = parent
	}
}

func (a arr) swap(i, j int) {
	a[i], a[j] = a[j], a[i]
}

func (a arr) less(i, j int) bool {
	return a[i] > a[j]
}

func (a *arr) pop() int {
	if len(*a) == 0 {
		panic("no element")
	}

	old := *a
	a.swap(0, len(old)-1)
	v := old[len(old)-1]
	*a = old[:len(old)-1]

	a.down(0)
	return v
}

func (a arr) len() int {
	return len(a)
}

func (a arr) down(i int) {
	for {
		l := 2*i + 1
		r := 2*i + 2
		curr := i

		if l < a.len() && a.less(curr, l) {
			curr = l
		}
		if r < a.len() && a.less(curr, r) {
			curr = r
		}

		if curr == i {
			return
		}

		a.swap(i, curr)
		i = curr
	}
}

func top(nums []int, n int) []int {
	myheap := make(arr, 0, n)
	for i := range nums {
		myheap.push(nums[i])
		if myheap.len() > n {
			myheap.pop()
		}
	}
	result := make([]int, min(myheap.len(), n))
	for myheap.len() > 0 {
		v := myheap.pop()
		result[myheap.len()] = v
	}
	return result
}

type element struct {
	id   int
	cnt  int
	name string
}

type elements []element

func (e *elements) push(data element) {
	*e = append(*e, data)
	e.up(e.len() - 1)
}

func (e elements) len() int {
	return len(e)
}

func (e elements) less(i, j int) bool {
	return e[i].cnt > e[j].cnt
}

func (e elements) swap(i, j int) {
	e[i], e[j] = e[j], e[i]
}

func (e *elements) up(curr int) {
	for curr > 0 {
		parent := (curr - 1) / 2
		if e.less(curr, parent) {
			return
		}
		e.swap(curr, parent)
		curr = parent
	}
}

func (e *elements) pop() element {
	if e.len() == 0 {
		panic("no elements")
	}

	e.swap(0, e.len()-1)

	clone := *e
	v := clone[clone.len()-1]
	*e = clone[:clone.len()-1]

	// down
	e.down(0)

	return v
}

func (e elements) down(curr int) {
	for {
		l := 2*curr + 1
		r := 2*curr + 2
		i := curr
		if l < e.len() && e.less(i, l) {
			i = l
		}

		if r < e.len() && e.less(i, r) {
			i = r
		}

		if i == curr {
			return
		}

		e.swap(curr, i)
		curr = i
	}
}

func (e elements) Len() int {
	return e.len()
}

func (e elements) Top(n int) []element {
	if e.len() < n {
		n = e.len()
	}

	sort.Slice(e, func(i, j int) bool {
		return e[i].cnt > e[j].cnt
	})

	return e[:n]
}

func topElements(es []element, n int) []element {
	myheap := make(elements, 0, n)
	for i := range es {
		myheap.push(es[i])
		if myheap.len() > n {
			myheap.pop()
		}
	}
	result := make([]element, min(myheap.len(), n))
	for myheap.len() > 0 {
		v := myheap.pop()
		result[myheap.len()] = v
	}
	return result
}
