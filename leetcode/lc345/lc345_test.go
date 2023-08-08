package lc345

import (
	"fmt"
	"strings"
	"testing"
)

func reverseVowels(s string) string {
	bytes := []byte(s)
	length := len(s)
	vowels := "aeiouAEIOU"
	i, j := 0, len(s)-1

	for i < j {
		for i < length && !strings.Contains(vowels, string(bytes[i])) {
			i++
		}
		for j >= 0 && !strings.Contains(vowels, string(bytes[j])) {
			j--
		}
		if i < j {
			bytes[i], bytes[j] = bytes[j], bytes[i]
			i++
			j--
		}
	}
	return string(bytes)
}

func TestReverse(t *testing.T) {
	s := "hello"
	fmt.Println(reverseVowels(s))
}
