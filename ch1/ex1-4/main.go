//练习 1.4： 修改 dup2 ，出现重复的行时打印文件名称。
// +build ignore
// "args": ["test1-4-1.txt", "test1-4-2.txt"],
package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	counts := make(map[string]int)
	files := os.Args[1:]
	if len(files) == 0 {
		countLines(os.Stdin, counts)
	} else {
		for _, arg := range files {
			f, err := os.Open(arg)
			if err != nil {
				fmt.Fprintf(os.Stderr, "dup2: %v\n", err)
				continue
			}
			countLines(f, counts)
			for _, n := range counts {
				if n > 1 {
					fmt.Println(arg)
					break
				}
			}
			counts = make(map[string]int)
			f.Close()
		}
	}
}

func countLines(f *os.File, counts map[string]int) {
	input := bufio.NewScanner(f)
	for input.Scan() {
		counts[input.Text()]++
	}
}


