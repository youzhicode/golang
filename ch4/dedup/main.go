package main

import (
	"bufio"
	"os"
	"fmt"
)

func main() {
	res := make(map[string]bool)
	input := bufio.NewScanner(os.Stdin)
	
	for input.Scan() {
		line := input.Text()
		if !res[line] {
			res[line] = true
			fmt.Println(line)
		}
	}
	
	if err:=input.Err(); err != nil {
		fmt.Fprintf(os.Stderr, "Have error %v\n", err)
		os.Exit(1)
	}

}

