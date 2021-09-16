package main

import "fmt"


func add(x int, y int) int {
	return x + y
}

func add2(x, y int) int {
	return x + y
}

func evenNumbersAndOddNumbers(number int) bool {
	if number % 2 == 0 {
		return true
	}
	return false
}

func filterOddNumbers (array []int, callback func(int) bool) ([]int, []int) {
	even_list, odd_list := []int{}, []int{}
	for _, element := range array {
		result := callback(element)
		if result == true {
			even_list = append(even_list, element)
		} else {
			odd_list = append(odd_list, element)
		}
	}
	return even_list, odd_list
}

var numbers = []int{1, 2, 3, 4, 5, 6}


func main() {
	res := add(42, 13)
	res2 := add2(42, 13)
	fmt.Println(res)
	fmt.Println(res2)
	evenNumbers , oddNumbers := filterOddNumbers(numbers, evenNumbersAndOddNumbers)
	fmt.Println(evenNumbers)
	fmt.Println(oddNumbers)
}