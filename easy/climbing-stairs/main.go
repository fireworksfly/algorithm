package main

import "fmt"

// 递归算法，超时
//func climbStairs(n int) int {
//	var num int
//	num = 0
//	choose(n, &num)
//	return num
//}
//
//func choose(n int, num *int) {
//	if n == 0 {
//		*num++
//		return
//	}
//
//	if n > 0 {
//		choose(n-1, num)
//		choose(n-2, num)
//	}
//}

func climbStairs(n int) int {
	if n == 1 {
		return 1
	} else if n == 2 {
		return 2
	}
	num1 := 1
	num2 := 2
	var returnNum int
	for i := 3; i <= n; i++ {
		returnNum = num1 + num2
		num1 = num2
		num2 = returnNum
	}
	return returnNum
}

func main() {
	kk := climbStairs(3)
	fmt.Println(kk)
}
