package main

import "fmt"

func sequares() func() int {
	var x int
	return func() int {
		x++
		return x * x
	}
}

func main() {
	f := sequares()
	fmt.Println(f())
	fmt.Println(f())
	fmt.Println(f())
	fmt.Println(f())
	fmt.Println(f())
}
