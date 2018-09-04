package main

import (
	"fmt"
)

var brackets = map[rune]rune{'(': ')', '{': '}', '[': ']'}

// https://leetcode-cn.com/problems/valid-parentheses/description/
// ([)不算有效括号
// 利用stack做题
func isValid(s string) bool {
	sCount := len(s)
	if sCount == 0 {
		return true
	}
	if sCount == 1 {
		return false
	}
	var stack []rune
	for _, char := range s {
		count := len(stack)
		if char == '(' || char == '{' || char == '[' {
			stack = append(stack, brackets[char])
		} else if count > 0 && char == stack[count-1] {
			if count == 1 {
				stack = nil
			} else {
				stack = stack[0 : count-1]
			}
		} else {
			stack = append(stack, brackets[char])
		}
	}
	return len(stack) == 0
}

func main() {
	fmt.Println(isValid("()"))
}
