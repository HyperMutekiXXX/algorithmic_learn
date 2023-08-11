package lc443

import (
	"fmt"
	"math"
	"testing"
)

func compress(chars []byte) int {
	start, res, length := 0, 0, len(chars)
	var tmp byte = 0
	for start < length {
		if chars[start] != tmp {
			res++
			end := start
			count := 0
			for end < length && chars[end] == chars[start] {
				end++
				count++
			}
			if count != 1 {
				res += int(math.Log10(float64(count))) + 1
			}
			tmp = chars[start]
			start = end
		} else {
			start++
		}
	}
	return res
}

func TestCompress(t *testing.T) {
	chars := []byte{'a', 'a', 'b', 'b', 'c', 'c', 'c'}
	fmt.Println(compress(chars))
}
