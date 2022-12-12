package LeetCode

import (
	"math/rand"
	"sort"
)

type Solution struct {
	preSum []int
}

func Constructor(w []int) Solution {
	for i := 1; i < len(w); i++ {
		w[i] += w[i-1]
	}
	return Solution{w}
}

// PickIndex Step1: preSum
// Step2: binary search
func (s *Solution) PickIndex() int {
	x := rand.Intn(s.preSum[len(s.preSum)-1]) + 1
	return sort.SearchInts(s.preSum, x)
}

/**
 * Your Solution object will be instantiated and called as such:
 * obj := Constructor(w);
 * param_1 := obj.PickIndex();
 */
