package main

import (
	"fmt"
	"reflect"
)

/**
Types
*/

/*
*
Integer

	int,  int8,  int16,  int32,  int64
	uint, uint8, uint16, uint32, uint64
*/
func integer() {
	var num uint8 = 255
	fmt.Println(num)
	fmt.Println(reflect.TypeOf(float64(num) * 5.5))
}

/**
Float
	float32, float64
*/

/**
String
	string
*/
// Single Package should not create mutiple main function
/*
func main() {
	integer()
}
*/
