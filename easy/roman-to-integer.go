package main

import (
	"fmt"
)

var romanMap = map[string]int{
	"I": 1,
	"V": 5,
	"X": 10,
	"L": 50,
	"C": 100,
	"D": 500,
	"M": 1000,
}

// https://leetcode-cn.com/problems/roman-to-integer/
func romanToInt(s string) int {
	count := len(s)
	res := 0
	if count == 1 {
		res = romanMap[s]
	} else {
		for i := 0; i < count; i++ {
			roman := string(s[i])
			if i == 0 {
				nextRoman := string(s[i+1])
				if romanMap[roman] >= romanMap[nextRoman] {
					res += romanMap[roman]
				}
			} else if i > 0 && i < count-1 {
				prevRoman := string(s[i-1])
				nextRoman := string(s[i+1])
				if romanMap[roman] >= romanMap[nextRoman] && romanMap[roman] <= romanMap[prevRoman] {
					res += romanMap[roman]
				} else if romanMap[roman] > romanMap[prevRoman] {
					res += romanMap[roman] - romanMap[prevRoman]
				}
			} else {
				prevRoman := string(s[i-1])
				if romanMap[roman] > romanMap[prevRoman] {
					res += romanMap[roman] - romanMap[prevRoman]
				} else {
					res += romanMap[roman]
				}
			}
		}
	}
	return res
}

func main() {
	fmt.Println(romanToInt("LVIII"))
}
