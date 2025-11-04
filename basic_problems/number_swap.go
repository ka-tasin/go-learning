package main
import "fmt"

func main() {
	var a int
	fmt.Scan(&a)
	var b int
	fmt.Scan(&b)

	a = a + b
	b = a - b
	a = a - b

	fmt.Println(a, b)
}