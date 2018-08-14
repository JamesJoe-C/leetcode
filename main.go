package main

import (
	"fmt"
)

func main() {
	test_arr := []int{1, 23, 4}
	fmt.Println(test_arr)
	test(test_arr)
	fmt.Println(test_arr)
}

func test(tt []int) {
	tt[0] = 999
}
