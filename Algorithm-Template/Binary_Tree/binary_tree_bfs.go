package Binary_Tree

import "container/list"

// Breadth-First Traversal
func levelOrder(root *TreeNode) []int {
	queue := list.New()
	queue.PushBack(root)

	nums := make([]int, 0)
	for queue.Len() > 0 {
		node := queue.Remove(queue.Front()).(*TreeNode)
		nums = append(nums, node.Val)

		if node.Left != nil {
			queue.PushBack(node.Left)
		}

		if node.Right != nil {
			queue.PushBack(node.Right)
		}
	}
	return nums
}
