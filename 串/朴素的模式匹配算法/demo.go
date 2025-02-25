package main

import "fmt"

func main() {
	S := "goodgoogle"
	T := "google"

	fmt.Println("==============================")
	fmt.Printf("%#v\n", Index(S, T))
	fmt.Println("==============================")
}

func Index(S, T string) int {
	i := 0 // 主串匹配的起始位置(下标)
	j := 0 // 子串匹配的起始位置(下标)

	for {
		if i > len(S)-1 || j > len(T)-1 {
			break
		}
		if S[i] == T[j] {
			i++
			j++
		} else {
			i = i - j + 1 // 主串回到上次匹配首位的下一位
			j = 0         // 子串回到首位
		}
	}
	if j == len(T) {
		return i - len(T)
	} else {
		return -1
	}
}
