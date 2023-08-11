package lc643

import (
	"fmt"
	"testing"
)

func findMaxAverage(nums []int, k int) float64 {
	length, sum := len(nums), 0
	for i := 0; i < k; i++ {
		sum += nums[i]
	}
	max := sum
	for i := k; i < length; i++ {
		sum = sum + nums[i] - nums[i-k]
		if sum > max {
			max = sum
		}
	}
	return float64(max) / float64(k)
}

func TestFindMaxAverage(t *testing.T) {
	nums := []int{1, 12, -5, -6, 50, 3}
	k := 4
	fmt.Println(findMaxAverage(nums, k))
}
