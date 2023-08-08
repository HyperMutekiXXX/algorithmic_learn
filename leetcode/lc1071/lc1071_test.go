package lc1071

import (
	"fmt"
	"testing"
)

func gcdOfStrings(str1 string, str2 string) string {
	tmp1 := str1 + str2
	tmp2 := str2 + str1
	if tmp1 != tmp2 {
		return ""
	}
	return str1[:gcd(len(str1), len(str2))]
}

func gcd(a, b int) int {
	remainder := a % b
	for remainder != 0 {
		a = b
		b = remainder
		remainder = a % b
	}
	return b
}

func TestGcd(t *testing.T) {
	str1 := "ABABAB"
	str2 := "ABAB"
	fmt.Println(gcdOfStrings(str1, str2))
}
