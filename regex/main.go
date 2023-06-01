package main

import (
	"fmt"
	"regexp"
)

const text = `My email is youmaicai168@163.com wen188fhaof@165.com
ddd@camel.com.cn`

func main() {

	re := regexp.MustCompile(`(\S+)@(\S+\.\S+)`)
	match := re.FindAllStringSubmatch(text, -1)
	//fmt.Println(match)

	for _, m := range match {
		fmt.Println(m)
	}
}
