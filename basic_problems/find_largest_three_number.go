package find_largest_number

import "fmt"

func Find_largest_number()  {
	var num1, num2, num3 int;

	fmt.Printf("First number: ")
	fmt.Scan(&num1)
	fmt.Printf("Second number: ")
	fmt.Scan(&num2)
	fmt.Printf("Third number: ")
	fmt.Scan(&num3)

	largest := num1

	if(num2 > largest){
		largest = num2
	}
	if(num3 > largest) {
		largest = num3
	}


	fmt.Println("The largest number is: ", largest)
}