package codesignal

import (
	"strings"
)

/*

- Description

Given a string str and array of pairs that indicates which indices in the string can be swapped,
return the lexicographically largest string that results from doing the allowed swaps.

You can swap indices any number of times.

- Example

For str = "abdc" and pairs = [[1, 4], [3, 4]], the output should be
solution(str, pairs) = "dbca".

By swapping the given indices, you get the strings: "cbda", "cbad", "dbac", "dbca".
The lexicographically largest string in this list is "dbca".

- Input/Output

[execution time limit] 4 seconds (go)

[memory limit] 1 GB

[input] string str

A string consisting only of lowercase English letters.

Guaranteed constraints:
1 ≤ str.length ≤ 104.

[input] array.array.integer pairs

An array containing pairs of indices that can be swapped in str (1-based). This means that for each pairs[i], you can swap elements in str that have the indices pairs[i][0] and pairs[i][1].

Guaranteed constraints:
0 ≤ pairs.length ≤ 5000,
pairs[i].length = 2.

[output] string

HINT:
1. find the edges/connections for swapping on each chars
2. build the sets of counts on each chars by 1
3. make the result from 2
*/

func swapLexOrderv48(str string, pairs [][]int) string {
	conns := make([][]int, len(str))
	for _, pair := range pairs {
		v1, v2 := pair[0]-1, pair[1]-1
		conns[v1] = append(conns[v1], v2)
		conns[v2] = append(conns[v2], v1)
	}

	sets := make([]map[rune]int, len(str))
	var buildSets func(v int, curr map[rune]int)
	buildSets = func(v int, curr map[rune]int) {
		if len(sets[v]) > 0 {
			return
		}

		curr[rune(str[v])]++
		sets[v] = curr
		for _, u := range conns[v] {
			buildSets(u, curr)
		}
	}

	for i := range str {
		if len(sets[i]) == 0 {
			buildSets(i, map[rune]int{})
		}
	}

	result := ""
	for i := range str {
		for j := 'z'; j >= 'a'; j-- {
			if sets[i][j] > 0 {
				result += string(j)
				sets[i][j]--
				break
			}
		}
	}
	return result
}

func swapLexOrderv47(str string, pairs [][]int) string {
	conns := make([][]int, len(str))
	for _, pair := range pairs {
		v1, v2 := pair[0]-1, pair[1]-1
		conns[v1] = append(conns[v1], v2)
		conns[v2] = append(conns[v2], v1)
	}

	sets := make([]map[rune]int, len(str))
	var buildSets func(v int, curr map[rune]int)
	buildSets = func(v int, curr map[rune]int) {
		if len(sets[v]) > 0 {
			return
		}
		curr[rune(str[v])]++
		sets[v] = curr
		for _, u := range conns[v] {
			buildSets(u, curr)
		}
	}

	for i := range str {
		if len(sets[i]) == 0 {
			buildSets(i, map[rune]int{})
		}
	}

	result := ""
	for i := range str {
		for j := 'z'; j >= 'a'; j-- {
			if sets[i][j] > 0 {
				result += string(j)
				sets[i][j]--
				break
			}
		}
	}
	return result
}

func swapLexOrderv46(str string, pairs [][]int) string {
	conns := make([][]int, len(str))
	for _, pair := range pairs {
		v1, v2 := pair[0]-1, pair[1]-1
		conns[v1] = append(conns[v1], v2)
		conns[v2] = append(conns[v2], v1)
	}

	sets := make([]map[rune]int, len(str))
	var buildSets func(v int, curr map[rune]int)
	buildSets = func(v int, curr map[rune]int) {
		if len(sets[v]) > 0 {
			return
		}
		curr[rune(str[v])]++
		sets[v] = curr
		for _, u := range conns[v] {
			buildSets(u, curr)
		}
	}

	for i := range str {
		if len(sets[i]) == 0 {
			buildSets(i, map[rune]int{})
		}
	}

	result := ""

	for i := range str {
		for j := 'z'; j >= 'a'; j-- {
			if sets[i][j] > 0 {
				result += string(j)
				sets[i][j]--
				break
			}
		}
	}

	return result
}

func swapLexOrderv45(str string, pairs [][]int) string {
	conns := make([][]int, len(str))
	for _, pair := range pairs {
		v1, v2 := pair[0]-1, pair[1]-1
		conns[v1] = append(conns[v1], v2)
		conns[v2] = append(conns[v2], v1)
	}

	sets := make([]map[rune]int, len(str))
	var buildSets func(v int, curr map[rune]int)
	buildSets = func(v int, curr map[rune]int) {
		if len(sets[v]) > 0 {
			return
		}

		curr[rune(str[v])]++
		sets[v] = curr
		for _, u := range conns[v] {
			buildSets(u, curr)
		}
	}

	for i := range str {
		if len(sets[i]) == 0 {
			buildSets(i, map[rune]int{})
		}
	}

	result := ""
	for i := range str {
		for j := 'z'; j >= 'a'; j-- {
			if sets[i][j] > 0 {
				result += string(j)
				sets[i][j]--
				break
			}
		}
	}
	return result
}

func swapLexOrderv44(str string, pairs [][]int) string {
	conns := make([][]int, len(str))
	for _, pair := range pairs {
		v1, v2 := pair[0]-1, pair[1]-1
		conns[v1] = append(conns[v1], v2)
		conns[v2] = append(conns[v2], v1)
	}
	sets := make([]map[rune]int, len(str))
	var buildSets func(v int, curr map[rune]int)
	buildSets = func(v int, curr map[rune]int) {
		if len(sets[v]) > 0 {
			return
		}

		curr[rune(str[v])]++
		sets[v] = curr
		for _, u := range conns[v] {
			buildSets(u, curr)
		}
	}

	for i := range str {
		if len(sets[i]) == 0 {
			buildSets(i, map[rune]int{})
		}
	}

	result := ""

	for i := range str {
		for j := 'z'; j >= 'a'; j-- {
			if sets[i][j] > 0 {
				sets[i][j]--
				result += string(j)
				break
			}
		}
	}

	return result
}

func swapLexOrderv43(str string, pairs [][]int) string {
	conns := make([][]int, len(str))
	for _, pair := range pairs {
		v1, v2 := pair[0]-1, pair[1]-1
		conns[v1] = append(conns[v1], v2)
		conns[v2] = append(conns[v2], v1)
	}

	sets := make([]map[rune]int, len(str))
	var buildSets func(v int, curr map[rune]int)
	buildSets = func(v int, curr map[rune]int) {
		if len(sets[v]) > 0 {
			return
		}

		curr[rune(str[v])]++
		sets[v] = curr
		for _, u := range conns[v] {
			buildSets(u, curr)
		}
	}

	for i := range str {
		if len(sets[i]) == 0 {
			buildSets(i, map[rune]int{})
		}
	}

	result := ""

	for i := range str {
		for j := 'z'; j >= 'a'; j-- {
			if sets[i][j] > 0 {
				sets[i][j]--
				result += string(j)
				break
			}
		}
	}

	return result
}

func swapLexOrderv42(str string, pairs [][]int) string {
	conns := make([][]int, len(str))
	for _, pair := range pairs {
		v1, v2 := pair[0]-1, pair[1]-1
		conns[v1] = append(conns[v1], v2)
		conns[v2] = append(conns[v2], v1)
	}

	sets := make([]map[rune]int, len(str))
	var buildSets func(v int, curr map[rune]int)
	buildSets = func(v int, curr map[rune]int) {
		if len(sets[v]) > 0 {
			return
		}

		curr[rune(str[v])]++
		sets[v] = curr
		for _, u := range conns[v] {
			buildSets(u, curr)
		}
	}

	for i := range str {
		if len(sets[i]) == 0 {
			buildSets(i, map[rune]int{})
		}
	}

	result := ""

	for i := range str {
		for j := 'z'; j >= 'a'; j-- {
			if sets[i][j] > 0 {
				sets[i][j]--
				result += string(j)
				break
			}
		}
	}

	return result
}

func swapLexOrderv41(str string, pairs [][]int) string {
	conns := make([][]int, len(str))
	for _, pair := range pairs {
		v1, v2 := pair[0]-1, pair[1]-1
		conns[v1] = append(conns[v1], v2)
		conns[v2] = append(conns[v2], v1)
	}

	sets := make([]map[rune]int, len(str))
	var buildSets func(v int, curr map[rune]int)
	buildSets = func(v int, curr map[rune]int) {
		if len(sets[v]) > 0 {
			return
		}

		curr[rune(str[v])]++
		sets[v] = curr
		for _, u := range conns[v] {
			buildSets(u, curr)
		}
	}

	for i := range str {
		if len(sets[i]) == 0 {
			buildSets(i, map[rune]int{})
		}
	}

	result := ""

	for i := range str {
		for j := 'z'; j >= 'a'; j-- {
			if sets[i][j] > 0 {
				sets[i][j]--
				result += string(j)
				break
			}
		}
	}

	return result
}

func swapLexOrderv40(str string, pairs [][]int) string {
	conns := make([][]int, len(str))
	for _, pair := range pairs {
		v1, v2 := pair[0]-1, pair[1]-1
		conns[v1] = append(conns[v1], v2)
		conns[v2] = append(conns[v2], v1)
	}
	sets := make([]map[rune]int, len(str))
	var buildSets func(v int, curr map[rune]int)
	buildSets = func(v int, curr map[rune]int) {
		if len(sets[v]) > 0 {
			return
		}

		curr[rune(str[v])]++
		sets[v] = curr
		for _, u := range conns[v] {
			buildSets(u, curr)
		}
	}

	for i := range str {
		if len(sets[i]) == 0 {
			buildSets(i, map[rune]int{})
		}
	}

	result := ""

	for i := range str {
		for j := 'z'; j >= 'a'; j-- {
			if sets[i][j] > 0 {
				sets[i][j]--
				result += string(j)
				break
			}
		}
	}

	return result
}

func swapLexOrderv39(str string, pairs [][]int) string {
	conns := make([][]int, len(str))
	for _, pair := range pairs {
		v1, v2 := pair[0]-1, pair[1]-1
		conns[v1] = append(conns[v1], v2)
		conns[v2] = append(conns[v2], v1)
	}

	sets := make([]map[rune]int, len(str))
	var buildSets func(v int, curr map[rune]int)
	buildSets = func(v int, curr map[rune]int) {
		if len(sets[v]) > 0 {
			return
		}

		curr[rune(str[v])]++
		sets[v] = curr
		for _, u := range conns[v] {
			buildSets(u, curr)
		}
	}

	for i := range str {
		if len(sets[i]) == 0 {
			buildSets(i, map[rune]int{})
		}
	}

	result := ""
	for i := range str {
		for j := 'z'; j >= 'a'; j-- {
			if sets[i][j] > 0 {
				result += string(j)
				sets[i][j]--
				break
			}
		}
	}
	return result
}

func swapLexOrderv38(str string, paris [][]int) string {
	conns := make([][]int, len(str))
	for _, pair := range paris {
		v1, v2 := pair[0]-1, pair[1]-1
		conns[v1] = append(conns[v1], v2)
		conns[v2] = append(conns[v2], v1)
	}

	sets := make([]map[rune]int, len(str))
	var buildSet func(v int, curr map[rune]int)
	buildSet = func(v int, curr map[rune]int) {
		if len(sets[v]) > 0 {
			return
		}

		curr[rune(str[v])]++
		sets[v] = curr
		for _, u := range conns[v] {
			buildSet(u, curr)
		}
	}

	for i := range str {
		if len(sets[i]) == 0 {
			buildSet(i, map[rune]int{})
		}
	}
	result := ""
	for i := range str {
		for j := 'z'; j >= 'a'; j-- {
			if sets[i][j] > 0 {
				sets[i][j]--
				result += string(j)
				break
			}
		}
	}
	return result
}

func swapLexOrderv37(str string, pairs [][]int) string {
	conns := make([][]int, len(str))
	for _, pair := range pairs {
		v1, v2 := pair[0]-1, pair[1]-1
		conns[v1] = append(conns[v1], v2)
		conns[v2] = append(conns[v2], v1)
	}

	sets := make([]map[rune]int, len(str))
	var buildSets func(v int, curr map[rune]int)
	buildSets = func(v int, curr map[rune]int) {
		if len(sets[v]) > 0 {
			return
		}

		curr[rune(str[v])]++
		sets[v] = curr
		for _, u := range conns[v] {
			buildSets(u, curr)
		}
	}

	for i := range str {
		if len(sets[i]) == 0 {
			buildSets(i, map[rune]int{})
		}
	}

	result := ""

	for i := range str {
		for j := 'z'; j >= 'a'; j-- {
			if sets[i][j] > 0 {
				sets[i][j]--
				result += string(j)
				break
			}
		}
	}

	return result
}

func swapLexOrderv36(str string, pairs [][]int) string {
	conns := make([][]int, len(str))
	for _, pair := range pairs {
		v1, v2 := pair[0]-1, pair[1]-1
		conns[v1] = append(conns[v1], v2)
		conns[v2] = append(conns[v2], v1)
	}
	sets := make([]map[rune]int, len(str))
	var buildSets func(v int, curr map[rune]int)
	buildSets = func(v int, curr map[rune]int) {
		if len(sets[v]) > 0 {
			return
		}

		curr[rune(str[v])]++
		sets[v] = curr
		for _, u := range conns[v] {
			buildSets(u, curr)
		}
	}

	for i := range str {
		if len(sets[i]) == 0 {
			buildSets(i, map[rune]int{})
		}
	}

	result := ""
	for i := range str {
		for j := 'z'; j >= 'a'; j-- {
			if sets[i][j] > 0 {
				result += string(j)
				sets[i][j]--
				break
			}
		}
	}

	return result
}

func swapLexOrderv35(str string, pairs [][]int) string {
	conns := make([][]int, len(str))
	for _, pair := range pairs {
		v1, v2 := pair[0]-1, pair[1]-1
		conns[v1] = append(conns[v1], v2)
		conns[v2] = append(conns[v2], v1)
	}

	sets := make([]map[rune]int, len(str))
	var buildSets func(v int, curr map[rune]int)
	buildSets = func(v int, curr map[rune]int) {
		if len(sets[v]) > 0 {
			return
		}

		curr[rune(str[v])]++
		sets[v] = curr
		for _, u := range conns[v] {
			buildSets(u, curr)
		}
	}

	for i := range str {
		if len(sets[i]) == 0 {
			buildSets(i, map[rune]int{})
		}
	}

	result := ""
	for i := range str {
		for j := 'z'; j >= 'a'; j-- {
			if sets[i][j] > 0 {
				sets[i][j]--
				result += string(j)
				break
			}
		}
	}
	return result
}

func swapLexOrderv34(str string, pairs [][]int) string {
	conns := make([][]int, len(str))
	for _, pair := range pairs {
		v1, v2 := pair[0]-1, pair[1]-1
		conns[v1] = append(conns[v1], v2)
		conns[v2] = append(conns[v2], v1)
	}

	sets := make([]map[rune]int, len(str))
	var buildSets func(v int, curr map[rune]int)
	buildSets = func(v int, curr map[rune]int) {
		if len(sets[v]) > 0 {
			return
		}

		curr[rune(str[v])]++
		sets[v] = curr
		for _, u := range conns[v] {
			buildSets(u, curr)
		}
	}

	for i := range str {
		if len(sets[i]) == 0 {
			buildSets(i, map[rune]int{})
		}
	}

	result := ""

	for i := range str {
		for j := 'z'; j >= 'a'; j-- {
			if sets[i][j] > 0 {
				sets[i][j]--
				result += string(j)
				break
			}
		}
	}

	return result
}

func swapLexOrderv33(str string, pairs [][]int) string {
	conns := make([][]int, len(str))
	for _, pair := range pairs {
		v1, v2 := pair[0]-1, pair[1]-1
		conns[v1] = append(conns[v1], v2)
		conns[v2] = append(conns[v2], v1)
	}

	sets := make([]map[rune]int, len(str))
	var buildSets func(v int, curr map[rune]int)
	buildSets = func(v int, curr map[rune]int) {
		if len(sets[v]) > 0 {
			return
		}

		curr[rune(str[v])]++
		sets[v] = curr
		for _, u := range conns[v] {
			buildSets(u, curr)
		}
	}

	for i := range str {
		if len(sets[i]) == 0 {
			buildSets(i, map[rune]int{})
		}
	}

	result := ""
	for i := range str {
		for j := 'z'; j >= 'a'; j-- {
			if sets[i][j] > 0 {
				sets[i][j]--
				result += string(j)
				break
			}
		}
	}
	return result
}

func swapLexOrderv32(str string, pairs [][]int) string {
	conns := make([][]int, len(str))
	for _, pair := range pairs {
		v1, v2 := pair[0]-1, pair[1]-1
		conns[v1] = append(conns[v1], v2)
		conns[v2] = append(conns[v2], v1)
	}

	sets := make([]map[rune]int, len(str))
	var buildSets func(v int, curr map[rune]int)
	buildSets = func(v int, curr map[rune]int) {
		if len(sets[v]) > 0 {
			return
		}
		curr[rune(str[v])]++
		sets[v] = curr
		for _, u := range conns[v] {
			buildSets(u, curr)
		}
	}

	for i := range str {
		if len(sets[i]) == 0 {
			buildSets(i, map[rune]int{})
		}
	}

	result := ""
	for i := range str {
		for j := 'z'; j >= 'a'; j-- {
			if sets[i][j] > 0 {
				sets[i][j]--
				result += string(j)
				break
			}
		}
	}
	return result
}

func swapLexOrderv31(str string, pairs [][]int) string {
	conns := make([][]int, len(str))
	for _, pair := range pairs {
		v1, v2 := pair[0]-1, pair[1]-1
		conns[v1] = append(conns[v1], v2)
		conns[v2] = append(conns[v2], v1)
	}

	sets := make([]map[rune]int, len(str))
	var buildSets func(v int, curr map[rune]int)
	buildSets = func(v int, curr map[rune]int) {
		if len(sets[v]) > 0 {
			return
		}

		curr[rune(str[v])]++
		sets[v] = curr
		for _, u := range conns[v] {
			buildSets(u, curr)
		}
	}

	for i := range str {
		if len(sets[i]) == 0 {
			buildSets(i, map[rune]int{})
		}
	}

	result := ""

	for i := range str {
		for j := 'z'; j >= 'a'; j-- {
			if sets[i][j] > 0 {
				sets[i][j]--
				result += string(j)
				break
			}
		}
	}

	return result
}

func swapLexOrderv30(str string, pairs [][]int) string {
	conns := make([][]int, len(str))
	for _, pair := range pairs {
		v1, v2 := pair[0]-1, pair[1]-1
		conns[v1] = append(conns[v1], v2)
		conns[v2] = append(conns[v2], v1)
	}

	sets := make([]map[rune]int, len(str))
	var buildSets func(v int, curr map[rune]int)
	buildSets = func(v int, curr map[rune]int) {
		if len(sets[v]) > 0 {
			return
		}
		curr[rune(str[v])]++
		sets[v] = curr
		for _, u := range conns[v] {
			buildSets(u, curr)
		}
	}

	for i := range str {
		if len(sets[i]) == 0 {
			buildSets(i, map[rune]int{})
		}
	}

	result := ""

	for i := range str {
		for j := 'z'; j >= 'a'; j-- {
			if sets[i][j] > 0 {
				sets[i][j]--
				result += string(j)
				break
			}
		}
	}

	return result
}

func swapLexOrderv29(str string, pairs [][]int) string {
	conns := make([][]int, len(str))
	for _, pair := range pairs {
		v1, v2 := pair[0]-1, pair[1]-1
		conns[v1] = append(conns[v1], v2)
		conns[v2] = append(conns[v2], v1)
	}
	sets := make([]map[rune]int, len(str))
	var buildSets func(v int, curr map[rune]int)
	buildSets = func(v int, curr map[rune]int) {
		if len(sets[v]) > 0 {
			return
		}

		curr[rune(str[v])]++
		sets[v] = curr
		for _, u := range conns[v] {
			buildSets(u, curr)
		}
	}

	for i := range str {
		if len(sets[i]) == 0 {
			buildSets(i, map[rune]int{})
		}
	}

	result := ""
	for i := range str {
		for j := 'z'; j >= 'a'; j-- {
			if sets[i][j] > 0 {
				sets[i][j]--
				result += string(j)
				break
			}
		}
	}
	return result
}

func swapLexOrderv28(str string, pairs [][]int) string {
	conns := make([][]int, len(str))
	for _, pair := range pairs {
		v1, v2 := pair[0]-1, pair[1]-1
		conns[v1] = append(conns[v1], v2)
		conns[v2] = append(conns[v2], v1)
	}

	sets := make([]map[rune]int, len(str))
	var buildSets func(v int, curr map[rune]int)
	buildSets = func(v int, curr map[rune]int) {
		if len(sets[v]) > 0 {
			return
		}

		curr[rune(str[v])]++
		sets[v] = curr
		for _, u := range conns[v] {
			buildSets(u, curr)
		}
	}

	for i := range str {
		if len(sets[i]) == 0 {
			buildSets(i, map[rune]int{})
		}
	}

	result := ""
	for i := range str {
		for j := 'z'; j >= 0; j-- {
			if sets[i][j] > 0 {
				sets[i][j]--
				result += string(j)
				break
			}
		}
	}
	return result
}

func swapLexOrderv27(str string, pairs [][]int) string {
	conns := make([][]int, len(str))
	for _, pair := range pairs {
		v1, v2 := pair[0]-1, pair[1]-1
		conns[v1] = append(conns[v1], v2)
		conns[v2] = append(conns[v2], v1)
	}

	sets := make([]map[rune]int, len(str))
	var buildSets func(v int, curr map[rune]int)
	buildSets = func(v int, curr map[rune]int) {
		if len(sets[v]) > 0 {
			return
		}
		curr[rune(str[v])]++
		sets[v] = curr
		for _, u := range conns[v] {
			buildSets(u, curr)
		}
	}

	for i := range str {
		if len(sets[i]) == 0 {
			buildSets(i, map[rune]int{})
		}
	}

	result := ""
	for i := range str {
		for j := 'z'; j >= 'a'; j-- {
			if sets[i][j] > 0 {
				sets[i][j]--
				result += string(j)
				break
			}
		}
	}
	return result
}

func swapLexOrderv26(str string, pairs [][]int) string {
	conns := make([][]int, len(str))
	for _, pair := range pairs {
		v1, v2 := pair[0]-1, pair[1]-1
		conns[v1] = append(conns[v1], v2)
		conns[v2] = append(conns[v2], v1)
	}

	sets := make([]map[rune]int, len(str))
	var buildSets func(v int, curr map[rune]int)
	buildSets = func(v int, curr map[rune]int) {
		if len(sets[v]) > 0 {
			return
		}

		curr[rune(str[v])]++
		sets[v] = curr
		for _, u := range conns[v] {
			buildSets(u, curr)
		}
	}

	for i := range str {
		if len(sets[i]) == 0 {
			buildSets(i, map[rune]int{})
		}
	}

	result := ""
	for i := range str {
		for j := 'z'; j >= 'a'; j-- {
			if sets[i][j] > 0 {
				sets[i][j]--
				result += string(j)
				break
			}
		}
	}
	return result
}

func swapLexOrderv25(str string, pairs [][]int) string {
	conns := make([][]int, len(str))
	for _, pair := range pairs {
		v1, v2 := pair[0]-1, pair[1]-1
		conns[v1] = append(conns[v1], v2)
		conns[v2] = append(conns[v2], v1)
	}
	sets := make([]map[rune]int, len(str))
	var buildSets func(v int, curr map[rune]int)
	buildSets = func(v int, curr map[rune]int) {
		if len(sets[v]) > 0 {
			return
		}
		curr[rune(str[v])]++
		sets[v] = curr
		for _, u := range conns[v] {
			buildSets(u, curr)
		}
	}

	for i := range str {
		if len(sets[i]) == 0 {
			buildSets(i, map[rune]int{})
		}
	}
	result := ""
	for i := range str {
		for j := 'z'; j >= 'a'; j-- {
			if sets[i][j] > 0 {
				sets[i][j]--
				result += string(j)
				break
			}
		}
	}
	return result
}

func swapLexOrderv24(str string, pairs [][]int) string {
	conns := make([][]int, len(str))
	for _, pair := range pairs {
		v1, v2 := pair[0]-1, pair[1]-1
		conns[v1] = append(conns[v1], v2)
		conns[v2] = append(conns[v2], v1)
	}

	sets := make([]map[rune]int, len(str))
	var buildSets func(v int, curr map[rune]int)
	buildSets = func(v int, curr map[rune]int) {
		if len(sets[v]) > 0 {
			return
		}

		curr[rune(str[v])]++
		sets[v] = curr
		for _, u := range conns[v] {
			buildSets(u, curr)
		}
	}

	for i := range str {
		if len(sets[i]) == 0 {
			buildSets(i, map[rune]int{})
		}
	}

	result := ""
	for i := range str {
		for j := 'z'; j >= 'a'; j-- {
			if sets[i][j] > 0 {
				sets[i][j]--
				result += string(j)
				break
			}
		}
	}
	return result
}

func swapLexOrderv23(str string, pairs [][]int) string {
	conns := make([][]int, len(str))
	for _, pair := range pairs {
		v1, v2 := pair[0]-1, pair[1]-1
		conns[v1] = append(conns[v1], v2)
		conns[v2] = append(conns[v2], v1)
	}

	sets := make([]map[rune]int, len(str))
	var buildSets func(v int, curr map[rune]int)
	buildSets = func(v int, curr map[rune]int) {
		if len(sets[v]) > 0 {
			return
		}
		curr[rune(str[v])]++
		sets[v] = curr
		for _, u := range conns[v] {
			buildSets(u, curr)
		}
	}

	for i := range str {
		if len(sets[i]) == 0 {
			buildSets(i, map[rune]int{})
		}
	}

	result := ""
	for i := range str {
		for j := 'z'; j >= 'a'; j-- {
			if sets[i][j] > 0 {
				sets[i][j]--
				result += string(j)
				break
			}
		}
	}

	return result
}

func swapLexOrderv22(str string, paris [][]int) string {
	conns := make([][]int, len(str))
	for _, pair := range paris {
		v1, v2 := pair[0]-1, pair[1]-1
		conns[v1] = append(conns[v1], v2)
		conns[v2] = append(conns[v2], v1)
	}

	sets := make([]map[rune]int, len(str))
	var buildSets func(v int, curr map[rune]int)
	buildSets = func(v int, curr map[rune]int) {
		if len(sets[v]) > 0 {
			return
		}

		curr[rune(str[v])]++
		sets[v] = curr
		for _, u := range conns[v] {
			buildSets(u, curr)
		}
	}
	for i := range str {
		if len(sets[i]) == 0 {
			buildSets(i, map[rune]int{})
		}
	}

	result := ""

	for i := range str {
		for j := 'z'; j >= 'a'; j-- {
			if sets[i][j] > 0 {
				result += string(j)
				sets[i][j]--
				break
			}
		}
	}

	return result
}

func swapLexOrderv21(str string, pairs [][]int) string {
	conns := make([][]int, len(str))
	for _, pair := range pairs {
		v1, v2 := pair[0]-1, pair[1]-1
		conns[v1] = append(conns[v1], v2)
		conns[v2] = append(conns[v2], v1)
	}
	sets := make([]map[rune]int, len(str))
	var buildSets func(v int, curr map[rune]int)
	buildSets = func(v int, curr map[rune]int) {
		if len(sets[v]) > 0 {
			return
		}
		curr[rune(str[v])]++
		sets[v] = curr
		for _, u := range conns[v] {
			buildSets(u, curr)
		}
	}

	for i := range str {
		if len(sets[i]) == 0 {
			buildSets(i, map[rune]int{})
		}
	}

	result := ""

	for i := range str {
		for j := 'z'; j >= 'a'; j-- {
			if sets[i][j] > 0 {
				result += string(j)
				sets[i][j]--
				break
			}
		}
	}

	return result
}

func swapLexOrderv20(str string, pairs [][]int) string {
	conns := make([][]int, len(str))
	for _, pair := range pairs {
		v1, v2 := pair[0]-1, pair[1]-1
		conns[v1] = append(conns[v1], v2)
		conns[v2] = append(conns[v2], v1)
	}

	sets := make([]map[rune]int, len(str))
	var buildSets func(v int, curr map[rune]int)
	buildSets = func(v int, curr map[rune]int) {
		if len(sets[v]) > 0 {
			return
		}

		curr[rune(str[v])]++
		sets[v] = curr
		for _, u := range conns[v] {
			buildSets(u, curr)
		}
	}

	for i := range str {
		if len(sets[i]) == 0 {
			buildSets(i, map[rune]int{})
		}
	}

	result := ""

	for i := range str {
		for j := 'z'; j >= 'a'; j-- {
			if sets[i][j] > 0 {
				sets[i][j]--
				result += string(j)
				break
			}
		}
	}
	return result
}

func swapLexOrderv19(str string, pairs [][]int) string {
	conns := make([][]int, len(str))
	for _, pair := range pairs {
		v1, v2 := pair[0]-1, pair[1]-1
		conns[v1] = append(conns[v1], v2)
		conns[v2] = append(conns[v2], v1)
	}

	sets := make([]map[rune]int, len(str))
	var buildSets func(v int, curr map[rune]int)
	buildSets = func(v int, curr map[rune]int) {
		if len(sets[v]) > 0 {
			return
		}

		curr[rune(str[v])]++
		sets[v] = curr
		for _, u := range conns[v] {
			buildSets(u, curr)
		}
	}

	for i := range str {
		if len(sets[i]) == 0 {
			buildSets(i, map[rune]int{})
		}
	}

	result := ""

	for i := range str {
		for j := 'z'; j >= 'a'; j-- {
			if sets[i][j] > 0 {
				sets[i][j]--
				result += string(j)
				break
			}
		}
	}

	return result
}

func swapLexOrderv18(str string, pairs [][]int) string {
	conns := make([][]int, len(str))
	for _, pair := range pairs {
		v1, v2 := pair[0]-1, pair[1]-1
		conns[v1] = append(conns[v1], v2)
		conns[v2] = append(conns[v2], v1)
	}

	sets := make([]map[rune]int, len(str))
	var buildSets func(v int, curr map[rune]int)
	buildSets = func(v int, curr map[rune]int) {
		if len(sets[v]) > 0 {
			return
		}
		curr[rune(str[v])]++
		sets[v] = curr

		for _, u := range conns[v] {
			buildSets(u, curr)
		}
	}

	for i := range str {
		if len(sets[i]) == 0 {
			buildSets(i, map[rune]int{})
		}
	}

	result := ""

	for i := range str {
		for j := 'z'; j >= 'a'; j-- {
			if sets[i][j] > 0 {
				sets[i][j]--
				result += string(j)
				break
			}
		}
	}

	return result
}

func swapLexOrderv17(str string, pairs [][]int) string {
	conns := make([][]int, len(str))
	for _, pair := range pairs {
		v1, v2 := pair[0]-1, pair[1]-1
		conns[v1] = append(conns[v1], v2)
		conns[v2] = append(conns[v2], v1)
	}

	sets := make([]map[rune]int, len(str))
	var buildSets func(v int, curr map[rune]int)
	buildSets = func(v int, curr map[rune]int) {
		if len(sets[v]) > 0 {
			return
		}

		curr[rune(str[v])]++
		sets[v] = curr
		for _, u := range conns[v] {
			buildSets(u, curr)
		}
	}

	for i := range str {
		if len(sets[i]) == 0 {
			buildSets(i, map[rune]int{})
		}
	}

	result := ""

	for i := range str {
		for j := 'z'; j >= 'a'; j-- {
			if sets[i][j] > 0 {
				sets[i][j]--
				result += string(j)
				break
			}
		}
	}

	return result
}

func swapLexOrderv16(str string, pairs [][]int) string {
	conns := make([][]int, len(str))
	for _, pair := range pairs {
		v1, v2 := pair[0]-1, pair[1]-1
		conns[v1] = append(conns[v1], v2)
		conns[v2] = append(conns[v2], v1)
	}

	sets := make([]map[rune]int, len(str))
	var buildSets func(v int, curr map[rune]int)
	buildSets = func(v int, curr map[rune]int) {
		if len(sets[v]) > 0 {
			return
		}

		curr[rune(str[v])]++
		sets[v] = curr

		for _, u := range conns[v] {
			buildSets(u, curr)
		}
	}

	for i := range str {
		if len(sets[i]) == 0 {
			buildSets(i, map[rune]int{})
		}
	}

	result := ""
	for i := range str {
		for j := 'z'; j >= 'a'; j-- {
			if sets[i][j] > 0 {
				result += string(j)
				sets[i][j]--
				break
			}
		}
	}

	return result
}

func swapLexOrderv15(str string, pairs [][]int) string {
	conns := make([][]int, len(str))
	for _, pair := range pairs {
		v1, v2 := pair[0]-1, pair[1]-1
		conns[v1] = append(conns[v1], v2)
		conns[v2] = append(conns[v2], v1)
	}

	sets := make([]map[rune]int, len(str))
	var buildSets func(v int, curr map[rune]int)
	buildSets = func(v int, curr map[rune]int) {
		if len(sets[v]) > 0 {
			return
		}
		curr[rune(str[v])]++
		sets[v] = curr
		for _, u := range conns[v] {
			buildSets(u, curr)
		}
	}

	for i := range str {
		if len(sets[i]) == 0 {
			buildSets(i, map[rune]int{})
		}
	}

	result := ""
	for i := range str {
		for j := 'z'; j >= 'a'; j-- {
			if sets[i][j] > 0 {
				sets[i][j]--
				result += string(j)
				break
			}
		}
	}

	return result
}

func swapLexOrderv14(str string, pairs [][]int) string {
	edges := make([][]int, len(str))
	for _, pair := range pairs {
		v1, v2 := pair[0]-1, pair[1]-1
		edges[v1] = append(edges[v1], v2)
		edges[v2] = append(edges[v2], v1)
	}

	sets := make([]map[rune]int, len(str))
	var buildSets func(v int, curr map[rune]int)
	buildSets = func(v int, curr map[rune]int) {
		if len(sets[v]) > 0 {
			return
		}

		curr[rune(str[v])]++
		sets[v] = curr
		for _, u := range edges[v] {
			buildSets(u, curr)
		}
	}

	for i := range str {
		if len(sets[i]) == 0 {
			buildSets(i, map[rune]int{})
		}
	}

	result := ""

	for i := range str {
		for j := 'z'; j >= 'a'; j-- {
			if sets[i][j] > 0 {
				sets[i][j]--
				result += string(j)
				break
			}
		}
	}

	return result
}

func swapLexOrderv13(str string, pairs [][]int) string {
	edges := make([][]int, len(str))
	for _, pair := range pairs {
		v1, v2 := pair[0]-1, pair[1]-1
		edges[v1] = append(edges[v1], v2)
		edges[v2] = append(edges[v2], v1)
	}

	sets := make([]map[rune]int, len(str))
	var buildSets func(v int, curr map[rune]int)
	buildSets = func(v int, curr map[rune]int) {
		if len(sets[v]) > 0 {
			return
		}

		curr[rune(str[v])]++
		sets[v] = curr
		for _, u := range edges[v] {
			buildSets(u, curr)
		}
	}

	for i := range str {
		if len(sets[i]) == 0 {
			buildSets(i, map[rune]int{})
		}
	}

	result := ""

	for i := range str {
		for j := 'z'; j >= 'a'; j-- {
			if sets[i][j] > 0 {
				sets[i][j]--
				result += string(j)
				break
			}
		}
	}

	return result
}

func swapLexOrderv12(str string, pairs [][]int) string {
	edges := make([][]int, len(str))
	for _, pair := range pairs {
		v1, v2 := pair[0]-1, pair[1]-1
		edges[v1] = append(edges[v1], v2)
		edges[v2] = append(edges[v2], v1)
	}

	sets := make([]map[rune]int, len(str))
	var buildSets func(v int, curr map[rune]int)
	buildSets = func(v int, curr map[rune]int) {
		if len(sets[v]) > 0 {
			return
		}

		curr[rune(str[v])]++
		sets[v] = curr
		for _, u := range edges[v] {
			buildSets(u, curr)
		}
	}

	for i := range str {
		if len(sets[i]) == 0 {
			buildSets(i, map[rune]int{})
		}
	}

	result := ""
	for i := range str {
		for j := 'z'; j >= 'a'; j-- {
			if sets[i][j] > 0 {
				result += string(j)
				sets[i][j]--
				break
			}
		}
	}
	return result
}

func swapLexOrderv11(str string, pairs [][]int) string {
	edges := make([][]int, len(str))
	for _, pair := range pairs {
		v1, v2 := pair[0]-1, pair[1]-1
		edges[v1] = append(edges[v1], v2)
		edges[v2] = append(edges[v2], v1)
	}

	sets := make([]map[rune]int, len(str))
	var buildSets func(v int, curr map[rune]int)
	buildSets = func(v int, curr map[rune]int) {
		if len(sets[v]) > 0 {
			return
		}

		curr[rune(str[v])]++
		sets[v] = curr
		for _, u := range edges[v] {
			buildSets(u, curr)
		}
	}

	for i := range str {
		if len(sets[i]) == 0 {
			buildSets(i, map[rune]int{})
		}
	}

	result := ""

	for i := range str {
		for j := 'z'; j >= 'a'; j-- {
			if sets[i][j] > 0 {
				result += string(j)
				sets[i][j]--
				break
			}
		}
	}

	return result
}

func swapLexOrderv10(str string, pairs [][]int) string {
	edges := make([][]int, len(str))
	for _, pair := range pairs {
		v1, v2 := pair[0]-1, pair[1]-1
		edges[v1] = append(edges[v1], v2)
		edges[v2] = append(edges[v2], v1)
	}

	sets := make([]map[rune]int, len(str))
	var buildSets func(v int, curr map[rune]int)
	buildSets = func(v int, curr map[rune]int) {
		if len(sets[v]) > 0 {
			return
		}

		curr[rune(str[v])]++
		sets[v] = curr

		for _, u := range edges[v] {
			buildSets(u, curr)
		}
	}

	for i := range str {
		if len(sets[i]) == 0 {
			buildSets(i, map[rune]int{})
		}
	}

	result := ""

	for i := range str {
		for j := 'z'; j >= 'a'; j-- {
			if sets[i][j] > 0 {
				result += string(j)
				sets[i][j]--
				break
			}
		}
	}

	return result
}

func swapLexOrderv9(str string, pairs [][]int) string {
	edges := make([][]int, len(str))
	for _, pair := range pairs {
		v1, v2 := pair[0]-1, pair[1]-1
		edges[v1] = append(edges[v1], v2)
		edges[v2] = append(edges[v2], v1)
	}

	sets := make([]map[rune]int, len(str))
	var buildSets func(v int, curr map[rune]int)
	buildSets = func(v int, curr map[rune]int) {
		if len(sets[v]) > 0 {
			return
		}
		curr[rune(str[v])]++
		sets[v] = curr
		for _, u := range edges[v] {
			buildSets(u, curr)
		}
	}

	for i := range str {
		if len(sets[i]) == 0 {
			buildSets(i, map[rune]int{})
		}
	}

	result := ""
	for i := range str {
		for j := 'z'; j >= 'a'; j-- {
			if sets[i][j] > 0 {
				sets[i][j]--
				result += string(j)
				break
			}
		}
	}

	return result
}

func swapLexOrderv8(str string, pairs [][]int) string {
	edges := make([][]int, len(str))
	for _, pair := range pairs {
		v1, v2 := pair[0]-1, pair[1]-1
		edges[v1] = append(edges[v1], v2)
		edges[v2] = append(edges[v2], v1)
	}

	sets := make([]map[rune]int, len(str))
	var buildSets func(v int, curr map[rune]int)
	buildSets = func(v int, curr map[rune]int) {
		if len(sets[v]) > 0 {
			return
		}

		curr[rune(str[v])]++
		sets[v] = curr
		for _, u := range edges[v] {
			buildSets(u, curr)
		}
	}

	for i := range str {
		if len(sets[i]) == 0 {
			buildSets(i, map[rune]int{})
		}
	}

	result := ""
	for i := range str {
		for j := 'z'; j >= 'a'; j-- {
			if sets[i][j] > 0 {
				sets[i][j]--
				result += string(j)
				break
			}
		}
	}

	return result
}

func swapLexOrderv7(str string, pairs [][]int) string {
	edges := make([][]int, len(str))
	for _, pair := range pairs {
		v1, v2 := pair[0]-1, pair[1]-1
		edges[v1] = append(edges[v1], v2)
		edges[v2] = append(edges[v2], v1)
	}

	sets := make([]map[rune]int, len(str))
	var buildSets func(v int, curr map[rune]int)
	buildSets = func(v int, curr map[rune]int) {
		if len(sets[v]) > 0 {
			return
		}

		curr[rune(str[v])]++
		sets[v] = curr

		for _, u := range edges[v] {
			buildSets(u, curr)
		}
	}

	for i := range str {
		if len(sets[i]) == 0 {
			buildSets(i, map[rune]int{})
		}
	}

	result := ""

	for i := range str {
		for j := 'z'; j >= 'a'; j-- {
			if sets[i][j] > 0 {
				result += string(j)
				sets[i][j]--
				break
			}
		}
	}

	return result
}

func swapLexOrderv6(str string, pairs [][]int) string {
	edges := make([][]int, len(str))
	for _, pair := range pairs {
		v1, v2 := pair[0]-1, pair[1]-1
		edges[v1] = append(edges[v1], v2)
		edges[v2] = append(edges[v2], v1)
	}

	sets := make([]map[rune]int, len(str))
	var buildSets func(v int, curr map[rune]int)
	buildSets = func(v int, curr map[rune]int) {
		if len(sets[v]) > 0 {
			return
		}

		sets[v] = curr
		curr[rune(str[v])]++
		for _, u := range edges[v] {
			buildSets(u, curr)
		}
	}

	for i := range str {
		if len(sets[i]) == 0 {
			buildSets(i, map[rune]int{})
		}
	}

	result := ""

	for i := range str {
		for j := 'z'; j >= 'a'; j-- {
			if sets[i][j] > 0 {
				sets[i][j]--
				result += string(j)
				break
			}
		}
	}
	return result
}

func swapLexOrderv5(str string, pairs [][]int) string {
	edges := make([][]int, len(str))
	for _, pair := range pairs {
		v1, v2 := pair[0]-1, pair[1]-1
		edges[v1] = append(edges[v1], v2)
		edges[v2] = append(edges[v2], v1)
	}

	sets := make([]map[rune]int, len(str))
	var buildSets func(v int, curr map[rune]int)
	buildSets = func(v int, curr map[rune]int) {
		if len(sets[v]) != 0 {
			return
		}

		sets[v] = curr
		curr[rune(str[v])]++
		for _, u := range edges[v] {
			buildSets(u, curr)
		}
	}

	for i := range str {
		if len(sets[i]) == 0 {
			buildSets(i, map[rune]int{})
		}
	}

	result := ""

	for i := range str {
		for j := 'z'; j >= 'a'; j-- {
			if sets[i][j] > 0 {
				result += string(j)
				sets[i][j]--
				break
			}
		}
	}

	return result
}

func swapLexOrderv4(str string, pairs [][]int) string {
	edges := make([][]int, len(str))
	for _, pair := range pairs {
		v1, v2 := pair[0]-1, pair[1]-1
		edges[v1] = append(edges[v1], v2)
		edges[v2] = append(edges[v2], v1)
	}
	sets := make([]map[rune]int, len(str))
	var build func(i int, curr map[rune]int)
	build = func(i int, curr map[rune]int) {
		// base case: if the sets[i] exists then skip it
		if len(sets[i]) > 0 {
			return
		}

		sets[i] = curr
		curr[rune(str[i])]++
		for _, u := range edges[i] {
			build(u, curr)
		}
	}

	for i := range str {
		if len(sets[i]) == 0 {
			build(i, map[rune]int{})
		}
	}

	result := ""

	for i := range str {
		for j := 'z'; j >= 'a'; j-- {
			if sets[i][j] > 0 {
				result += string(j)
				sets[i][j]--
				break
			}
		}
	}

	return result
}

func swapLexOrderv3(str string, pairs [][]int) string {
	edges := make([][]int, len(str))
	for _, pair := range pairs {
		v1, v2 := pair[0]-1, pair[1]-1
		edges[v1] = append(edges[v1], v2)
		edges[v2] = append(edges[v2], v1)
	}

	sets := make([]map[rune]int, len(str))
	var buildSets func(i int, curr map[rune]int)
	buildSets = func(v int, curr map[rune]int) {
		if len(sets[v]) > 0 {
			return
		}
		sets[v] = curr
		curr[rune(str[v])]++
		for _, u := range edges[v] {
			buildSets(u, curr)
		}
	}
	for i := range str {
		if len(sets[i]) == 0 {
			buildSets(i, map[rune]int{})
		}
	}
	result := ""
	for i := range str {
		set := sets[i]
		for c := 'z'; c >= 'a'; c-- {
			if set[rune(c)] > 0 {
				result += string(c)
				set[rune(c)]--
				break
			}
		}
	}

	return result
}

func swapLexOrderv2(str string, pairs [][]int) string {
	edges := make([][]int, len(str))
	for _, pair := range pairs {
		v1, v2 := pair[0]-1, pair[1]-1
		edges[v1] = append(edges[v1], v2)
		edges[v2] = append(edges[v2], v1)
	}

	var buildSets func(s string, v int, tmp map[byte]int, ss []map[byte]int, e [][]int)
	buildSets = func(s string, v int, tmp map[byte]int, ss []map[byte]int, e [][]int) {
		if len(ss[v]) > 0 {
			return
		}
		ss[v] = tmp
		tmp[s[v]]++
		for _, u := range e[v] {
			buildSets(s, u, tmp, ss, e)
		}
	}

	sets := make([]map[byte]int, len(str))
	for i := range str {
		if len(sets[i]) == 0 {
			buildSets(str, i, map[byte]int{}, sets, edges)
		}
	}

	result := ""
	for i := range str {
		set := sets[i]
		for j := 'z'; j >= 'a'; j-- {
			if set[byte(j)] > 0 {
				result += string(j)
				set[byte(j)]--
				break
			}
		}
	}

	return result
}

func swapLexOrderv1(str string, pairs [][]int) string {
	graph := make([][]int, len(str))
	for _, pair := range pairs {
		graph[pair[0]-1] = append(graph[pair[0]-1], pair[1]-1)
		graph[pair[1]-1] = append(graph[pair[1]-1], pair[0]-1)
	}
	sets := make([]map[rune]int, len(str))
	for i := range str {
		if sets[i] == nil {
			dfs(str, graph, i, map[rune]int{}, sets)
		}
	}
	var builder strings.Builder
	builder.Grow(len(str))
	for i := range str {
		set := sets[i]
		// save the largest char by scanning the set from 'z' -> 'a'
		for c := 'z'; c >= 'a'; c-- {
			if set[c] > 0 {
				set[c]--
				builder.WriteRune(c)
				break
			}
		}
	}
	return builder.String()
}

func dfs(
	str string,
	graph [][]int,
	v int,
	set map[rune]int,
	sets []map[rune]int,
) {
	if sets[v] != nil {
		return
	}
	sets[v] = set
	set[rune(str[v])]++
	for _, u := range graph[v] {
		dfs(str, graph, u, set, sets)
	}
}
