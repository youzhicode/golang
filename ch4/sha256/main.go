package main

import "fmt"

const MSG_LEN = 64

func main() {
	// len 函数返回的是字节数长度，需要乘以8，得到位数
	var str string = "hello world"
	var strlen int = len(str)
	var block [MSG_LEN]byte
	// 判断是否需要补位
	var need int = strlen % MSG_LEN & 1
	var blocks = strlen / MSG_LEN + ((strlen % MSG_LEN) & 1)
	for i:=0; i<blocks; i++ {
		for j:=0; j<MSG_LEN; j++ {
			idx := i * MSG_LEN + j
			if idx < strlen {
				block[j] = str[idx]
				continue 
			}
			if need == 1 {
				block[j] = byte(128)
				break
			}
		}
	}
	
	
	
	
}