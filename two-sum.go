package main

import "fmt"

func twoSum(nums []int, target int) []int {
	result := make(map[int]int)
	for index := 0; index < len(nums); index++ {
		num := nums[index]
		resultNum := target - num
		if result[resultNum] != 0 {
			return []int{result[resultNum] - 1, index}
		}
		result[num] = index + 1
	}
	fmt.Println(result)
	return []int{}
}

func main() {
	fmt.Println(twoSum([]int{2, 7, 11, 15}, 9))
}
