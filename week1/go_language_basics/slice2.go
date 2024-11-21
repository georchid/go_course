package main

import "fmt"

func main() {
	buf := []int{1, 2, 3, 4, 5}
	fmt.Println(buf)

	// получение среза, указывающего на ту же память
	sl1 := buf[1:4]
	sl2 := buf[:2]
	sl3 := buf[2:]
	fmt.Println(sl1, sl2, sl3)

	newBuf := buf[:]
	newBuf[0] = 9 //buf[0] = 9 тк указывают на одну и ту же память
	// newBuf теперь указывает на другие данные при превышении cap
	newBuf = append(newBuf, 6)
	newBuf[0] = 1
	fmt.Println("buf", buf)
	fmt.Println("newBuf", newBuf)

	//копирование одного слайса в другой
	var emptyBuf []int //len=0, cap=0

	copied := copy(emptyBuf, buf) //copied=0
	fmt.Println(copied, emptyBuf)

	newBuf = make([]int, len(buf), cap(buf))
	copy(newBuf, buf)
	fmt.Println(newBuf)

	// можно копировать в часть существующего массива
	ints := []int{1, 2, 3, 4}
	copy(ints[1:3], []int{5, 6})
	fmt.Println(ints)

}
