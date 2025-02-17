package main

import (
	"encoding/json"
	"fmt"
)

type User struct {
	ID       int
	Username string
	phone    string
}

var jsonString = `{"id": 42, "username": "George", "phone": "123" }`

func main() {
	data := []byte(jsonString)
	u := &User{}
	json.Unmarshal(data, u)
	fmt.Printf("struct: \n\t%#v\n\n", u)
	u.phone = "345"

	result, err := json.Marshal(u)
	if err != nil {
		panic(err)
	}
	fmt.Printf("json string:\n\t%s\n", string(result))
}

// Поскольку в го объекты именованные с маленькой буквы видны только в пределах пакета, то поле phone недоступно
// для пакета json
