package main

func wordBreak(s string, wordDict []string) bool {
	length := len(s)
	dp := make([]bool, length+1)
	wordDictMap := make(map[string]bool, len(wordDict))
	for _, words := range wordDict {
		wordDictMap[words] = true
	}
	dp[0] = true
	for i := 0; i < length; i++ {
		for j := 0; j <= i; j++ {
			if dp[j] && wordDictMap[s[j:i+1]] {
				dp[i+1] = true
			}
		}
	}
	return dp[len(s)]
}

func main() {

}
