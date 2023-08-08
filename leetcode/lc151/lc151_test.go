package lc151

import (
	"fmt"
	"strings"
	"testing"
)

func reverseWords(s string) string {
	strList := strings.Fields(s)
	newStrList := make([]string, 0)
	for i := len(strList) - 1; i >= 0; i-- {
		newStrList = append(newStrList, strList[i])
	}
	return strings.Join(newStrList, " ")
}

func TestReverseWord(t *testing.T) {
	s := "a good   example"
	fmt.Println(reverseWords(s))
}
