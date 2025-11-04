package main

import "fmt"

const x = 10

var y = 20


func outer() func() {
	// this is stored in heap by escape analysis after the function in run
	money := 2000

	// this is stored in heap by escape analysis after the function in run
	show := func() {
		money = money + x + y
		fmt.Println(money)
	}
	
	return show
}

func callOuter() {
	incr1:= outer()
	incr1()
	incr1()

	incr2 := outer()
	incr2()
	incr2()
}

func main() {
	callOuter()
}