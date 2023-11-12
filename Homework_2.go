package main

import (
	"fmt"
)

func main() {
	var a int
	var b int
	var word string
	var year int
	fmt.Println("Введите первое число для проверки(проверка суммы и положительности текстом, печать слова):")
	fmt.Scan(&a)
	fmt.Println("Введите второе число для проверки(проверка суммы и положительности булеаном):")
	fmt.Scan(&b)
	fmt.Println("Введите слово для проверки:")
	fmt.Scan(&word)
	fmt.Println("Введите год для проверки:")
	fmt.Scan(&year)
	fmt.Println(checkToTwentySum(a, b))
	checkToPositively(a)
	fmt.Println(checkToPositivelyBool(b))
	printWord(a, word)
	fmt.Println(checkYear(year))
}

func checkToTwentySum(x, y int) bool {
	if x+y >= 10 && x+y <= 20 {
		return true
	}
	return false
}

func checkToPositively(x int) {
	switch {
	case x >= 0:
		fmt.Println("Число положительное")
	default:
		fmt.Println("Число отрицательное")
	}
}

func checkToPositivelyBool(x int) bool {
	if x >= 0 {
		return true
	}
	return false
}

func printWord(x int, s string) {
	for i := 0; i != x; i++ {
		fmt.Println(s)
	}
}

func checkYear(x int) bool {
	var b bool
	if x%4 != 0 {
		b = false
	} else if x%100 == 0 {
		if x%400 == 0 {
			b = true
		} else {
			b = false
		}
	} else {
		b = true
	}

	return b
}
