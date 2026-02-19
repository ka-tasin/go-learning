package main

import (
	"fmt"
)

type User struct {
	name  string
	email string
}

func (usr User) receiver() {
	fmt.Println("Name: ", usr.name)
	fmt.Println("Email: ", usr.email)
	usr.name = "doe"
	fmt.Println("Receiver function changed name: ", usr.name)
}

func (usr *User) pointerReceiver() {
	fmt.Println("Name: ", usr.name)
	fmt.Println("Email: ", usr.email)
	usr.name = "doe"
	fmt.Println("Pointer receiver function changed name: ", usr.name)
}

func main() {
	user1 := User{
		name: "John",
		email: "john@gmail.com",
	}
	user2 := User{
		name: "Isac",
		email: "isac@gmail.com",
	}

	user1.receiver()
	user2.pointerReceiver()

	// Receiver function only changes the local value
	fmt.Println("main function reciever name: ", user1.name)

	// Pointer receiver function changes the global value
	fmt.Println("main function pointer reciever name: ", user2.name)
}
