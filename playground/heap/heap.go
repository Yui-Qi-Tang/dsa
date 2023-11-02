package heap

func parent(i int) int { return i / 2 }
func left(i int) int   { return (2 * i) + 1 }
func right(i int) int  { return (2 * i) + 2 }

type ints []int

func (i *ints) len() int {
	return len(*i)
}

func (i *ints) push(values ...int) {
	for _, v := range values {
		*i = append(*i, v)
		i.up(i.len() - 1)
	}
}

func (i ints) up(idx int) {
	for idx > 0 {
		p := parent(idx)
		if !less(i[p], i[idx]) {
			return
		}
		i[p], i[idx] = i[idx], i[p]
		idx = p
	}
}

func (i *ints) pop() int {
	old := *i
	old[0], old[old.len()-1] = old[old.len()-1], old[0]
	v := old[old.len()-1]
	*i = old[:old.len()-1]
	i.down()
	return v
}

func (i *ints) down() {
	idx := 0
	for {
		pointer := idx
		l := left(idx)
		r := right(idx)

		if l < i.len() && less((*i)[pointer], (*i)[l]) {
			pointer = l
		}

		if r < i.len() && less((*i)[pointer], (*i)[r]) {
			pointer = r
		}

		if pointer == idx {
			return
		}

		(*i)[idx], (*i)[pointer] = (*i)[pointer], (*i)[idx]
		idx = pointer
	}
}

func less(i, j int) bool {
	return i > j
}

func push(arr *[]int, v ...int) {
	for _, x := range v {
		*arr = append(*arr, x)
		up(*arr, len(*arr)-1)
	}
}

// up moves the value up with min-heapify condition
func up(arr []int, idx int) {
	for idx > 0 {
		p := parent(idx)
		if !less(arr[p], arr[idx]) {
			return
		}
		arr[p], arr[idx] = arr[idx], arr[p]
		idx = p
	}
}

func pop(arr *[]int) int {
	if len(*arr) == 0 {
		panic("empty")
	}

	old := *arr
	old[0], old[len(old)-1] = old[len(old)-1], old[0]
	v := old[len(old)-1]
	*arr = old[:len(old)-1]
	// why down if pop?
	down(*arr, 0)
	//up(*arr, len(*arr)-1)

	return v
}

func down(arr []int, idx int) {
	for {
		curr := idx
		l := left(curr)
		r := right(curr)

		// handle min heapify (root, left child)
		if l < len(arr) && less(arr[curr], arr[l]) {
			curr = l
		}

		// handle min heapify (root, right child)
		if r < len(arr) && less(arr[curr], arr[r]) {
			curr = r
		}

		// return if arr[curr] statisfies the heapify
		if curr == idx {
			return
		}

		arr[curr], arr[idx] = arr[idx], arr[curr]
		idx = curr
	}
}
