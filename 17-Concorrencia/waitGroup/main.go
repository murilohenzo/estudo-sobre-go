package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {

	var waitGroup sync.WaitGroup

	// adiciona duas go routine ao grupo de espera
	waitGroup.Add(2)

	go func() {
		write("Go routine 1")
		waitGroup.Done() // -1
	}()

	go func() {
		write("Go routine 2")
		waitGroup.Done() // -1
	}()

	waitGroup.Wait() // 0
}

func write(text string) {
	for i := 0; i < 5; i++ {
		fmt.Println(text)
		time.Sleep(time.Second)
	}
}
