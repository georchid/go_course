package main

import (
	"fmt"
)

func main() {
	ch1 := make(chan int)
	ch2 := make(chan int)
	select {
	case val := <-ch1:
		fmt.Println("ch1 val: ", val)
	case ch2 <- 42:
		fmt.Println("put val to ch2")
	default:
		fmt.Println("default case")
	}
}
