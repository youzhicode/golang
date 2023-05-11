package main

import (
	"strconv"
	"os"
	"fmt"

	tc "github.com/youzhicode/golang/ch2/tempconv" // 导入一个包，并且取一个别名
)




func main() {
	for _, arg := range os.Args[1:] {
		t, err := strconv.ParseFloat(arg, 64)
		if err != nil {
			fmt.Fprintf(os.Stderr, "cf:%v\n", err)
		}
		f := tc.Fahrenheit(t)
		c := tc.Celsius(t)

		fmt.Printf("%s = %s, %s = %s\n", f, tc.FToC(f), c, tc.CToF(c))
	}
}
