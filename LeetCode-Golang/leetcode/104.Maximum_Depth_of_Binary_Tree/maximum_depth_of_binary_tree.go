package LeetCode

import (
	"github.com/CharlRanx/Zeus/LeetCode-Golang/structures"
)

type TreeNode = structures.TreeNode

func maxDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}
	return max(maxDepth(root.Left), maxDepth(root.Right)) + 1

}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
