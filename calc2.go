package main

import (
	"fmt"
	"strconv"
	"strings"
)

func romanToArabic(roman string) int {
	romanMap := map[string]int{
		"I":    1,
		"II":   2,
		"III":  3,
		"IV":   4,
		"V":    5,
		"VI":   6,
		"VII":  7,
		"VIII": 8,
		"IX":   9,
		"X":    10,
	}

	return romanMap[roman]
}

func calculate(a, b int, operator string) int {
	switch operator {
	case "+":
		return a + b
	case "-":
		return a - b
	case "*":
		return a * b
	case "/":
		return a / b
	default:
		return 0
	}
}

func main() {
	var input string
	fmt.Println("Введите выражение (например, 5 + 3 или V - III):")
	fmt.Scanln(&input)

	input = strings.ReplaceAll(input, " ", "")
	var a, b int
	var operator string

	if strings.ContainsAny(input, "+-*/") {
		for i, char := range input {
			if char == '+' || char == '-' || char == '*' || char == '/' {
				aStr := input[:i]
				bStr := input[i+1:]
				operator = string(char)

				if num, err := strconv.Atoi(aStr); err == nil {
					if num >= 1 && num <= 10 {
						a = num
					} else {
						fmt.Println("Введите числа от 1 до 10.")
						return
					}
				} else {
					a = romanToArabic(aStr)
				}

				if num, err := strconv.Atoi(bStr); err == nil {
					if num >= 1 && num <= 10 {
						b = num
					} else {
						fmt.Println("Введите числа от 1 до 10.")
						return
					}
				} else {
					b = romanToArabic(bStr)
				}

				result := calculate(a, b, operator)
				if result == 0 {
					fmt.Println("Ошибка. В римской системе счисления нет 0")
				} else {
					fmt.Printf("Результат: %d\n", result)
					return
				}
			}
		}
	} else {
		fmt.Println("Неверный формат выражения.")
	}
}
