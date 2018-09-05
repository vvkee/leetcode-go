package main

import (
	"fmt"
	"math"
)

func lengthOfLongestSubstring(s string) int {
	var as, i float64
	keyMap := map[rune]float64{}
	for j, str := range s {

		if keyMap[str] > 0 {
			i = math.Max(keyMap[str], i)
		}
		as = math.Max(as, float64(j)-i+1)
		keyMap[str] = float64(j) + 1
	}
	return int(as)
}

func main() {
	fmt.Println(lengthOfLongestSubstring("abcabcbb"))
}
