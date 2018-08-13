package main

import (
	"fmt"
	"strconv"
)

func main() {
	fmt.Println("计算结果：", countAndSay(6))
}

func countAndSay(n int) string {
	return_str := "1"

	for i := 1; i < n; i++ {
		//计算出当前位置的值
		len_str := len(return_str)
		temp := ""
		//计算重复项
		jj := 1
		for j := 0; j < len_str; j++ {
			if j < len_str-1 && return_str[j] == return_str[j+1] {
				jj++
			} else {
				temp += strconv.Itoa(jj) + string(return_str[j])
				jj = 1
			}
		}
		return_str = temp
	}
	return return_str
}
