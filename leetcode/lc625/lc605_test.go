package lc625

import (
	"fmt"
	"testing"
)

func canPlaceFlowers(flowerbed []int, n int) bool {
	length := len(flowerbed)
	if n == 0 {
		return true
	}
	if length == 1 {
		if flowerbed[0] == 0 {
			return true
		} else {
			return false
		}
	}

	for i := 0; i < length; i++ {
		if n == 0 {
			return true
		}
		if flowerbed[i] == 0 {
			if i == 0 {
				if flowerbed[1] == 0 {
					flowerbed[i] = 1
					n--
				}
			} else if i == length-1 {
				if flowerbed[i-1] == 0 {
					flowerbed[i] = 1
					n--
				}
			} else if flowerbed[i+1] == 0 && flowerbed[i-1] == 0 {
				flowerbed[i] = 1
				n--
			}
		}
	}
	return n == 0
}

func TestName(t *testing.T) {
	var arr = []int{0, 0, 0, 0, 1}
	n := 2
	fmt.Println(canPlaceFlowers(arr, n))
}
