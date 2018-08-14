package main

import (
	"fmt"
)

func main() {
	fmt.Println(lengthOfLastWord("Hello World"))
	fmt.Println(lengthOfLastWord("a "))
	fmt.Println(lengthOfLastWord("b   a    "))
}

func lengthOfLastWord(s string) int {

	len_s := len(s)
	end_len := 0

	prev_str := false

	for i := 0; i < len_s; i++ {

		if string(s[i]) == " " {
			prev_str = true
		} else {
			if prev_str {
				end_len = 0
				prev_str = false
			}

			end_len++
		}
	}
	return end_len
}
