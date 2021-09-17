package main

import "fmt"

func main() {
	// chaves e valores sao todos do mesmo tipo
	user := map[string]string {
		"name": "John Doe",
		"lastname": "Doe",
	}

	fmt.Println(user)
	fmt.Println(user["name"])
	fmt.Println(user["lastname"])

	// maps aninhados
	user2 := map[string]map[string]string {
		"name": {
			"firstname": "John Doe",
			"lastname": "Doe",
		},
	}

	fmt.Println(user2)
	fmt.Println(user2["name"]["firstname"])
	fmt.Println(user2["name"]["lastname"])

	delete(user2["name"], "lastname")

	fmt.Println(user2)

	user2["language"] = map[string]string {
		"name": "English",
	}

	fmt.Println(user2)

}