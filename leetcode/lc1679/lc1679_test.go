package lc1679

import (
	"fmt"
	"testing"
)

type Stack struct {
	data []int
}

func (s *Stack) Push(item int) {
	s.data = append(s.data, item)
}

func (s *Stack) Pop() int {
	if len(s.data) == 0 {
		return -1
	}
	lastIndex := len(s.data) - 1
	item := s.data[lastIndex]
	s.data = s.data[:lastIndex]
	return item
}

func (s *Stack) Peek() int {
	if len(s.data) == 0 {
		return -1
	}
	return s.data[len(s.data)-1]
}

func (s *Stack) IsEmpty() bool {
	return len(s.data) == 0
}

// 进阶：得出下标数组
// 使用栈+哈希表
func maxOperationsIndex(nums []int, k int) [][]int {
	length := len(nums)
	numMap := make(map[int]*Stack, length)
	indexList := make([][]int, 0)
	for i, num := range nums {
		if stack, ok := numMap[k-num]; ok && !stack.IsEmpty() {
			indexList = append(indexList, []int{stack.Pop(), i})

		} else if _, ok := numMap[num]; !ok {
			s := &Stack{}
			s.Push(i)
			numMap[num] = s
		} else {
			numMap[num].Push(i)
		}

	}
	return indexList
}

// 单纯得出替换次数
func maxOperations(nums []int, k int) int {
	count := 0
	m := make(map[int]int)
	for _, num := range nums {
		if m[k-num] > 0 {
			count++
			m[k-num]--
		} else {
			m[num]++
		}
	}
	return count
}

func TestMaxOperations(t *testing.T) {
	nums := []int{2, 5, 4, 4, 1, 3, 4, 4, 1, 4, 4, 1, 2, 1, 2, 2, 3, 2, 4, 2}
	k := 3
	fmt.Println(maxOperationsIndex(nums, k))
	fmt.Println(maxOperations(nums, k))
}
