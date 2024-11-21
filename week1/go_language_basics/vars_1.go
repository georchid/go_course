package main

import "fmt"

func main() {
	// значение по умолчанию
	var num0 int

	// значение при инициализации
	var num1 int = 1

	// пропуск типа
	var num2 = 20

	fmt.Println(num0, num1, num2)

	// короткое объявление переменной
	num := 30
	// только для новых переменных
	// num := 20 вызовет ошибку поскольку переменная уже есть

	num += 1
	fmt.Println("+=", num)

	num++
	fmt.Println("++", num)

	//CamelCase - принятый стиль
	UserIndex := 10
	//under_score - не принято
	user_index := 10
	fmt.Println(UserIndex, user_index)

	// объявление нескольких переменных
	var weight, height int = 10, 20

	// присваиваниев существующие переменные
	weight, height = 30, 40

	// короткое присваивание
	// хотя бы одна переменная должна быть новой
	weight, age := 31, 18

	fmt.Println(weight, height, age)

}
