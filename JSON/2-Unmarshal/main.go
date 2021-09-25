package main

import (
	"encoding/json"
	"fmt"
	"log"
)

type IUser struct {
	Name    string            `json:"name"`
	Age     uint8             `json:"age"`
	Courses map[string]string `json:"courses"`
}

func main() {
	userInJSON := `{"name":"John Doe","age":21,"courses":{"name":"Golang"}}`

	var user IUser

	if erro := json.Unmarshal([]byte(userInJSON), &user); erro != nil {
		log.Fatal(erro)
	}
	fmt.Println(user)

}
