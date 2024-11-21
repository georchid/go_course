package main

import "fmt"

func main() {
	// простое условие
	boolVal := false
	if boolVal {
		fmt.Println("boolVal is true!")
	}

	mapVal := map[string]string{"name": "Vasily"}

	// условие с блоком инициализации
	if keyValue, keyExist := mapVal["name"]; keyExist {
		fmt.Println("name = ", keyValue)
	}
	// получаем только признак существования ключа
	if _, keyExist := mapVal["name"]; keyExist {
		fmt.Println("key 'name' exist")
	}

	cond := 1
	// множественные if else
	if cond == 1 {
		fmt.Println("cond is 1")
	} else if cond == 2 {
		fmt.Println("cond is 2")
	} else {
		fmt.Println("cond digger =O")
	}

	// switch по 1 переменной
	strVal := "name"
	switch strVal {
	case "name":
		fallthrough //явно провалиться в следующее условие (как по дефлдту в си++ без break)
	case "test", "lastName":
		//some work
	default:
		//some work
	}

	// switch как замена логики if else
	var val1, val2 = 2, 2
	switch {
	case val1 > 1 || val2 < 11:
		fmt.Println("first block")
	case val2 > 10:
		fmt.Println("second block")
	}

	// выход из цикла находясь внутри switch
Loop:
	for key, val := range mapVal {
		println("switch in loop", key, val)
		switch {
		case key == "lastName":
			break
			println("dont print this")
		case key == "firstName" && val == "Vasily":
			println("switch - break loop here")
			break Loop
		}
	} //конец for
}
