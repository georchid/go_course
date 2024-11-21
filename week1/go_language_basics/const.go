package main

import "fmt"

const pi = 3.141
const (
	hello = "Привет"
	e     = 2.718
)
const (
	//
	zero  = iota
	_     // пропуск, автоинкремента iota
	three // 3
)
const (
	_         = iota
	KB uint64 = 1 << (10 * iota) //1024
	MB                           //1048576
)
const (
	year          = 2017 // нетипизированная константа
	yearTyped int = 2017 //типизированная константа
)

func main() {
	var month int32 = 13
	fmt.Println(month + year)

	// month + yearTyped(mismatched types int32 and int)
	// fmt.Println(month + year)

}
