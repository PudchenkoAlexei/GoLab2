package Lab_2

import (
	"fmt"
	"strconv"
	"strings"
)

func isOperator(c string) bool {
	return len(c) == 1 && strings.ContainsAny(c, "+-*/^")
}

func isValidOperand(c string) bool {
	_, err := strconv.Atoi(c)
	return err == nil
}

func reverseSlice(slice []string) []string {
	for i, j := 0, len(slice)-1; i < j; i, j = i+1, j-1 {
		slice[i], slice[j] = slice[j], slice[i]
	}
	return slice
}

func PrefixToPostfix(prefixExpr string) (string, error) {
	tokens := strings.Split(prefixExpr, " ")

	for _, token := range tokens {
		if isOperator(token) || isValidOperand(token) {
			continue
		}
		return "", fmt.Errorf("invalid input")
	}

	expr := reverseSlice(tokens)
	var stack []string

	for _, token := range expr {
		if isOperator(token) {
			if len(stack) < 2 {
				return "", fmt.Errorf("invalid expression")
			}
			op1 := stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			op2 := stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			stack = append(stack, op1+" "+op2+" "+token)
		} else {
			stack = append(stack, token)
		}
	}

	if len(stack) != 1 {
		return "", fmt.Errorf("invalid expression")
	}

	return stack[0], nil
}
