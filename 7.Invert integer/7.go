package main

import (
	"fmt"
	"strconv"
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
	reverse(-123)
}

func reverse(x int) int {
	xx := []byte(strconv.Itoa(x))
	//fmt.Println(xx)
	//fmt.Println(string(xx))
	len := len(xx)
	y := make([]byte, len, len)
	j := 0
	for i := len - 1; i >= 0; i-- {
		//fmt.Println(string(xx[i]))
		if string(xx[i]) != "-" {
			y[j] = xx[i]
		}
		j++
	}
	fmt.Println(string(y))
	yy := string(y)
	if x < 0 {
		yy = "-" + yy
	}
	fmt.Println(yy)
	rs, _ := strconv.Atoi(yy)
	fmt.Println(rs)
	return rs
}
