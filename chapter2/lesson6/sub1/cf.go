package main

import (
	"fmt"
	"main/chapter2/lesson6/tmpcnv"
	"os"
	"strconv"
)

func main() {
	for _, arg := range os.Args[1:] {
		t, err := strconv.ParseFloat(arg, 64)
		if err != nil {
			fmt.Fprintf(os.Stderr, "cf: %v\n", err)
			os.Exit(1)
		}
		f := tmpcnv.Fahrenheit(t)
		c := tmpcnv.Celsius(t)
		fmt.Printf("%s = %s, %s = %s", f, tmpcnv.FToC(f), c, tmpcnv.CToF(c))
	}
}
