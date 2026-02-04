package main

import (
	"fmt"
	"sync"
)

func main() {
	fmt.Println("Range sobre inteiros:")
	for i := range 5 {
		fmt.Println("i:", i)
	}

	fmt.Println("\nVariáveis de range por iteração:")
	var wg sync.WaitGroup
	for i := range 3 {
		wg.Add(1)
		go func() {
			defer wg.Done()
			fmt.Println("goroutine:", i)
		}()
	}
	wg.Wait()
}
