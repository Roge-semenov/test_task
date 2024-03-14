package main

import (
	"bytes"
	"fmt"
	"os"
	"strconv"
)

func isOperator(c rune) bool {
	return c == '+' || c == '-' || c == '*' || c == '/'
}

func precedence(c rune) int {
	switch c {
	case '+', '-':
		return 1
	case '*', '/':
		return 2
	}
	return 0
}

func infixToPostfix(expression string) []string {
	var result bytes.Buffer
	var postfix []string
	stack := []rune{}

	for _, c := range expression {
		if c >= '0' && c <= '9' {
			result.WriteRune(c)
		} else {
			if result.Len() > 0 {
				postfix = append(postfix, result.String())
				result.Reset()
			}
			if c == '(' {
				stack = append(stack, c)
			} else if c == ')' {
				for len(stack) > 0 && stack[len(stack)-1] != '(' {
					postfix = append(postfix, string(stack[len(stack)-1]))
					stack = stack[:len(stack)-1]
				}
				stack = stack[:len(stack)-1] // Pop '('
			} else if isOperator(c) {
				for len(stack) > 0 && precedence(c) <= precedence(stack[len(stack)-1]) {
					postfix = append(postfix, string(stack[len(stack)-1]))
					stack = stack[:len(stack)-1]
				}
				stack = append(stack, c)
			}
		}
	}

	if result.Len() > 0 {
		postfix = append(postfix, result.String())
	}

	for len(stack) > 0 {
		postfix = append(postfix, string(stack[len(stack)-1]))
		stack = stack[:len(stack)-1]
	}

	return postfix
}

func evaluatePostfix(postfix []string) int {
	var stack []int

	for _, elem := range postfix {
		if num, err := strconv.Atoi(elem); err == nil {
			stack = append(stack, num)
		} else {
			right, left := stack[len(stack)-1], stack[len(stack)-2]
			stack = stack[:len(stack)-2]

			switch elem {
			case "+":
				stack = append(stack, left+right)
			case "-":
				stack = append(stack, left-right)
			case "*":
				stack = append(stack, left*right)
			case "/":
				stack = append(stack, left/right)
			}
		}
	}

	return stack[0]
}

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Введите выражение в виде: go run calc.go \"(1+2)-3\"")
		return
	}

	expression := os.Args[1]
	postfixArray := infixToPostfix(expression)
	fmt.Println("Результат:", evaluatePostfix(postfixArray))
}
