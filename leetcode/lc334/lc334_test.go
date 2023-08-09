package lc334

import (
	"fmt"
	"math"
	"testing"
)

func increasingTriplet(nums []int) bool {
	first, second := nums[0], math.MaxInt
	for i := 1; i < len(nums); i++ {
		if nums[i] > second {
			return true
		} else if first > nums[i] {
			first = nums[i]
		} else if first < nums[i] {
			second = nums[i]
		}
	}
	return false
}

func TestIncreasingTriplet(t *testing.T) {
	nums := []int{1, 2, 3, 4, 5}
	fmt.Println(increasingTriplet(nums))
}
