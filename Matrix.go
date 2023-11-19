package main

import "fmt"

func main() {
	matrix1()
	fmt.Println("")
	matrix2()
	fmt.Println("")
	matrix3()
}

func matrix1() {
	for i := 0; i < 6; i++ {
		for j := 0; j < 6; j++ {
			if i == 0 {
				fmt.Print("0 ")
			} else if i == 5 {
				fmt.Print("0 ")
			} else {
				if j == 0 {
					fmt.Print("0 ")
				} else if j == 5 {
					fmt.Print("0 ")
				} else {
					fmt.Print("* ")
				}
			}
		}
		fmt.Println("")
	}
}

func matrix2() {
	for i := 0; i < 6; i++ {
		for j := 0; j < 6; j++ {
			if i%2 == 0 {
				if j%2 == 0 {
					fmt.Print("* ")
				} else {
					fmt.Print("0 ")
				}
			} else {
				if j%2 == 0 {
					fmt.Print("0 ")
				} else {
					fmt.Print("* ")
				}
			}
		}
		fmt.Println("")
	}
}

func matrix3() {
	for i := 0; i < 6; i++ {
		for j := 0; j < 5; j++ {
			if i == 0 {
				fmt.Print("* ")
			} else if i == 1 {
				if j == 0 {
					fmt.Print("0 ")
				} else {
					fmt.Print("* ")
				}
			} else if i == 2 {
				if j <= 1 {
					fmt.Print("0 ")
				} else {
					fmt.Print("* ")
				}
			} else if i == 3 {
				if j <= 2 {
					fmt.Print("0 ")
				} else {
					fmt.Print("* ")
				}
			} else if i == 4 {
				if j <= 3 {
					fmt.Print("0 ")
				} else {
					fmt.Print("* ")
				}
			} else {
				fmt.Print("0 ")
			}
		}
		fmt.Println("")
	}
}
