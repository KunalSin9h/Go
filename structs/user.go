package main

import (
	"fmt"
)

/*
User
It's Like C's Struct i.e Their main purpose is Encapsulation
User is user/developer defined type
*/
type User struct {
	ID                         int
	Email, FirstName, LastName string
}

/*
Group
Nested Structs
After all struct is just a user defined type
users -> []User is a `Slice` not an `Array` coz we haven't given the size and arrays are fixed sized
*/
type Group struct {
	role           string
	users          []User
	newestUser     User
	spaceAvailable bool
}

func describeUser(user User) (data string) {
	data = fmt.Sprintf("Hello, meet %s, his email is %s", user.FirstName, user.Email)
	return
}

func describeGroup(gr Group) (data string) {

	var aval bool

	if len(gr.users) < 2 {
		aval = true
	} else {
		aval = false
	}

	data = fmt.Sprintf("This user groups has %d users\nThe newest user is %s, Accepting new users: %t",
		len(gr.users), gr.newestUser.FirstName, aval)
	return
}

func main() {

	var u1 User = User{
		ID:        1,
		Email:     "kunal@gmail.com",
		FirstName: "Kunal",
		LastName:  "Singh",
	}

	var u2 User = User{
		ID:        2,
		Email:     "kunal@gmail.com",
		FirstName: "Kunal",
		LastName:  "Singh",
	}

	/*	var x = make([]User, 2)
		x[0] = u1
		x[1] = u2*/

	var grp Group = Group{
		role: "normal",
		//users:          x,
		users:          []User{u1, u2}, // literal
		newestUser:     u2,
		spaceAvailable: false,
	}

	/*	greet := describeUser(u1)
		fmt.Println(greet)*/

	fmt.Println(describeGroup(grp))

	var a []int

	for i := 1; i <= 10; i++ {
		a = append(a, i)
	}

	fmt.Println(a)

}
