package main

import "fmt"

/**
	给定一个字符串，找出不含有重复字符的最长子串的长度。
	示例：
	给定 "abcabcbb" ，没有重复字符的最长子串是 "abc" ，那么长度就是3。
	给定 "bbbbb" ，最长的子串就是 "b" ，长度是1。
	给定 "pwwkew" ，最长子串是 "wke" ，长度是3。请注意答案必须是一个子串，"pwke" 是 子序列  而不是子串。
**/

func main() {
	s1 := "abcabcbb"
	s2 := "bbbbb"
	s3 := "pwwkew"
	fmt.Println(lengthOfLongestSubstring(s1))
	fmt.Println(lengthOfLongestSubstring(s2))
	fmt.Println(lengthOfLongestSubstring(s3))
}

func lengthOfLongestSubstring(s string) int {
	maxLength := 0

	for i:=0;i<len(s)-1;i++ {
		for j:=i+1;j<len(s)-1;j++ {
			if(s[i] == s[j]){
				if(j-i >= maxLength){
					maxLength = j -i
				}
				i++
			}else {
				continue
			}
		}
	}

	return maxLength
}