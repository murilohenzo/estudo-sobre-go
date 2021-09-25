package main

import (
	"bytes"
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
	user := IUser{
		"John Doe",
		21,
		map[string]string{
			"name": "Golang",
		},
	}
	userToJSON, erro := json.Marshal(user)
	if erro != nil {
		log.Fatal(erro)
	}
	fmt.Println(userToJSON)

	fmt.Println(bytes.NewBuffer(userToJSON))
}
