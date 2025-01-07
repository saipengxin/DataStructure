package main

import (
	"fmt"
	"strconv"
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

// 将一个计算公式拆分(只是简单的实现原理，没有实现对负数，小数和其他复杂计算的支持)
func tokenize(s string) []string {
	var tokens []string // 存储拆分后的计算公式
	var current string  // 如果是两位及以上的数字，做字符串连接用

	for i := 0; i < len(s); i++ {
		item := s[i]
		if unicode.IsDigit(rune(item)) {
			current = current + string(item)
			continue // 如果是数字，记录之后继续向后遍历，记录下完整的数字
		}
		if current != "" {
			tokens = append(tokens, current)
		}
		tokens = append(tokens, string(item))
		current = ""
	}
	// 最后如果current不为空需要将他写入tokens中
	if current != "" {
		tokens = append(tokens, current)
	}
	return tokens
}

// 判断是否是计算符号
func isOperator(s string) bool {
	_, ok := priority[s]
	return ok
}

func precedenceOf(op string) int {
	return priority[op]
}

// 解析一个数学公式为后缀表达式（逆波兰表达式）
func getRPN(tokens []string) []string {
	expression := []string{} // 后缀表达式存储
	stack := []string{}      // 栈，存储计算符号
	for _, token := range tokens {
		// 第一个字符是数字那整个token一定是一个数字
		if unicode.IsDigit(rune(token[0])) {
			expression = append(expression, token) // 数字直接输出
			continue
		}

		if token == "(" {
			stack = append(stack, token) // 左括号压入栈中
			continue
		}

		if token == ")" {
			// 右括号输出栈中的计算符号，直到匹配到左括号
			for len(stack) > 0 && stack[len(stack)-1] != "(" {
				expression = append(expression, stack[len(stack)-1])
				stack = stack[:len(stack)-1]
			}
			if len(stack) == 0 || stack[len(stack)-1] != "(" {
				// 异常处理,走到这里栈中最少还有一个左括号，所以长度为0或者最后一个字符不是左括号说明匹配有问题
			}
			// 左括号出栈，不需要在保留
			stack = stack[:len(stack)-1]
			continue
		}

		if isOperator(token) {
			// 如果是计算符号，比较栈顶的符号和当前符号的优先级，如果栈顶符号的优先级大于当前符号优先级,则依次出栈
			for len(stack) > 0 && stack[len(stack)-1] != "(" && precedenceOf(stack[len(stack)-1]) >= precedenceOf(token) {
				expression = append(expression, stack[len(stack)-1])
				stack = stack[:len(stack)-1]
			}
			// 当前计算符号入栈
			stack = append(stack, token)
		}
	}

	// 遍历完成后，将栈中元素依次出栈方式表达式中
	for i := len(stack) - 1; i >= 0; i-- {
		if stack[i] == "(" || stack[i] == ")" {
			continue
		}
		expression = append(expression, stack[i])
	}
	return expression
}

// 计算逆波兰表达式 （只实现了简单的思路，没有支持四则运算以外的运算，也没有应有的异常处理）
func calculate(expression []string) int {
	stack := []string{}
	for _, token := range expression {
		// 数字直接入栈
		if unicode.IsDigit(rune(token[0])) {
			stack = append(stack, token)
			continue
		}

		// 非数字将栈顶的两个元素出栈并计算结果
		a, _ := strconv.Atoi(stack[len(stack)-1])
		stack = stack[:len(stack)-1]
		b, _ := strconv.Atoi(stack[len(stack)-1])
		stack = stack[:len(stack)-1]

		res := ""
		switch token {
		case "+":
			res = strconv.Itoa(b + a)
		case "-":
			res = strconv.Itoa(b - a)
		case "*":
			res = strconv.Itoa(b * a)
		case "/":
			res = strconv.Itoa(b / a)
		}
		stack = append(stack, res) // 计算结果入栈
	}
	result, _ := strconv.Atoi(stack[0])
	return result
}

func main() {
	result := calculate(getRPN(tokenize("9+(3-1)*3+10/2")))
	fmt.Println("==============================")
	fmt.Printf("%#v\n", result)
	fmt.Println("==============================")
}
