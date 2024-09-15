package main

import (
	"bytes"
	"fmt"
)

func intsToStrings(values []int) string {
	var buf bytes.Buffer
	buf.WriteByte('[')
	for i, v := range values {
		if i > 0 {
			buf.WriteString(", ")
		}
		fmt.Fprintf(&buf, "%d", v)
	}
	buf.WriteByte(']')
	return buf.String()
}

func main() {
	fmt.Println(intsToStrings([]int{1, 2, 3}))
	fmt.Println(basename1("c/d.g.go"))
	fmt.Println(basename2("c/d.g.go"))
	fmt.Println(comma("12345"))
}
