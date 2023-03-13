package LeetCode

import (
	"github.com/CharlRanx/Zeus/LeetCode-Golang/structures"
)

type TreeNode = structures.TreeNode

func buildTree(preorder []int, inorder []int) *TreeNode {
	// 存储inorder中值到索引的映射
	valToIndex := make(map[int]int)
	for i, v := range inorder {
		valToIndex[v] = i
	}
	return build(preorder, 0, len(preorder)-1, inorder, 0, len(inorder)-1, valToIndex)
}

// 前序遍历数组为 preorder[preStart...preEnd]
// 中序遍历数组为 inorder[inStart...inEnd]
// 构造该二叉树并返回其根节点
func build(preorder []int, preStart int, preEnd int,
	inorder []int, inStart int, inEnd int, valToIndex map[int]int) *TreeNode {
	if preStart > preEnd {
		return nil
	}

	// root 值对应前序遍历数组的第一个元素
	rootVal := preorder[preStart]
	// root 在中序遍历数组中的index
	index := valToIndex[rootVal]
	// 左子树的size
	leftSize := index - inStart
	// 构造根节点
	root := &TreeNode{Val: rootVal}
	// 构造左右子树
	root.Left = build(preorder, preStart+1, preStart+leftSize, inorder, inStart, index-1, valToIndex)
	root.Right = build(preorder, preStart+leftSize+1, preEnd, inorder, index+1, inEnd, valToIndex)
	return root
}
