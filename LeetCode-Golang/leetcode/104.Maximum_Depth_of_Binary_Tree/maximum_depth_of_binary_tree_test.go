package LeetCode

import (
	"fmt"
	"math/rand"
	"testing"

	"github.com/CharlRanx/Zeus/LeetCode-Golang/structures"
)

func Test_Problem104(t *testing.T) {
	root := structures.Constructor(100)
	for i := 0; i < 10; i++ {
		n := rand.Intn(500)
		root.Insert(structures.Constructor(n))
	}

	maxDepth := maxDepth(root)
	fmt.Printf("binary tree max depth: %v\n", maxDepth)
}
