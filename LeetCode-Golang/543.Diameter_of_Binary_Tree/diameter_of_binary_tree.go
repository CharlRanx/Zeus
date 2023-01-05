package LeetCode

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

func Constructor(val int) *TreeNode {
	return &TreeNode{Val: val}
}

func (curNode *TreeNode) Insert(node *TreeNode) {
	if node.Val == curNode.Val {
		return
	}

	if node.Val > curNode.Val {
		if curNode.Right == nil {
			curNode.Right = node
		} else {
			curNode.Right.Insert(node)
		}
	} else {
		if curNode.Left == nil {
			curNode.Left = node
		} else {
			curNode.Left.Insert(node)
		}
	}
}

func diameterOfBinaryTree(root *TreeNode) int {
	var maxDepth func(node *TreeNode) int
	var maxDiameter int
	maxDepth = func(node *TreeNode) int {
		if node == nil {
			return 0
		}
		leftDepth := maxDepth(node.Left)
		rightDepth := maxDepth(node.Right)
		maxDiameter = max(leftDepth+rightDepth, maxDiameter)
		return max(leftDepth, rightDepth) + 1
	}
	maxDepth(root)
	return maxDiameter
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}
