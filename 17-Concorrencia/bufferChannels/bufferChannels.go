package main

import "fmt"

func main() {
	// canal com buffer so vai bloquear a operacao de canal, quando ocupar toda a capacidade dele
	channel := make(chan string, 2) //capacidade de armazenamento = 2

	fmt.Println("Capacidade = ", cap(channel))

	channel <- "Hello World"
	channel <- "Hello World 2"

	fmt.Println(<-channel)
	fmt.Println(<-channel)
}
