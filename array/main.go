package main

import "fmt"

func main() {
	/* 
		- array size is fixed cannot grow or shrink
		- array cannot be const
	*/
	var arr = []int{1, 3, 4, 6}
	fmt.Println(arr)
	fmt.Println(arr[3])

}