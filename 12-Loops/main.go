package main

import (
	"fmt"
	"time"
)

func main() {

	for i:= 0; i < 5; i++ {
		fmt.Println(i)
		time.Sleep(time.Second)
	}

	var numbers = [5]uint8{1, 2, 3, 4, 5}

	for index, value := range numbers {
		fmt.Println("Index [",index,"] = ",value)
	}

	user := map[string]string {
		"name": "John Doe",
		"email": "johndoe@gmail.com",
	}

	for key, value := range user {
		fmt.Println("Key [",key,"] = ",value)
	}

	type IUser struct {
		name string
		email string
	}

	user2 := IUser{name: "John Doe", email: "johndoe@gmail.com"}
	fmt.Println(user2)
	// for key, value := range user2 {
	// 	fmt.Println("Key [",key,"] = ",value)
	// }
}