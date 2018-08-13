package main

import (
	"fmt"
)

func main() {
	fmt.Println(maxSubArray([]int{-2, 1, -3, 4, -1, 2, 1, -5, 4}))
	fmt.Println(maxSubArray([]int{-2, 1}))
}

func maxSubArray(nums []int) int {
	len_n := len(nums)
	if len_n == 0 {
		return 0
	}

	return_int := nums[0]

	for i := 0; i < len_n; i++ {
		temp := 0
		for j := i; j < len_n; j++ {
			temp += nums[j]
			if temp > return_int {
				return_int = temp
			}
		}
	}
	return return_int
}
