package main

import (
	"fmt"
	"strings"
)

func isAng(s1 string, s2 string) bool {
	dict := []byte(s1)
	for i := range dict {
		if !strings.Contains(s2, string(dict[i])) {
			return false
		}
	}
	return true
}

func main() {
	fmt.Println(isAng("abcd", "dbca"))
}
