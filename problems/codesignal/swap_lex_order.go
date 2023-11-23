package codesignal

import (
	"fmt"
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
*/

func swapLexOrderv2(str string, pairs [][]int) string {
	edges := make([][]int, len(str))
	for _, pair := range pairs {
		v1, v2 := pair[0]-1, pair[1]-1
		edges[v1] = append(edges[v1], v2)
		edges[v2] = append(edges[v2], v1)
	}

	fmt.Println(edges)
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