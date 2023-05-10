package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"
	"time"
)

func main() {
	start := time.Now()
	ch := make(chan string)
	for _, url := range os.Args[1:] {
		go fetch(url, ch)
	}

	for range os.Args[1:] {
		fmt.Println(<-ch)
	}

	secs := time.Since(start).Seconds()
	fmt.Printf("%.2f elapsed\n", secs)
}

func fetch(url string, ch chan<- string) {
	start := time.Now()
	resp, err := http.Get(url)
	if err != nil {
		ch <- fmt.Sprintf("while reading url:%s have error: %v", url, err)
		return
	}
	var fileName string
	if strings.HasPrefix(url, "https://") {
		fileName = url[8:]
	}
	writeFile, err := os.OpenFile("./"+fileName, os.O_CREATE|os.O_TRUNC|os.O_WRONLY, 6)
	if err != nil {
		ch <- fmt.Sprintf("open file of write is fail %v", err)
		return
	}
	nbytes, err := io.Copy(writeFile, resp.Body)
	resp.Body.Close()
	writeFile.Close()
	if err != nil {
		ch <- fmt.Sprintf("url:%s, error:%v", url, err)
		return
	}
	secs := time.Since(start).Seconds()
	ch <- fmt.Sprintf("%.2f, %7d, %s", secs, nbytes, url)

}
