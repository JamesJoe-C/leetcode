package main

import (
	"fmt"
)

func main() {
	//fmt.Println(isValid("()[]{}"))
	fmt.Println(isValid("["))
	fmt.Println(isValid("]"))
}

//未完成
func isValid(s string) bool {
	if s == "" {
		return true
	}
	temp := make(map[int]string)
	len_s := len(s)
	j := 0
	for i := 0; i < len_s; i++ {
		temp_s := string(s[i])
		//fmt.Println("temp_s", temp_s)
		if temp_s == "(" || temp_s == "{" || temp_s == "[" {
			j++
			temp[j] = temp_s
			//fmt.Println("temp:", temp, temp_s)
			continue
		}
		//fmt.Println("i,j,temp[j]=", i, j, string(temp[j]))
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
		case "":
			return false
		}

	}
	//fmt.Println("end temp:", temp)
	if j > 1 || temp[1] != "" {
		return false
	}
	return true
}
