package main

import (
	"fmt"
	"strings"
)


func basename(s string) string {
	startIdx := strings.LastIndex(s, "/")
	s = s[startIdx+1:]
	endIdx := strings.LastIndex(s, ".")
	s = s[:endIdx]
	return s
}

func main() {
	fmt.Println(basename("a/b/c.go"))
}

