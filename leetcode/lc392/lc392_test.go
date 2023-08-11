package lc392

import (
	"fmt"
	"testing"
)

func isSubsequence(s string, t string) bool {
	i, j, sLength, tLength := 0, 0, len(s), len(t)
	if sLength == 0 {
		return true
	}
	if tLength == 0 {
		return false
	}
	for i < tLength {
		if s[j] == t[i] {
			j++
		}
		if j >= sLength {
			return true
		}
		i++
	}
	return j >= sLength
}

func TestIsSubsequence(t *testing.T) {
	s := "abc"
	tt := "ahbgdc"
	fmt.Println(isSubsequence(s, tt))
}
