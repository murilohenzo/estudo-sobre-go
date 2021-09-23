package main

import (
	"fmt"
	"time"
)

func main() {

	channel1, channel2 := make(chan string), make(chan string)

	go func() {
		for {
			time.Sleep(time.Millisecond * 500)
			channel1 <- "Channel 1"
		}
	}()

	go func() {
		for {
			time.Sleep(time.Second * 2)
			channel2 <- "Channel 2"
		}
	}()

	for {
		select {
		// caso o canal 1 esteja pronto para receber o valor do canal 1, printa o valor
		case <-channel1:
			fmt.Println(<-channel1)
		case <-channel2:
			fmt.Println(<-channel2)
		}

	}

}
