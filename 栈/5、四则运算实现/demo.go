package main

import (
	"fmt"
	"unicode"
)

// 四则运算实现的简易实现

// 定义一个map，记录计算符号的优先级
var priority = map[string]int{
	"*": 2,
	"/": 2,
	"+": 1,
	"-": 1,
}

// 解析一个数学公式为后缀表达式（逆波兰表达式）
func getRPN(str string) []string {
	// 初始化一个栈
	stack := NewList()
	// 初始化一个切片，用于存储逆波兰表达式
	expression := []string{}
	var e interface{}
	for _, char := range []rune(str) {
		if unicode.IsDigit(char) {
			// 如果是一个数字，直接输出
			expression = append(expression, string(char))
		} else {
			// 不是数字，是一个计算符号
			// 获取栈顶元素
			stack.Pop(&e)
			e1, _ := e.(string)
			// 左括号，第一个符号元素，或者优先级比栈顶元素高，入栈
			if string(char) == "(" || e == nil || priority[string(char)] > priority[e1] {
				stack.Push(string(char))
				continue
			}

			// 右括号，出栈直到碰到左括号为止
			if string(char) == ")" {
				for true {
					ok := stack.Pop(&e)
					if ok == -1 {
						break
					}
					e1, _ = e.(string)
					fmt.Println("==============================")
					fmt.Printf("%#v\n", e1)
					fmt.Println("==============================")
					if e1 == "(" {
						break
					}
					expression = append(expression, string(char))
				}
				continue
			}

			// 符号优先级小于等于栈顶元素，出栈
			for true {
				if e1 != "(" {
					expression = append(expression, e1)
				}
				if priority[string(char)] <= priority[e1] {
					ok := stack.Pop(&e)
					if ok == -1 {
						break
					}
					e1, _ = e.(string)
				} else {
					break
				}
			}

		}
	}

	for true {
		ok := stack.Pop(&e)
		if ok == -1 {
			break
		}
		e1, _ := e.(string)
		expression = append(expression, e1)
	}
	return expression
}

func main() {
	expression := getRPN("9+(3-1)*3+10/2")
	fmt.Println("==============================")
	fmt.Printf("%#v\n", expression)
	fmt.Println("==============================")
}
