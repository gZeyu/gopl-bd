package main

import (
	"fmt"
	"gopl-bd/ch02/ex02-02/lengthconv"
	"os"
	"strconv"
)

func main() {
	if len(os.Args) <= 1 {
		for {
			var l float64
			fmt.Scanf("%f\n", &l)
			m := lengthconv.Meter(l)
			f := lengthconv.Foot(l)
			fmt.Printf("%s = %s\n%s = %s\n",
				m, lengthconv.MToF(m), f, lengthconv.FToM(f))
		}
	}
	for _, arg := range os.Args[1:] {
		l, err := strconv.ParseFloat(arg, 64)
		if err != nil {
			fmt.Fprintf(os.Stderr, "mf: %v\n", err)
			os.Exit(1)
		}
		m := lengthconv.Meter(l)
		f := lengthconv.Foot(l)
		fmt.Printf("%s = %s\n%s = %s\n",
			m, lengthconv.MToF(m), f, lengthconv.FToM(f))
	}

}
