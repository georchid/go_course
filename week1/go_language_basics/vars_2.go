package main

import "fmt"

func main() {
	var i int = 10
	var autoInt = -10

	//int8 int16 int32 int64, аналогично uint
	var bigInt int64 = 1<<32 - 1
	var unsignedInt uint = 100500

	fmt.Println(i, autoInt, bigInt, unsignedInt)

	// float32 float 64
	var pi float32 = 3.141
	var e float64 = 2.718
	goldenRatio := 1.618

	fmt.Println(pi, e, goldenRatio)

	//bool
	//по умолчанию false
	var b bool
	var isOk bool = true
	var success = true
	cond := true

	fmt.Println(b, isOk, success, cond)

	//complex64, complex128
	var c complex128 = -1.1 + 7.12i
	c2 := -1.1 - 7.12i

	fmt.Println(c, c2)

}
