package main

import (
	"fmt"
	"math"
)

/*
给定一个 32 位有符号整数，将整数中的数字进行反转。

示例 1:

输入: 123
输出: 321
 示例 2:

输入: -123
输出: -321
示例 3:

输入: 120
输出: 21
注意:

假设我们的环境只能存储 32 位有符号整数，其数值范围是 [−231,  231 − 1]。根据这个假设，如果反转后的整数溢出，则返回 0。
*/
func main() {
	fmt.Println(reverse(-123))
	fmt.Println(reverse(123))
	fmt.Println(reverse(-33))
	fmt.Println(reverse(33))
	fmt.Println(reverse(-9999))
	fmt.Println(reverse(9999))
	fmt.Println(reverse(-100010))
	fmt.Println(reverse(100010))
	fmt.Println(reverse(-6553565535111))
	fmt.Println(reverse(6553565535111))

}

func reverse(x int) int {
	num := 0
	for x != 0 {
		num = num*10 + x%10
		x = x / 10
	}
	if num > math.MaxInt32 || num < math.MinInt32 {
		return 0
	}
	return num
}
