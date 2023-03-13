package LeetCode

import (
	"fmt"
	"math/rand"
	"testing"
)

func Test_Problem144_Recursion(t *testing.T) {
	root := Constructor(100)
	for i := 0; i < 10; i++ {
		n := rand.Intn(500)
		root.Insert(Constructor(n))
	}

	result := preorderTraversal(root)
	fmt.Printf("binary tree preorder result: %v\n", result)
}

func Test_Problem144_Iteration(t *testing.T) {
	root := Constructor(100)
	for i := 0; i < 10; i++ {
		n := rand.Intn(500)
		root.Insert(Constructor(n))
	}

	result := preorderTraversalByIteration(root)
	fmt.Printf("binary tree preorder result: %v\n", result)
}
