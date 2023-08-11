package lc1572

import (
	"fmt"
	"testing"
)

func diagonalSum(mat [][]int) int {
	length := len(mat)
	sum := 0
	for i := 0; i < length; i++ {
		if 2*i == length-1 {
			sum += mat[i][i]
		} else {
			sum += mat[i][i] + mat[i][length-i-1]
		}
	}
	return sum
}

func TestDiagonalSum(t *testing.T) {
	//mat := [][]int{
	//	{1, 2, 3},
	//	{4, 5, 6},
	//	{7, 8, 9},
	//}
	mat := [][]int{
		{6, 3, 1, 10, 7, 4},
		{4, 10, 1, 9, 5, 10},
		{5, 5, 7, 3, 8, 5},
		{2, 7, 6, 4, 7, 6},
		{7, 9, 6, 1, 8, 5},
		{1, 7, 9, 5, 8, 4},
	}
	fmt.Println(diagonalSum(mat))
}
