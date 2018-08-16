package main

import (
	"fmt"
	"strconv"
)

func main() {
	//fmt.Println(addBinary("111", "101"))
	//fmt.Println(addBinary("1010", "1011"))
	//fmt.Println(addBinary("1", "111"))
	fmt.Println(addBinary("100", "110010"))
}

func addBinary(a string, b string) string {
	len_b := len(b)

	for i := len_b - 1; i >= 0; i-- {

		temp, _ := strconv.Atoi(string(b[i]))
		len_a := len(a)
		//fmt.Println("a长度：", len_a)
		plus_i := i
		if len_b > len_a {
			//b长度大与a，减去a-b的长度就是当前计算位
			//fmt.Println("i=", i, plus_i)
			plus_i -= len_b - len_a
			//fmt.Println("当前位值；", plus_i)
		} else {
			//a长度大与b,需要加上a-b的差值，即为当前计算位(若等于，则加0，无影响)
			plus_i += len_a - len_b
		}
		a = plus(a, temp, plus_i)
		fmt.Println("【addBinary】加", string(b[i]), "，当前位置：", plus_i, "结果：", a)
	}
	return a
}

//2进制加法运算
func plus(a string, plusNum int, i int) string {
	//fmt.Println(a, plusNum, i)
	if i < 0 {
		//进位超出范围，需要在头部+plusNum
		return strconv.Itoa(plusNum) + a
	}
	//fmt.Println(i, a)
	temp, _ := strconv.Atoi(string(a[i]))
	temp += plusNum
	aa := []byte(a)
	//进位
	if temp >= 2 {
		aa[i] = []byte("0")[0]
		a = string(aa)
		fmt.Println("递归增加")
		a = plus(a, plusNum, i-1)
		fmt.Println("【进位】：", "，当前位置：", i, "增加值：", plusNum, "结果：", a)
	} else {
		aa[i] = []byte(strconv.Itoa(temp))[0]
		a = string(aa)
	}
	//fmt.Println("【plus】", a, "，当前位置：", i, "增加值：", plusNum, "结果：", a)

	return a
}
