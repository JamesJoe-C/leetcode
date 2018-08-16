package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("结果：", mySqrt(984230921))
	fmt.Println("结果：", mySqrt(-16))
}

//暴力穷举
func mySqrt(x int) int {
	t1 := time.Now() // get current time
	return_int := 0
	if x >= 0 {
		for i := 0; i <= x; i++ {
			//fmt.Println(i)
			//赋值一个比原数小的数
			if i*i <= x {
				//fmt.Println(i)
				return_int = i
			}
			//大于x说明上一个数字最为贴近
			if i*i > x {
				break
			}
		}
	} else {
		for i := 0; i >= x; i-- {
			//fmt.Println(i)
			//赋值一个比原数小的数
			if i*i >= x {
				//fmt.Println(i)
				return_int = i
			}
			//大于x说明上一个数字最为贴近
			if i*i < x {
				break
			}
		}
	}
	elapsed := time.Since(t1)
	fmt.Println("App elapsed1: ", elapsed)
	return return_int
}
