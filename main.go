package main

import "fmt"

func greet() {
	fmt.Println("Hello, World!")
}

func main() {
	x := greet

	x()
}
