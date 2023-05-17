package main

import (
	"fmt"
	"log"
	"time"
)

func bigSlowOperation() {

	defer trace("bigSlowOperation")()

	time.Sleep(10 * time.Second)
}

func trace(msg string) func() {
	start := time.Now()
	log.Printf("enter %s", start)

	return func() {
		log.Printf("exit %s (%s)", msg, time.Since(start))
	}
}

func double(x int) (result int) {
	defer func() { fmt.Println(result) }()
	return 2 * x
}

func triple(x int) (result int) {
	defer func() { result += 4 }()
	return double(x)
}

func main() {
	//bigSlowOperation()

	_ = double(2)

	fmt.Println(triple(4))

}
