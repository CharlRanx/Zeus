package structures

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

func PreorderTraversal(root *TreeNode) []int {
	var preorder func(*TreeNode)
	result := make([]int, 0)
	preorder = func(node *TreeNode) {
		if node == nil {
			return
		}
		result = append(result, node.Val)
		preorder(node.Left)
		preorder(node.Right)
	}
	preorder(root)
	return result
}
