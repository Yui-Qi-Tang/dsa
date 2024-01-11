package b75

/*
1282. Group the People Given the Group Size They Belong To

There are n people that are split into some unknown number of groups.
Each person is labeled with a unique ID from 0 to n - 1.

You are given an integer array groupSizes, where groupSizes[i] is the size of the group that person i is in. For example, if groupSizes[1] = 3, then person 1 must be in a group of size 3.

Return a list of groups such that each person i is in a group of size groupSizes[i].

Each person should appear in exactly one group, and every person must be in a group. If there are multiple answers, return any of them. It is guaranteed that there will be at least one valid solution for the given input.



Example 1:

Input: groupSizes = [3,3,3,3,3,1,3]
Output: [[5],[0,1,2],[3,4,6]]
Explanation:
The first group is [5]. The size is 1, and groupSizes[5] = 1.
The second group is [0,1,2]. The size is 3, and groupSizes[0] = groupSizes[1] = groupSizes[2] = 3.
The third group is [3,4,6]. The size is 3, and groupSizes[3] = groupSizes[4] = groupSizes[6] = 3.
Other possible solutions are [[2,1,6],[5],[0,4,3]] and [[5],[0,6,2],[4,3,1]].
Example 2:

Input: groupSizes = [2,1,3,3,3,2]
Output: [[1],[0,5],[2,3,4]]


Constraints:

groupSizes.length == n
1 <= n <= 500
1 <= groupSizes[i] <= n

HINT: I should learn some better way to solve this problem
*/

func groupthepeoplev12(groupSizes []int) [][]int {
	m := make(map[int][]int)

	for person, size := range groupSizes {
		m[size] = append(m[size], person)
	}

	result := make([][]int, 0)
	for size, persons := range m {
		p := 0
		for p < len(persons) {
			result = append(result, persons[p:p+size])
			p += size
		}
	}
	return result
}

func groupthepeoplev11(groupSizes []int) [][]int {
	m := make(map[int][]int) // size: people

	for p, s := range groupSizes {
		m[s] = append(m[s], p)
	}

	result := [][]int{}

	for s, pset := range m {
		pp := 0
		for pp < len(pset) {
			result = append(result, pset[pp:pp+s])
			pp += s
		}
	}

	return result
}

func groupthepeoplev10(groupSizes []int) [][]int {
	m := make(map[int][]int) // size: people

	for p, s := range groupSizes {
		m[s] = append(m[s], p)
	}

	result := [][]int{}

	for s, p := range m {
		pp := 0
		for pp < len(p) {
			result = append(result, p[pp:pp+s])
			pp += s
		}
	}

	return result
}

func groupthepeoplev9(groupSizes []int) [][]int {
	m := make(map[int][]int) // size: {p1, p2,...}

	for p, s := range groupSizes {
		m[s] = append(m[s], p)
	}

	result := [][]int{}
	for s, p := range m {
		pp := 0
		for pp < len(p) {
			result = append(result, p[pp:pp+s])
			pp += s
		}
	}
	return result
}

func groupthepeoplev8(groupSizes []int) [][]int {
	m := make(map[int][]int) // size:{people}

	for p, s := range groupSizes {
		m[s] = append(m[s], p)
	}

	result := [][]int{}

	for s, p := range m {
		pp := 0
		for pp < len(p) {
			result = append(result, p[pp:pp+s])
			pp += s
		}
	}

	return result
}

func groupthepeoplev7(groupSizes []int) [][]int {
	m := make(map[int][]int)

	for p, s := range groupSizes {
		m[s] = append(m[s], p)
	}

	result := make([][]int, 0)

	for s, p := range m {
		pp := 0
		for pp < len(p) {
			result = append(result, p[pp:pp+s])
			pp += s
		}
	}

	return result
}

func groupthepeoplev6(groupSizes []int) [][]int {
	m := make(map[int][]int, 0) // size: peoples

	for p, s := range groupSizes {
		m[s] = append(m[s], p)
	}

	result := make([][]int, 0)
	for s, p := range m {
		pp := 0 // start from people 0
		for pp < len(p) {
			result = append(result, p[pp:pp+s])
			pp += s
		}
	}

	return result
}

func groupthepeoplev5(groupSizes []int) [][]int {
	m := make(map[int][]int) // size: people set
	for p, s := range groupSizes {
		m[s] = append(m[s], p)
	}

	result := make([][]int, 0)
	for s, p := range m {
		pp := 0
		for pp < len(p) { // collect people if the there is reamining
			result = append(result, p[pp:pp+s])
			pp += s
		}
	}

	return result
}

func groupthepeoplev4(groupSizes []int) [][]int {

	// size: people
	m := make(map[int][]int)
	for p, s := range groupSizes {
		m[s] = append(m[s], p)
	}

	result := make([][]int, 0)

	for s, p := range m {
		pp := 0
		for pp < len(p) {
			result = append(result, p[pp:pp+s])
			pp += s
		}
	}

	return result
}

func groupthepeoplev3(groupSizes []int) [][]int {
	sizePeople := make(map[int][]int) // size: people1, people2,...
	for people, size := range groupSizes {
		sizePeople[size] = append(sizePeople[size], people)
	}
	// according to example 1, the size[3] = {0,1,2,3,4,6}, size[1] = {5}
	result := make([][]int, 0)
	for size, people := range sizePeople {
		pp := 0                // the people that is prepared to allocate to group; start from 0
		for pp < len(people) { // for each pp is less than the the size, collect them into the group
			result = append(result, people[pp:pp+size])
			pp += size // for split
		}

	}
	return result
}

func groupthepeoplev2(groupSizes []int) [][]int {
	n := len(groupSizes)
	if n == 0 {
		return nil
	}

	result := make([][]int, 0)
	p := make(map[int]bool)
	for i := 0; i < n; i++ {
		p[i] = false
	}

	for i := 0; i < n; i++ {
		if p[i] {
			continue
		}
		size := groupSizes[i]
		grouped := make([]int, 0)

		for j := i; j < n; j++ {
			if size == 0 {
				break
			}

			if groupSizes[i] == groupSizes[j] {
				grouped = append(grouped, j)
				p[j] = true
				size--
			}
		}

		if len(grouped) > 0 {
			result = append(result, grouped)
		}
	}

	return result
}

func groupthepeoplev1(groupSizes []int) [][]int {

	n := len(groupSizes)
	if n == 0 {
		return nil
	}

	result := make([][]int, 0)

	p := make(map[int]bool) // people: selected

	for i := range groupSizes {
		p[i] = false
	}

	for i := range groupSizes {
		if p[i] {
			continue
		}
		gs := groupSizes[i]
		grouped := make([]int, 0)
		for j := i; j < n; j++ {
			if gs == 0 {
				break
			}
			if groupSizes[i] == groupSizes[j] {
				grouped = append(grouped, j)
				p[j] = true
				gs--
			}
		}
		if len(grouped) > 0 {
			p[i] = true
			result = append(result, grouped)
		}

	}

	return result
}
