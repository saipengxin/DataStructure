package main

import "fmt"

func Fbi(i int) int {
	if i == 0 {
		return 0
	}
	if i == 1 {
		return 1
	}
	return Fbi(i-1) + Fbi(i-2)
}

func main() {
	fmt.Println(Fbi(5))
}
