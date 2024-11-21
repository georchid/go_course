package main

import (
	"fmt"
)

func deferTest() {
	defer func() {
		if err := recover(); err != nil {
			fmt.Println("panic happened FIRST:", err)
		}
	}()

	defer func() {
		if err := recover(); err != nil {
			fmt.Println("panic happend SECOND:", err)
			panic(2)
		}
	}()
	fmt.Println("some useful work")
	panic("something bad happened")
	return
}

func main() {
	deferTest()
	return
}
