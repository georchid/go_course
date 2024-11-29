package main

import "fmt"

func main() {
	ch := make(chan int)

	go func(out chan<- int) {
		for i := 0; i <= 4; i++ {
			fmt.Println("before ", i)
			out <- i
			fmt.Println("after ", i)
		}
		close(out)
		fmt.Println("generation finish")
	}(ch)

	for i := range ch {
		fmt.Println("\tget", i)
	}

}
