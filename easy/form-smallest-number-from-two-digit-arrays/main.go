package main

import "sort"

func minNumber(nums1 []int, nums2 []int) int {
	sort.Ints(nums1)
	sort.Ints(nums2)
	for i := 0; i < len(nums1); i++ {
		for j := 0; j < len(nums2); j++ {
			if nums1[i] == nums2[j] {
				return nums1[i]
			}
		}
	}
	return Min(nums1[0], nums2[0])
}

func Min(x, y int) int {
	if x > y {
		return 10*y + x
	} else {
		return 10*x + y
	}
}

func main() {

}
