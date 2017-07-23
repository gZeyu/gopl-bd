//练习 1.3： 做实验测量潜在低效的版本和使用了 strings.Join 的版本的运行时间差异。（1.6
//节讲解了部分 time 包，11.4节展示了如何写标准测试程序，以得到系统性的性能评测。）
// +build ignore

package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	for idx, arg := range os.Args[1:]{
		fmt.Println(strconv.Itoa(idx + 1) + " "+ arg)
	}
}
