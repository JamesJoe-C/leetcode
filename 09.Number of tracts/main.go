package main

import (
	"fmt"
)

func main() {
	fmt.Println(isPalindrome(121))
	fmt.Println(isPalindrome(123))
	fmt.Println(isPalindrome(334))
	fmt.Println(isPalindrome(333))
	fmt.Println(isPalindrome(111111114))
	fmt.Println(isPalindrome(123129389))
	fmt.Println(isPalindrome(123321))
	fmt.Println(isPalindrome(123456654321))
}

func isPalindrome(x int) bool {
	num := 0
	y := x
	for x > 0 {
		num = num*10 + x%10
		x = x / 10
	}
	if y < 0 || num != y {
		return false
	}
	return true
}
