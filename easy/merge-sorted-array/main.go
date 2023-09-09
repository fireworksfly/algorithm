package main

import "sort"

func merge(nums1 []int, m int, nums2 []int, n int) {
	length1 := len(nums1)
	length2 := len(nums2)
	minus := length1 - length2
	for i := 0; i < length2; i++ {
		nums1[minus+i] = nums2[i]
	}
	sort.Ints(nums1)
}

func main() {

}
