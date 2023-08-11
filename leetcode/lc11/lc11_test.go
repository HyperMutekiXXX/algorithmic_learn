package lc11

import (
	"fmt"
	"testing"
)

func maxArea(height []int) int {
	max, i, j := 0, 0, len(height)-1
	for i < j {
		minHeight := Min(height[i], height[j])
		area := minHeight * (j - i)
		if max < area {
			max = area
		}
		if height[i] <= height[j] {
			i++
		} else {
			j--
		}
	}
	return max
}

func Min(i, j int) int {
	if i < j {
		return i
	}
	return j
}

func TestMaxArea(t *testing.T) {
	height := []int{1, 8, 6, 2, 5, 4, 8, 3, 7}
	fmt.Println(maxArea(height))
}
