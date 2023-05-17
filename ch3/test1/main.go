package main

import (
	"bytes"
	"fmt"
)

func comma(s string) string {
	var buf bytes.Buffer
	var sLen = len(s)
	
	//1234567
	for i:=sLen;i>0;i-- {
		fmt.Println(i, sLen-i, i%3)
		buf.WriteByte(s[sLen-i])
		if i%3==1 && i>3 {
			buf.WriteString(",")
		}
	}
	
	return buf.String()
}

func main() {
	fmt.Println(comma("1234567"))
}