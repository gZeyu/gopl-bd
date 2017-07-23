//找一个数据量比较大的网站，用本小节中的程序调研网站的缓存策略，对每个
// URL执行两遍请求，查看两次时间是否有较大的差别，并且每次获取到的响应内容是否一
// 致，修改本节中的程序，将响应结果输出，以便于进行对比。

// Fetchall fetches URLs in parallel and reports their times and sizes.
package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"time"
)

func main() {
	start := time.Now()
	ch := make(chan string)
	for _, url := range os.Args[1:] {
		go fetch(0, url, ch)
		go fetch(1, url, ch)
	}
	for range os.Args[1:] {
		fmt.Println(<-ch)
		fmt.Println(<-ch)
	}
	fmt.Printf("%.2fs elapsed\n", time.Since(start).Seconds())
}

func fetch(id int, url string, ch chan<- string) {
	start := time.Now()
	resp, err := http.Get(url)
	if err != nil {
		ch <- fmt.Sprint(err) // send to channel ch
		return
	}

	nbytes, err := io.Copy(ioutil.Discard, resp.Body)
	resp.Body.Close() // don't leak resources
	if err != nil {
		ch <- fmt.Sprintf("[%d] while reading %s: %v", id, url, err)
		return
	}
	secs := time.Since(start).Seconds()
	ch <- fmt.Sprintf("[%d] %.2fs  %7d  %s", id, secs, nbytes, url)
}

//!-
