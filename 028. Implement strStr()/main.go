package main

import (
	"fmt"
	"log"
)

func main() {
	log.Println(strStr("mississippi", "issip"))

}

func strStr(haystack string, needle string) int {
	if needle == "" {
		return 0
	}
	len_hs := len(haystack)
	len_nd := len(needle)
	if len_nd > len_hs {
		return -1
	}
	for i := 0; i < len_hs; i++ {
		if haystack[i] == needle[0] && (len_hs-i) >= len_nd {
			for j := 0; j < len_nd; j++ {
				fmt.Println(string(needle[j]), "==", string(haystack[i+j]))
				if needle[j] != haystack[i+j] {
					break
				} else if j == len_nd-1 {
					return i
				}
			}
		}
	}

	return -1
}
