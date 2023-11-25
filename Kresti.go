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

func movimientoDelHombre() {
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

func progresoDeLaComputadora() {
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
	if gameMap[MapSize-1][MapSize-1] == "x" {
		fmt.Println("El hombre gano")
		return true
	} else if gameMap[MapSize-1][MapSize-1] == "0" {
		fmt.Println("La computadora gano")
		return true
	}
	return false
}
