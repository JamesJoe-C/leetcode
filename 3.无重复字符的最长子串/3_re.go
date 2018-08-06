package main

////////////////////////////////////抄的，没完成，后续再搞

/*
给定一个字符串，找出不含有重复字符的最长子串的长度。

示例：

给定 "abcabcbb" ，没有重复字符的最长子串是 "abc" ，那么长度就是3。

给定 "bbbbb" ，最长的子串就是 "b" ，长度是1。

给定 "pwwkew" ，最长子串是 "wke" ，长度是3。请注意答案必须是一个子串，"pwke" 是 子序列  而不是子串。
*/

func lengthOfLongestSubstring(s string) int {
	s_byte := []byte(s)
	len := len(s_byte)
	no_length := 0
	no_j := 0
	temp_length := 0
	if len <= 1 {
		no_length = len
		return no_length
	}
	for i := 0; i < len; i++ {

		for j := i + 1; j < len; j++ {
			temp_length++
			if s_byte[i] == s_byte[j] {
				if temp_length > no_length {
					no_j = i //截取子串的起始点
					no_length = temp_length
				}
				//fmt.Println(string(s_byte[i:j]))
				break
			}

		}

		temp_length = 0
	}
	_ = no_j
	//fmt.Println("长度；", no_length)
	//fmt.Println("起始点：", no_j)
	//fmt.Println(string(s_byte[no_j:no_length]))
	return no_length
}

func lengthOfLongestSubstring_demo(s string) int {
	// location[s[i]] == j 表示：
	// s中第i个字符串，上次出现在s的j位置，所以，在s[j+1:i]中没有s[i]
	// location[s[i]] == -1 表示： s[i] 在s中第一次出现
	location := [256]int{} // 只有256长是因为，假定输入的字符串只有ASCII字符
	for i := range location {
		location[i] = -1 // 先设置所有的字符都没有见过
	}

	maxLen, left := 0, 0

	for i := 0; i < len(s); i++ {
		// 说明s[i]已经在s[left:i+1]中重复了
		// 并且s[i]上次出现的位置在location[s[i]]
		if location[s[i]] >= left {
			left = location[s[i]] + 1 // 在s[left:i+1]中去除s[i]字符及其之前的部分
		} else if i+1-left > maxLen {
			// fmt.Println(s[left:i+1])
			maxLen = i + 1 - left
		}
		location[s[i]] = i
	}

	return maxLen
}

// func main() {

// 	fmt.Println("结果应3=", lengthOfLongestSubstring("abcabcbb"))
// 	fmt.Println("结果应1=", lengthOfLongestSubstring("c"))
// 	fmt.Println("结果应0=", lengthOfLongestSubstring(""))
// 	fmt.Println("解果应为：", lengthOfLongestSubstring("bikdmd"))

// 	fmt.Println("------------------")

// 	fmt.Println("结果应3=", lengthOfLongestSubstring_demo("abcabcbb"))
// 	fmt.Println("结果应1=", lengthOfLongestSubstring_demo("c"))
// 	fmt.Println("结果应0=", lengthOfLongestSubstring_demo(""))
// 	fmt.Println("解果应为=", lengthOfLongestSubstring_demo("bikdmd"))
// }
