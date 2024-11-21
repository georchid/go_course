package main

type UserID int

func main() {
	idx := 1
	var uid UserID = 2

	// даже еслибазовый тип одинаковы, разные типы несовместимы
	// cannot use uid (type UserID) as type int64 in assignmen
	// myID := idx

	myID := UserID(idx)
}
