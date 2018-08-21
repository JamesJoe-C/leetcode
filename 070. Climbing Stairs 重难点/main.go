package main

import (
	"fmt"
)

func main() {
	fmt.Println(climbStairs(3))
	fmt.Println(climbStairs(4))
	fmt.Println(climbStairs(5))
	fmt.Println(climbStairs(6))
	fmt.Println(climbStairs(7))

}

func climbStairs(n int) int {
	if n <= 3 {
		return n
	}
	nn := 0
	f1 := 1
	f2 := 2
	for i := 3; i <= n; i++ {
		nn = f1 + f2
		f1 = f2
		f2 = nn
	}

	return nn
}
