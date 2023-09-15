package main

import "math"

func giveGem(gem []int, operations [][]int) int {
	for _, operation := range operations {
		send := operation[0]
		receive := operation[1]
		gem[receive] = gem[receive] + gem[send]/2
		gem[send] = gem[send] - gem[send]/2
	}
	return Max(gem) - Min(gem)
}

func Max(k []int) int {
	var maxNum int
	for _, every := range k {
		if every > maxNum {
			maxNum = every
		}
	}
	return maxNum
}

func Min(k []int) int {
	minNum := int(math.MaxInt16)
	for _, every := range k {
		if every < minNum {
			minNum = every
		}
	}
	return minNum
}
