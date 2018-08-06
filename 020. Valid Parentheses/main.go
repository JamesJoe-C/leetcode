package main

import (
	"fmt"
)

func main() {
	fmt.Println(isValid("{}"))
}

//未完成
func isValid(s string) bool {
	if s == "" {
		return false
	}
	temp := make([]string, 1, 2)
	len_s := len(s)
	j := 0
	fmt.Println("temp start:", len(temp), temp)
	for i := 0; i < len_s; i++ {
		temp_s := string(s[i])
		fmt.Println("temp_s", temp_s)
		if temp_s == "(" || temp_s == "{" || temp_s == "[" {
			j++
			temp = append(temp, temp_s)
			fmt.Println("temp:", temp)
			continue
		}
		fmt.Println(j, string(temp[j]))
		switch string(temp[j]) {
		case "(":
			if s[i] != ")"[0] {
				return false
			}
			temp[j] = ""
			j--
		case "{":
			if s[i] != "}"[0] {
				return false
			}
			temp[j] = ""
			j--
		case "[":
			if s[i] != "]"[0] {
				return false
			}
			temp[j] = ""
			j--
		}

	}
	return true
}
