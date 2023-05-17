package main


import (
	"fmt"
	"time"
)

var pc[256]byte = func()(pc [256]byte) {
	for i := range pc {
		pc[i] = pc[i/2] + byte(i&1)
	}
	return
}()

/*func init() {
	for i := range pc{
		pc[i] = pc[i/2] + byte(i&1)
	}
}*/

func Popcount(x uint64) int {
	return int(pc[byte(x >> (0*8))] +
		pc[byte(x >> (1*8))] +
		pc[byte(x >> (2*8))] +
		pc[byte(x >> (3*8))] +
		pc[byte(x >> (4*8))] +
		pc[byte(x >> (5*8))] +
		pc[byte(x >> (6*8))] +
		pc[byte(x >> (7*8))])
}

/*
	使用循环代替单一表达式
*/
func Popcount2(x uint64) int {
	var sum int = 0;
	for i:=0; i<8; i++ {
		sum += int(pc[byte(x >> (i*1))])
	}
	return sum
}

func Popcount3(x uint64) int {
	start := time.Now()
	var sum int = 0
	for i:=0; i<64; i++ {
		sum += int((x >> i) & 1)
	}
	fmt.Println(time.Since(start).Seconds())
	return sum
}

func Popcount4(x uint64) int {
	var sum int = 0
	for x > 0 {
		x = x & (x-1)
		sum++
	}
	return sum
}

func main() {
	fmt.Println(Popcount3(255))

}

