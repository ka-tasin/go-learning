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
}

func main() {
	user1 := User{
		name: "John",
		email: "john@gmail.com",
	}

	user1.receiver()
}
