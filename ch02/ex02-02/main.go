// 练习 2.2： 写一个通用的单位转换程序，用类似cf程序的方式从命令行读取参数，如果缺省的
// 话则是从标准输入读取参数，然后做类似Celsius和Fahrenheit的单位转换，长度单位可以对
// 应英尺和米，重量单位可以对应磅和公斤等。
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
