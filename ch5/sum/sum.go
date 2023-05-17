package main

import "fmt"

/*
可变参数的示例
这里vals可看成是一个[]int类型的切片
*/
func sum(vals ...int) int {
	var total int
	for _, val := range vals {
		total += val
	}
	return total
}

func main() {
	// 这里调用者隐式创建了一个数组，然后将参数逐个复制到数组中，然后返回该数组的一个切片传入到函数中
	fmt.Println(sum(1, 2, 3, 4))
	// 如果参数本身就是一个切片，该如何传值?
	var nums = []int{1, 2, 3, 4, 5}
	fmt.Println(sum(nums...))
}
