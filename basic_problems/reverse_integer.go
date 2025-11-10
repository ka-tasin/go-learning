package main

import "fmt"

func reverse_integer(n int) int {
	reverse := 0

	for n != 0 {
		reverse = (reverse * 10 ) + n % 10
		n = n / 10
	}

	return reverse
}


func main() {
	var res int = reverse_integer(1234567)

	fmt.Println(res)

}