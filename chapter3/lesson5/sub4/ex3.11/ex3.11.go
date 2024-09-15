package main

import (
	"bytes"
	"strings"
)

func comma(s string) string {
	b := &bytes.Buffer{}
	signIndex := 0
	if s[0] == '+' || s[0] == '-' {
		b.WriteByte(s[0])
		signIndex = 1
	}
	commaIndex := strings.LastIndex(s, ".")
	if commaIndex == -1 {
		commaIndex = len(s)
	}
	text := s[signIndex:commaIndex]
	n := len(s) % 3
	if n > 0 {
		b.Write([]byte(text[:n]))
		if len(text) > n {
			b.WriteString(",")
		}
	}
	for i, c := range text[n:] {
		if i%3 == 0 && i != 0 {
			b.WriteRune(',')
		}
		b.WriteRune(c)
	}
	b.WriteString(s[commaIndex:])
	return b.String()
}
