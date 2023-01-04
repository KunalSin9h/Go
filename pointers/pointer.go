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
}
