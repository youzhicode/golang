package main

import (
	"fmt"
	"os"
)

/*
	声明变量的几种方式
	s := ""
	var s string
	var s = ""
	var s string = ""
*/

func main() {

	var s, sep string

	for _, val := range os.Args[1:] {
		s += sep + val
		sep = " "
	}
	fmt.Println(s)
}
