package LeetCode

import (
	"fmt"
	"testing"

	"github.com/CharlRanx/Zeus/LeetCode-Golang/structures"
)

func Test_Problem105(t *testing.T) {
	preOrder := []int{3, 9, 20, 15, 7}
	inOrder := []int{9, 3, 15, 20, 7}

	root := buildTree(preOrder, inOrder)
	result := structures.PreorderTraversal(root)
	fmt.Printf("binary tree preorder result: %v\n", result)
}
