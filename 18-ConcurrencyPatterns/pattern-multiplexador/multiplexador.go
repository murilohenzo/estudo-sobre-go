package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	channel := multiplexador(write("Hello World"), write("Programming in Golang"))

	for i := 0; i < 10; i++ {
		fmt.Println(<-channel)
	}
}

// vai passar varios canais e retorn um canal
func multiplexador(inputChannel1 <-chan string, inputChannel2 <-chan string) <-chan string {
	outputChannel := make(chan string)

	go func() {
		for {
			select {
			case message := <-inputChannel1:
				outputChannel <- message
			case message := <-inputChannel2:
				outputChannel <- message
			}
		}
	}()

	return outputChannel
}

// encapsula a chamada de uma goroutine e retorna um canal de comunicacao para ser feita com essa goroutine
func write(text string) <-chan string {
	channel := make(chan string)

	go func() {
		for {
			channel <- fmt.Sprintf("Valor recebido %s", text)
			time.Sleep(time.Millisecond * time.Duration(rand.Intn(2000)))
		}
	}()
	return channel
}
