package main

import (
	"math"
	"sort"
)

//	func repairCars(ranks []int, cars int) int64 {
//		sort.Ints(ranks)
//		nums := make([]int64, len(ranks))
//		fmt.Println(nums)
//		for i := 0; i < cars; i++ {
//			currentMin := int64(math.MaxInt64)
//			indexMin := 0
//			for index, every := range nums {
//				tempNum := int64(ranks[index]) * (every + 1) * (every + 1)
//				if tempNum < currentMin {
//					indexMin = index
//					currentMin = tempNum
//				}
//			}
//			nums[indexMin] = nums[indexMin] + 1
//		}
//		timeMax := int64(0)
//		for index, every := range nums {
//			tempMax := int64(ranks[index]) * every * every
//			if tempMax > timeMax {
//				timeMax = tempMax
//			}
//		}
//		return timeMax
//	}
func repairCars(ranks []int, cars int) int64 {
	bestTime := sort.Search(ranks[0]*cars*cars, func(i int) bool {
		cnt := 0
		for _, every := range ranks {
			cnt += int(math.Sqrt(float64(i / every)))
		}
		return cnt >= cars
	})
	return int64(bestTime)
}

func main() {
	sort.Search()
}
