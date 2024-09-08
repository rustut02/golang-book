package main

import (
	"fmt"
	"main/chapter2/lesson6/weightconv"
	"os"
	"strconv"
)

func main() {
	for _, arg := range os.Args[1:] {
		m, err := strconv.ParseFloat(arg, 64)
		if err != nil {
			fmt.Fprintf(os.Stderr, "kp: %v\n", err)
			os.Exit(1)
		}
		k := weightconv.Kilos(m)
		p := weightconv.Pounds(m)
		fmt.Printf("%s = %s, %s = %s", k, weightconv.KToP(k), p, weightconv.PToK(p))
	}
}
