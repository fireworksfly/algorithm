package main

import "fmt"

func main() {
	idx := map[int]int{}
	idx[1] = 1
	k, ok := idx[1]
	fmt.Println(k, ok)
}
