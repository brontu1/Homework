package main

import (
	"fmt"
	"math/rand"
)

const (
	MapSize = 3
)

var gameMap [MapSize][MapSize]string

func main() {
	createMap()
	for {
		movimientoDelHombre()
		printMap()
		if checkWin() == true {
			break
		}
		progresoDeLaComputadora()
		printMap()
		if checkWin() == true {
			break
		}
	}
}

func createMap() {
	initializeMap()
	printMap()
}

func initializeMap() {
	for i := 0; i < MapSize; i++ {
		for j := 0; j < MapSize; j++ {
			gameMap[i][j] = "_"
		}
	}
}

func printMap() {
	for _, i := range gameMap {
		for _, j := range i {
			fmt.Printf("|%s|", j)
		}
		fmt.Println()
	}
}

func emptyFieldCheck(x, y int) bool {
	if gameMap[x][y] == "_" {
		return true
	}
	return false
}

func movimientoDelHombre() { //Ход мясного мешка
	var x, y int
	fmt.Println("Movimiento del hombre, seleccione coordenadas de campo (x, y)")
	fmt.Scan(&x, &y)
	if !emptyFieldCheck(x-1, y-1) {
		fmt.Println("Este campo no está vacío")
		movimientoDelHombre()
		return
	}
	gameMap[x-1][y-1] = "x"
}

func progresoDeLaComputadora() { //Ход компьютера
	x := rand.Intn(MapSize)
	y := rand.Intn(MapSize)
	if !emptyFieldCheck(x, y) {
		progresoDeLaComputadora()
		return
	}
	fmt.Println("Progreso de la computadora:")
	gameMap[x][y] = "0"
}

func checkWin() bool {
	var countHS int                                                           //Переменная счетчика для проверки мясного мешка по вертикали
	var countCS int                                                           //Переменная счетчика для проверки комьютера по вертикали
	if gameMap[0][0] == "x" && gameMap[1][1] == "x" && gameMap[2][2] == "x" { //Проверка по основной диагонали
		fmt.Println("El hombre gano")
		return true
	} else if gameMap[0][0] == "0" && gameMap[1][1] == "0" && gameMap[2][2] == "0" {
		fmt.Println("La computadora gano")
		return true
	} else if gameMap[0][2] == "x" && gameMap[1][1] == "x" && gameMap[2][0] == "x" { //Проверка по вторичной диагонали
		fmt.Println("El hombre gano")
		return true
	} else if gameMap[0][2] == "0" && gameMap[1][1] == "0" && gameMap[2][0] == "0" {
		fmt.Println("La computadora gano")
		return true
	}
	for i := 0; i < MapSize; i++ {
		var countH int //Переменная счетчика для проверки мясного мешка по горизонтали
		var countC int //Переменная счетчика для проверки компьютера по горизонтали
		var y int      //Костыль-счётчик для сохранения j-ого значения для проверки по вертикали
		if gameMap[i][y] == "x" {
			countHS++
			if countHS == 3 {
				fmt.Println("El hombre gano")
				return true
			}
		} else if gameMap[i][y] == "0" {
			countCS++
			if countCS == 3 {
				fmt.Println("La computadora gano")
				return true
			}
		}

		for j := 0; j < MapSize; j++ {
			if gameMap[i][j] == "x" {
				countH++
				if countH == 3 {
					fmt.Println("El hombre gano")
					return true
				}
			} else if gameMap[i][j] == "0" {
				countC++
				if countC == 3 {
					fmt.Println("La computadora gano")
					return true
				}
			}
			y++
		}
	}
	return false
}
