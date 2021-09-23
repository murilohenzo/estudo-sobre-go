package main

import (
	"fmt"
	"time"
)

func main() {
	channel := write("Hello world !")

	for i := 0; i < 5; i++ {
		fmt.Println(<-channel)
	}
}

// encapsula a chamada de uma goroutine e retorna um canal de comunicacao para ser feita com essa goroutine
func write(text string) <-chan string {
	channel := make(chan string)

	go func() {
		for {
			channel <- fmt.Sprintf("Valor recebido %s", text)
			time.Sleep(time.Millisecond * 500)
		}
	}()
	return channel
}
