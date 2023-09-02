package main

import "fmt"

func removeDuplicates(nums []int) int {
	slow := 0
	fast := 1
	length := len(nums)
	for ; fast < length; fast++ {
		if nums[slow] != nums[fast] {
			nums[slow+1] = nums[fast]
			slow++
		}
	}
	return slow + 1
}

func main() {
	nums_1 := []int{1, 1, 2, 2, 3}
	kk := removeDuplicates(nums_1)
	fmt.Println(kk)
}
