package main

func rob(nums []int) int {
	length := len(nums)
	if length == 1 {
		return nums[0]
	}
	first := nums[0]
	second := Max(nums[0], nums[1])
	for i := 2; i < length; i++ {
		first, second = second, Max(second, first+nums[i])
	}
	return Max(first, second)
}
func Max(x, y int) int {
	if x > y {
		return x
	}
	return y
}
func main() {

}
