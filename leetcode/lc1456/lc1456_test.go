package lc1456

import (
	"fmt"
	"testing"
)

func yuanShenQiDong(yuan byte) int {
	if yuan == 'a' || yuan == 'e' || yuan == 'i' || yuan == 'o' || yuan == 'u' {
		return 1
	}
	return 0
}

func maxVowels(s string, k int) int {
	bytes := []byte(s)
	length := len(s)
	max := 0
	for i := 0; i < k; i++ {
		max += yuanShenQiDong(bytes[i])
	}
	count := max
	for i := 1; i+k-1 < length; i++ {
		count = count - yuanShenQiDong(bytes[i-1]) + yuanShenQiDong(bytes[i+k-1])
		if max < count {
			max = count
		}
	}
	return max
}

func TestMaxVowels(t *testing.T) {
	s := "abciiidef"
	k := 3
	fmt.Println(maxVowels(s, k))
}
