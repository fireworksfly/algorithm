package main

import (
	"fmt"
)

func longestCommonPrefix(strs []string) string {
	num := len(strs)
	if num == 0 {
		return ""
	}
	prefix := strs[0]
	for i := 1; i < num; i++ {
		prefix = lcp(prefix, strs[i])
		if len(prefix) == 0 {
			break
		}
	}
	return prefix

}

func lcp(str1, str2 string) string {
	length := Min(len(str1), len(str2))
	index := 0
	for index < length && str1[index] == str2[index] {
		index++
	}

	return str1[:index]
}

func Min(num1, num2 int) int {
	if num1 < num2 {
		return num1
	}
	return num2
}
func main() {
	strs := []string{"flower", "flow", "flight"}
	final := longestCommonPrefix(strs)
	fmt.Println(final)
}
