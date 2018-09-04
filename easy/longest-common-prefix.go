package main

import (
	"bytes"
	"fmt"
)

// https://leetcode-cn.com/problems/longest-common-prefix/description/
func longestCommonPrefix(strs []string) string {
	count := len(strs)
	if count == 0 {
		return ""
	}
	if count == 1 {
		return strs[0]
	}
	var res bytes.Buffer
	pivot := strs[0]
	res.Grow(len(pivot))
	for i := 0; i < len(pivot); i++ {
		letter := pivot[i]
		isCommon := true
		for _, str := range strs {
			if i > len(str)-1 || letter != str[i] {
				isCommon = false
			}
		}
		if isCommon {
			res.WriteString(string(letter))
		} else {
			break
		}
	}
	return res.String()
}

func main() {
	fmt.Println(longestCommonPrefix([]string{}))
}
