package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

/*
	读取全部的文件内容，然后再分割文件，
*/

func main() {

	var counts map[string]int = make(map[string]int)
	if len(os.Args) > 1 {
		for _, file := range os.Args[1:] {
			data, err := ioutil.ReadFile(file)
			if err != nil {
				fmt.Fprintf(os.Stderr, "Have error %v\n", err)
				continue
			}
			for _, line := range strings.Split(string(data), "\n") {
				counts[line]++
			}
		}
		for line, number := range counts {
			if number > 1 {
				fmt.Printf("%q have duplication\n", line)
			}
		}
	}

}
