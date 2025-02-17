package main

import (
	"encoding/json"
	"fmt"
)

type User struct {
	ID       int `json:"user_id, string"`
	Username string
	Address  string `json:",omitempty"`
	Company  string `json:"-"`
}

func main() {
	u := &User{
		ID:       42,
		Username: "George",
		Address:  "Saint-Petersburg",
		Company:  "Bla",
	}
	result, _ := json.Marshal(u)
	fmt.Printf("json string: \n%s\n", result)
}
