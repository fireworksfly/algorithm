package main

import "fmt"

func waysToBuyPensPencils(total int, cost1 int, cost2 int) int64 {
	var maxCount int
	var loss int
	var num int64 = 0
	maxCount = total / cost1
	for i := 0; i <= maxCount; i++ {
		loss = total - i*cost1
		num = num + int64(loss/cost2) + 1
	}
	return num
}

func main() {
	kk := waysToBuyPensPencils(10, 2, 5)
	fmt.Println(kk)
}
