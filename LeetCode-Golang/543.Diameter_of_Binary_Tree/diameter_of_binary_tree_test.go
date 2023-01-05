package LeetCode

import (
	"fmt"
	"math/rand"
	"testing"
)

func Test_Problem543(t *testing.T) {
	root := Constructor(100)
	for i := 0; i < 10; i++ {
		n := rand.Intn(500)
		root.Insert(Constructor(n))
	}

	maxDiameter := diameterOfBinaryTree(root)
	fmt.Printf("binary tree max diameter: %v\n", maxDiameter)
}
