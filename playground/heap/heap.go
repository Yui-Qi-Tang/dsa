package heap

func parent(i int) int { return i / 2 }
func left(i int) int   { return (2 * i) + 1 }
func right(i int) int  { return (2 * i) + 2 }

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
