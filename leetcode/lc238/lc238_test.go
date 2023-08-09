package lc238

import (
	"fmt"
	"testing"
)

func productExceptSelf(nums []int) []int {
	length := len(nums)
	ans := make([]int, length)
	ans[0] = 1
	for i := 1; i < length; i++ {
		ans[i] = ans[i-1] * nums[i-1]
	}
	num := nums[length-1]
	for i := length - 2; i >= 0; i-- {
		ans[i] *= num
		num *= nums[i]
	}
	return ans
}

func TestProductExceptSelf(t *testing.T) {
	var ints = []int{1, 5, 5, 6, 3}
	fmt.Println(productExceptSelf(ints))
}
