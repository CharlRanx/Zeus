package Binary_Tree

import (
	"fmt"
	"testing"
)

func Test_levelOrder(t *testing.T) {
	n1 := NewTreeNode(1)
	n2 := NewTreeNode(2)
	n3 := NewTreeNode(3)
	n4 := NewTreeNode(4)
	n5 := NewTreeNode(5)
	n1.Left = n2
	n1.Right = n3
	n2.Left = n4
	n2.Right = n5

	result := levelOrder(n1)
	fmt.Printf("Breadth-First Traversal result: %v\n", result)
}
