package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func lcaDeepestLeaves(root *TreeNode) *TreeNode {
	_, treeNode := dfs(root, 0)
	return treeNode
}

func dfs(root *TreeNode, currentDepth int) (int, *TreeNode) {
	if root == nil {
		return currentDepth, &TreeNode{}
	}
	currentDepth = currentDepth + 1
	leftDepth, leftTreeNode := dfs(root.Left, currentDepth)
	rightDepth, rightTreeNode := dfs(root.Right, currentDepth)
	if leftDepth == rightDepth {
		return leftDepth, root
	} else if leftDepth > rightDepth {
		return leftDepth, leftTreeNode
	} else {
		return rightDepth, rightTreeNode
	}
}

func main() {

}
