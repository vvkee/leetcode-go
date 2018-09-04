package main

import "fmt"

// 双指针的方法
// https://leetcode-cn.com/problems/remove-duplicates-from-sorted-array/description/
func removeDuplicates(nums []int) int {
	j := 0
	for index := 1; index < len(nums); index++ {
		if nums[index] != nums[j] {
			j++
			nums[j] = nums[index]
		}
	}
	return j + 1
}

func main() {
	nums := []int{0, 0, 1, 1, 1, 2, 2, 3, 3, 4}
	fmt.Println(removeDuplicates(nums))
	fmt.Println(nums)
}
