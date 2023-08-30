package main

import (
	"fmt"
)

func isValid(s string) bool {
	length := len(s)
	if length%2 == 1 {
		return false
	}
	var stack []byte
	pairs := map[byte]byte{
		'(': ')',
		'[': ']',
		'{': '}',
	}
	for _, char := range s {
		if pairs[byte(char)] > 0 {
			stack = append(stack, byte(char))
		} else if len(stack) == 0 {
			return false
		} else {
			if pairs[stack[len(stack)-1]] == byte(char) {
				stack = stack[:len(stack)-1]
			} else {
				return false
			}
		}
	}
	if len(stack) > 0 {
		return false
	}
	return true
}

func main() {
	s := "){"
	gg := isValid(s)
	fmt.Println(gg)
}
