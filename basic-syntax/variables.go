package main

import (
	"fmt"
	"reflect"
)

var a, b, c = "A", "B", "C"

func main() {
	//var name = "Kunal Singh" /*Type Inference*/
	var name string
	var age int
	fmt.Println(name)
	fmt.Println(age)
	fmt.Println(reflect.TypeOf(name))
	fmt.Println(reflect.TypeOf(age))

	fmt.Println(a, b, c)

	// In Function (Not Global) we can ignore `var` and if we are providing values then
	// we can ignore types also (Go Type Inference)

	x := 10
	fmt.Println(x)

	var a, b = 1, "Kunal"
	fmt.Println(a, b)

}
