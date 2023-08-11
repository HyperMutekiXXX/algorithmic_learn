package lc283

import (
	"fmt"
	"testing"
)

func moveZeroes(nums []int) {
	zero := 0
	for i, num := range nums {
		if num != 0 {
			nums[i], nums[zero] = nums[zero], nums[i]
			zero++
		}
	}
	fmt.Println(nums)
}

func TestMoveZeroes(t *testing.T) {
	nums := []int{0, 1, 0, 3, 12}
	moveZeroes(nums)
}
