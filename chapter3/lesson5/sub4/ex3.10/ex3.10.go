package main

import (
	"bytes"
)

func comma(s string) string {
	b := &bytes.Buffer{}
	n := len(s) % 3
	if n == 0 {
		n = 3
	}
	b.WriteString(s[:n])
	for i := n; i < len(s); i += 3 {
		b.WriteByte(',')
		b.WriteString(s[i : i+3])
	}
	return b.String()
}
