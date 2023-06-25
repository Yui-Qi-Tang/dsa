package b75

import (
	"fmt"
	"strconv"
	"strings"
)

/*
297. Serialize and Deserialize Binary Tree
Hard
8.8K
317
Companies
Serialization is the process of converting a data structure or object into a sequence of bits so that it can be stored in a file or memory buffer, or transmitted across a network connection link to be reconstructed later in the same or another computer environment.

Design an algorithm to serialize and deserialize a binary tree.
There is no restriction on how your serialization/deserialization algorithm should work.
You just need to ensure that a binary tree can be serialized to
a string and this string can be deserialized to the original tree structure.

serialized(Tree) to string
decserialized(string) to Tree

Clarification: The input/output format is the same as how LeetCode serializes a binary tree. You do not necessarily need to follow this format,
so please be creative and come up with different approaches yourself.



Example 1:


Input: root = [1,2,3,null,null,4,5]
Output: [1,2,3,null,null,4,5]
Example 2:

Input: root = []
Output: []


Constraints:

The number of nodes in the tree is in the range [0, 104].
-1000 <= Node.val <= 1000
*/

/**
 * Definition for a binary tree node.
 */
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func constructorv26() codecv26 {
	return codecv26{null: "n", join: ","}
}

type codecv26 struct {
	null, join string
}

func (c codecv26) serialize(root *TreeNode) string {
	if root == nil {
		return ""
	}

	result := []string{}
	queue := []*TreeNode{root}
	for len(queue) > 0 {
		n := queue[0]
		queue = queue[1:]
		if n == nil {
			result = append(result, c.null)
		} else {
			result = append(result, strconv.Itoa(n.Val))
			queue = append(queue, n.Left, n.Right)
		}
	}

	return strings.Join(result, c.join)
}

func (c codecv26) deserialize(data string) *TreeNode {
	if len(data) == 0 {
		return nil
	}

	in := strings.Split(data, c.join)
	toInt := func(v string) int {
		x, _ := strconv.Atoi(v)
		return x
	}
	root := &TreeNode{Val: toInt(in[0])}
	queue := []*TreeNode{root}

	for i := 1; i < len(in); {
		n := queue[0]
		queue = queue[1:]
		if in[i] != c.null {
			n.Left = &TreeNode{Val: toInt(in[i])}
			queue = append(queue, n.Left)
		}
		i++

		if i < len(in) && in[i] != c.null {
			n.Right = &TreeNode{Val: toInt(in[i])}
			queue = append(queue, n.Right)
		}
		i++
	}

	return root
}

func constructorv25() codecv25 {
	return codecv25{null: "n", join: ","}
}

type codecv25 struct {
	null, join string
}

func (c codecv25) serialize(root *TreeNode) string {
	if root == nil {
		return ""
	}
	result := []string{}
	queue := []*TreeNode{root}
	for len(queue) > 0 {
		n := queue[0]
		queue = queue[1:]
		if n == nil {
			result = append(result, c.null)
		} else {
			result = append(result, strconv.Itoa(n.Val))
			queue = append(queue, n.Left, n.Right)
		}
	}
	return strings.Join(result, c.join)
}

func (c codecv25) deserialize(data string) *TreeNode {
	if len(data) == 0 {
		return nil
	}

	in := strings.Split(data, c.join)
	toInt := func(x string) int {
		n, _ := strconv.Atoi(x)
		return n
	}
	root := &TreeNode{Val: toInt(in[0])}
	queue := []*TreeNode{root}
	for i := 1; i < len(in); {
		n := queue[0]
		queue = queue[1:]
		if in[i] != c.null {
			n.Left = &TreeNode{Val: toInt(in[i])}
			queue = append(queue, n.Left)
		}
		i++

		if i < len(in) && in[i] != c.null {
			n.Right = &TreeNode{Val: toInt(in[i])}
			queue = append(queue, n.Right)
		}
		i++
	}

	return root
}

func constructorv24() codecv24 {
	return codecv24{null: "n", join: ","}
}

type codecv24 struct {
	null, join string
}

func (c codecv24) serialize(root *TreeNode) string {
	if root == nil {
		return ""
	}

	result := []string{}
	queue := []*TreeNode{root}
	for len(queue) > 0 {
		n := queue[0]
		queue = queue[1:]
		if n == nil {
			result = append(result, c.null)
		} else {
			result = append(result, strconv.Itoa(n.Val))
			queue = append(queue, n.Left, n.Right)
		}
	}

	return strings.Join(result, c.join)
}

func (c codecv24) deserialize(data string) *TreeNode {
	if len(data) == 0 {
		return nil
	}

	in := strings.Split(data, c.join)
	toInt := func(x string) int {
		n, _ := strconv.Atoi(x)
		return n
	}

	root := &TreeNode{Val: toInt(in[0])}
	queue := []*TreeNode{root}
	for i := 1; i < len(in); {
		n := queue[0]
		fmt.Println(n.Val)
		queue = queue[1:]
		if in[i] != c.null {
			n.Left = &TreeNode{Val: toInt(in[i])}
			queue = append(queue, n.Left)
		}
		i++

		if i < len(in) && in[i] != c.null {
			n.Right = &TreeNode{Val: toInt(in[i])}
			queue = append(queue, n.Right)
		}
		i++
	}

	return root
}

func constructorv23() codecv23 {
	return codecv23{null: "n", join: ","}
}

type codecv23 struct {
	null, join string
}

func (c codecv23) serialize(root *TreeNode) string {
	if root == nil {
		return ""
	}

	result := make([]string, 0)
	queue := []*TreeNode{root}
	for len(queue) > 0 {
		n := queue[0]
		queue = queue[1:]
		if n == nil {
			result = append(result, c.null)
		} else {
			result = append(result, strconv.Itoa(n.Val))
			queue = append(queue, n.Left, n.Right)
		}
	}

	return strings.Join(result, c.join)
}

func (c codecv23) deserialize(data string) *TreeNode {
	if len(data) == 0 {
		return nil
	}

	in := strings.Split(data, c.join)
	toInt := func(x string) int {
		n, _ := strconv.Atoi(x)
		return n
	}
	root := &TreeNode{Val: toInt(in[0])}
	queue := []*TreeNode{root}
	for i := 1; i < len(in); {
		n := queue[0]
		queue = queue[1:]
		if in[i] != c.null {
			n.Left = &TreeNode{Val: toInt(in[i])}
			queue = append(queue, n.Left)
		}
		i++
		if i < len(in) && in[i] != c.null {
			n.Right = &TreeNode{Val: toInt(in[i])}
			queue = append(queue, n.Right)
		}
		i++
	}
	return root
}

func constructorv22() codecv22 {
	return codecv22{null: "n", join: ","}
}

type codecv22 struct {
	null, join string
}

func (c codecv22) serialize(root *TreeNode) string {
	if root == nil {
		return ""
	}

	result := []string{}
	queue := []*TreeNode{root}
	for len(queue) > 0 {
		n := queue[0]
		queue = queue[1:]
		if n == nil {
			result = append(result, c.null)
		} else {
			result = append(result, strconv.Itoa(n.Val))
			queue = append(queue, n.Left, n.Right)
		}
	}

	return strings.Join(result, c.join)
}

func (c codecv22) deserialize(data string) *TreeNode {
	if len(data) == 0 {
		return nil
	}

	in := strings.Split(data, c.join)
	toInt := func(s string) int {
		n, _ := strconv.Atoi(s)
		return n
	}
	root := &TreeNode{Val: toInt(in[0])}
	queue := []*TreeNode{root}
	for i := 1; i < len(in); {
		n := queue[0]
		queue = queue[1:]
		if in[i] != c.null {
			n.Left = &TreeNode{Val: toInt(in[i])}
			queue = append(queue, n.Left)
		}
		i++
		if i < len(in) && in[i] != c.null {
			n.Right = &TreeNode{Val: toInt(in[i])}
			queue = append(queue, n.Right)
		}
		i++
	}

	return root
}

func constructorv21() codecv21 {
	return codecv21{null: "n", split: ","}
}

type codecv21 struct {
	null, split string
}

func (c codecv21) serialize(root *TreeNode) string {
	if root == nil {
		return ""
	}
	result := []string{}
	queue := []*TreeNode{root}
	for len(queue) > 0 {
		n := queue[0]
		queue = queue[1:]
		if n == nil {
			result = append(result, c.null)
		} else {
			result = append(result, strconv.Itoa(n.Val))
			queue = append(queue, n.Left, n.Right)
		}
	}
	return strings.Join(result, c.split)
}

func (c codecv21) deserialize(data string) *TreeNode {
	if len(data) == 0 {
		return nil
	}

	in := strings.Split(data, c.split)
	toInt := func(x string) int {
		num, _ := strconv.Atoi(x)
		return num
	}
	root := &TreeNode{Val: toInt(in[0])}
	queue := []*TreeNode{root}
	for i := 1; i < len(in); {
		n := queue[0]
		queue = queue[1:]
		if in[i] != c.null {
			n.Left = &TreeNode{Val: toInt(in[i])}
			queue = append(queue, n.Left)
		}
		i++
		if i < len(in) && in[i] != c.null {
			n.Right = &TreeNode{Val: toInt(in[i])}
			queue = append(queue, n.Right)
		}
		i++
	}

	return root
}

func constructorv20() codecv20 {
	return codecv20{null: "n", split: ","}
}

type codecv20 struct {
	null, split string
}

func (c codecv20) serialize(root *TreeNode) string {
	if root == nil {
		return ""
	}

	result := []string{}
	queue := []*TreeNode{root}
	for len(queue) > 0 {
		n := queue[0]
		queue = queue[1:]
		if n == nil {
			result = append(result, c.null)
		} else {
			result = append(result, strconv.Itoa(n.Val))
			queue = append(queue, n.Left, n.Right)
		}

	}

	return strings.Join(result, c.split)
}

func (c codecv20) deserialize(data string) *TreeNode {
	if len(data) == 0 {
		return nil
	}

	toInt := func(s string) int {
		num, _ := strconv.Atoi(s)
		return num
	}
	in := strings.Split(data, c.split)
	root := &TreeNode{Val: toInt(in[0])}
	queue := []*TreeNode{root}
	for i := 1; i < len(in); {
		n := queue[0]
		queue = queue[1:]
		if in[i] != c.null {
			n.Left = &TreeNode{Val: toInt(in[i])}
			queue = append(queue, n.Left)
		}
		i++

		if i < len(in) && in[i] != c.null {
			n.Right = &TreeNode{Val: toInt(in[i])}
			queue = append(queue, n.Right)
		}
		i++
	}

	return root
}

func constructorv19() codecv19 {
	return codecv19{null: "n", split: ","}
}

type codecv19 struct {
	null, split string
}

func (c codecv19) serialize(root *TreeNode) string {
	if root == nil {
		return ""
	}
	result := []string{}
	queue := []*TreeNode{root}
	for len(queue) > 0 {
		n := queue[0]
		queue = queue[1:]
		if n == nil {
			result = append(result, c.null)
		} else {
			result = append(result, strconv.Itoa(n.Val))
			queue = append(queue, n.Left, n.Right)
		}
	}

	return strings.Join(result, c.split)
}

func (c codecv19) deserialize(data string) *TreeNode {
	if len(data) == 0 {
		return nil
	}

	toInt := func(x string) int {
		num, _ := strconv.Atoi(x)
		return num
	}
	in := strings.Split(data, c.split)
	root := &TreeNode{Val: toInt(in[0])}
	queue := []*TreeNode{root}

	for i := 1; i < len(in); {
		n := queue[0]
		queue = queue[1:]
		if in[i] != c.null {
			n.Left = &TreeNode{Val: toInt(in[i])}
			queue = append(queue, n.Left)
		}
		i++

		if i < len(in) && in[i] != c.null {
			n.Right = &TreeNode{Val: toInt(in[i])}
			queue = append(queue, n.Right)
		}
		i++
	}

	return root
}

func constructorv18() codecv18 {
	return codecv18{null: "n", split: ","}
}

type codecv18 struct {
	null, split string
}

func (c codecv18) serialize(root *TreeNode) string {
	if root == nil {
		return ""
	}

	result := []string{}
	queue := []*TreeNode{root}
	for len(queue) > 0 {
		n := queue[0]
		queue = queue[1:]

		if n == nil {
			result = append(result, c.null)
		} else {
			result = append(result, strconv.Itoa(n.Val))
			queue = append(queue, n.Left, n.Right)
		}

	}

	return strings.Join(result, c.split)
}

func (c codecv18) deserialize(data string) *TreeNode {
	if len(data) == 0 {
		return nil
	}

	in := strings.Split(data, c.split)
	toInt := func(x string) int {
		num, _ := strconv.Atoi(x)
		return num
	}
	root := &TreeNode{Val: toInt(in[0])}
	queue := []*TreeNode{root}
	i := 1
	for i < len(in) {
		n := queue[0]
		queue = queue[1:]
		if in[i] != c.null {
			n.Left = &TreeNode{Val: toInt(in[i])}
			queue = append(queue, n.Left)
		}
		i++
		if i < len(in) && in[i] != c.null {
			n.Right = &TreeNode{Val: toInt(in[i])}
			queue = append(queue, n.Right)
		}
		i++
	}

	return root
}

func constructorv17() codecv17 {
	return codecv17{null: "n", split: ","}
}

type codecv17 struct {
	null, split string
}

func (c codecv17) serialize(root *TreeNode) string {
	if root == nil {
		return ""
	}

	result := []string{}
	queue := []*TreeNode{root}
	for len(queue) > 0 {
		n := queue[0]
		queue = queue[1:]
		if n == nil {
			result = append(result, c.null)
		} else {
			result = append(result, strconv.Itoa(n.Val))
			queue = append(queue, n.Left, n.Right)
		}
	}

	return strings.Join(result, c.split)
}

func (c codecv17) deserialize(data string) *TreeNode {
	if len(data) == 0 {
		return nil
	}

	in := strings.Split(data, c.split)
	toInt := func(s string) int {
		num, _ := strconv.Atoi(s)
		return num
	}
	root := &TreeNode{Val: toInt(in[0])}
	queue := []*TreeNode{root}
	i := 1

	for i < len(in) {
		n := queue[0]
		queue = queue[1:]

		if in[i] != c.null {
			n.Left = &TreeNode{Val: toInt(in[i])}
			queue = append(queue, n.Left)
		}
		i++

		if i < len(in) && in[i] != c.null {
			n.Right = &TreeNode{Val: toInt(in[i])}
			queue = append(queue, n.Right)
		}
		i++
	}

	return root
}

func constructorv16() codecv16 {
	return codecv16{null: "n", split: ","}
}

type codecv16 struct {
	null, split string
}

func (c codecv16) serialize(root *TreeNode) string {
	if root == nil {
		return ""
	}

	result := []string{}
	queue := []*TreeNode{root}
	for len(queue) > 0 {
		n := queue[0]
		queue = queue[1:]
		if n == nil {
			result = append(result, c.null)
		} else {
			result = append(result, strconv.Itoa(n.Val))
			queue = append(queue, n.Left, n.Right)
		}
	}

	return strings.Join(result, c.split)
}

func (c codecv16) deserialize(data string) *TreeNode {
	if len(data) == 0 {
		return nil
	}

	in := strings.Split(data, c.split)
	toInt := func(s string) int {
		num, _ := strconv.Atoi(s)
		return num
	}
	root := &TreeNode{Val: toInt(in[0])}
	queue := []*TreeNode{root}
	i := 1

	for i < len(in) {

		n := queue[0]
		queue = queue[1:]

		if in[i] != c.null {
			n.Left = &TreeNode{Val: toInt(in[i])}
			queue = append(queue, n.Left)
		}
		i++

		if i < len(in) && in[i] != c.null {
			n.Right = &TreeNode{Val: toInt(in[i])}
			queue = append(queue, n.Right)
		}
		i++
	}

	return root
}

func constructorv15() codecv15 {
	return codecv15{null: "n", split: ","}
}

type codecv15 struct {
	null, split string
}

func (c codecv15) serialize(root *TreeNode) string {
	if root == nil {
		return ""
	}
	result := []string{}

	queue := []*TreeNode{root}

	for len(queue) > 0 {
		n := queue[0]
		queue = queue[1:]

		if n == nil {
			result = append(result, "n")
		} else {
			result = append(result, strconv.Itoa(n.Val))
			queue = append(queue, n.Left, n.Right)
		}
	}

	return strings.Join(result, c.split)
}

func (c codecv15) deserialize(data string) *TreeNode {
	if len(data) == 0 {
		return nil
	}

	in := strings.Split(data, c.split)
	toInt := func(x string) int {
		num, _ := strconv.Atoi(x)
		return num
	}
	root := &TreeNode{Val: toInt(in[0])}
	queue := []*TreeNode{root}
	i := 1

	for i < len(in) {
		n := queue[0]
		queue = queue[1:]
		if in[i] != c.null {
			n.Left = &TreeNode{Val: toInt(in[i])}
			queue = append(queue, n.Left)
		}
		i++

		if i < len(in) && in[i] != c.null {
			n.Right = &TreeNode{Val: toInt(in[i])}
			queue = append(queue, n.Right)
		}
		i++
	}

	return root
}

func constructorv14() codecv14 {
	return codecv14{null: "n", split: ","}
}

type codecv14 struct {
	null, split string
}

func (c codecv14) serialize(root *TreeNode) string {
	if root == nil {
		return ""
	}

	result := []string{}
	queue := []*TreeNode{root}

	for len(queue) > 0 {
		n := queue[0]
		queue = queue[1:]

		if n == nil {
			result = append(result, c.null)
		} else {
			result = append(result, strconv.Itoa(n.Val))
			queue = append(queue, n.Left, n.Right)
		}

	}

	return strings.Join(result, c.split)
}

func (c codecv14) deserialize(data string) *TreeNode {
	if len(data) == 0 {
		return nil
	}

	in := strings.Split(data, c.split)
	toInt := func(x string) int {
		num, _ := strconv.Atoi(x)
		return num
	}

	root := &TreeNode{Val: toInt(in[0])}
	queue := []*TreeNode{root}
	i := 1
	for i < len(in) {
		n := queue[0]
		queue = queue[1:]

		if in[i] != c.null {
			n.Left = &TreeNode{Val: toInt(in[i])}
			queue = append(queue, n.Left)
		}
		i++

		if i < len(in) && in[i] != c.null {
			n.Right = &TreeNode{Val: toInt(in[i])}
			queue = append(queue, n.Right)
		}
		i++
	}

	return root
}

func constructorv13() codecv13 {
	return codecv13{null: "n", split: ","}
}

type codecv13 struct {
	null, split string
}

func (codec codecv13) serialize(root *TreeNode) string {
	if root == nil {
		return ""
	}

	result := []string{}
	queue := []*TreeNode{root}

	for len(queue) > 0 {
		n := queue[0]
		queue = queue[1:]

		if n == nil {
			result = append(result, codec.null)
		} else {
			result = append(result, strconv.Itoa(n.Val))
			queue = append(queue, n.Left, n.Right)
		}
	}

	return strings.Join(result, codec.split)
}

func (codec codecv13) deserialize(data string) *TreeNode {
	if len(data) == 0 {
		return nil
	}

	in := strings.Split(data, codec.split)
	toInt := func(x string) int {
		num, _ := strconv.Atoi(x)
		return num
	}
	root := &TreeNode{Val: toInt(in[0])}
	queue := []*TreeNode{root}
	i := 1

	for i < len(in) {
		n := queue[0]
		queue = queue[1:]

		if in[i] != codec.null {
			n.Left = &TreeNode{Val: toInt(in[i])}
			queue = append(queue, n.Left)
		}
		i++
		if i < len(in) && in[i] != codec.null {
			n.Right = &TreeNode{Val: toInt(in[i])}
			queue = append(queue, n.Right)
		}
		i++
	}

	return root
}

func constructorv12() codecv12 {
	return codecv12{null: "n", split: ","}
}

type codecv12 struct {
	null, split string
}

func (codec codecv12) serialize(root *TreeNode) string {
	if root == nil {
		return ""
	}
	result := []string{}

	queue := []*TreeNode{root}

	for len(queue) > 0 {
		n := queue[0]
		queue = queue[1:]
		if n == nil {
			result = append(result, codec.null)
		} else {
			result = append(result, strconv.Itoa(n.Val))
			queue = append(queue, n.Left, n.Right)
		}

	}

	return strings.Join(result, codec.split)
}

func (codec codecv12) deserialize(data string) *TreeNode {
	if len(data) == 0 {
		return nil
	}

	in := strings.Split(data, codec.split)
	toInt := func(x string) int {
		num, _ := strconv.Atoi(x)
		return num
	}
	root := &TreeNode{Val: toInt(in[0])}
	queue := []*TreeNode{root}
	i := 1
	for i < len(in) {
		n := queue[0]
		queue = queue[1:]
		if in[i] != codec.null {
			n.Left = &TreeNode{Val: toInt(in[i])}
			queue = append(queue, n.Left)
		}
		i++
		if i < len(in) && in[i] != codec.null {
			n.Right = &TreeNode{Val: toInt(in[i])}
			queue = append(queue, n.Right)
		}
		i++
	}
	return root
}

func constructorv11() *codecv11 {
	return &codecv11{
		null: "n", split: ",",
	}
}

type codecv11 struct {
	null, split string
}

func (codec *codecv11) deserialize(data string) *TreeNode {
	if len(data) == 0 {
		return nil
	}

	in := strings.Split(data, codec.split)
	toInt := func(x string) int {
		num, _ := strconv.Atoi(x)
		return num
	}

	root := &TreeNode{Val: toInt(in[0])}
	queue := []*TreeNode{root}
	i := 1
	for i < len(in) {
		n := queue[0]
		queue = queue[1:]

		if in[i] != codec.null {
			n.Left = &TreeNode{Val: toInt(in[i])}
			queue = append(queue, n.Left)
		}
		i++

		if i < len(in) && in[i] != codec.null {
			n.Right = &TreeNode{Val: toInt(in[i])}
			queue = append(queue, n.Right)
		}
		i++
	}

	return root
}

func (codec *codecv11) serialize(root *TreeNode) string {
	if root == nil {
		return ""
	}

	result := []string{}
	queue := []*TreeNode{root}

	for len(queue) > 0 {
		n := queue[0]
		queue = queue[1:]

		if n == nil {
			result = append(result, codec.null)
		} else {
			result = append(result, strconv.Itoa(n.Val))
			queue = append(queue, n.Left, n.Right)
		}
	}

	return strings.Join(result, codec.split)
}

func constructorv10() *codecv10 {
	return &codecv10{null: "n", split: ","}
}

type codecv10 struct {
	null, split string
}

func (codec *codecv10) deserialize(data string) *TreeNode {
	if len(data) == 0 {
		return nil
	}

	in := strings.Split(data, codec.split)
	toInt := func(x string) int {
		num, _ := strconv.Atoi(x)
		return num
	}
	root := &TreeNode{Val: toInt(in[0])}
	queue := []*TreeNode{root}
	i := 1

	for i < len(in) {
		n := queue[0]
		queue = queue[1:]
		if in[i] != codec.null {
			n.Left = &TreeNode{Val: toInt(in[i])}
			queue = append(queue, n.Left)
		}
		i++

		if i < len(in) && in[i] != codec.null {
			n.Right = &TreeNode{Val: toInt(in[i])}
			queue = append(queue, n.Right)
		}
		i++
	}

	return root
}

func (codec *codecv10) serialize(root *TreeNode) string {
	if root == nil {
		return ""
	}

	result := make([]string, 0)
	queue := []*TreeNode{root}

	for len(queue) > 0 {
		n := queue[0]
		queue = queue[1:]
		if n == nil {
			result = append(result, codec.null)
		} else {
			result = append(result, strconv.Itoa(n.Val))
			queue = append(queue, n.Left, n.Right)
		}

	}

	return strings.Join(result, codec.split)
}

func constructorv9() *codecv9 {
	return &codecv9{null: "n", split: ","}
}

type codecv9 struct {
	null, split string
}

func (codec *codecv9) serialize(root *TreeNode) string {
	if root == nil {
		return ""
	}

	result := make([]string, 0)
	queue := []*TreeNode{root}

	for len(queue) > 0 {
		n := queue[0]
		queue = queue[1:]

		if n == nil {
			result = append(result, codec.null)
		} else {
			result = append(result, strconv.Itoa(n.Val))
			queue = append(queue, n.Left, n.Right)
		}
	}

	return strings.Join(result, codec.split)
}

func (codec *codecv9) deserialize(data string) *TreeNode {
	if len(data) == 0 {
		return nil
	}

	in := strings.Split(data, codec.split)
	toInt := func(x string) int {
		num, _ := strconv.Atoi(x)
		return num
	}
	root := &TreeNode{Val: toInt(in[0])}
	queue := []*TreeNode{root}

	i := 1
	for i < len(in) {

		n := queue[0]
		queue = queue[1:]
		if in[i] != codec.null {
			n.Left = &TreeNode{Val: toInt(in[i])}
			queue = append(queue, n.Left)
		}
		i++

		if i < len(in) && in[i] != codec.null {
			n.Right = &TreeNode{Val: toInt(in[i])}
			queue = append(queue, n.Right)
		}
		i++

	}

	return root
}

func constructorv8() *codecv8 {
	return &codecv8{null: "n", split: ","}
}

type codecv8 struct {
	null, split string
}

func (codec *codecv8) serialize(root *TreeNode) string {
	if root == nil {
		return ""
	}

	result := []string{}
	queue := []*TreeNode{root}
	for len(queue) > 0 {
		n := queue[0]
		queue = queue[1:]
		if n == nil {
			result = append(result, codec.null)
		} else {
			result = append(result, strconv.Itoa(n.Val))
			queue = append(queue, n.Left)
			queue = append(queue, n.Right)
		}
	}

	return strings.Join(result, codec.split)
}

func (codec *codecv8) deserialize(data string) *TreeNode {
	if len(data) == 0 {
		return nil
	}

	toInt := func(x string) int {
		num, _ := strconv.Atoi(x)
		return num
	}
	in := strings.Split(data, codec.split)
	root := &TreeNode{Val: toInt(in[0])}
	queue := []*TreeNode{root}
	i := 1
	for i < len(in) {
		n := queue[0]
		queue = queue[1:]
		if in[i] != codec.null {
			n.Left = &TreeNode{Val: toInt(in[i])}
			queue = append(queue, n.Left)
		}
		i++

		if i < len(in) && in[i] != codec.null {
			n.Right = &TreeNode{Val: toInt(in[i])}
			queue = append(queue, n.Right)
		}
		i++
	}

	return root
}

type codecv7 struct {
	null, split string
}

func constructorv7() *codecv7 {
	return &codecv7{
		null:  "n",
		split: ",",
	}
}

func (codec *codecv7) serialize(root *TreeNode) string {
	if root == nil {
		return ""
	}

	queue := []*TreeNode{root}
	result := []string{}

	for len(queue) > 0 {
		qLen := len(queue)
		for qLen > 0 {

			n := queue[0]
			queue = queue[1:]

			if n == nil {
				result = append(result, codec.null)
			} else {
				result = append(result, strconv.Itoa(n.Val))
				queue = append(queue, n.Left)
				queue = append(queue, n.Right)
			}

			qLen--
		}
	}

	return strings.Join(result, codec.split)
}

func (codec *codecv7) deserialize(data string) *TreeNode {
	if len(data) == 0 {
		return nil
	}

	toInt := func(x string) int {
		num, _ := strconv.Atoi(x)
		return num
	}
	in := strings.Split(data, codec.split)
	root := &TreeNode{Val: toInt(in[0])}
	queue := []*TreeNode{root}
	i := 1

	for i < len(in) {
		n := queue[0]
		queue = queue[1:]

		if in[i] != codec.null {
			n.Left = &TreeNode{Val: toInt(in[i])}
			queue = append(queue, n.Left)
		}
		i++

		if i < len(in) && in[i] != codec.null {
			n.Right = &TreeNode{Val: toInt(in[i])}
			queue = append(queue, n.Right)
		}
		i++
	}

	return root
}

func constructorv6() *codecv6 {
	return &codecv6{
		null:  "n",
		split: ",",
	}
}

type codecv6 struct {
	null  string
	split string
}

func (codec *codecv6) serialize(root *TreeNode) string {
	if root == nil {
		return ""
	}

	result := []string{}

	queue := []*TreeNode{root}

	for len(queue) > 0 {
		qLen := len(queue)
		for qLen > 0 {
			qLen--
			n := queue[0]
			queue = queue[1:]

			if n == nil {
				result = append(result, codec.null)
			} else {
				result = append(result, strconv.Itoa(n.Val))
				queue = append(queue, n.Left)
				queue = append(queue, n.Right)
			}
		}
	}

	return strings.Join(result, codec.split)
}

func (codec *codecv6) deserialize(data string) *TreeNode {
	if len(data) == 0 {
		return nil
	}

	in := strings.Split(data, codec.split)

	toInt := func(x string) int {
		v, _ := strconv.Atoi(x)
		return v
	}

	root := &TreeNode{Val: toInt(in[0])}
	queue := []*TreeNode{root}
	i := 1
	for i < len(in) {
		n := queue[0]
		queue = queue[1:]
		if in[i] != codec.null {
			n.Left = &TreeNode{Val: toInt(in[i])}
			queue = append(queue, n.Left)
		}
		i++

		if i < len(in) && in[i] != codec.null {
			n.Right = &TreeNode{Val: toInt(in[i])}
			queue = append(queue, n.Right)
		}
		i++
	}

	return root
}

func constructorv5() *codecv5 {
	return &codecv5{null: "n"}
}

type codecv5 struct {
	null string
}

func (codec *codecv5) serialize(root *TreeNode) string {
	if root == nil {
		return ""
	}

	result := []string{}

	queue := []*TreeNode{root}
	for len(queue) > 0 {
		n := queue[0]
		queue = queue[1:]
		if n == nil {
			result = append(result, codec.null)
		} else {
			result = append(result, strconv.Itoa(n.Val))
			queue = append(queue, n.Left)
			queue = append(queue, n.Right)
		}
	}

	return strings.Join(result, ",")
}

func (codec *codecv5) deserialize(data string) *TreeNode {
	if len(data) == 0 {
		return nil
	}

	in := strings.Split(data, ",")
	toInt := func(v string) int {
		num, _ := strconv.Atoi(v)
		return num
	}

	root := &TreeNode{Val: toInt(in[0])}
	queue := []*TreeNode{root}
	i := 1
	for i < len(in) {

		n := queue[0]
		queue = queue[1:]
		if in[i] != codec.null {
			n.Left = &TreeNode{Val: toInt(in[i])}
			queue = append(queue, n.Left)
		}
		i++

		if i < len(in) && in[i] != codec.null {
			n.Right = &TreeNode{Val: toInt(in[i])}
			queue = append(queue, n.Right)
		}
		i++
	}

	return root
}

func constructorv4() *codecv4 {
	return &codecv4{
		null: "n",
	}
}

type codecv4 struct {
	null string
}

func (codec *codecv4) serialize(root *TreeNode) string {

	if root == nil {
		return ""
	}

	result := []string{}

	queue := []*TreeNode{root}

	for len(queue) > 0 {
		qLen := len(queue)
		for qLen > 0 {
			qLen--
			n := queue[0]
			queue = queue[1:]

			if n == nil {
				result = append(result, codec.null)
			} else {
				result = append(result, strconv.Itoa(n.Val))
				queue = append(queue, n.Left)
				queue = append(queue, n.Right)
			}
		}
	}

	return strings.Join(result, ",")
}

func (codec *codecv4) deserialize(data string) *TreeNode {
	if len(data) == 0 {
		return nil
	}

	in := strings.Split(data, ",")
	toInt := func(x string) int {
		num, _ := strconv.Atoi(x)
		return num
	}
	root := &TreeNode{Val: toInt(in[0])}
	queue := []*TreeNode{root}
	i := 1
	for i < len(in) {
		n := queue[0]
		queue = queue[1:]
		if in[i] != codec.null {
			n.Left = &TreeNode{Val: toInt(in[i])}
			queue = append(queue, n.Left)
		}
		i++

		if i < len(in) && in[i] != codec.null {
			n.Right = &TreeNode{Val: toInt(in[i])}
			queue = append(queue, n.Right)
		}
		i++
	}
	return root
}

func constructorv3() *codecv3 {
	return &codecv3{null: "n"}
}

type codecv3 struct {
	null string
}

func (codec *codecv3) serialize(root *TreeNode) string {
	if root == nil {
		return ""
	}

	result := []string{}
	queue := []*TreeNode{root}
	for len(queue) > 0 {
		qLen := len(queue)
		for qLen > 0 {
			qLen--

			n := queue[0]
			queue = queue[1:]

			if n == nil {
				result = append(result, codec.null)
			} else {
				result = append(result, strconv.Itoa(n.Val))
				queue = append(queue, n.Left)
				queue = append(queue, n.Right)
			}
		}
	}

	return strings.Join(result, ",")
}

func (codec *codecv3) deserialize(data string) *TreeNode {
	if len(data) == 0 {
		return nil
	}

	in := strings.Split(data, ",")
	toInt := func(n string) int {
		num, _ := strconv.Atoi(n)
		return num
	}

	root := &TreeNode{Val: toInt(in[0])}
	queue := []*TreeNode{root}
	i := 1
	for i < len(in) {
		n := queue[0]
		queue = queue[1:]

		if in[i] != codec.null {
			n.Left = &TreeNode{Val: toInt(in[i])}
			queue = append(queue, n.Left)
		}
		i++

		if i < len(in) && in[i] != codec.null {
			n.Right = &TreeNode{Val: toInt(in[i])}
			queue = append(queue, n.Right)
		}
		i++
	}
	return root
}

func constructorv2() *codecv2 {
	return &codecv2{null: "n"}
}

type codecv2 struct {
	null string
}

// 1,2,n...
func (codec *codecv2) serialize(root *TreeNode) string {
	if root == nil {
		return ""
	}
	result := []string{}
	queue := []*TreeNode{root}
	for len(queue) > 0 {
		qLen := len(queue)
		for qLen > 0 {
			n := queue[0]
			queue = queue[1:]
			qLen--

			if n == nil {
				result = append(result, codec.null)
			} else {
				queue = append(queue, n.Left)
				queue = append(queue, n.Right)
				result = append(result, strconv.Itoa(n.Val)) // i forgot to add the normal value...
			}
		}
	}

	return strings.Join(result, ",")
}
func (codec *codecv2) deserialize(data string) *TreeNode {
	if len(data) == 0 {
		return nil
	}

	toint := func(s string) int {
		num, _ := strconv.Atoi(s)
		return num
	}

	in := strings.Split(data, ",")
	root := &TreeNode{Val: toint(in[0])}
	queue := []*TreeNode{root}

	i := 1
	for i < len(in) {
		n := queue[0]
		queue = queue[1:]
		if in[i] != codec.null {
			n.Left = &TreeNode{Val: toint(in[i])}
			queue = append(queue, n.Left)
		}
		i++

		if i < len(in) && in[i] != codec.null {
			n.Right = &TreeNode{Val: toint(in[i])}
			queue = append(queue, n.Right)
		}
		i++
	}

	return root
}

func Constructorv1() *codecv1 {
	return &codecv1{null: "n"}
}

type codecv1 struct {
	null string
}

func (c *codecv1) serialize(root *TreeNode) string {

	if root == nil {
		return ""
	}

	result := []string{}

	queue := []*TreeNode{root}
	for len(queue) > 0 {

		qLen := len(queue)
		for qLen > 0 {

			n := queue[0]
			queue = queue[1:]
			qLen--
			if n != nil {
				result = append(result, strconv.Itoa(n.Val))
				queue = append(queue, n.Left)
				queue = append(queue, n.Right)
			} else {
				result = append(result, "n")
			}
		}
	}

	return strings.Join(result, ",")
}

// data:= "1,2,2,n..."
func (c *codecv1) deserialize(data string) *TreeNode {
	if len(data) == 0 {
		return nil
	}
	in := strings.Split(data, ",")

	toInt := func(s string) int {
		num, _ := strconv.Atoi(s)
		return num
	}

	root := &TreeNode{Val: toInt(in[0])}
	queue := []*TreeNode{root}
	i := 1
	for i < len(in) {
		n := queue[0]
		queue = queue[1:]
		if in[i] != c.null {
			n.Left = &TreeNode{Val: toInt(in[i])}
			queue = append(queue, n.Left)
		}
		i++

		if i < len(in) && in[i] != c.null {
			n.Right = &TreeNode{Val: toInt(in[i])}
			queue = append(queue, n.Right)
		}
		i++
	}

	return root
}

type Codec struct {
	null string
}

func Constructor() Codec {
	return Codec{null: "n"}
}

// Serializes a tree to a single string.
func (this *Codec) serialize(root *TreeNode) string {
	if root == nil {
		return ""
	}
	queue := []*TreeNode{root}
	result := make([]string, 0)
	for len(queue) > 0 {
		qLen := len(queue)
		for qLen > 0 {
			n := queue[0]
			queue = queue[1:]
			qLen--

			if n != nil {
				result = append(result, strconv.Itoa(n.Val))
				queue = append(queue, n.Left)
				queue = append(queue, n.Right)
			} else {
				result = append(result, "n")
			}
		}

	}

	// remove the remainder null values (optional)
	for i := len(result) - 1; i >= 0; i-- {
		if result[i] != "n" {
			return strings.Join(result[:i+1], ",")
		}
	}
	return strings.Join(result, ",")
}

// Deserializes your encoded data to tree.
func (this *Codec) deserialize(data string) *TreeNode {
	if len(data) == 0 {
		return nil
	}

	toInt := func(s string) int {
		num, _ := strconv.Atoi(s)
		return num
	}

	nums := strings.Split(data, ",")

	root := &TreeNode{Val: toInt(nums[0])}
	queue := []*TreeNode{root}

	i := 1
	for i < len(nums) {
		n := queue[0]
		queue = queue[1:]

		if nums[i] != this.null {
			n.Left = &TreeNode{Val: toInt(nums[i])}
			queue = append(queue, n.Left)
		}
		i++

		if i < len(nums) && nums[i] != this.null {
			n.Right = &TreeNode{Val: toInt(nums[i])}
			queue = append(queue, n.Right)
		}
		i++
	}
	return root
}

/**
 * Your Codec object will be instantiated and called as such:
 * ser := Constructor();
 * deser := Constructor();
 * data := ser.serialize(root);
 * ans := deser.deserialize(data);
 */
