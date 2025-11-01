package main

import "fmt"

var num1 int = 13
var num2 int = 45

func add(num1 int, num2 int) int {
	return num1 + num2
}

func main() {
	var a int = 2
	var b int = 4
	fmt.Println("Hello World")

	fmt.Println(add(num1, num2))

	fmt.Println(add(34, 745))

	fmt.Println(add(a, b))
}