package main

import (
	"fmt"
)

// buildNext 计算模式串 pattern 的 next 数组（部分匹配表）
// next[i] 表示 pattern[0:i+1] 内，前缀和后缀的最长公共部分的长度
func buildNext(pattern string) []int {
	m := len(pattern)
	next := make([]int, m)
	j := 0 // 表示当前匹配的前缀长度
	// next[0] 固定为 0
	next[0] = 0
	// 从 i=1 开始计算 next 数组
	for i := 1; i < m; i++ {
		fmt.Println("==============================")
		fmt.Printf("%#v\n", pattern)
		fmt.Printf("%#v\n", i)
		fmt.Printf("%c\n", pattern[i])
		fmt.Printf("%c\n", pattern[j])
		fmt.Println("==============================")
		// 当不匹配时，通过 next 数组回退
		for j > 0 && pattern[i] != pattern[j] {
			j = next[j-1]
		}
		// 如果匹配，前缀长度增加
		if pattern[i] == pattern[j] {
			j++
		}
		next[i] = j
	}
	return next
}

// kmpSearch 使用 KMP 算法在文本 text 中查找模式 pattern
// 如果找到，则返回匹配开始的索引；否则返回 -1
func kmpSearch(text, pattern string) int {
	n, m := len(text), len(pattern)
	if m == 0 {
		return 0
	}
	// 计算模式串的 next 数组
	next := buildNext(pattern)
	fmt.Println("=============next=================")
	fmt.Printf("%#v\n", next)
	fmt.Println("==============================")
	j := 0 // pattern 指针
	for i := 0; i < n; i++ {
		// 当不匹配时，利用 next 数组回溯模式串指针 j
		for j > 0 && text[i] != pattern[j] {
			j = next[j-1]
		}
		if text[i] == pattern[j] {
			j++
		}
		// 当 j 达到模式串长度，表示匹配成功
		if j == m {
			return i - m + 1
		}
	}
	return -1
}

func main() {
	text := "abcababca"
	pattern := "abcabx"
	index := kmpSearch(text, pattern)
	if index != -1 {
		fmt.Printf("在文本中找到匹配，起始索引为：%d\n", index)
	} else {
		fmt.Println("未找到匹配的子串")
	}
}
