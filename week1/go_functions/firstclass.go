package main

import "fmt"

func doNothing() {
	fmt.Println("i'm regular function")
}

func main() {
	// анонимная функция
	func(in string) {
		fmt.Println("anon func out:", in)
	}("nobody")
	// присваивание анонимной функции в переменную
	printer := func(in string) {
		fmt.Println("printer outs:", in)
	}

	printer("as variable")

	// определяем тип функции
	type strTypeFunc func(string)

	// функция принимает коллбек - переменную
	worker := func(callback strTypeFunc) {
		callback("as callback")
	}

	worker(printer)

	// функция возвращает замыкание
	// функцию, хранящую состояние функции, генерирующей ее
	prefixer := func(prefix string) strTypeFunc {
		return func(in string) {
			fmt.Printf("[%s] %s\n", prefix, in)
		}
	}
	successLogger := prefixer("SUCCESS")
	successLogger("expected behaviour")
}
