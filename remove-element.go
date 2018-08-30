package main

import "fmt"

func removeElement(nums []int, val int) int {
	j := 0
	for index := 0; index < len(nums); index++ {
		if nums[index] != val {
			nums[j] = nums[index]
			j++
		}
	}
	return j
}

func main() {
	nums := []int{0, 1, 2, 2, 3, 0, 4, 2}
	val := 2
	fmt.Println(removeElement(nums, val))
	fmt.Println(nums)
}
