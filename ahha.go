package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func calc(a string) string {
	var numc int
	numc = -999999
	var splittedA []string
	pat := `[\+\-\*/]`
	re := regexp.MustCompile(pat)
	splittedB := re.Split(a, -1)
	for i, part := range splittedB {
		if i > 0 {
			splittedA = append(splittedA, re.FindAllString(a, -1)[i-1])
		}
		splittedA = append(splittedA, part)
	}
	pattern := `^\d+$`
	reg, _ := regexp.Compile(pattern)
	matched := reg.MatchString(splittedA[2])
	if matched {
		numc, _ = strconv.Atoi(splittedA[2])
	}
	contains := strings.Contains(splittedA[0], "\"")
	if contains {
		splittedA[0] = strings.Replace(splittedA[0], "\"", "", -1)
	} else {
		panic("Первый элемент не строчка")
	}
	if len(splittedA[0]) > 10 {
		panic("Количество символо больше 10")
	}
	contain2 := strings.Contains(splittedA[2], "\"")
	if contain2 {
		splittedA[2] = strings.Replace(splittedA[2], "\"", "", -1)
		if len(splittedA[2]) > 10 {
			panic("Количество символо больше 10")
		}
	} else {
		numc, _ = strconv.Atoi(splittedA[2])
		if numc > 10 {
			panic("Число больше 10")
		}
	}
	var result string
	switch splittedA[1] {
	case "+":
		if numc != -999999 {
			panic("Нельзя прибавлять str")
		}
		result = splittedA[0] + splittedA[2]
	case "-":
		if numc != -999999 {
			panic("Нельзя вычитать str")
		}
		result = strings.Replace(splittedA[0], splittedA[2], "", -1)
	case "*":
		if numc == -999999 {
			panic("Нельзя умножать на str")
		}
		split := splittedA[0]
		for i := 0; i < numc-1; i++ {
			splittedA[0] = splittedA[0] + split
		}
		result = splittedA[0]
	case "/":
		if numc == -999999 {
			panic("Нельзя делить на str")
		}
		ru := len(splittedA[0]) / numc
		result = splittedA[0][:ru]
	}
	if len(result) > 40 {
		return (splittedA[0][:41] + "...")
	}
	return result
}
func main() {
	scanner := bufio.NewScanner(os.Stdin)
	if scanner.Scan() {
		a := scanner.Text()
		fmt.Println(calc(a))
	}
}
