package main

import "fmt"

func main() {
	print("Добро пожаловать на охуенную викторину\nПравила простые, ответы на вопросы: Да или нет\nПогнали\n")
	counter := 0
	print("Чей Крым?:\n1-Да, 2-Нет\n")
	var answer int
	fmt.Scan(&answer)
	switch answer {
	case 1:
		println("Чел харош")
		counter++
	case 2:
		println("Чел нехарош")
	}
	print("Сколько?:\n1-Да, 2-Нет\n")
	fmt.Scan(&answer)
	switch answer {
	case 1:
		println("Чел харош")
		counter++
	case 2:
		println("Чел нехарош")
	}
	print("Ы?:\n1-Да, 2-Нет\n")
	fmt.Scan(&answer)
	switch answer {
	case 1:
		println("Чел харош")
		counter++
	case 2:
		println("Чел нехарош")
	}
	if counter == 3 {
		println("Чел, поздравляю, мегахарош")
	} else {
		println("Чел, мегаплох, хошь еще?\n1-Да, 2-Нет\n")
		fmt.Scan(&answer)
		switch answer {
		case 1:
			println("Чел харош")
			restart()
		case 2:
			println("Чел нехарош")
			break
		}
	}
}

func restart() {
	main()
}
