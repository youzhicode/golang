package main

/*
	slice demo

*/

import "fmt"

func rev(s []int) {
	
	for i, j:=0, len(s)-1; i<j; i, j=i+1,j-1 {
		s[i], s[j] = s[j], s[i]
	}
}

// GO语言针对byte类型提供了一个比较byte类型的slice的方法，但是其他类型，需要自定义
func equal(x, y []string) bool {
	if len(x) != len(y) {
		return false
	}
	
	for i := range x {
		if x[i] != y[i] {
			return false
		}
	}
	return true
}


func main() {
	
	s := [...]int{1, 2, 3, 4, 5} // 定义的是一个数组
	rev(s[:]) 
	fmt.Println(s)
	s1 := []int{1, 2, 3, 4, 5, 6, 7, 8}
	rev(s1[:2])
	rev(s1[2:])
	rev(s1)
	fmt.Println(s1)
	
	// slice 唯一合法的比较是跟nil的比较
	// 一个零值的slice等于nil
	var s2 []int
	fmt.Println(s2 == nil, len(s2))
	// 如果需要判断一个slice是否为空，可以通过
	if len(s2) == 0 {
		fmt.Println("slice 为空")
	}
	
	// 内置make可以创建一个slice, make([]T, len, [cap]) ，通过指定类型，长度，容量创建一个slice
	// 如果省略容量，那么容量默认等于长度
	// 在底层，make创建了一个匿名的数组，然后返回一个slice，只有通过返回slice，才能引用底层的数组
	// 第一种写法，slice就是底层数组的一个view
	s3 := make([]int, 10)
	// 第二种写法，slice只是涵盖了底层数组前len个元素，但是容量包含了了整个数组,额外的元素留着给未来做增长
	s4 := make([]int, 10, 20)
	fmt.Println(s3 == nil)
	fmt.Println(s4 == nil)
	
	
	
}