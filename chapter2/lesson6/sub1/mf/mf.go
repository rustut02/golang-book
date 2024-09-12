package main

import (
	"fmt"
	"main/chapter2/lesson6/sub1/lengthconv"
	"os"
	"strconv"
)

func main() {
	for _, arg := range os.Args[1:] {
		len, err := strconv.ParseFloat(arg, 64)
		if err != nil {
			fmt.Fprintf(os.Stderr, "mf: %v\n", err)
			os.Exit(1)
		}
		m := lengthconv.Meter(len)
		f := lengthconv.Foot(len)
		fmt.Printf("%s = %s, %s = %s", m, lengthconv.MToF(m), f, lengthconv.FToM(f))
	}
}
