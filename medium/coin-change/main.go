package main

import "fmt"

func coinChange(coins []int, amount int) int {
	dp := make([]int, amount+1)
	dp[0] = 0
	for i := 1; i < len(dp); i++ {
		dp[i] = -1
		for _, coin := range coins {
			temp := i - coin
			if temp >= 0 && dp[temp] != -1 {
				minTemp := dp[temp] + 1
				if dp[i] == -1 {
					dp[i] = minTemp
				}
				if dp[i] > minTemp {
					dp[i] = minTemp
				}
			}
		}
	}
	fmt.Println(dp)
	return dp[amount]
}

func main() {
	coins := []int{1, 2, 5}
	gg := coinChange(coins, 11)
	fmt.Println(gg)
}
