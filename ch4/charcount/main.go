package main

import (
	"unicode/utf8"
	"os"
	"io"
	"bufio"
	"fmt"
	"unicode"
)

func main() {
	counts := make(map[rune]int)
	var utflen [utf8.UTFMax + 1]int
	var invalid int = 0
	
	in :=bufio.NewReader(os.Stdin)
	for {
		r, n, err := in.ReadRune()
		if err == io.EOF {
			break
		}
		if err != nil {
			fmt.Fprintf(os.Stderr, "Have error %v\n", err)
			os.Exit(1)
		}
		if r == unicode.ReplacementChar && n==1 {
			invalid++
		}
		counts[r]++
		utflen[n]++
	}
	
	fmt.Println("rune counts:")
	for c, n := range counts {
		fmt.Printf("%q\t%d\n", c, n)
	}
	
	for l, n := range utflen {
		if l > 0 {
			fmt.Printf("%q\t%d\n", l, n)
		}
	}
	
	if invalid > 0 {
		fmt.Printf("\n%d invalid utf-8 characters\n", invalid)
	}
}