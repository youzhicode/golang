package main

import (
	"fmt"
)

type Currency int

const (
	USD Currency = iota
	EUR
	GBP
	RMB
)

var pc [32]byte = func()(pc [32]byte) {
	for i := range pc {
		pc[i] = pc[i/2] + byte(i&1)
	}
	return 
}()

func bitDiff(a, b byte) int {
	var sum int = 0
	for i:=0; i<8; i++ {
		a = a >> i
		b = b >> i
		if a&1 != b&1 {
			sum++
		}
	}
	return sum
}

func bitDiff2(a, b [32]byte) int {
	
	var res, sum int
	for i := range a {
		res = int(a[i]) ^ int(b[i])
		sum += int(pc[res])
	}
	
	return sum
}

func sha256Diff(sha1,sha2 [32]byte) int {
	hLen := len(sha1)
	var sum int
	for i:=0; i<hLen; i++ {
		for j:=0; j<8; j++ {
			sha1[i] = sha1[i] >> j
			sha2[i] = sha2[i] >> j
			if sha1[i]&1 != sha2[i]&1 {
				sum += 1
			}
		}
	}
	return sum
}

func main() {
	var a1 [32]byte
	var a2 [32]byte
	
	for i:=0; i<32; i++ {
		a1[i] = '1'
	}
	for i:=0; i<32; i++ {
		a2[i] = '0'
	}
//	fmt.Println(sha256Diff(a1, a2))

	// 数组声明
	var a [3]int
	fmt.Println(a[0]) // 打印第一个元素, 数组会被默认初始化为数组所属类型的零值
	fmt.Println(a[len(a) - 1]) // 打印最后一个元素
	
	// 遍历数组
	for i, v := range a {
		fmt.Printf("%d %d\n", i, v)
	}
	
	// 只遍历数组值
	for _, v := range a {
		fmt.Printf("%d\n", v)
	}
	
	// 使用数字字面值初始化
	var q [3]int = [3]int{1, 2, 3}
	for i, v := range q {
		fmt.Printf("%d, %d\n", i, v)
	}
	// 可以不指定数组大小
	var r = [...]int{1, 2, 3} // r := [...]{1, 2, 3}
	for i, v := range r {
		fmt.Printf("%d, %d\n", i, v)
	}
	
	// 使用索引和对应值的初始化方式
	symbol := [...]string{USD:"$", EUR:"€", GBP: "￡", "￥"}
	fmt.Println(RMB, symbol[RMB])
	
	fmt.Println(bitDiff2(a1, a2))
	
}

