package main

import "fmt"

func dayOfTheWeek(number int) string {
	switch number {
	case 1:
		return "Sunday"
	case 2:
		return "Monday"
	case 3:
		return "Tuesday"
	case 4:
		return "Wednesday"
	case 5:
		return "Thursday"
	case 6:
		return "Friday"
	case 7:
		return "Saturday"
	default:
		return "Unknown"
	}
}

func dayOfTheWeek2(number int) string {
	var dayOfTheWeek string

	switch {
	case number == 1:
		dayOfTheWeek = "Sunday"
		fallthrough // joga o codigo para a proxima condicao
	case number == 2:
		dayOfTheWeek = "Monday"
	case number == 3:
		dayOfTheWeek = "Tuesday"
	case number == 4:
		dayOfTheWeek = "Wednesday"
	case number == 5:
		dayOfTheWeek = "Thursday"
	case number == 6:
		dayOfTheWeek = "Friday"
	case number == 7:
		dayOfTheWeek = "Saturday"
	default:
		dayOfTheWeek = "Unknown"
	}
	return dayOfTheWeek
}

func main() {
	fmt.Println(dayOfTheWeek(2))
	fmt.Println(dayOfTheWeek2(1))
}