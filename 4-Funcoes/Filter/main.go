package main

import "fmt"

type FilterCallback = func(value int) bool

func filter(array []int, callback FilterCallback) []int {
	newArray := []int{}

	for i := 0; i < len(array); i++ {
		if callback(array[i]) {
			newArray = append(newArray, array[i])
		}
	}
	return newArray
}

var array = []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

func main() {
	fmt.Println(filter(array, func(number int) bool {
		return number%2 == 0
	}))

	fmt.Println(filter(array, func(number int) bool {
		return number%2 != 0
	}))
}