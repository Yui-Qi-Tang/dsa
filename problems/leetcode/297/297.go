package b75

import (
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
