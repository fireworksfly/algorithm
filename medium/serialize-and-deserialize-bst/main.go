package main

import (
	"fmt"
	"sort"
	"strconv"
	"strings"
)

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

type Codec struct {
}

func Constructor() Codec {
	return Codec{}
}

// Serializes a tree to a single string.
func (this *Codec) serialize(root *TreeNode) string {
	var valSlice []int
	var returnString string
	preSlice := preRoot(root, valSlice)
	for _, every := range preSlice {
		returnString = returnString + strconv.Itoa(every) + ","
	}
	returnString = strings.TrimRight(returnString, ",")
	return returnString
}

// Deserializes your encoded data to tree.
func (this *Codec) deserialize(data string) *TreeNode {
	var preSlice []int
	if data != "" {
		for _, every := range strings.Split(data, ",") {
			num, _ := strconv.Atoi(every)
			preSlice = append(preSlice, num)
		}
	}
	orderSlice := make([]int, len(preSlice))
	copy(orderSlice, preSlice)
	sort.Ints(orderSlice)
	treeNode := buildTree(preSlice, orderSlice)
	return treeNode
}

func buildTree(preSlice []int, orderSlice []int) *TreeNode {
	var i int
	if len(preSlice) == 0 || len(orderSlice) == 0 {
		return nil
	}
	rootVal := preSlice[0]
	root := &TreeNode{Val: rootVal}
	for index, every := range orderSlice {
		if every == rootVal {
			i = index
			break
		}
	}
	root.Left = buildTree(preSlice[1:i+1], orderSlice[:i])
	root.Right = buildTree(preSlice[i+1:], orderSlice[i+1:])
	return root
}
func preRoot(root *TreeNode, valSlice []int) []int {
	if root != nil {
		valSlice = append(valSlice, root.Val)
		valSlice = preRoot(root.Left, valSlice)
		valSlice = preRoot(root.Right, valSlice)
	}
	return valSlice
}

/**
 * Your Codec object will be instantiated and called as such:
 * ser := Constructor()
 * deser := Constructor()
 * tree := ser.serialize(root)
 * ans := deser.deserialize(tree)
 * return ans
 */

func main() {
	kk := strings.Split("23,2", ",")
	fmt.Println(len(kk))
}
