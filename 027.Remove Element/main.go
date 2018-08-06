package main

import (
	"fmt"
)

func main() {
	a := []int{0, 1, 2, 2, 3, 0, 4, 2}
	fmt.Println(removeElement(a, 2))
	b := []int{3, 2, 2, 3}
	fmt.Println(removeElement(b, 3))
	c := []int{1}
	fmt.Println(removeElement(c, 1))
	fmt.Println(removeElement([]int{3, 3}, 3))
	fmt.Println(removeElement([]int{4, 5}, 5))
}

func removeElement(nums []int, val int) int {
	len := len(nums)
	i := 0
	for i < len {
		if nums[i] == val {
			len--
			nums[i] = nums[len]
		} else {
			i++
		}
	}
	return len
}
