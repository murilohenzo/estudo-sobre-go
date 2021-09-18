package main

import "fmt"

func generic( interf interface{}) {
	fmt.Println(interf)
}

func main() {
	generic("Hello")
	generic(1)
	generic(true)
}