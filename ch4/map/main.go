package main

import (
	"fmt"
	"sort"
)

func equals(x, y map[string]int) bool {
	// 和slice一样，map也不能用==来比较
	// 需要自定义比较函数
	if len(x) != len(y) {
		return false
	}
	for k, xv := range x {
		if yv, ok := y[k]; !ok || yv!=xv {
			return false
		}
	}
	return true
}

func main() {
	// 定义一个map的三种方式
	// 使用make
	map1 := make(map[string]int)
	map1["zhangsan"] = 21
	map1["lisi"] = 22
	
	// 使用字面量
	map2 := map[string]int {
		"guanyu": 31,
		"zhangfei": 20,
	}
	
	// 定义一个空map
	map3 := map[string]int{}
	map3["liubei"] = 31
	
	// 对map中的元素值进行修改
	map3["liubei"]++
	
	
	fmt.Printf("%v\n", map1)
	fmt.Printf("%v\n", map2)
	fmt.Printf("%v\n", map3)
	
	// 遍历元素
	// 使用range
	for k, v := range map1 {
		fmt.Printf("k=%s v=%d\n", k, v)
	}
	
	
	// map的每次遍历都是无序的，这是故意位置，为了强制让程序不依赖哈希函数的实现
	// 如果需要按照顺序遍历，则需要显式对key进行排序
	var names []string
	// var names := make([]string, 0, len(map1)) 这样做更有效
	for name,_ := range map1 {
		names = append(names, name)
	}
	sort.Strings(names)
	for _,name := range names {
		fmt.Printf("%s=>%v\n", name, map1[name])
	}
	
	// map的各种操作都是安全的，如果我们访问一个不存在的key，那么将得到value对应的零值。
	// 有时候我们需要知道对应key的值是真的0还是对应的类型零值。可以这样做:
	// 在下列这两种场景中，map的下标语法将产生两个返回值，第二个返回值是一个bool类型
	// 用于告诉我们该key是否存在 bool变量一般命名为ok
	_, ok := map1["test"]
	if !ok {
		fmt.Println("key 不存在")
	}
	
	// 或者可以合并起来写
	if _, ok := map1["test"]; !ok {
		fmt.Println("key 不存在")
	}
	
	// 比较两个map是否相等
	fmt.Println(equals(map1, map2))
	
	
	
	
}

