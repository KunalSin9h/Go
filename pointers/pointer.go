package main

import (
	"fmt"
	"strings"
)

/*
Pass By Reference Achieved By Pointers only
*/
func changeName(name *string) {
	*name = strings.ToTitle(*name)
}

type User struct {
	ID          int
	Name, Email string
}

func updateEmail(user *User, email string) {
	fmt.Println(user.Name)
	user.Email = email
}

func main() {
	var name string
	var namePointer *string

	name = "Kunal"
	namePointer = &name
	var nameValue = *namePointer

	fmt.Println(name)
	fmt.Println(namePointer)
	fmt.Println(&name)
	fmt.Println(&namePointer)
	fmt.Println(nameValue)

	n := "kunal"
	changeName(&n)
	fmt.Println(n)

	u := User{
		ID:    1,
		Name:  "Kunal",
		Email: "kunal@knl.sh",
	}

	updateEmail(&u, "x@knl.sh")

	fmt.Println(u)

}
