package main

import (
	"fmt"
)

func main() {
	//fmt.Println(longestCommonPrefix([]string{"flower", "flow", "flight"}))
	//fmt.Println(longestCommonPrefix([]string{"dog", "racecar", "car"}))
	fmt.Println(longestCommonPrefix([]string{}))
}

func longestCommonPrefix(strs []string) string {
	if len(strs) == 0 {
		return ""
	}
	str_len := len(strs[0])

	return_str := ""
	for i := 0; i < str_len; i++ {
		temp := strs[0][i]
		for _, str := range strs {

			if i >= len(str) {
				// fmt.Println("i超长")
				return return_str
			}

			if str[i] != temp {
				// fmt.Println(123123)
				// fmt.Println(string(str[i]))
				return return_str
			}
		}
		return_str += string(strs[0][i])
	}

	return return_str
}
