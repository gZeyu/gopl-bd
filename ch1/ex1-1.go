//练习 1.1： 修改 echo 程序，使其能够打印 os.Args[0] ，即被执行命令本身的名字。
// +build ignore

package main

import (
	"fmt"
	"os"
)

func main() {
	fmt.Println(os.Args[0])
}
