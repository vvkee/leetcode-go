package main

import (
	"fmt"
)

// https://leetcode-cn.com/problems/palindrome-number/description/
func isPalindrome(x int) bool {
	if x < 0 {
		return false
	}
	palindrome := 0
	copyX := x
	for copyX != 0 {
		pow := copyX % 10
		palindrome = palindrome*10 + pow
		copyX /= 10
	}
	return palindrome == x
}

func isPalindrome1(x int) bool {
	if x < 0 || (x%10 == 0 && x != 0) {
		return false
	}
	palindrome := 0
	// 处理数字的一半，因为是回文，所以存在1221，12，12这样的数字
	// 我们的目标是要判断是不是回文，不需要像第一个方法那样，把所有的都组合起来
	for palindrome < x {
		palindrome = palindrome*10 + x%10
		x /= 10
	}
	return x == palindrome || x == palindrome/10
}

func main() {
	fmt.Println(isPalindrome1(121))
}
