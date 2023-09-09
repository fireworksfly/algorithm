package main

import "sort"

func eliminateMaximum(dist []int, speed []int) int {
	var num int = 0
	var timesList []int
	current := 0
	for index, every := range dist {
		if every%speed[index] == 0 {
			timesList = append(timesList, every/speed[index])
		} else {
			timesList = append(timesList, every/speed[index]+1)
		}
	}
	sort.Ints(timesList)
	for i := 0; i < len(timesList); i++ {
		if timesList[i] <= current {
			return num
		}
		num++
		current++
	}
	return num
}

func Min(dist []int) (int, int) {
	minNum := dist[0]
	minIndex := 0
	for index, every := range dist {
		if every <= minNum {
			minIndex = index
			minNum = dist[index]
		}
	}
	return minNum, minIndex
}
func removeElement(dist []int, index int) []int {
	return append(dist[:index], dist[index+1:]...)
}

func main() {

}
