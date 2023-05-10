package main

import (
	"bufio"
	"fmt"
	"os"
)

/*
	从文件或者标准输入中获取数据，并统计重复行

*/

func main() {

	counts := make(map[string]int)
	if len(os.Args) < 2 {
		input(os.Stdin, counts)
	} else {
		for _, file := range os.Args[1:] {
			fileHandle, err := os.Open(file)
			if err != nil {
				fmt.Fprintf(os.Stderr, "Have error %v\n", err)
				continue
			}
			input(fileHandle, counts)
			fileHandle.Close()
		}
	}

	for line, number := range counts {
		if number > 1 {
			fmt.Printf("line: %q is dumplication %d\n", line, number)
		}
	}

}

func input(source *os.File, counts map[string]int) {
	input := bufio.NewScanner(source)

	for input.Scan() {
		counts[input.Text()]++
	}
}
