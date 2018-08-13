package main

import (
	"fmt"
	"log"
)

func main() {
	nums := []int{1, 3, 5, 6}
	log.Println(searchInsert(nums, 7))
	fmt.Println(nums)
}

func searchInsert(nums []int, target int) int {

	len_nums := len(nums)
	for i := 0; i < len_nums; i++ {
		if nums[i] == target {
			return i
		} else if target < nums[i] {
			var temp []int
			temp = append(temp, target)
			temp = append(temp, nums...)
			nums = temp
			return i
		} else if i < len_nums-1 && nums[i+1] > target {
			a := nums[:i]
			b := nums[i:]
			var temp []int
			temp = append(temp, a...)
			temp = append(temp, target)
			temp = append(b)
			nums = temp

			return i + 1
		} else if i == len_nums-1 && target > nums[i] {
			nums = append(nums, target)
			return len_nums
		}
	}

	return 0
}
