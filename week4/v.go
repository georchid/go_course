package main

import "fmt"

func main() {
	v := 1
	inc := func() {
		v++
	}

	fmt.Println(v)
	inc()
	fmt.Println(v)
}
