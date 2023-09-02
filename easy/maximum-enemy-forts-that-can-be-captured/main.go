package main

func captureForts(forts []int) int {
	var maxNum int
	var start int = 0
	var flag int = 0
	var startType int
	for index, every := range forts {
		if (every == -1 || every == 1) && flag == 0 {
			start = index
			startType = every
			flag = 1

		} else if (every == -1 || every == 1) && flag == 1 {
			if index-start > maxNum && startType != every {
				maxNum = index - start
			}
			start = index
			startType = every
		}
	}
	return maxNum - 1
}
func main() {

}
