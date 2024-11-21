package main

import "fmt"

func main() {
	// цикл без условия while(true) или for(;;;)
	for {
		fmt.Println("Loop iteration")
		break
	}

	// цикл с одиночным условие while(isrun)
	isRun := true
	for isRun {
		fmt.Println("Loop iteration with condition")
		isRun = false
	}

	// цикл с условием и блоком инициализации
	for i := 0; i < 2; i++ {
		fmt.Println("Loop iteration", i)
		if i == 1 {
			continue
		}
	}

	// операции по slice
	sl := []int{1, 2, 3}
	idx := 0

	for idx < len(sl) {
		fmt.Println("While style loop, idx: ", idx, "value: ", sl[idx])
		idx++
	}

	for i := 0; i < len(sl); i++ {
		fmt.Println("c-style loop", i, sl[i])
	}

	for i := range sl {
		fmt.Println("range slice by index", i)
	}

	for idx, val := range sl {
		fmt.Println("range slice by idx-value", idx, val)
	}

	// операции по map
	profile := map[int]string{
		1: "Vasily",
		2: "Romanov",
	}

	for key := range profile {
		fmt.Println("range map by key", key)
	}

	for key, val := range profile {
		fmt.Println("range map by key val", key, val)
	}

	for _, val := range profile {
		fmt.Println("range map by val", val)
	}

	str := "Привет Мир!"
	for pos, char := range str {
		fmt.Printf("%#U at pos %d\n", char, pos)
	}
}
