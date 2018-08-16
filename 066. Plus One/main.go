package main

import (
	"fmt"
)

func main() {
	fmt.Println(plusOne([]int{1, 2, 3}))
	fmt.Println(plusOne([]int{9, 9}))
	fmt.Println(plusOne([]int{4, 3, 2, 1}))
	fmt.Println(plusOne([]int{1, 2, 3}))

}

func plusOne(digits []int) []int {
	len_d := len(digits)
	if len_d == 0 {
		return []int{}
	}
	return carry(digits, len_d-1, 1)

	//return digits
}

//进位，剩余的值继续做递归调用
/*
@param digits 需要递归的数组
@param i 当前位key
@param add 增加的值(这里在固定加一情况下冗余，但是为可处理+n的情况，故保留)
*/
func carry(digits []int, i int, add int) []int {
	end_sum := digits[i] + add
	if end_sum >= 10 {
		digits[i] = end_sum % 10
		if i < 1 {
			//处理首位的进位
			digits = append([]int{1}, digits...)
			return digits
		}
		//递归调用处理进位
		digits = carry(digits, i-1, end_sum/10)
	} else {
		digits[i] = end_sum
	}
	//fmt.Println(digits)
	return digits
}
