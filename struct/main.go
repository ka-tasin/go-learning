package main
import "fmt"

type User struct {
	name string
	age int
}

func main() {
	var user1 User //instantiate
	user1.name = "tasin"
	user1.age = 27

	// Instance or Object
	user2 := User{
		name : "dh",
		age : 36,
	}

	fmt.Println(user1)
	fmt.Println(user2)
}