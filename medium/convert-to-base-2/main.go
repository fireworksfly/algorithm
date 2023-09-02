package main

import (
	"fmt"
	"strconv"
)

// 如果余数为0，那么字符串就上0，如果余数为1
func baseNeg2(n int) string {
	var base2String string
	if n == 0 {
		return "0"
	}
	for n != 0 {
		a := n / (-2)
		if n-a*(-2) < 0 {
			a++
		}
		b := n - a*(-2)
		base2String = strconv.Itoa(b) + base2String
		n = a
	}
	return base2String
}

func main() {
	gg := baseNeg2(6)
	fmt.Println(gg)
}
