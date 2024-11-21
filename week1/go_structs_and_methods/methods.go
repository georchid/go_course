package main

import "fmt"

type Person struct {
	Id   int
	Name string
}

// не изменит оригинальной структуры, для которой вызван
// в метод передается копия объекта
func (p Person) UpdateName(name string) {
	p.Name = name
}

// изменит оригинальную структуру
// в метод передается адресс объекта типа
func (p *Person) SetName(name string) {
	p.Name = name
}

type Account struct {
	Id   int
	Name string
	Person
}

type MySlice []int

func (sl *MySlice) Add(val int) {
	*sl = append(*sl, val)
}

func (sl *MySlice) Count() int {
	return len(*sl)
}

func main() {
	pers := Person{1, "Vasily"}
	//pers := new(Person)
	pers.SetName("Vasily Romanov")
	// (&pers).SetName("Vasily Romanov")
	fmt.Printf("updated person: %#v\n", pers)

	// при встраивании методы наследуются
	var acc Account = Account{
		Id:   1,
		Name: "rvasily",
		Person: Person{
			Id:   2,
			Name: "George Chernyshev",
		},
	}
	acc.SetName("george.chernyshev")

	fmt.Printf("%#v\n", acc)

	sl := MySlice([]int{1, 2, 3, 4})
	sl.Add(3)
	fmt.Println(sl.Count(), sl)
}
