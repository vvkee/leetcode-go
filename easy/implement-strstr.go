package main

import "fmt"

// https://leetcode-cn.com/problems/implement-strstr/description/
// 根据needle长度来判断
func strStr(haystack string, needle string) int {
	index := -1
	if len(needle) == 0 || haystack == needle {
		return 0
	}
	needleCount := len(needle)
	haystackCount := len(haystack)
	if haystackCount < needleCount {
		return index
	}
	for i := 0; i < haystackCount; i++ {
		if i+needleCount > haystackCount {
			return index
		}
		chars := haystack[i : i+needleCount]
		if string(chars) == needle {
			return i
		}
	}
	return index
}

func main() {
	fmt.Println(strStr("aaaaa", "bba"))
}
