package main

import "fmt"

func getSomeVars() string {
	fmt.Println("getSomeVars executinon")
	return "getSomeVars result"
}

func main() {
	// отложенное выполнение функции
	defer fmt.Println("After work")

	// аргументы вычисляются в момент объявления defer
	//defer fmt.Println(getSomeVars())
	fmt.Println("Some useful work")

	// если хотим выполнения не в момент объявления defer
	defer func() {
		fmt.Println(getSomeVars())
	}()
}
