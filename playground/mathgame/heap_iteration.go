package mathgame

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
