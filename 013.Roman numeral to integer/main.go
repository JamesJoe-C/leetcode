package main

import (
	"fmt"
)

func main() {
	fmt.Println(3, romanToInt("III"))
	fmt.Println(4, romanToInt("IV"))
	fmt.Println(9, romanToInt("IX"))
	fmt.Println(58, romanToInt("LVIII"))
	fmt.Println(1994, romanToInt("MCMXCIV"))

}

func romanToInt(s string) int {
	number_index := map[string]int{
		"I": 1,
		"V": 5,
		"X": 10,
		"L": 50,
		"C": 100,
		"D": 500,
		"M": 1000,
	}
	len := len(s)
	number := 0
	//从右向左遍历字符串
	for i := len - 1; i >= 0; i-- {
		//fmt.Println(i, string(s[i]))
		a, b := 0, 0
		a = number_index[string(s[i])]
		if i >= 1 {
			b = number_index[string(s[i-1])]
			if b < a {
				a = a - b
				i--
			}
		}
		number += a

	}
	return number
}
