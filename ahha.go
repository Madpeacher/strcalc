package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

// Функция для обрезки строки, если она длиннее 40 символов
func truncateString(s string) string {
	if len(s) > 40 {
		return s[:40] + "..."
	}
	return s
}

// Функция для сложения строк с пробелом
func addStrings(a, b string) string {
	return a + " " + b
}

// Функция для вычитания строки из строки
func subtractStrings(a, b string) string {
	return strings.Replace(a, b, "", -1)
}

// Функция для умножения строки на число
func multiplyString(a string, n int) string {
	var result strings.Builder
	for i := 0; i < n; i++ {
		result.WriteString(a) // добавление строки к sB
	}
	return result.String()
}

// Функция для деления строки на число
func divideString(a string, n int) string {
	if n <= 0 || len(a) == 0 {
		panic("Неверный ввод")
	}
	if len(a)%n != 0 {
		panic("Длина строки не делится на число")
	}
	partLength := len(a) / n
	return a[:partLength]
}

// Основная функция
func main() {
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		input := scanner.Text()
		input = strings.TrimSpace(input)

		// Регулярные выражения для анализа
		stringOpRegex := regexp.MustCompile(`^"(.*?)"\s*(\+|-)\s*"(.*?)"$`)
		numOpRegex := regexp.MustCompile(`^"(.*?)"\s*([\*|/])\s*(\d+)$`)

		if match := stringOpRegex.FindStringSubmatch(input); match != nil {
			a, op, b := match[1], match[2], match[3]
			switch op {
			case "+":
				result := addStrings(a, b)
				fmt.Println(truncateString(result))
			case "-":
				result := subtractStrings(a, b)
				fmt.Println(truncateString(result))
			default:
				panic("Неверный оператор")
			}
		} else if match := numOpRegex.FindStringSubmatch(input); match != nil {
			a, op, numStr := match[1], match[2], match[3]
			num, err := strconv.Atoi(numStr)
			if err != nil || num < 1 || num > 10 {
				panic("Неверное число")
			}
			switch op {
			case "*":
				result := multiplyString(a, num)
				fmt.Println(truncateString(result))
			case "/":
				result := divideString(a, num)
				fmt.Println(truncateString(result))
			default:
				panic("Неверная операция")
			}
		} else {
			panic("Неверный формат")
		}
	}
}
