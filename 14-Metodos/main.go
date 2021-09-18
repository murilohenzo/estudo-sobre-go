package main

import "fmt"

type UserDto struct {
	name string
	lastName  string
	age uint8
}

// metodo com parametros
func (user UserDto) create(name string, lastName string, age uint8) UserDto {
	user.name = name
	user.lastName = lastName
	user.age = age
	return user
}	

func (user UserDto) fullName() {
	fmt.Println(user.name + " " + user.lastName)
}	

// metodo com ponteiro
func (user *UserDto) incrementAge() {
	user.age++
}


func main() {

	user := UserDto{}
	createUser := user.create("murilo", "henzo", 22)
	fmt.Println(createUser)
	createUser.incrementAge()
	fmt.Println(createUser)


	user2 := UserDto{"John", "Doe", 21}
	fmt.Println(user2)
	user2.fullName()
	user2.incrementAge()
	fmt.Println(user2)
}