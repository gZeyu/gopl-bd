//练习 1.2： 修改 echo 程序，使其打印每个参数的索引和值，每个一行。
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
