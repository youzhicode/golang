package main

import (
	"bufio"
	"fmt"
	"os"
)

/*
	go 语言中的verb
	%d   十进制整数
	%x, %o, %b   十六进制，八进制，二进制整数
	%f, %g, %e   浮点数
	%t           bool true or false
	%c			 字符(rune)(unicode码点)
	%s			 字符串
	%q           带双引号的字符串"abc" 或者 带单引号的字符'c'
	%v           变量类型的自然形式，适用于任何类型的变量
	%T 			 变量的类型
	%%			 字面上的百分号标志(无操作数)

*/

func main() {
	inputMap := make(map[string]int)
	input := bufio.NewScanner(os.Stdin)

	for input.Scan() {
		inputMap[input.Text()]++
	}

	for text, number := range inputMap {
		if number > 1 {
			fmt.Println(text)
		}
	}
}
