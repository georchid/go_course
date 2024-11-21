package main

import "fmt"

type Person struct {
	Id      int
	Name    string
	Address string
}

type Account struct {
	Id      int
	Name    string
	Cleaner func(string) string
	Owner   Person
	Person  // часть самой структуры
}

func main() {
	// полное объявление структуры
	// можем пропускать поля
	// пропущенные получают значения по умолчанию
	var acc Account = Account{
		Id:   1,
		Name: "rvasily",
		Person: Person{
			Name:    "Василий",
			Address: "Moscow",
		},
	}

	fmt.Printf("%#v\n", acc)
	// короткое объявление структуры
	// поля пропускать не можем
	acc.Owner = Person{2, "Romanov Vasily", "Moscow"}

	fmt.Printf("%#v\n", acc)

	//втроили структуру в структуру
	fmt.Println(acc.Address)

	//приоритет к наиболее верхнему полю если есть совпадающие поля
	fmt.Println(acc.Name)
	fmt.Println(acc.Person.Name)
}
