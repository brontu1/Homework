package main

import "fmt"

func main() {
	var value1, value2 int
	println("Введите числа для функции проверки числа:")
	fmt.Scan(&value1, &value2)
	var gayColor int
	println("Введите индекс цвета:")
	fmt.Scan(&gayColor)
	var compare1, compare2 int
	println("Сравним две хуйни:")
	fmt.Scan(&compare1, &compare2)
	printThreeWords()
	checkSumSign(value1, value2)
	LGBTColor(gayColor)
	print(compareNumbers(compare1, compare2))
}

func printThreeWords() {
	print("Orange\nBanana\nApple\n")
}

// Знаю, что в задании нужно проинициализировать две переменные внутри функции, но я решил сделать элегантнее *ВинниПух_вКостюме.жпг*
func checkSumSign(value1, value2 int) {
	if value1+value2 > 0 {
		println("Сумма положительная")
	} else if value1+value2 == 0 {
		println("Сумма нулевая")
	} else {
		println("Сумма отрицательная")
	}
}

func LGBTColor(gayColor int) {
	if gayColor <= 0 {
		println("Это красный, братан")
	} else if gayColor > 0 && 100 >= gayColor {
		println("Это, браток, цвет мочи")
	} else {
		println("Зелёный, бля")
	}
}

func compareNumbers(a, b int) string {
	if a >= b {
		return "a >= b"
	} else {
		return "a < b"
	}
}
