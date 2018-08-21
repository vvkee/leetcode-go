package main

import (
	"fmt"
)

func reverse(x int) int {
	rev := 0
	for x != 0 {
		pow := x % 10
		x = x / 10
		rev = rev*10 + pow
	}
	if rev > 1<<31-1 || rev <= -1<<31 {
		return 0
	}
	return rev
}

func main() {
	fmt.Println(reverse(-234239408239048))
}
