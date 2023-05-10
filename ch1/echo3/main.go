package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	var s string = ""
	s = strings.Join(os.Args[1:], ",")
	fmt.Println(s)
}
