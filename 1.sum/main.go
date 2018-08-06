package main

import "fmt"

// func main() {

// 	//twoSum([]int{3,3}, 6)
// 	//fmt.Printf("reverse:%d", reverse(120))
// }

/*

给定一个整数数列，找出其中和为特定值的那两个数。

你可以假设每个输入都只会有一种答案，同样的元素不能被重用。

示例:

给定 nums = [2, 7, 11, 15], target = 9

因为 nums[0] + nums[1] = 2 + 7 = 9
所以返回 [0, 1]

*/
func twoSum(nums []int, target int) []int {
	//var i, qyy int
	var res []int
	numslen := len(nums)
	fmt.Printf("target=%d \n", target)
	for i := 0; i < numslen; i++ {
		fmt.Printf("i=%d %d \n", i, nums[i])
		for j := i + 1; j < numslen; j++ {
			if target == nums[i]+nums[j] {
				res := []int{i, j}
				fmt.Printf("%d %d \n", res[0], res[1])
				return res
			}
		}
	}

	return res
}

/*
给定一个范围为 32 位 int 的整数，将其颠倒。

例 1:

输入: 123
输出:  321


例 2:

输入: -123
输出: -321


例 3:

输入: 120
输出: 21


注意:

假设我们的环境只能处理 32 位 int 范围内的整数。根据这个假设，如果颠倒后的结果超过这个范围，则返回 0。
*/

// func reverse(x int) int {
// 	var y int
// 	i := 1
// 	for ; x > 0; i++ {
// 		var z int
// 		z = x % 10

// 		x = x / 10
// 		//if(z==0 && x>0){
// 		//	fmt.Print("asdfasdfas\n")
// 		//	continue
// 		//}

// 		cifang := 1
// 		for j := i - 1; j > 0; j-- {
// 			cifang = cifang * 10
// 		}

// 		y += cifang * z
// 	}
// 	return y
// }
