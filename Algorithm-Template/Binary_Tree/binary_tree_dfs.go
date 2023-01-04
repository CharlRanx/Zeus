package Binary_Tree

import "fmt"

func preOrder(node *TreeNode) {
	if node == nil {
		return
	}

	fmt.Printf("%v ", node.Val)
	preOrder(node.Left)
	preOrder(node.Right)
}

func inOrder(node *TreeNode) {
	if node == nil {
		return
	}

	inOrder(node.Left)
	fmt.Printf("%v ", node.Val)
	inOrder(node.Right)
}

func postOrder(node *TreeNode) {
	if node == nil {
		return
	}

	postOrder(node.Left)
	postOrder(node.Right)
	fmt.Printf("%v ", node.Val)
}
