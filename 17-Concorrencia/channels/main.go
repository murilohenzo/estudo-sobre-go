package main

import (
	"fmt"
	"time"
)

// Os canais são os tubos que conectam goroutines simultâneas.
// Você pode enviar valores para canais de uma goroutine e receber esses valores em outra goroutine.

func main() {

	// Create a new channel with make(chan val-type). Channels are typed by the values they convey.
	channel := make(chan string)

	go write("Hello World", channel)

	fmt.Println("Depois da funcao escrever e a comecar a ser executada")

	// The <-channel syntax receives a value from the channel.
	// Here we’ll receive the "ping" message we sent above and print it out.

	for message := range channel {
		fmt.Println(message)
	}
	fmt.Println("Fim do programa !")
}

func write(text string, channel chan string) {
	for i := 0; i < 5; i++ {
		// Send a value into a channel using the channel <- syntax.
		// Here we send "ping" to the messages channel we made above, from a new goroutine.
		channel <- text
		time.Sleep(time.Second)
	}

	close(channel)
}
