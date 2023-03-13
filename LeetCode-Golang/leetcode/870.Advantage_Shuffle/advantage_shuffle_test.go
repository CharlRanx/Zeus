package LeetCode

import (
	"fmt"
	"testing"
)

func Test_Problem870(t *testing.T) {
	nums1 := []int{3, 4, 0, 0, 1}
	nums2 := []int{1, 4, 5, 0, 0}
	//Output:[]int{4, 0, 0, 1, 3 }
	fmt.Printf("target slice: %v\n", advantageCount(nums1, nums2))
}
