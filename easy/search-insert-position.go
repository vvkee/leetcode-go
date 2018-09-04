package main

import "fmt"

func searchInsert(nums []int, target int) int {
	index := 0
	for i := 0; i < len(nums); i++ {
		if nums[i] == target {
			return i
		}
		if nums[i] < target {
			index = i
			index++
		}
	}
	return index
}

func main() {
	fmt.Println(searchInsert([]int{1, 3, 5, 6}, 5))
}
