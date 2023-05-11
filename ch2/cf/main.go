package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"strconv"

	tc "github.com/youzhicode/golang/ch2/tempconv" // 导入一个包，并且取一个别名
)

func cal(p string, val float64) {
	if p == "1" {
		f := tc.Fahrenheit(val)
		c := tc.Celsius(val)
		fmt.Printf("%s = %s, %s = %s\n", f, tc.FToC(f), c, tc.CToF(c))
		return
	}
	if p == "2" {
		foot := tc.LengthForFoot(val)
		m := tc.LengthForMeter(val)
		fmt.Printf("%s = %s, %s = %s\n", foot, tc.FToM(foot), m, tc.MToF(m))
		return
	}
	if p == "3" {
		p := tc.WeightForPound(val)
		k := tc.WeightForKG(val)
		fmt.Printf("%s = %s, %s = %s\n", p, tc.PToK(p), k, tc.KToP(k))
		return
	}

}

var p = flag.Int("product", 1, "You sould input a value of temp or length or weight")
var v = flag.Float64("value", 0.0, "A value of temp or length or weight")

func main() {
	if len(os.Args[1:]) < 2 {
		stdin := bufio.NewScanner(os.Stdin)
		var title string = "您需要使用的功能, 1)温度转换 2)长度转换 3)重量转换:"
		fmt.Print(title)
		pmap := map[string]string {"1":"温度", "2":"长度", "3":"重量"}

		for stdin.Scan() {
			product := stdin.Text()
			if _, ok := pmap[product]; ok {
				fmt.Printf("请输入一个%s值:", pmap[product])
				stdin.Scan()
				inputVal, err:= strconv.ParseFloat(stdin.Text(), 64)
				if err != nil {
					fmt.Fprintf(os.Stderr, "Have a error %v\n", err)
					return
				}
				cal(product, inputVal)
				break
			}
			fmt.Print(title)
		}
	} else {
		flag.Parse()
		p := flag.Args()[1]
		v := flag.Args()[3]
		vv, err := strconv.ParseFloat(v, 64)
		if err != nil {
			fmt.Fprintf(os.Stderr, "cf:%v\n", err)
			return
		}
		cal(p, vv)
	}

}
