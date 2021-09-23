package main

import "fmt"

// padrao utilizado para resolver uma fila de tarefas grandes para serem executadas

func main() {
	tasks := make(chan int, 40)
	results := make(chan int, 40)

	go worker(tasks, results)
	go worker(tasks, results)
	go worker(tasks, results)
	go worker(tasks, results)
	go worker(tasks, results)
	go worker(tasks, results)
	go worker(tasks, results)
	go worker(tasks, results)

	for i := 0; i < 40; i++ {
		tasks <- i
	}
	close(tasks)

	for i := 0; i < 40; i++ {
		fmt.Println(<-results)
	}
}

func worker(tasks <-chan int, results chan<- int) {
	for taskNumber := range tasks {
		results <- fibonacci(taskNumber)
	}
}

func fibonacci(position int) int {
	if position <= 1 {
		return position
	}

	return fibonacci(position-2) + fibonacci(position-1)
}
